package main

type CTE struct {
	Name            string          `@Ident "AS" `
	QueryExpression QueryExpression `"(" @@ ")"`
}

type WithStatement struct {
	CTE []*CTE `("WITH" @@ ("," @@)*)?`
}
