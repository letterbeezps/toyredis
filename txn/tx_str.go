package txn

import (
	"errors"

	"github.com/letterbeezps/toyredis/basedb"
)

func (tx *Tx) get(key string) (value string, err error) {
	if ok, v := tx.db.StringStorage.Get(key); ok {
		value = v
		err = nil
		return
	}
	err = errors.New("not exist")
	return
}

func (tx *Tx) Get(key string) (value string, err error) {
	value, err = tx.get(key)

	return
}

func (tx *Tx) Set(key, value string) error {
	e := basedb.NewRecord([]byte(key), []byte(value), basedb.StringOper, basedb.StringSet)
	tx.addRecord(e)

	return nil
}
