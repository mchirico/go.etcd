package etcdutils

import (
	"fmt"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	e := NewETC()

	r := e.EtcdRun()
	fmt.Println(r)
}

func TestETC_Put(t *testing.T) {
	e := NewETC()
	now := time.Now()
	e.Put("/testing/TestETC_Put", now.String())

	result, _ := e.Get("/testing/TestETC_Put")
	if string(result.Kvs[0].Value) != now.String() {
		t.Fatalf("%s\n", result.Kvs[0].Value)
	}

}
