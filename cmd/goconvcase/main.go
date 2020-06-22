package main

import (
	"fmt"

	conv "github.com/kita127/goconvcase"
)

func main() {

	src := `package hoge
var SNAKE_CASE_VAR int`

	res, err := conv.ConvertCase(src, conv.UpperSnake, conv.UpperCamel)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
