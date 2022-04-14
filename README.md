![Go](https://github.com/mchirico/go.etcd/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/mchirico/go.etcd/branch/main/graph/badge.svg?token=1UpZxvESjW)](https://codecov.io/gh/mchirico/go.etcd)

# go.etcd


This project is replaced by https://github.com/cwxstat/go.etcd

## Usage

```bash
export GO111MODULE=on
go mod init


```

## Sample Put and Get

```go
package main

import (
	"github.com/mchirico/go.etcd/pkg/etcdutils"
	"log"
	"time"
)

func main() {

	e, cancel := etcdutils.NewETC("test")
	defer cancel()

	e.DeleteWithPrefix("/testing")

	now := time.Now()

	e.Put("/testing/TestETC_Put ... more", now.String())
	e.PutWithLease("/testing/a", now.String(), 3)

	result, _ := e.GetWithPrefix("/testing/")

	if len(result.Kvs) != 2 {
		log.Printf("Number of keys: %d\n", len(result.Kvs))
	}

	for i, v := range result.Kvs {
		log.Printf("result.Kvs[%d]: %s, ver: %d,  lease: %d\n", i, v.Value, v.Version, v.Lease)
	}
}



```






```
Next version:

git tag -fa v0.0.15 -m "Describe what you did"
git push origin v0.0.15 --force


```
