package main

import (
	"fmt"
	"encoding/json"
)

func main(){

	type Person struct{
		Name string `json:"name"`
		Age int	`json:"age"`
		Add []string `json:"omitempty"`
	}
	type Student struct{
		School string
		p *Person
	}

	joly := Student{
		School:"bupt",
	}
	var per Person
	per.Name = "joly"
	per.Age = 2
	joly.p = &per

	jsonstr := `{
		"Name":"tom",
		"age":4
	}`
	var ap Person
	json.Unmarshal([]byte(jsonstr), &ap)


	paramJson, err := json.Marshal(joly)
	fmt.Println("parmJson", string(paramJson))
	jper,_ := json.Marshal(per)

	if err != nil{
		fmt.Println("err"+err.Error())
	}
	var tom Student
	var per2 Person
	json.Unmarshal(paramJson, &tom)
	json.Unmarshal(jper, & per2)

	mapIns := make(map[string]interface{})
	mapIns["name"] = "adam"
	mapIns["age"] = 12

	jmap, _ := json.Marshal(mapIns)
	fmt.Println(string(jmap))

	aper1 := Person{"adam1", 12, []string{"bj","sh",}}

	ja1, _ := json.Marshal(aper1)
	var va1 map[string]interface{}
	json.Unmarshal(ja1, &va1)
	fmt.Println("ja1  ",va1)
	for k, v := range va1{
		fmt.Println("k>>>>", k, "v>>>>>>>", v)
	}





}
