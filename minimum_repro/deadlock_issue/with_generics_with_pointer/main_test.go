package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestT1(t *testing.T) {
	_ = T1[any]{}
}
