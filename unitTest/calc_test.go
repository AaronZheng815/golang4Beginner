package main
import (
	"testing"
)
//必须以Test开头,函数的参数固定
func TestAdd(t *testing.T){
	r := add(2,4)
	if r != 6{
		t.Fatalf("add(2,4) error, expect:%d, actuall:%d",6,r)
	}
	t.Logf("test add successfull")
}