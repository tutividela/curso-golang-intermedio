package main

import (
	"fmt"
	"time"
)

func sum(values ... int) int {
	sum :=0
	for _,v := range values {
		sum+=v
	}
	return sum
}

func mult(x int) (d,t,q int) {
	d =x*2
	t=x*3
	q=x*4
	return
}
func main() {
	x := 5
	y := func() int {
		return x * 2
	}()
	fmt.Println(y)

	c := make(chan int)

	go func (){
		fmt.Println("Starting Function")
		time.Sleep(5 *time.Second)

		fmt.Println("End")
		c <- 1
	}()
	<-c
}