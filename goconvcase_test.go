package goconvcase

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/token"
	"reflect"
	"testing"
)

func TestConvertCase(t *testing.T) {
	testTbl := []struct {
		comment string
		src     string
		expect  string
	}{
		{"test convert 1",
			`package hoge

var HOGE_VAR int

const FUGA_CONST int = 0
`,
			`package hoge

var HogeVar int

const FugaConst int = 0
`,
		},
	}

	for _, tt := range testTbl {
		got, err := ConvertCase(tt.src, UpperSnake, UpperCamel)
		if err != nil {
			t.Error(err)
		}
		if got != tt.expect {
			t.Errorf("got=%v, expect=%v", got, tt.expect)
		}
	}
}

func TestNewConverter(t *testing.T) {
	testTbl := []struct {
		comment string
		from    CaseType
		to      CaseType
		expect  *Converter
	}{
		{"test NewConverter 1",
			UpperSnake,
			UpperCamel,
			&Converter{&USnake{}, &UCamel{}},
		},
		{"test NewConverter 2",
			UpperCamel,
			UpperSnake,
			&Converter{&UCamel{}, &USnake{}},
		},
	}

	for _, tt := range testTbl {
		got := NewConverter(tt.from, tt.to)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("got=%v, expect=%v", got, tt.expect)
		}
	}
}

func TestConvertIdentifireUStoUC(t *testing.T) {
	testTbl := []struct {
		comment string
		src     string
		expect  string
	}{
		{"test 1",
			`package hoge

var HOGE_VAR int
`,
			`package hoge

var HogeVar int
`,
		},
	}

	for _, tt := range testTbl {
		c := NewConverter(UpperSnake, UpperCamel)

		fset := token.NewFileSet() // positions are relative to fset

		// Parse src but stop after processing the imports.
		node, err := parser.ParseFile(fset, "", tt.src, parser.ParseComments)
		if err != nil {
			t.Error(err)
		}

		got := c.convertIdentifire(node)

		var buf bytes.Buffer
		err = format.Node(&buf, fset, got)
		if err != nil {
			t.Error(err)
		}

		if buf.String() != tt.expect {
			t.Errorf("got=%s, expect=%s", buf.String(), tt.expect)
		}
	}
}
