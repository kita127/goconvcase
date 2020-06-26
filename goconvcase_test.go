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
	xxx_ := 20
	_yyy := 30
}
`,
			`package hoge

var hogeVar int
var fugaVar int
var camelVar int

const constVar = 999

func silenceKid() {
	hiyoko := "ひよこ"
	xxx_ := 20
	_yyy := 30
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
	XXX_ := 20
	_YYY := 30
}
`,
			`package hoge

var HogeVar int
var FugaVar int
var camelVar int

const ConstVar = 999

func SilenceKid() {
	HIYOKO := "ひよこ"
	XXX_ := 20
	_YYY := 30
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

func TestConvertCaseUCtoUS(t *testing.T) {
	testTbl := []struct {
		comment string
		src     string
		expect  string
	}{
		{"test convert 1",
			`package hoge

var UPPER_SNAKE_VAR int
var lower_snake_var int
var UpperCamelVar int
var lowerCamelVar int

const UPPER_SNAKE_CONST int = 0
const lower_snake_const int = 0
const UpperCamelConst int = 0
const lowerCamelConst int = 0

func UPPER_SNAKE_FUNC() {
	LOCAL_VAR := 0
}

func lower_snake_func() {
	local_var := 0
}

func UpperCamelFunc() {
	LocalVar := 0
}

func lowerCamelFunc() {
	localVar := 0
}
`,
			`package hoge

var UPPER_SNAKE_VAR int
var lower_snake_var int
var UPPER_CAMEL_VAR int
var lowerCamelVar int

const UPPER_SNAKE_CONST int = 0
const lower_snake_const int = 0
const UPPER_CAMEL_CONST int = 0
const lowerCamelConst int = 0

func UPPER_SNAKE_FUNC() {
	LOCAL_VAR := 0
}

func lower_snake_func() {
	local_var := 0
}

func UPPER_CAMEL_FUNC() {
	LOCAL_VAR := 0
}

func lowerCamelFunc() {
	localVar := 0
}
`,
		},
	}

	for _, tt := range testTbl {
		got, err := ConvertCase(tt.src, UpperCamel, UpperSnake)
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

func TestUSnakeEncode(t *testing.T) {
	testTbl := []struct {
		comment string
		input   *InterCode
		expect  string
	}{
		{"test 1",
			&InterCode{[]string{"snake", "case", "var"}},
			`SNAKE_CASE_VAR`,
		},
	}

	for _, tt := range testTbl {
		c := &USnake{}
		got := c.Encode(tt.input)
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

func TestUCamelDecode(t *testing.T) {
	testTbl := []struct {
		comment string
		input   string
		expect  *InterCode
	}{
		{"test 1",
			`CamelCaseVar`,
			&InterCode{[]string{"camel", "case", "var"}},
		},
	}

	for _, tt := range testTbl {
		c := &UCamel{}
		got := c.Decode(tt.input)
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

func TestIsThisCase(t *testing.T) {
	inputTbl := []string{
		"SNAKE_CASE_VAR",
		"snake_case_var",
		"CamelCaseVar",
		"camelCaseVar",
		"UPPER",
		"lower",
		"HOGE_VAR_",
		"_HOGE_VAR",
		"hoge_var_",
		"_hoge_var",
		"C",
		"c",
		"_"}
	testTbl := []struct {
		comment string
		cs      Case
		expects []bool
	}{
		{
			"Is this UpperCamel?",
			&UCamel{},
			[]bool{
				false, //   "SNAKE_CASE_VAR"
				false, //   "snake_case_var"
				true,  //   "CamelCaseVar"
				false, //   "camelCaseVar"
				false, //   "UPPER"
				false, //   "lower"
				false, //   "HOGE_VAR_"
				false, //   "_HOGE_VAR"
				false, //   "hoge_var_"
				false, //   "_hoge_var"
				false, //   "C"
				false, //   "c"
				false, //   "_"
			},
		},
	}

	for _, tt := range testTbl {
		c := tt.cs
		for j, inTb := range inputTbl {
			got := c.IsThisCase(inTb)
			if got != tt.expects[j] {
				t.Logf("%s", tt.comment)
				t.Errorf("input=%s, got=%v, expect=%v", inTb, got, tt.expects[j])
			}
		}
	}

}
