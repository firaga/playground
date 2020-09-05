package main

import (
	"encoding/json"
	"fmt"
)

type SelfMarshal struct {
	Name string
	Age  int
	City string
}

func (self SelfMarshal) MarshalJSON() ([]byte, error) {
	result := fmt.Sprintf("name:--%s,age:--%d,city:--%s", self.Name, self.Age, self.City)
	if !json.Valid([]byte(result)) {
		fmt.Println("invalid")
		return json.Marshal(result)
	}
	return []byte(result), nil
}

func main() {

	var self = SelfMarshal{}
	self.Age = 20
	self.Name = "XieWei"
	self.City = "HangZhou"

	selfJsonMarshal, err := json.Marshal(self)
	fmt.Println(err, string(selfJsonMarshal))
	selfJsonMarshal, err = self.MarshalJSON()
	fmt.Println(err, string(selfJsonMarshal))
}
