package main

import (
	"fmt"
	"testing"
)

func TestUpdate(t *testing.T) {
	Update()
}

func TestStatus(t *testing.T) {
	result,err := Status()
	if err != nil {
		t.Fatalf("No status")
	}
	fmt.Println(result)
}