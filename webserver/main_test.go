package main

import (
	"testing"
	"fmt"
)

func testPrint(t *testing.T){
	sum := print1to10()
	fmt.Print("Sucess!!")
	if sum != 55{
		t.Errorf("WORRY!")
	}
}

func testPrint2(t *testing.T){
	sum := print1to10()
	sum ++
	if sum !=56{
		t.Errorf("WORRT!!!!!")
	}
}


func TestAll(t *testing.T){
	t.Run("t1",testPrint)
	t.Run("t2",testPrint2)
}
func TestMain(m *testing.M){
	fmt.Print("TESTING BEGIN!!")
	m.Run()
}

func BenchmarkAll(b *testing.B){
	for n :=0;n < b.N;n++{
		print1to10()
	}
}