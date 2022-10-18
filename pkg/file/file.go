package file

import (
	"fmt"

	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/format"
	"github.com/acorn-io/aml"
	"github.com/acorn-io/amllsp/pkg/parser"
	"github.com/acorn-io/amllsp/pkg/stdlib"
	"github.com/tliron/kutil/logging"
)

type Acornfile struct {
	path       string
	astContent *ast.File
	log        logging.Logger
	index      acornIndex
}

var std aml.StdDef

func New(file string, logger logging.Logger) (*Acornfile, error) {
	stdlib, err := stdlib.GetStdData()
	if err != nil {
		return nil, err
	}

	functions := map[string]bool{}
	for _, e := range stdlib.Decls[1].(*ast.LetClause).Expr.(*ast.StructLit).Elts {
		functions[e.(*ast.Field).Label.(*ast.Ident).Name] = true
	}

	std.Imports = stdlib.Imports
	std.Unresolved = stdlib.Unresolved
	std.Decls = stdlib.Decls
	std.Functions = functions

	content, err := aml.ParseFile(file, nil, &std)
	if err != nil {
		return nil, err
	}

	idx := indexAcornfile(content)
	return &Acornfile{
		path:       file,
		astContent: content,
		log:        logging.NewScopeLogger(logger, "file"),
		index:      idx,
	}, nil
}

func (a *Acornfile) GetDocDefinitionMarkDown(line, character int) (string, error) {
	r := a.GetRange(line, character)
	topLevel, err := a.index.findTopLevelDef(line, r)
	if err != nil {
		return "", err
	}
	node, err := parser.GetNode(topLevel)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("#### Type\n## %s\n", customerFormatNode(node, 0)), nil
}

func (a *Acornfile) ContentChanges([]interface{}) error {
	return nil
}

func (a *Acornfile) GetRange(line, character int) Range {
	for _, item := range a.index[line] {
		if item.InRange(character) {
			return item
		}
	}
	return Range{}
}

func customerFormatNode(node ast.Node, depth int) string {
	var doc string

	formatNode := func(n ast.Node) string {
		display, err := format.Node(n, format.Simplify())
		if err == nil {
			return string(display)
		}
		return "unknown"
	}

	switch n := node.(type) {
	case *ast.Field:
		switch v := n.Value.(type) {
		case *ast.Ident:
			name := fmt.Sprintf("%s", n.Label)
			if depth == 0 {
				return fmt.Sprintf("%s: %s", name, v)
			}

			return formatNode(n)

		case *ast.UnaryExpr, *ast.BinaryExpr:
			return formatNode(n)
		case *ast.StructLit:
			return fmt.Sprintf("%s: %s", n.Label, formatNode(v))
		case *ast.BasicLit:
			return fmt.Sprintf("%s: %s", n.Label, formatNode(v))
		default:
			doc += fmt.Sprintf("%s: {\n%s}", n.Label, customerFormatNode(v, depth+1))
		}
	case *ast.StructLit:
		for _, d := range n.Elts {
			doc += customerFormatNode(d, depth+1)
		}
	}

	return doc
}
