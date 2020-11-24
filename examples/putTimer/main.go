package main

import (
	"github.com/mchirico/go.etcd/pkg/etcdutils"
	"log"
	"time"
)

func Update() {

	e, cancel := etcdutils.NewETC()
	defer cancel()

	now := time.Now()

	e.PutWithLease("gopi.service/update", now.String(), 120)
	e.PutWithLease("gopi.service/addr", Ifconfig(), 120)

	result, _ := e.GetWithPrefix("gopi.service")

	for i, v := range result.Kvs {
		log.Printf("result.Kvs[%d]: %s, ver: %d,  lease: %d\n", i, v.Value, v.Version, v.Lease)
	}

}

func main() {

	for {
		Update()
		time.Sleep(70 * time.Second)
	}

}
