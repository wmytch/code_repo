package leveldb_batch

import (
	"leveldb"
)

type ReadBatch struct {
	db *leveldb.DB
}

func NewReadBatch(db *leveldb.DB) *ReadBatch {
	return &ReadBatch{db: db}
}

/********************
参数说明 
* 0 遍历全部数据
* n 正整数，返回前n个记录
* -n 负整数，返回后n个记录
返回值
返回两个[][]byte类型的slice，分别为key和value
********************/
func (btch *ReadBatch) Read(limit int) ([][]byte, [][]byte) {
	iter := btch.db.NewIterator(nil, nil)
	defer iter.Release()
	count := 1
	key := make([][]byte, 0)
	value := make([][]byte, 0)
	if limit >= 0 {
		for iter.Next() {
			tmpKey := make([]byte, len(iter.Key()))
			tmpValue := make([]byte, len(iter.Value()))
			copy(tmpKey, iter.Key())
			copy(tmpValue, iter.Value())
			key = append(key, tmpKey)
			value = append(value, tmpValue)
            count++
			if limit > 0 && count > limit {
				break
			}
		}
	} else {
		for iter.Last(); ; {
			tmpKey := make([]byte, len(iter.Key()))
			tmpValue := make([]byte, len(iter.Value()))
			copy(tmpKey, iter.Key())
			copy(tmpValue, iter.Value())
			key = append(key, tmpKey)
			value = append(value, tmpValue)
			count++
			if !iter.Prev() || count > -limit {
				break
			}
		}
	}
	return key, value
}

