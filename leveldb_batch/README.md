# Leveldb 批处理  

### 批量写 writebatch.go 
- `NewWriteBatch(db *leveldb.DB)` 函数创建一个WriteBatch对象 
- `Put`函数实现`Putter`接口，将数据加入到WriteBatch中
- `Delete`函数实现`Deleter`接口，实际上只是删除WriteBatch中的数据，在Write之前不会影响数据库文件中的数据  
- `Reset`函数重置WriteBatch.btch,也就是数据库提供的Batch
- `Write()` 函数执行批量写任务，执行完成之后数据才真正写入到数据库文件中 
    
### 批量读 readbatch.go

- `NewReadBatch`函数创建一个ReadBatch对象
- `Read`函数执行批量读,并不保证按某种顺序返回。其参数`limit`
    - 为`0`时表示返回数据库全部数据
    - 为某个正整数`n`时表示返回前n个记录，如果n大于数据库中已有数据数量，则返回数据库中所有数据
    - 为某个负整数`-n`时表示返回后n个记录，如果n大于数据库中已有数据数量，则返回数据库中所有数据
    - 

### 单元测试程序 leveldb_test.go  
- 批量写测试
- 批量读测试
- 删除测试
