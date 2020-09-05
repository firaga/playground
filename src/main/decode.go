package main

import (
	"encoding/json"
	"fmt"
)

//type paramsStruct struct {
//	Id       uint64    `json:"id,string"`
//	Title    string    `json:"name"`
//	NotExist int       `json:"notExist"`
//	Etc      etcStruct `json:"etc"`
//}
//
//type etcStruct struct {
//	Content string `json:"content"`
//}

type res struct {
	Id  int    `json:"id,string,omitempty"`
	Num string `json:"num"`
}

type myStruct struct {
	Id  int `json:"id"`
	Num int `json:"num,string"`
	Str string `json:"str,integer"`
}

type myStructOut struct {
	Id  int `json:"id"`
	Num int `json:"num,string"`
	Str string `json:"str,int"`
}

func main() {
	a := myStruct{
		Id:  3,
		Num: 0,
		Str: "4",
	}
	b, err := json.Marshal(a)
	fmt.Println(string(b))
	fmt.Println(err)
	var newA myStructOut
	err = json.Unmarshal([]byte(b), &newA)
	fmt.Println(newA)
	fmt.Println(err)

	c, err := json.Marshal(newA)
	fmt.Println(string(c))
	fmt.Println(err)

}
