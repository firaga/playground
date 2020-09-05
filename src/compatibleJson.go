package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *Person) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	var name string
	err = json.Unmarshal(*objMap["name"], &name)
	if err != nil {
		return err
	}

	var ageInt int
	err = json.Unmarshal(*objMap["age"], &ageInt)
	if err != nil {
		fmt.Println(err)
		// age  is string
		var ageString string
		err = json.Unmarshal(*objMap["age"], &ageString)
		if err != nil {
			err.Error()
			return err
		}
		aI, err := strconv.Atoi(ageString)
		if err != nil {
			return err
		}
		p.Age = aI

	} else {
		p.Age = ageInt
	}

	p.Name = name

	fmt.Printf("%+v", *p)

	return nil
}

func main() {
	p := `{"name": "John", "age": "10"}`
	//	p := `{"name": "John", "age": 10}`


	newP := Person{}
	err := newP.UnmarshalJSON([]byte(p))
	fmt.Println(newP)
	if err != nil {
		fmt.Printf("Error %+v", err)
	}

}
