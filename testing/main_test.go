package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(5,5)

	if total != 10 {
		t.Errorf("Sum was incorrect , got %d expected %d",total,10)
	}
}

func TestMax(t *testing.T){
	max := GetMax(12,45)

	if max != 45 {
		t.Errorf("Max was incorrect , got %d expected %d",max,45)
	}
}

func TestFibonacci(t *testing.T){
	tables := []struct{
		a int
		b int
	}{
		{1,1},
		{80,21},
		{50,12586269025},
}
	for _ , v := range tables{
		fib := Fibonacci(v.a)
		if v.a != fib{
			t.Errorf("Fibonacci was incorrect , got %d expected %d",fib,v.a)
		}
	}
}