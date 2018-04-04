package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	FirstName string
	LastName string
	wontBeVisible string
	WontBeVisible string `json:"-"`
	Fw12f1w2fw12f1wf2 string `json:"job"`
}

func main() {
	value, _ := json.Marshal(Person{"First", "Last", "nop", "nop2", "Dev"})
	fmt.Println(string(value))
}
