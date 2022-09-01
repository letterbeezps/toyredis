package toyredis

import (
	"github.com/letterbeezps/toyredis/basedb"
	"github.com/letterbeezps/toyredis/txn"
)

type (
	ToyRedis struct {
		db *basedb.BaseDb
	}
)

func NewToyRedis() *ToyRedis {
	return &ToyRedis{
		db: basedb.NewBaseDb(),
	}
}

func (toyredis *ToyRedis) View(fn txn.TxFunc) error {
	return txn.ManageTxn(false, fn, toyredis.db)
}

func (toyredis *ToyRedis) Update(fn txn.TxFunc) error {
	return txn.ManageTxn(true, fn, toyredis.db)
}
