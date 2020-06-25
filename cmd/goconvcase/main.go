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
	from = kingpin.Flag("from", "from case").Short('f').Required().String()
	to   = kingpin.Flag("to", "to case").Short('t').Required().String()
)

func main() {
	kingpin.Parse()

	src, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Fatal(err)
	}

	f, err := validateCase(*from)
	if err != nil {
		log.Fatal(err)
	}

	t, err := validateCase(*to)
	if err != nil {
		log.Fatal(err)
	}

	res, err := conv.ConvertCase(string(src), f, t)
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

func validateCase(s string) (conv.CaseType, error) {
	switch s {
	case "US":
		return conv.UpperSnake, nil
	case "UC":
		return conv.UpperCamel, nil
	}

	return -1, fmt.Errorf("invalid case. please enter : goconvcase --list")
}
