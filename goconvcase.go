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
	"time"
)

// これは関数
func bar() {

    // 関数の中身

	fmt.Println(time.Now())
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
