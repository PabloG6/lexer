package ast

import "lexer/lexer"

type UnaryExpr struct {
	Expr
	operator lexer.TOKEN
	right    *Expr
}

func NewUnaryExpr(operator lexer.TOKEN, right *Expr) *UnaryExpr {
	return &UnaryExpr{
		operator: operator,
		right:    right,
	}

}
