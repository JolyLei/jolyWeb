package main

import (
	"fmt"
)

type Bird struct{
	Name string
	LifeTime uint32
}

func (b *Bird) Fly(){
	fmt.Println("i can fly" )
}

func AAmain(){

}