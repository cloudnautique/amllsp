package stdlib

import (
	"embed"
	"fmt"

	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/parser"
)

type Function struct {
	Name   string
	Params []string
}

var (
	//go:embed std.cue
	fs embed.FS
	//std aml.StdDef
)

func Functions() ([]Function, error) {
	stdData, err := GetStdData()
	if err != nil {
		return nil, err
	}

	functions := []Function{}
	for _, e := range stdData.Decls[1].(*ast.LetClause).Expr.(*ast.StructLit).Elts {
		functions = append(functions, Function{
			Name:   e.(*ast.Field).Label.(*ast.Ident).Name,
			Params: []string{},
		})
	}

	return functions, nil
}

func GetStdData() (*ast.File, error) {
	data, err := fs.ReadFile("std.cue")
	if err != nil {
		return nil, err
	}

	return parser.ParseFile("std.cue", data)
}

func (f *Function) Markdown() string {
	return fmt.Sprintf("# std.%s()", f.Name)
}
