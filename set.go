package main

type SetOperator struct {
	Operator string `"UNION" ( "ALL" | "DISTINCT" ) | "INTERSECT DISTINCT" | "EXCEPT DISTINCT"`
}

// Operator *Value `"UNION" ( "ALL" | "DISTINCT" ) | "INTERSECT DISTINCT" | "EXCEPT DISTINCT"`
type SetOperation struct {
	QueryExpressionPre  *QueryExpression `@@`
	SetOperator         *SetOperator     `@@`
	QueryExpressionPost *QueryExpression `@@`
}
