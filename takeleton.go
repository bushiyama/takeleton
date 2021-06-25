package takeleton

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "takeleton is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "takeleton",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
		(*ast.StructType)(nil),
		(*ast.TypeSpec)(nil),
	}

	var oyaName string
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.TypeSpec:
			oyaName = n.Name.Name
		case *ast.StructType:
			for _, v := range n.Fields.List {
				for _, vv := range v.Names {
					if strings.HasPrefix(vv.Name, oyaName) {
						pass.Reportf(n.Pos(), "identifier is oyaName")
					}
				}
			}
		}
	})

	return nil, nil
}
