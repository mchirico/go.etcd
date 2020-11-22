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
	defer e.Cancel()

	e.DeleteWithPrefix("/testing")

	now := time.Now()
	e.Put("/testing/TestETC_Put", now.String())

	result, _ := e.Get("/testing/TestETC_Put")
	if string(result.Kvs[0].Value) != now.String() {
		t.Fatalf("%s\n", result.Kvs[0].Value)
	}

}

func TestETC_Delete(t *testing.T) {
	e := NewETC()
	defer e.Cancel()
	e.DeleteWithPrefix("/testing")

	now := time.Now()

	e.Put("/testing/TestETC_Put ... more", now.String())
	e.PutWithLease("/testing/a", now.String(), 3)

	e.DeleteWithPrefix("/testing")
	result, _ := e.GetWithPrefix("/testing/")

	if len(result.Kvs) != 0 {
		t.Fatalf("Number of keys should be 0. You got: %d\n", len(result.Kvs))
	}

}

func TestETC_GetWithPrefix(t *testing.T) {
	e := NewETC()
	defer e.Cancel()

	e.DeleteWithPrefix("/testing")

	now := time.Now()

	e.Put("/testing/TestETC_Put ... more", now.String())
	e.PutWithLease("/testing/a", now.String(), 3)

	result, _ := e.GetWithPrefix("/testing/")

	if len(result.Kvs) != 2 {
		t.Fatalf("Number of keys: %d\n", len(result.Kvs))
	}

	for i, v := range result.Kvs {
		t.Logf("result.Kvs[%d]: %s, ver: %d,  lease: %d\n", i, v.Value, v.Version, v.Lease)
	}

}
