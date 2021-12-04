package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Book struct {
	Name   string `xml:"Name"`
	Author string `xml:"Author"`
}

func main() {
	p := &Person{"Jack", 22}

	v, _ := xml.MarshalIndent(p, "", " ")

	fmt.Println(string(v))

	s := `
        <Book>
         <Name>War And Peace</Name>
         <Author>Leo Tolstoy</Author>
        </Book>
    `

	b := &Book{}

	xml.Unmarshal([]byte(s), b)

	fmt.Println(b)
}
