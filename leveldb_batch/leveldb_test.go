package leveldb_batch_test

import (
	//	"sync"
    "os"
	"leveldb_batch"
	"leveldb"
	"testing"
)

var test_values = []string{"", "a", "1251", "汉字", "\x00123\x00","sdfaewrqerf汉大幅度发顺丰"}

func TestLDB_Batch(t *testing.T) {
	db, err := leveldb.OpenFile("tmp/test.db", nil)
	if err != nil {
		t.Fatalf("Failed to create databas %v", err)
	}
	defer testLDB_Close(db, t)
    t.Logf("*************数据库批量读写测试*************");
    t.Logf("*************测试数据*****************");
    for _,str:=range test_values{
        t.Logf("[%s]",str);
    }
    t.Logf("*************测试数据*****************");
    testBatchWriteRead(db,t)
}

func testLDB_Close(db *leveldb.DB, t *testing.T) {
	err := db.Close()
    os.RemoveAll("tmp")
	if err == nil {
		t.Logf("Database closed")

	} else {
		t.Errorf("Failed to close database: %v", err)
	}
}

func testRead(db *leveldb.DB,t *testing.T,limit int){
	readBatch := leveldb_batch.NewReadBatch(db)
    key,value:= readBatch.Read(limit)
    for index,_:=range key{
        t.Logf("index[%d] key[%s] value[%s]",index,string(key[index]),string(value[index]))    
    }
}

