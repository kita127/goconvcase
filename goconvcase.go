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
	Decode(ident string) *InterCode
	Encode(ic *InterCode) string
}

// USnake struct
type USnake struct{}

// Decode *USnake.Decode method
func (c *USnake) Decode(ident string) *InterCode {
	return nil
}

// Encode *USnake.Encode method
func (c *USnake) Encode(ic *InterCode) string {
	return ""
}

// UCamel struct
type UCamel struct{}

// Decode *UCamel.Decode method
func (c *UCamel) Decode(ident string) *InterCode {
	return nil
}

// Encode *UCamel.Encode method
func (c *UCamel) Encode(ic *InterCode) string {
	return ""
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

//     0  *ast.GenDecl {
//     1  .  TokPos: 2:1
//     2  .  Tok: var
//     3  .  Lparen: -
//     4  .  Specs: []ast.Spec (len = 1) {
//     5  .  .  0: *ast.ValueSpec {
//     6  .  .  .  Names: []*ast.Ident (len = 1) {
//     7  .  .  .  .  0: *ast.Ident {
//     8  .  .  .  .  .  NamePos: 2:5
//     9  .  .  .  .  .  Name: "HOGE_VAR"
//    10  .  .  .  .  .  Obj: *ast.Object {
//    11  .  .  .  .  .  .  Kind: var
//    12  .  .  .  .  .  .  Name: "HOGE_VAR"
//    13  .  .  .  .  .  .  Decl: *(obj @ 5)
//    14  .  .  .  .  .  .  Data: 0
//    15  .  .  .  .  .  }
//    16  .  .  .  .  }
//    17  .  .  .  }
//    18  .  .  .  Type: *ast.Ident {
//    19  .  .  .  .  NamePos: 2:14
//    20  .  .  .  .  Name: "int"
//    21  .  .  .  }
//    22  .  .  }
//    23  .  }
//    24  .  Rparen: -
//    25  }
// convertIdentifire method
func (c *Converter) convertIdentifire(node *ast.File) *ast.File {
	for _, d := range node.Decls {
		switch d.(type) {
		case *ast.GenDecl:
			for _, s := range d.(*ast.GenDecl).Specs {
				switch s.(type) {
				case *ast.ValueSpec:
					for _, ident := range s.(*ast.ValueSpec).Names {
						ident.Name = "HogeVar"
					}
				}
			}
		}
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
