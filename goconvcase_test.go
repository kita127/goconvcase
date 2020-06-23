package goconvcase

import (
	"reflect"
	"testing"
)

func TestConvertCaseUStoUC(t *testing.T) {
	testTbl := []struct {
		comment string
		src     string
		expect  string
	}{
		{"test convert 1",
			`package hoge

var HOGE_VAR int
var FUGA_VAR int
var camelVar int

const CONST_VAR = 999

func SILENCE_KID() {
	HIYOKO := "ひよこ"
}
`,
			`package hoge

var HogeVar int
var FugaVar int
var camelVar int

const ConstVar = 999

func SilenceKid() {
	HIYOKO := "ひよこ"
}
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

func TestUSnakeDecode(t *testing.T) {
	testTbl := []struct {
		comment string
		input   string
		expect  *InterCode
	}{
		{"test 1",
			`SNAKE_CASE_VAR`,
			&InterCode{[]string{"snake", "case", "var"}},
		},
	}

	for _, tt := range testTbl {
		c := &USnake{}
		got := c.Decode(tt.input)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("got=%v, expect=%v", got, tt.expect)
		}
	}

}

func TestUCamelEncode(t *testing.T) {
	testTbl := []struct {
		comment string
		input   *InterCode
		expect  string
	}{
		{"test 1",
			&InterCode{[]string{"snake", "case", "var"}},
			`SnakeCaseVar`,
		},
	}

	for _, tt := range testTbl {
		c := &UCamel{}
		got := c.Encode(tt.input)
		if !reflect.DeepEqual(got, tt.expect) {
			t.Errorf("got=%v, expect=%v", got, tt.expect)
		}
	}

}

func TestUSnakeIsThisCase(t *testing.T) {
	testTbl := []struct {
		comment string
		input   string
		expect  bool
	}{
		{"test 1",
			`SNAKE_CASE_VAR`,
			true,
		},
		{"test 2",
			`snake_case_var`,
			false,
		},
		{"test 3",
			`CamelCaseVar`,
			false,
		},
		{"test 4",
			`camelCaseVar`,
			false,
		},
		{"test 5",
			`UPPER`,
			false,
		},
		{"test 6",
			`lower`,
			false,
		},
	}

	for _, tt := range testTbl {
		c := &USnake{}
		got := c.IsThisCase(tt.input)
		if got != tt.expect {
			t.Errorf("got=%v, expect=%v", got, tt.expect)
		}
	}

}
