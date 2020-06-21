package goconvcase

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

// case type values
const (
	UpperSnake = iota
	UpperCamel
)

// CaseType type
type CaseType int

// Sample function
func Sample() {
	src := `package foo

import (
	"fmt"
)

const (
    constVar = 100
)

// これは関数
func bar() {

    localVar := 99

    // 関数の中身

	fmt.Println(localVar + constVar)
}`

	fset := token.NewFileSet() // positions are relative to fset

	// Parse src but stop after processing the imports.
	node, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	ast.Print(fset, node)

	var buf bytes.Buffer
	err = format.Node(&buf, fset, node)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(buf.String())

}

// ConvertCase function
func ConvertCase(src string, from, to CaseType) (string, error) {

	res := `package hoge
var HogeVar int`

	return res, nil
}
