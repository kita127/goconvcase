package goconvcase

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

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

func ConvertCase(src string) (string, error) {
	return src, nil
}
