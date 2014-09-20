package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type CalcsData struct {
	XMLName xml.Name
	Funcs   []FuncData `xml:"func"`
}

type FuncData struct {
	XMLName    xml.Name
	Name       string     `xml:"name,attr"`
	ReturnType string     `xml:"return_type,attr"`
	Comment    string     `xml:"comment"`
	Arg        []Argument `xml:"arg"`
	Calc       string     `xml:"calc"`
}

type Argument struct {
	ValueName string `xml:",chardata"`
	TypeName  string `xml:"type,attr"`
}

func main() {
	readFile, err := os.Open("xml/data.xml")
	if err != nil {
		fmt.Println("error: %v", err)
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(readFile)
	if err != nil {
		fmt.Println("error: %v", err)
		os.Exit(1)
	}

	v := CalcsData{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Println("error: %v", err)
		os.Exit(1)
	}

	fmt.Println(v)
}
