


[![Build Status](https://travis-ci.org/mchirico/go.etcd.svg?branch=master)](https://travis-ci.org/mchirico/go.etcd)
[![codecov](https://codecov.io/gh/mchirico/go.etcd/branch/master/graph/badge.svg)](https://codecov.io/gh/mchirico/go.etcd)

[![Build Status](https://mchirico.visualstudio.com/go.etcd/_apis/build/status/mchirico.go.etcd?branchName=master)](https://mchirico.visualstudio.com/go.etcd/_build/latest?definitionId=9&branchName=master)


# go.etcd

## Usage

```bash
go get github.com/mchirico/go.etcd
```


```go
package main

import (
	"github.com/mchirico/go-etcd/pkg/etcdutils"
	"time"
)

func main() {

	e, cancel := etcdutils.NewETC()
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


```




```
Next version:

git tag -fa v0.0.1 -m "first version"
git push origin v0.0.1 --force

```