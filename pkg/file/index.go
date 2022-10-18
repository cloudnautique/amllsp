package file

import (
	"fmt"

	"cuelang.org/go/cue/ast"
)

type acornIndex map[int][]Range

var acornTopLevelKeys = map[string]string{
	"containers":  "#Container",
	"jobs":        "#Job",
	"acorns":      "#Acorn",
	"secrets":     "#Secret",
	"volumes":     "#Volume",
	"images":      "#Image",
	"routers":     "#Router",
	"labels":      "#Label",
	"annotations": "#Annotation",
}

func indexAcornfile(f *ast.File) acornIndex {
	index := map[int][]Range{}

	ast.Walk(f, func(node ast.Node) bool {
		switch v := node.(type) {
		case *ast.Ident:
			index[v.Pos().Line()] = append(index[v.Pos().Line()], Range{v.Name, v.Pos(), v.End()})
		}
		return true
	}, nil)

	return index
}

func (i *acornIndex) findTopLevelDef(line int, r Range) (string, error) {
	v := *i
	if key, ok := acornTopLevelKeys[v[line][0].name]; ok {
		return key, nil
	}
	return i.checkLineForTopLevelKey(line)
}

func (i *acornIndex) checkLineForTopLevelKey(line int) (string, error) {
	v := *i
	if line == 0 {
		if key, ok := acornTopLevelKeys[v[line][0].name]; ok {
			return key, nil
		}
		return "", fmt.Errorf("no key found")
	}
	for _, item := range v[line] {
		if key, ok := acornTopLevelKeys[item.name]; ok {
			return key, nil
		}
	}

	return i.checkLineForTopLevelKey(line - 1)
}
