package basedb

import (
	"fmt"
	"testing"
)

func TestRecord(t *testing.T) {
	rec := NewRecord([]byte("key"), []byte("value"), StringOper, StringSet)

	fmt.Println(rec.state)
	fmt.Println(rec.timestamp)
	fmt.Println(rec.meta)
}
