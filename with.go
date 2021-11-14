package main

type CTE struct {
	Name            *Value          `@@ "AS" `
	QueryExpression QueryExpression `"(" @@ ")"`
}

type WithStatement struct {
	CTE []*CTE `("WITH" @@ ("," @@)*)?`
}
