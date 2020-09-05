package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type paramsStruct struct {
	Id    int    `json:"id"`
	Title string `json:"name"`
	Etc   etcStruct `json:"etc"`
}

type etcStruct struct {
	Content string `json:"content"`
}

func main() {

	params := paramsStruct{
		Id:    3,
		Title: "ha",
		Etc: etcStruct{Content: "content"},

	}
	paramsMap := map[string]interface{}{
		"Id":    1,
		"title": "ha",
		"etc": map[string]interface{}{
			"content": "content",
		},
	}
	a, _ := json.Marshal(params)
	fmt.Println(string(a))
	b, _ := json.Marshal(paramsMap)
	fmt.Println(string(b))

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(params)
	fmt.Print(buf.String())
	var buf2 bytes.Buffer
	json.NewEncoder(&buf2).Encode(paramsMap)
	fmt.Print(buf2.String())
}

