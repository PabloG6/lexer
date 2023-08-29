package ast

import "lexer/lexer"

type BinaryExpr struct {
	Expr
	left     *Expr
	operator lexer.TOKEN
	right    *Expr
}

func NewBinaryExpr(left *Expr, operator lexer.TOKEN, right *Expr) *BinaryExpr {
	return &BinaryExpr{
		left:     left,
		right:    right,
		operator: operator,
	}

}
