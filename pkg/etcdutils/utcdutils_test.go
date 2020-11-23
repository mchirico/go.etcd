package etcdutils

import (
	"fmt"
	"github.com/etcd-io/etcd/clientv3"
	"testing"
	"time"
)

func TestETC_Put(t *testing.T) {
	e, cancel := NewETC()
	defer cancel()

	e.DeleteWithPrefix("/testing")

	now := time.Now()
	e.Put("/testing/TestETC_Put", now.String())

	result, _ := e.Get("/testing/TestETC_Put")
	if string(result.Kvs[0].Value) != now.String() {
		t.Fatalf("%s\n", result.Kvs[0].Value)
	}

}

func TestETC_Delete(t *testing.T) {
	e, cancel := NewETC()
	defer cancel()

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
	e, cancel := NewETC()
	defer cancel()

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

/*
For this you neeed:
    "github.com/etcd-io/etcd/clientv3"
*/
func TestETC_Txn(t *testing.T) {
	e, cancel := NewETC()
	defer cancel()

	tx := e.Txn()

	txresp, err := tx.If(
		clientv3.Compare(clientv3.Value("foo"), "=", "bar"),
	).Then(
		clientv3.OpPut("foo", "sanfoo"), clientv3.OpPut("newfoo", "newbar"),
	).Else(
		clientv3.OpPut("foo", "bar"), clientv3.OpDelete("newfoo"),
	).Commit()
	fmt.Println(txresp, err)

	result, _ := e.Get("foo")
	for i, v := range result.Kvs {
		t.Logf("result.Kvs[%d]: %s, ver: %d,  lease: %d\n", i, v.Value, v.Version, v.Lease)
	}

}
