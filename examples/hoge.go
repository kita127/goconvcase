package hoge

import "fmt"

var (
	UPPER_SNAKE_VAR int
	lower_snake_var int
	UpperCamelVar   int
	lowerCamelVar   int
)

const (
	UPPER_SNAKE_CONST int = 0
	lower_snake_const int = 0
	UpperCamelConst   int = 0
	lowerCamelConst   int = 0
)

func UPPER_SNAKE_FUNC() {
	LOCAL_VAR := 0
	fmt.Println(LOCAL_VAR)
}

func lower_snake_func() {
	local_var := 0
	fmt.Println(local_var)
}

func UpperCamelFunc() {
	LocalVar := 0
	fmt.Println(LocalVar)
}

func lowerCamelFunc() {
	localVar := 0
	fmt.Println(localVar)
}
