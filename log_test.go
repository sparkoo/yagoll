package yagoll

import (
	"testing"
)

func TestSimpleLog(t *testing.T) {
	Println("uoch", *t)
	Printf("euch %+v", *t)
}