# go.etcd

## Usage

```bash
export GO111MODULE=on
go get github.com/mchirico/go.etcd
```


```go
package main

import (
	"github.com/mchirico/go.etcd/pkg/etcdutils"
	"log"
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
		log.Printf("Number of keys: %d\n", len(result.Kvs))
	}

	for i, v := range result.Kvs {
		log.Printf("result.Kvs[%d]: %s, ver: %d,  lease: %d\n", i, v.Value, v.Version, v.Lease)
	}
}



```




```
Next version:

git tag -fa v0.0.5 -m "Put changes here."
git push origin v0.0.5 --force

```
