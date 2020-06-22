package goconvcase

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
)

// case type values
const (
	UpperSnake = iota
	UpperCamel
)

// CaseType type
type CaseType int

// Converter struct
type Converter struct {
	from Case
	to   Case
}

// InterCode struct
type InterCode struct {
	ss []string
}

// Case interface
type Case interface {
	Decode(name string) *InterCode
	Encode(ic *InterCode) string
}

// USnake struct
type USnake struct{}

// Decode *USnake.Decode method
func (c *USnake) Decode(name string) *InterCode {
	ic := &InterCode{}
	for _, s := range strings.Split(name, "_") {
		ic.ss = append(ic.ss, strings.ToLower(s))
	}
	return ic
}

// Encode *USnake.Encode method
func (c *USnake) Encode(ic *InterCode) string {
	return ""
}

// UCamel struct
type UCamel struct{}

// Decode *UCamel.Decode method
func (c *UCamel) Decode(name string) *InterCode {
	return nil
}

// Encode *UCamel.Encode method
func (c *UCamel) Encode(ic *InterCode) string {
	ss := []string{}
	for _, s := range ic.ss {
		ss = append(ss, strings.Title(s))
	}
	return strings.Join(ss, "")
}

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

// NewConverter function
func NewConverter(from, to CaseType) *Converter {
	c := &Converter{}

	switch from {
	case UpperSnake:
		c.from = &USnake{}
	case UpperCamel:
		c.from = &UCamel{}
	}

	switch to {
	case UpperCamel:
		c.to = &UCamel{}
	case UpperSnake:
		c.to = &USnake{}
	}

	return c
}

func (c *Converter) convertIdentifire(node ast.Node) ast.Node {
	switch node.(type) {
	case *ast.File:
		for _, n := range node.(*ast.File).Decls {
			c.convertIdentifire(n)
		}
	case *ast.GenDecl:
		for _, n := range node.(*ast.GenDecl).Specs {
			c.convertIdentifire(n)
		}
	case *ast.ValueSpec:
		for _, ident := range node.(*ast.ValueSpec).Names {
			ic := c.from.Decode(ident.Name)
			ident.Name = c.to.Encode(ic)
		}
	case *ast.FuncDecl:
		c.convertIdentifire(node.(*ast.FuncDecl).Name)
	case *ast.Ident:
		ident := node.(*ast.Ident)
		ic := c.from.Decode(ident.Name)
		ident.Name = c.to.Encode(ic)
	}
	return node
}

// Convert function
func (c *Converter) Convert(src string) (string, error) {
	fset := token.NewFileSet() // positions are relative to fset

	// Parse src but stop after processing the imports.
	node, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		return "", err
	}

	// TODO:
	// 最後けす
	ast.Print(fset, node)

	converted := c.convertIdentifire(node)

	var buf bytes.Buffer
	err = format.Node(&buf, fset, converted)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// ConvertCase function
func ConvertCase(src string, from, to CaseType) (string, error) {
	c := NewConverter(from, to)
	res, err := c.Convert(src)
	if err != nil {
		return "", err
	}

	return res, nil
}
