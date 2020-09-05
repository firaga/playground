package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type intOrEmpty int

func (ioe *intOrEmpty) UnmarshalJSON(data []byte) error {
	if string(data) == `""` {
		if ioe != nil {
			*ioe = 0
		}
		return nil
	}

	var i int
	err := json.Unmarshal(data, &i)
	if err != nil {
		var iStr string
		err := json.Unmarshal(data, &iStr)
		if err != nil {
			return err
		}
		i64, _ := strconv.ParseInt(iStr, 10, 32)
		i = int(i64)
	}
	p := (*int)(ioe)
	*p = i
	return nil
}

type S struct {
	A intOrEmpty
	B intOrEmpty `json:"b"`
}

func main() {
	//s := S{A: 1, B: 2}
	var s S

	err := json.Unmarshal([]byte(`{"A":1337,"B":"1"}`), &s)
	fmt.Printf("%#v\n", s)
	fmt.Println(err)

	//var i int = int(s.A)
	//fmt.Println(i)
}
