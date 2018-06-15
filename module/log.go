package main

import (
	"log"
	"os"
)


func init(){
	log.SetPrefix("[jolyweb]")
	log.SetFlags(log.Ldate| log.Lshortfile)
}

func main2(){
	var info *log.Logger

	info = log.New(os.Stdout, "info", log.Ldate)

	info.Println("hlele")
}
