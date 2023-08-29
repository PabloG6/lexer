package ast

type GroupingExpr struct {
	Expr
	expression Expr
}

func NewGroupingExpr(expression Expr) *GroupingExpr {
	return &GroupingExpr{}

}
