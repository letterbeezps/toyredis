<p align="center">
    <img src="./.github/WechatIMG3.jpeg">
</p>

# ToyRedis

使用go1.18的泛型特性来实现redis的底层数据结构，并它们为基础实现redis的API

## Example

```go
package main

import (
	"fmt"

	"github.com/letterbeezps/toyredis"
	"github.com/letterbeezps/toyredis/txn"
)

func main() {
	testRedis := toyredis.NewToyRedis()

	if err := testRedis.Update(func(tx *txn.Tx) error {
		err := tx.Set("letter", "name")
		if err != nil {
			return nil
		}
		return nil
	}); err != nil {
		fmt.Println(err)
	}

	if err := testRedis.View(func(tx *txn.Tx) error {
		v, err := tx.Get("letter")
		if err != nil {
			return nil
		}
		fmt.Println(v)
		return nil
	}); err != nil {
		fmt.Println(err)
	}
}

```
