package main

import (
	"fmt"
	"io/ioutil"
	"log"

	conv "github.com/kita127/goconvcase"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	file = kingpin.Arg("file", "go file").Required().String()
)

func main() {
	kingpin.Parse()

	src, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Fatal(err)
	}

	res, err := conv.ConvertCase(string(src), conv.UpperSnake, conv.UpperCamel)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
