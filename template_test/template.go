package main

import (
	"bytes"
	"io/ioutil"
	"text/template"
)

// テンプレートに適応するデータ用の構造体
type FuncData struct {
	FuncName   string
	ReturnType string
	ArgType    string
	ArgName    string
}

func main() {
	// テンプレートファイルを読み込む
	tmp_data, _ := ioutil.ReadFile("./tmpl/test.tmpl")

	// テンプレートファイルのデータをパース
	t, _ := template.New("test_tmpl").Parse(string(tmp_data))

	// テンプレートに適用するデータ生成
	func_data := FuncData{FuncName: "TestFunc", ReturnType: "int", ArgType: "int", ArgName: "aaa"}

	// 受け取り先のbytes.Bufferを用意して、データを適用して出力
	var out bytes.Buffer
	t.Execute(&out, func_data)

	// ファイル出力
	ioutil.WriteFile("./output/output.cs", out.Bytes(), 0644)
}
