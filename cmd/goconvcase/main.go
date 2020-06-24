package main

import (
	"fmt"
	"io/ioutil"
	"log"

	conv "github.com/kita127/goconvcase"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	wf   = kingpin.Flag("write", "write result to (source) file instead of stdout").Short('w').Bool()
	path = kingpin.Arg("path", "go file path").Required().String()
)

func main() {
	kingpin.Parse()

	src, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Fatal(err)
	}

	res, err := conv.ConvertCase(string(src), conv.UpperSnake, conv.UpperCamel)
	if err != nil {
		log.Fatal(err)
	}

	if *wf {
		err = ioutil.WriteFile(*path, []byte(res), 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println(res)
	}
}
