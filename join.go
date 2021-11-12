package main

type CrossJoinOperator struct {
	Operator string `("CROSS JOIN" | ",")`
}

type ConditionJoinOperator struct {
	Operator string `(("INNER")? "JOIN" | "FULL" ("OUTER")? "JOIN" | "LEFT" ("OUTER")? "JOIN" | "RIGHT" ("OUTER")? "JOIN")`
}

type JoinCondition struct {
	OnClause    string       `("ON" @Ident`
	UsingClause []Expression `| "USING" "(" @@ ("," @@)* ")")`
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

type JoinOperation struct {
	CrossJoinOperation     *CrossJoinOperation     `(@@`
	ConditionJoinOperation *ConditionJoinOperation `| @@)`
}
