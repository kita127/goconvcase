package goconvcase

import (
	"reflect"
	"testing"
)

const inputSrc string = `package hoge

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
`

func TestConvertCaseLStoLC(t *testing.T) {
	testTbl := []struct {
		comment string
		src     string
		expect  string
	}{
		{"test convert 1",
			inputSrc,
			`package hoge

var UPPER_SNAKE_VAR int
var lowerSnakeVar int
var UpperCamelVar int
var lowerCamelVar int

const UPPER_SNAKE_CONST int = 0
const lowerSnakeConst int = 0
const UpperCamelConst int = 0
const lowerCamelConst int = 0

func UPPER_SNAKE_FUNC() {
	LOCAL_VAR := 0
}

func lowerSnakeFunc() {
	localVar := 0
}

func UpperCamelFunc() {
	LocalVar := 0
}

func lowerCamelFunc() {
	localVar := 0
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
			inputSrc,
			`package hoge

var UpperSnakeVar int
var lower_snake_var int
var UpperCamelVar int
var lowerCamelVar int

const UpperSnakeConst int = 0
const lower_snake_const int = 0
const UpperCamelConst int = 0
const lowerCamelConst int = 0

func UpperSnakeFunc() {
	LocalVar := 0
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

func TestConvertCaseLCtoLS(t *testing.T) {
	testTbl := []struct {
		comment string
		src     string
		expect  string
	}{
		{"test convert 1",
			inputSrc,
			`package hoge

var UPPER_SNAKE_VAR int
var lower_snake_var int
var UpperCamelVar int
var lower_camel_var int

const UPPER_SNAKE_CONST int = 0
const lower_snake_const int = 0
const UpperCamelConst int = 0
const lower_camel_const int = 0

func UPPER_SNAKE_FUNC() {
	LOCAL_VAR := 0
}

func lower_snake_func() {
	local_var := 0
}

func UpperCamelFunc() {
	LocalVar := 0
}

func lower_camel_func() {
	local_var := 0
}
`,
		},
	}

	for _, tt := range testTbl {
		got, err := ConvertCase(tt.src, LowerCamel, LowerSnake)
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
			inputSrc,
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

func TestLSnakeEncode(t *testing.T) {
	testTbl := []struct {
		comment string
		input   *InterCode
		expect  string
	}{
		{"test 1",
			&InterCode{[]string{"snake", "case", "var"}},
			`snake_case_var`,
		},
	}

	for _, tt := range testTbl {
		c := &LSnake{}
		got := c.Encode(tt.input)
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

func TestLCamelDecode(t *testing.T) {
	testTbl := []struct {
		comment string
		input   string
		expect  *InterCode
	}{
		{"test 1",
			`camelCaseVar`,
			&InterCode{[]string{"camel", "case", "var"}},
		},
	}

	for _, tt := range testTbl {
		c := &LCamel{}
		got := c.Decode(tt.input)
		if !reflect.DeepEqual(got, tt.expect) {
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
		"Mix_Case_Var",
		"Mix_case_var",
		"mix_Case_Var",
		"_"}
	testTbl := []struct {
		comment string
		cs      Case
		expects []bool
	}{
		{
			"Is this UpperSnake?",
			&USnake{},
			[]bool{
				true,  //   "SNAKE_CASE_VAR"
				false, //   "snake_case_var"
				false, //   "CamelCaseVar"
				false, //   "camelCaseVar"
				false, //   "UPPER"
				false, //   "lower"
				false, //   "HOGE_VAR_"
				false, //   "_HOGE_VAR"
				false, //   "hoge_var_"
				false, //   "_hoge_var"
				false, //   "C"
				false, //   "c"
				false, //   "Mix_Case_Var",
				false, //   "Mix_case_var",
				false, //   "mix_Case_Var",
				false, //   "_"
			},
		},
		{
			"Is this LowerSnake?",
			&LSnake{},
			[]bool{
				false, //   "SNAKE_CASE_VAR"
				true,  //   "snake_case_var"
				false, //   "CamelCaseVar"
				false, //   "camelCaseVar"
				false, //   "UPPER"
				false, //   "lower"
				false, //   "HOGE_VAR_"
				false, //   "_HOGE_VAR"
				false, //   "hoge_var_"
				false, //   "_hoge_var"
				false, //   "C"
				false, //   "c"
				false, //   "Mix_Case_Var",
				false, //   "Mix_case_var",
				false, //   "mix_Case_Var",
				false, //   "_"
			},
		},
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
				false, //   "Mix_Case_Var",
				false, //   "Mix_case_var",
				false, //   "mix_Case_Var",
				false, //   "_"
			},
		},
		{
			"Is this LowerCamel?",
			&LCamel{},
			[]bool{
				false, //   "SNAKE_CASE_VAR"
				false, //   "snake_case_var"
				false, //   "CamelCaseVar"
				true,  //   "camelCaseVar"
				false, //   "UPPER"
				false, //   "lower"
				false, //   "HOGE_VAR_"
				false, //   "_HOGE_VAR"
				false, //   "hoge_var_"
				false, //   "_hoge_var"
				false, //   "C"
				false, //   "c"
				false, //   "Mix_Case_Var",
				false, //   "Mix_case_var",
				false, //   "mix_Case_Var",
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
