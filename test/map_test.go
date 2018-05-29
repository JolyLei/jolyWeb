package main

import (
	"testing"
	"fmt"
)

func Test_Add(t *testing.T){
	ret := Add(2,4)
	if ret != 6{
		t.Error("cuole")
	}
}
func Test_Add2(t *testing.T){
	ret := Add(2,4)
	if ret != 1{
		t.Error("cuole")
	}
}

func Test_test(t *testing.T){
	ret := Add(1,2)
	fmt.Println(ret)
	if ret != 3{
		t.Error("23")
	}
}