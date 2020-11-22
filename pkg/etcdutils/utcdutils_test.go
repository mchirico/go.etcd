package etcdutils

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	e := NewETC()

	r := e.EtcdRun()
	fmt.Println(r)
}
