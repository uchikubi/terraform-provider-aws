package asthelper

import (
	"go/ast"
	"strings"
)

// AstBasicLitStringValue fetches a known BasicLit string value from the Expr
// If the Expr is not BasicLit, returns an empty string.
func AstBasicLitStringValue(e ast.Expr) string {
	switch v := e.(type) {
	case *ast.BasicLit:
		return strings.Trim(v.Value, `"`)
	}

	return ""
}
