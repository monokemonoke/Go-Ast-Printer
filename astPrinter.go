package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
)

func Parser(filename string, w io.Writer) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.Mode(0))
	if err != nil {
		return err
	}

	for _, d := range f.Decls {
		ast.Fprint(w, fset, d, ast.NotNilFilter)
		fmt.Println()
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	err := Parser(os.Args[1], os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
}
