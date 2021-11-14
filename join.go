package main

type CrossJoinOperator struct {
	Operator string `("CROSS JOIN" | ",")`
}

type ConditionJoinOperator struct {
	Operator string `(("INNER")? "JOIN" | "FULL" ("OUTER")? "JOIN" | "LEFT" ("OUTER")? "JOIN" | "RIGHT" ("OUTER")? "JOIN")`
}

type JoinCondition struct {
	OnClause string `("ON" @Ident)`
	// UsingClause []Expression `| "USING" "(" @@ ("," @@)* ")")`
}

type ConditionJoinOperation struct {
	FromItemPre           *FromItem              `@@`
	ConditionJoinOperator *ConditionJoinOperator `@@`
	FromItemPost          *FromItem              `@@`
	JoinCondition         *JoinCondition         `@@`
}

type CrossJoinOperation struct {
	FromItemPre       *FromItem          `@@`
	CrossJoinOperator *CrossJoinOperator `@@`
	FromItemPost      *FromItem          `@@`
}

// type BoolExpression struct {
// 	Expression *Expression `@@`
// }

type JoinName struct {
	Name *Name `@@`
}

type JoinType struct {
	Type string `@"INNER" | @"CROSS" | @"FULL OUTER"? | @"LEFT OUTER"? | @"RIGHT OUTER"?`
}

type OnClause struct {
	// BoolExpression *BoolExpression `"ON" @@`
	Clause *Value `@@`
}

type UsingClause struct {
	JoinNames []*JoinName `"USING" "(" @@ ("," @@)* ")"`
}

type JoinOperation struct {
	// CrossJoinOperation     *CrossJoinOperation     `(@@`
	// ConditionJoinOperation *ConditionJoinOperation `| @@)`
	FromItemPre  *FromItem    `@@`
	JoinType     *JoinType    `@@?`
	FromItemPost *FromItem    `"JOIN" @@`
	OnClause     *OnClause    `( @@`
	UsingClause  *UsingClause `| @@ )`
	// UsingClause []Expression `| "USING" "(" @@ ("," @@)* ")")`
}
