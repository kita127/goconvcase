package goconvcase

import (
	"reflect"
	"testing"
)

func TestConvertCaseLStoLC(t *testing.T) {
	testTbl := []struct {
		comment string
		src     string
		expect  string
	}{
		{"test convert 1",
			`package hoge

var hoge_var int
var fuga_var int
var camelVar int

const const_var = 999

func silence_kid() {
	hiyoko := "ひよこ"
}
`,
			`package hoge

var hogeVar int
var fugaVar int
var camelVar int

const constVar = 999

func silenceKid() {
	hiyoko := "ひよこ"
}
`,
		},
	}

	for _, tt := range testTbl {
		got, err := ConvertCase(tt.src, LowerSnake, LowerCamel)
		if err != nil {
			t.Error(err)
		}
		if got != tt.expect {
			t.Errorf("got=%v, expect=%v", got, tt.expect)
		}
	}
}

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

func TestLSnakeDecode(t *testing.T) {
	testTbl := []struct {
		comment string
		input   string
		expect  *InterCode
	}{
		{"test 1",
			`snake_case_var`,
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
		{"SNAKE_CASE_VAR",
			`SNAKE_CASE_VAR`,
			true,
		},
		{"snake_case_var",
			`snake_case_var`,
			false,
		},
		{"CamelCaseVar",
			`CamelCaseVar`,
			false,
		},
		{"camelCaseVar",
			`camelCaseVar`,
			false,
		},
		{"UPPER",
			`UPPER`,
			false,
		},
		{"lower",
			`lower`,
			false,
		},
		{"_HOGE_VAR",
			`_HOGE`,
			false,
		},
		{"HOGE_VAR_",
			`HOGE_`,
			false,
		},
		{"_",
			`_`,
			false,
		},
	}

	for _, tt := range testTbl {
		c := &USnake{}
		got := c.IsThisCase(tt.input)
		if got != tt.expect {
			t.Logf("%s", tt.comment)
			t.Errorf("got=%v, expect=%v", got, tt.expect)
		}
	}

}

func TestLSnakeIsThisCase(t *testing.T) {
	testTbl := []struct {
		comment string
		input   string
		expect  bool
	}{
		{"SNAKE_CASE_VAR",
			`SNAKE_CASE_VAR`,
			false,
		},
		{"snake_case_var",
			`snake_case_var`,
			true,
		},
		{"CamelCaseVar",
			`CamelCaseVar`,
			false,
		},
		{"camelCaseVar",
			`camelCaseVar`,
			false,
		},
		{"UPPER",
			`UPPER`,
			false,
		},
		{"lower",
			`lower`,
			false,
		},
		{"hoge_var_",
			`hoge_var_`,
			false,
		},
		{"_hoge_var",
			`_hoge_var`,
			false,
		},
		{"_",
			`_`,
			false,
		},
	}

	for _, tt := range testTbl {
		c := &LSnake{}
		got := c.IsThisCase(tt.input)
		if got != tt.expect {
			t.Logf("%s", tt.comment)
			t.Errorf("got=%v, expect=%v", got, tt.expect)
		}
	}

}
