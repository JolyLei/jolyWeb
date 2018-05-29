package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println("主函数")
	sql := []string{"qqqq","wwww","eeee"}
	ret := strings.Join(sql,`','`)
	fmt.Println(ret)

}

