package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Datas struct {
	XMLName xml.Name
	ItemDatass   []Item    `xml:"item"`
	WeaponDatas  []Weapon  `xml:"weapon"`
	ProtectDatas []Protect `xml:"protect"`
}

type Item struct {
	// XMLName xml.Name
	Name    string `xml:"name,attr"`
	Kind    string `xml:"kind,attr"`
	Comment string `xml:",chardata"`
	Param   int32  `xml:"param"`
}

type Weapon struct {
	// XMLName xml.Name
	Name    string `xml:"name,attr"`
	Kind    string `xml:"kind,attr"`
	Comment string `xml:",chardata"`
	Attack  int32  `xml:"attack"`
	Defense int32  `xml:"defense"`
}

type Protect struct {
	// XMLName xml.Name
	Name    string `xml:"name,attr"`
	Kind    string `xml:"kind,attr"`
	Comment string `xml:",chardata"`
	Attack  int32  `xml:"attack"`
	Defense int32  `xml:"defense"`
}

func main() {

	// XMLファイルの読み込み
	readFile, err := os.Open("xml/data2.xml")
	if err != nil {
		fmt.Println("error: %v", err)
		os.Exit(1)
	}

	// ファイルから文字列を抽出
	data, err := ioutil.ReadAll(readFile)
	if err != nil {
		fmt.Println("error: %v", err)
		os.Exit(1)
	}

	// 文字列をパース
	v := Datas{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Println("error: %v", err)
		os.Exit(1)
	}
	
	// 表示
	fmt.Println(v)
}
