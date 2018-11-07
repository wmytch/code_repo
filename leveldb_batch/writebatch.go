package leveldb_batch

import (
	"leveldb"
)

type Putter interface {
	Put(key []byte, value []byte) error
}

type Deleter interface {
	Delete(key []byte) error
}

func NewWriteBatch(db *leveldb.DB) *WriteBatch{
    return &WriteBatch{db:db,btch:new(leveldb.Batch)}
}
type WriteBatch struct {
	db   *leveldb.DB
	btch *leveldb.Batch
	Putter
	Deleter
}

func (btch *WriteBatch) Put(key []byte, value []byte) error {
	btch.btch.Put(key, value)
	return nil
}

func (btch *WriteBatch) Delete(key []byte) error {
	btch.btch.Delete(key)
	return nil
}

func (btch *WriteBatch) Reset() {
	btch.btch.Reset()
}

func (btch *WriteBatch) Write() error {
	return btch.db.Write(btch.btch, nil)
}
