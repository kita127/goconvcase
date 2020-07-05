# goconvcase ![Go](https://github.com/kita127/goconvcase/workflows/Go/badge.svg?branch=kita127-patch-1)
goconvcase transforms the case of an identifier.

![image](https://github.com/kita127/goconvcase/blob/images/image.gif)

## Installation

    go get github.com/kita127/goconvcase/cmd/goconvcase


## Usage

    usage: goconvcase.exe [<flags>] [<path>]

    Flags:
          --help       Show context-sensitive help (also try --help-long and
                       --help-man).
      -w, --write      write result to (source) file instead of stdout
      -f, --from=FROM  from case
      -t, --to=TO      to case
      -l, --list       show valid cases

    Args:
      [<path>]  go file path


e.g.)
Convert upper snake case to upper camel case.

    $cat examples\hoge.go
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

    $goconvcase.exe --from us --to uc examples\hoge.go
    package hoge

    import "fmt"

    var (
            UpperSnakeVar   int
            lower_snake_var int
            UpperCamelVar   int
            lowerCamelVar   int
    )

    const (
            UpperSnakeConst   int = 0
            lower_snake_const int = 0
            UpperCamelConst   int = 0
            lowerCamelConst   int = 0
    )

    func UpperSnakeFunc() {
            LocalVar := 0
            fmt.Println(LocalVar)
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

The valid cases are displayed with `--list` option.

    $goconvcase.exe --list
    us : UPPER_SNAKE_CASE like this.
    uc : UpperCamelCase like this.
    ls : lower_snake_case like this.
    lc : lowerCamelCase like this.
