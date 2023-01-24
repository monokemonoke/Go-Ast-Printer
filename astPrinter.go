package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"strings"
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
	var filenames []string
	if len(os.Args) > 1 {
		filenames = []string{os.Args[1]}
	} else {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %s\n", err.Error())
			return
		}
		filenames = strings.Split(string(b), "\n")
		filenames = filenames[:len(filenames)-1]
	}

	for _, filename := range filenames {
		err := Parser(filename, os.Stdout)
		if err != nil {
			fmt.Println(err)
		}
	}
}
