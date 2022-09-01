package toyredis

import (
	"errors"
	"testing"

	"github.com/letterbeezps/toyredis/txn"
	"github.com/stretchr/testify/assert"
)

// go test -run TestGetSet
func TestGetSet(t *testing.T) {
	testRedis := NewToyRedis()

	if err := testRedis.Update(func(tx *txn.Tx) error {
		err := tx.Set("zp", "name")
		assert.Equal(t, nil, err)
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	if err := testRedis.View(func(tx *txn.Tx) error {
		v, err := tx.Get("zp")
		assert.Equal(t, nil, err)
		assert.Equal(t, "name", v)
		return nil
	}); err != nil {
		t.Fatal(err)
	}
}

func exampleFuncWithError() error {
	return errors.New("some error")
}

// go test -run TestGetSetWithErr
func TestGetSetWithErr(t *testing.T) {
	testRedis := NewToyRedis()

	// 模拟发生错误的函数，事务不会被提交
	testRedis.Update(func(tx *txn.Tx) error {
		err := tx.Set("zp", "name")
		assert.Equal(t, nil, err)
		err = exampleFuncWithError()
		return err
	})

	// 此处尝试读取key, 读取不到才是正常的
	if err := testRedis.View(func(tx *txn.Tx) error {
		v, err := tx.Get("zp")
		assert.Equal(t, nil, err)
		assert.Equal(t, "name", v)
		return nil
	}); err != nil {
		t.Fatal(err)
	}
}
