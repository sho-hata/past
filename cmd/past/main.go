package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("please specify file path")
		os.Exit(1)
	}

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, args[0], nil, 0)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	ast.Print(fset, f)
}
