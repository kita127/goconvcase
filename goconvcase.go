package goconvcase

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
	"unicode"
)

// case type values
const (
	UpperSnake = iota
	UpperCamel
	LowerSnake
	LowerCamel
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
	IsThisCase(name string) bool
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
	// TODO
	panic(fmt.Errorf("USnake.Encode() 未実装"))
	return ""
}

// IsThisCase *USnake.IsThisCase method
func (c *USnake) IsThisCase(name string) bool {
	if name == "_" {
		return false
	}
	ss := strings.Split(name, "_")
	if ss[0] == "" || ss[len(ss)-1] == "" {
		return false
	}
	if len(ss) > 1 {
		s := strings.Join(ss, "")
		for _, c := range []rune(s) {
			if !unicode.IsUpper(c) {
				return false
			}
		}
		return true
	}
	return false
}

// UCamel struct
type UCamel struct{}

// Decode *UCamel.Decode method
func (c *UCamel) Decode(name string) *InterCode {
	// TODO
	panic(fmt.Errorf("UCamel.Decode() 未実装"))
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

// IsThisCase *UCamel.IsThisCase method
func (c *UCamel) IsThisCase(name string) bool {
	// TODO
	panic(fmt.Errorf("UCamel.IsThisCase 未実装"))
	return false
}

// LSnake struct
type LSnake struct{}

// Decode *LSnake.Decode method
func (c *LSnake) Decode(name string) *InterCode {
	ic := &InterCode{}
	for _, s := range strings.Split(name, "_") {
		ic.ss = append(ic.ss, strings.ToLower(s))
	}
	return ic
}

// Encode *LSnake.Encode method
func (c *LSnake) Encode(ic *InterCode) string {
	// TODO
	panic(fmt.Errorf("LSnake.Encode() 未実装"))
	return ""
}

// IsThisCase *LSnake.IsThisCase method
func (c *LSnake) IsThisCase(name string) bool {
	if name == "_" {
		return false
	}
	ss := strings.Split(name, "_")
	if ss[0] == "" || ss[len(ss)-1] == "" {
		return false
	}
	if len(ss) > 1 {
		s := strings.Join(ss, "")
		for _, c := range []rune(s) {
			if !unicode.IsLower(c) {
				return false
			}
		}
		return true
	}
	return false
}

// LCamel struct
type LCamel struct{}

// Decode *LCamel.Decode method
func (c *LCamel) Decode(name string) *InterCode {
	// TODO
	panic(fmt.Errorf("LCamel.Decode() 未実装"))
	return nil
}

// Encode *LCamel.Encode method
func (c *LCamel) Encode(ic *InterCode) string {
	ss := []string{}
	ss = append(ss, ic.ss[0])
	for _, s := range ic.ss[1:] {
		ss = append(ss, strings.Title(s))
	}
	return strings.Join(ss, "")
}

// IsThisCase *LSnake.IsThisCase method
func (c *LCamel) IsThisCase(name string) bool {
	// TODO
	panic(fmt.Errorf("LCamel.IsThisCase 未実装"))
	return false
}

// NewConverter function
func NewConverter(from, to CaseType) *Converter {
	return &Converter{from: newCase(from), to: newCase(to)}
}

// newCase function
func newCase(c CaseType) Case {
	switch c {
	case UpperSnake:
		return &USnake{}
	case UpperCamel:
		return &UCamel{}
	case LowerSnake:
		return &LSnake{}
	case LowerCamel:
		return &LCamel{}
	}
	return nil
}

func (c *Converter) convertIdentifire(node ast.Node) ast.Node {
	ast.Inspect(node, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.Ident:
			ident := n.(*ast.Ident)
			if c.from.IsThisCase(ident.Name) {
				ic := c.from.Decode(ident.Name)
				ident.Name = c.to.Encode(ic)
			}
		}
		return true
	})
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
