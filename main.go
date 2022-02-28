package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int
	y := 7
	fmt.Println(x)
	fmt.Println(y)
	myVal,err :=strconv.ParseInt("NaN",0,64)
	if err != nil {
		fmt.Printf("%v\n",err)
	}else{
		fmt.Println(myVal)
	}

	m := make(map[string]int)

	m["Key"] = 6

	fmt.Println(m["Key"])

	s := []int{1,2,3}

	for i,v := range s {
		fmt.Printf("index %v , values %v\n",i,v)
	}

	s = append(s, 16)

	for i,v := range s {
		fmt.Printf("index %v , values %v\n",i,v)
	}

	//c:= make(chan int)
	//go doSomething(c)

	//<-c

	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

func doSomething(c chan int){
	time.Sleep(3*time.Second)
	fmt.Println("Done")
	c <- 1
}