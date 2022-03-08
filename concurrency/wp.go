package main

import "fmt"

func main() {
	tasks := []int{2,3,4,5,7,10,12,40}
	nWorkers := 3
	jobs := make(chan int,len(tasks))
	results := make(chan int,len(tasks))

//Cada worker ira sacando elementos a trabajar del buffer de jobs
	for i := 0; i < nWorkers; i++ {
		go Worker(i,jobs,results)
	}
//A medida que los jobs se van guardando en el buffer , seran removidos por los Workers del buffer
	for _,value := range tasks {
		jobs <- value
	}
	close(jobs)
	
//Este for es para el mostrado de cada numero de fibonacci calculado
	for r:= 0; r<len(tasks);r++ {
		<-results
	}
	close(results)
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started fib with %d\n",id,job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d finished job with value %d\n",id,fib)
		results <- fib
	}
}