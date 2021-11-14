package main

type SetOperator struct {
	Operator *Value `@@`
}

// Operator *Value `"UNION" ( "ALL" | "DISTINCT" ) | "INTERSECT DISTINCT" | "EXCEPT DISTINCT"`
type SetOperation struct {
	QueryExpressionPre  *QueryExpression `@@`
	SetOperator         *SetOperator     `@@`
	QueryExpressionPost *QueryExpression `@@`
}
