package ast

type LiteralExpr struct {
	Expr
	value *interface{}
}

func NewLiteralExpr(value *interface{}) *LiteralExpr {
	return &LiteralExpr{
		value: value,
	}

}
