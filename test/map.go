package main

import "fmt"

func doSomething(arg...int){
	fmt.Println(arg)
	b := arg[:3]
	fmt.Println(b)
}

func Add(x,y int) int{
	return x+y
}

func ABmain(){
	myMap := make(map[string]string)
	fmt.Println(myMap)
	aMap := map[int]int{
		3:2,
	}
	aMap[4] = 3
	fmt.Println(aMap)

	sum := 0
	for i:=0;i<15;i++{
		sum += i
	}

	for{
		sum --
		if sum < 0{
			break
		}
	}
	fmt.Println(sum)

	m1 := map[int]int{}
	fmt.Println(m1)

	doSomething(1,2,3,4,5)
}
