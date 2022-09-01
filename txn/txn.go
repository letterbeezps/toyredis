package txn

import (
	"context"
	"errors"

	"github.com/letterbeezps/toyredis/basedb"
)

type (
	TxFunc func(tx *Tx) error
	Tx     struct {
		db      *basedb.BaseDb
		isWrite bool
		context *txContex
	}
)

type txContex struct {
	records []*basedb.Record
	context context.Context
}

func (tx *Tx) addRecord(rec *basedb.Record) {
	tx.context.records = append(tx.context.records, rec)
}

func (tx *Tx) lock() {
	if tx.isWrite {
		tx.db.Mu.Lock()
	} else {
		tx.db.Mu.RLock()
	}
}

func (tx *Tx) unlock() {
	if tx.isWrite {
		tx.db.Mu.Unlock()
	} else {
		tx.db.Mu.RUnlock()
	}
}

func (tx *Tx) rollback() {
	tx.context.records = nil
}

func (tx *Tx) execRecord(recs []*basedb.Record) (err error) {
	for _, r := range recs {
		switch r.GetType() {
		case basedb.StringOper:
			{
				err = tx.db.ExecStringRecord(r)
			}
		}
	}
	return
}

func (tx *Tx) Commit() (err error) {
	if !tx.isWrite {
		err = errors.New("tx is not writeable")
	}
	err = tx.execRecord(tx.context.records)
	tx.unlock()
	return
}

func (tx *Tx) RollBack() error {
	if tx.isWrite {
		tx.rollback()
	}
	tx.unlock()
	tx.db = nil
	return nil
}

// 处理事务函数
func ManageTxn(isWrite bool, fn TxFunc, db *basedb.BaseDb) (err error) {
	var tx *Tx
	tx, err = NewTxn(isWrite, db)
	if err != nil {
		return
	}

	defer func() {
		// 如果执行出错，直接回滚
		if err != nil {
			tx.RollBack()
			return
		}
		// 写 事务，提交db
		if tx.isWrite {
			err = tx.Commit()
		} else {
			err = tx.RollBack()
		}
	}()

	// 先执行用户传入的函数
	err = fn(tx)
	return
}

func NewTxn(isWrite bool, db *basedb.BaseDb) (*Tx, error) {
	tx := &Tx{
		db:      db,
		isWrite: isWrite,
	}
	tx.lock()
	if isWrite {
		tx.context = &txContex{}
		tx.context.records = make([]*basedb.Record, 0, 1)
	}

	return tx, nil
}
