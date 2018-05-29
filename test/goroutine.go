package main

import "fmt"

func sum(arr []int, resChan chan int){
	sum := 0
	for _,val := range arr{
		sum += val
	}
	resChan <- sum
}

func amain() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8}
	resChan := make(chan int, 10)
	go sum(array[:len(array)-2],resChan)
	go sum(array[:len(array)/2],resChan)
	fmt.Println(<-resChan,<-resChan)
}