func testBatchWriteRead(db *leveldb.DB,t *testing.T) {
	t.Parallel()
	writeBatch := leveldb_batch.NewWriteBatch(db)
    seq:=1
    t.Logf(" ")
    t.Logf("+++++++++++++++++++批量写test %d: key:测试数据 value:nil+++++++++++++++++++++++++++++++",seq) 
	for _, k := range test_values {
		writeBatch.Put([]byte(k), nil)
	}
    t.Logf("*************批量写入前，应该没有数据*****************************");
    testRead(db,t,0)
	err := writeBatch.Write()
	if err != nil {
		t.Fatalf("Batch write failed: %v", err)
	}
    t.Logf("*************批量写入后,value应该都是空值*************************");
    testRead(db,t,0)
    t.Logf("++++++++++++++++++++批量写test %d:end+++++++++++++++++++++++++++++++++++++++++++++++++++",seq)
    seq++
    t.Logf(" ")

    t.Logf("+++++++++++++++++++批量写test %d: key:测试数据 value:与key同样的数据++++++++++++++++++++",seq) 
	for _, v := range test_values {
		writeBatch.Put([]byte(v), []byte(v))
	}
    t.Logf("*************批量写入前，打印数据应该与前一次测试相同*************");
    testRead(db,t,0)
	err = writeBatch.Write()
	if err != nil {
		t.Fatalf("Batch write failed: %v", err)
	}
    t.Logf("*************批量写入后，打印数据应该是新的测试数据***************");
    testRead(db,t,0)
    t.Logf("++++++++++++++++++++批量写test %d:end+++++++++++++++++++++++++++++++++++++++++++++++++++",seq)
    seq++
    t.Logf(" ")

    t.Logf("+++++++++++++++++++批量写test %d: key:测试数据 value:都是'?'++++++++++++++++++++++++++++",seq) 
	for _, v := range test_values {
		writeBatch.Put([]byte(v), []byte("?"))
	}
    t.Logf("*************批量写入前，打印数据应该与前一次测试相同*************");
    testRead(db,t,0)
	err = writeBatch.Write()
	if err != nil {
		t.Fatalf("Batch write failed: %v", err)
	}
    t.Logf("*************批量写入后，打印数据应该是新的测试数据***************");
    testRead(db,t,0)
    t.Logf("++++++++++++++++++++批量写test %d:end+++++++++++++++++++++++++++++++++++++++++++++++++++",seq)
    seq++
    t.Logf(" ")
    t.Logf("+++++++++++++++++++批量读test,首先写入数据  key:测试数据 value:与key同样的数据++++++++++++++++++++") 
	for _, v := range test_values {
		writeBatch.Put([]byte(v), []byte(v))
	}
	err = writeBatch.Write()
	if err != nil {
		t.Fatalf("Batch write failed: %v", err)
	}
    t.Logf("*************批量写入后，遍历数据库，打印数据应该是新的测试数据***************")
    testRead(db,t,0)
    t.Logf("*************批量写入后，获取数据库前5条记录***************")
    testRead(db,t,5)
    t.Logf("*************批量写入后，获取数据库后5条记录***************")
    testRead(db,t,-5)
    t.Logf("*************批量写入后，获取数据库前7条记录,因为已经超过数据所有数据数量，应该返回数据库所有记录***************")
    testRead(db,t,7)
    t.Logf("*************批量写入后，获取数据库后7条记录,因为已经超过数据所有数据数量，应该返回数据库所有记录***************")
    testRead(db,t,-7)
    t.Logf("++++++++++++++++++++批量读test:end+++++++++++++++++++++++++++++++++++++++++++++++++++++++")
    t.Logf(" ")
    t.Logf("++++++++++++++++++++数据删除test ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
    t.Logf("*************数据库原有数据*************");
    testRead(db,t,0)
	for _, v := range test_values {
		err = writeBatch.Delete([]byte(v))
		if err != nil {
			t.Fatalf("delete %q failed: %v", v, err)
		}
	}
    t.Logf("*************调用Delete后，调用Write前,数据库中数据应该不变**************");
    testRead(db,t,0)
    t.Logf("*************调用Write后,数据库中应该无数据**************");
	err = writeBatch.Write()
	if err != nil {
		t.Fatalf("Batch write failed: %v", err)
	}
    testRead(db,t,0)
    t.Logf("++++++++++++++++++++数据删除test:end ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

}


/*
func TestLDB_ParallelPutGet(t *testing.T) {
	db, err := leveldb.OpenFile("tmp/paratest.db",nil)
	defer testLDB_Close(db, t)
    if err!=nil {
        t.Fatalf("Failed to create databas:%v",err)
    }
	testParallelPutGet(db, t)
}

func testParallelPutGet(db *leveldb.DB, t *testing.T) {
	const n = 8
	var pending sync.WaitGroup

	pending.Add(n)
	for i := 0; i < n; i++ {
		go func(key []string) {
	        writeBatch := leveldb_batch.NewWriteBatch(db)
			defer pending.Done()
			for _, v := range test_values {
				writeBatch.Put([]byte(v), []byte("?"))
			}
			err := writeBatch.Write()
			if err != nil {
				t.Fatalf("Batch write failed: %v", err)
			}
		}(test_values)
	}
	pending.Wait()

	pending.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer pending.Done()
	        readBatch := leveldb_batch.NewReadBatch(db)
			key, value := readBatch.Read(0)
			for index, content := range key {
				t.Logf("Key:Index[%d]:Key[%s]", index, content)
			}
			for index, content := range value {
				t.Logf("Value:Index[%d]:Key[%s]", index, content)
			}
		}()
	}
	pending.Wait()

	pending.Add(n)
	for i := 0; i < n; i++ {
		go func(key []string) {
			defer pending.Done()
	        writeBatch := leveldb_batch.NewWriteBatch(db)
			for _, v := range test_values {
				err := writeBatch.Delete([]byte(v))
				if err != nil {
					t.Fatalf("delete %q failed: %v", v, err)
				}
			}
		}(test_values)
	}
	pending.Wait()

	pending.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer pending.Done()
	        readBatch := leveldb_batch.NewReadBatch(db)
			key, value := readBatch.Read(0)
			for index, content := range key {
				t.Logf("Key:Index[%d]:Key[%s]", index, content)
			}
			for index, content := range value {
				t.Logf("Value:Index[%d]:Key[%s]", index, content)
			}
		}()
	}
	pending.Wait()
}
*/
