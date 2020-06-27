package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	conv "github.com/kita127/goconvcase"
	"gopkg.in/alecthomas/kingpin.v2"
)

// constant value
const (
	US string = "us"
	UC string = "uc"
	LS string = "ls"
	LC string = "lc"
)

var (
	wf   = kingpin.Flag("write", "write result to (source) file instead of stdout").Short('w').Bool()
	path = kingpin.Arg("path", "go file path").String()
	from = kingpin.Flag("from", "from case").Short('f').String()
	to   = kingpin.Flag("to", "to case").Short('t').String()
	list = kingpin.Flag("list", "show valid cases").Short('l').Bool()
)

func main() {
	kingpin.Parse()

	if *list {
		putCaseList()
		os.Exit(0)
	}

	if *path == "" {
		log.Fatal(fmt.Errorf("goconvcase.exe: error: required argument 'path' not provided, try --help"))
	}

	f, err := validateCase(*from)
	if err != nil {
		log.Fatal(err)
	}

	t, err := validateCase(*to)
	if err != nil {
		log.Fatal(err)
	}

	src, err := ioutil.ReadFile(*path)
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

	os.Exit(0)
}

func validateCase(s string) (conv.CaseType, error) {
	switch s {
	case US:
		return conv.UpperSnake, nil
	case UC:
		return conv.UpperCamel, nil
	case LS:
		return conv.LowerSnake, nil
	case LC:
		return conv.LowerCamel, nil
	}

	return -1, fmt.Errorf("goconvcase.exe: error: invalid case , try --help or --list")
}

func putCaseList() {
	li := []struct {
		cs   string
		desc string
	}{
		{US, "UPPER_SNAKE_CASE like this."},
		{UC, "UpperCamelCase like this."},
		{LS, "lower_snake_case like this."},
		{LC, "lowerCamelCase like this."},
	}

	for _, v := range li {
		fmt.Printf("%s : %s\n", v.cs, v.desc)
	}
}
