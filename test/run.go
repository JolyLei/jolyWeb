package main

import(
	"fmt"
	"joly/services"
)

func add(x... int) {
	fmt.Println(x)
}

type A struct{
	a,b int
}

func (a *A) sprint(s1 string) (s11 string){
	s11 = s1+"11111111"
	return
}

func (a *A) ssprint(s2 string) (s22 string){
	temp := &A{}
	s22 = s2+temp.sprint("sssss")+"22222"
	return
}


func AAAmain(){
	add(1,2,3,4,5)
	v := &A{1,2}
	arr := []int{1,2,3,4,5}
	copy := make([]int,0)
	for _,i := range arr{
		var j int
		j = i
		copy = append(copy,j)
	}
	s := fmt.Sprintf(`ssssfe%s'\n'sssssssss`,"aaaa")

	fmt.Println(v,s)
	fmt.Println(copy)
	fmt.Println("ending...........")


	amap := make(map[int]int)
	arr = []int{1,2,3,4,4,3,2,5}
	for _,val := range arr{

		res,ok := amap[val]
		if ok{
			fmt.Println(res ,"is in the map")
			continue
		}else{
			amap[val] = val
		}
	}
	amap[2] = 3
	fmt.Println(amap)
	temp := &A{}
	fmt.Println(temp.ssprint("a"))

	services.Pt()
}