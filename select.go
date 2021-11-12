package main

type TypeCast struct {
	TypeName string `("AS" (@Ident|"STRUCT"|"VALUE"))?`
}

type Column struct {
	Name string `@Ident`
}

type SelectExcept struct {
	Columns []*Column `("EXCEPT" "(" @@ ("," @@)* ")")?`
}

type SelectReplace struct {
	OldColumnName string `@Ident ("AS")?`
	NewColumName  string `@Ident`
}

type GenericSelectExpression struct {
	Select  string         `(@Ident".")? @"*"`
	Except  *SelectExcept  `@@`
	Replace *SelectReplace `("REPLACE" "(" @@ ("," @@)* ")")?`
}

type AliasedSelectExpression struct {
	Expression string `@Ident`
	Alias      string `(("AS")? @Ident)?`
}

type RepeatableSelectExpression struct {
	GenericSelectExpression *GenericSelectExpression `@@`
	AliasedSelectExpression *AliasedSelectExpression `| @@`
}

type SelectExpression struct {
	TypeCast                   *TypeCast                     `@@?`
	Selection                  string                        `("ALL"|"DISTINCT")?`
	RepeatableSelectExpression []*RepeatableSelectExpression `@@ ("," @@)*`
}

type SelectStatement struct {
	SelectExpression SelectExpression   `"SELECT" @@`
	From             []*FromStatement   `("FROM" @@ ("," @@)*)?`
	Where            string             `("WHERE" @Ident)?`
	GroupBy          *GroupByExpression `("GROUP BY" @@ )?`
	Having           string             `("HAVING" @Ident)?`
	Qualify          string             `("QUALIFY" @Ident)?`
	Window           string             `("WINDOW" @Ident)?`
}
