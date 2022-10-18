package parser

import (
	"fmt"
	"path/filepath"
	"strings"

	"cuelang.org/go/cue/ast"
	cueparser "cuelang.org/go/cue/parser"
	"cuelang.org/go/cue/token"
	cue_mod "github.com/acorn-io/acorn/cue.mod"
	"github.com/acorn-io/acorn/pkg/appdefinition"
	"github.com/acorn-io/acorn/pkg/cue"
	"github.com/acorn-io/acorn/schema"
	"github.com/acorn-io/aml"
	"github.com/sirupsen/logrus"
)

const (
	Schema  = "github.com/acorn-io/acorn/schema/v1"
	AppType = "#App"
)

var std aml.StdDef

type Definitions map[string]Range

func LoadFileAndSchema(path string) (string, error) {
	data, err := cue.ReadCUE(path)
	if err != nil {
		return "", err
	}
	files := []cue.File{
		{
			// Need the .cue otherwise it looks for cue package
			Name: filepath.Base(path) + ".cue",
			Data: append(data, appdefinition.Defaults...),
			Parser: func(name string, src any) (*ast.File, error) {
				return parseFile(name, src)
			},
		},
	}
	ctx := cue.NewContext().
		WithNestedFS("schema", schema.Files).
		WithNestedFS("cue.mod", cue_mod.Files)
	ctx = ctx.WithFiles(files...)
	ctx = ctx.WithSchema(Schema, AppType)

	acornApp, err := ctx.Value()
	if err != nil {
		return "", err
	}

	v := acornApp.Value()
	logrus.Errorf("v is: %#v", v)

	return "", err
}

func parseFile(name string, src any) (f *ast.File, err error) {
	return aml.ParseFile(name, src, &std)
}

func ParseAcornDefinitions() (Definitions, error) {
	defs := Definitions{}

	f, err := getSchemaFile()
	if err != nil {
		return defs, err
	}

	parseDefinitions(&defs, f)
	return defs, nil
}

func parseDefinitions(defs *Definitions, f *ast.File) {
	ast.Walk(f, func(node ast.Node) bool {
		switch v := node.(type) {
		// case: #Def
		case *ast.Ident:
			if IsDefinition(v.Name) {
				defs.AppendRange(v.Name, v.Pos(), v.End())
			}
		}

		return true
	}, nil)
}

func getSchemaFile() (*ast.File, error) {
	content, err := schema.Files.ReadFile("v1/app.cue")
	if err != nil {
		return nil, err
	}

	// Currently this is valid CUE.
	return cueparser.ParseFile("app.cue", content)
}

func (def Definitions) String() string {
	str := fmt.Sprintf("%s\n", "")
	for _, r := range def {
		//str += fmt.Sprintf("Name %s:", name)
		str += fmt.Sprintf("\t%s\n", r)
	}
	return str
}

// AppendRange add a new definition to the line.
// Due to AST structure, it will always be sorted in ascending order by
// start.Column()
func (def Definitions) AppendRange(name string, start token.Pos, end token.Pos) {
	def[name] = Range{start, end, name}
}

// Find will search for a definition in the Definition object following line
// and column
// It will return the definition's name if found, or an error if not found
// Find function has a complexity of O(log(n)) thanks Definitions data
// structure that his a map.
func (def Definitions) Find(name string) (string, error) {
	if r, ok := def[name]; ok {
		return r.Name(), nil
	}
	return "", fmt.Errorf("definition not found")
}

// IsDefinition returns true if the current name is a CUE definition
// Pattern detected are:
// - #Foo
// - _#Foo
// It returns false if it's not a definition
func IsDefinition(name string) bool {
	return strings.HasPrefix(name, "#") || strings.HasPrefix(name, "_#")
}

func GetNode(def string) (ast.Node, error) {
	f, err := getSchemaFile()
	if err != nil {
		return nil, err
	}
	for _, node := range f.Decls {
		switch n := node.(type) {
		case *ast.Field:
			if fmt.Sprintf("%s", n.Label) == def {
				return node, nil
			}
		}
	}
	return nil, fmt.Errorf("node %s not found", def)
}
