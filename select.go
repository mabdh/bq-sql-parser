package main

type TypeCast struct {
	TypeName string `("AS" (@Ident|"STRUCT"|"VALUE"))?`
}

type Column struct {
	Name string `@Ident`
}

type SelectReplace struct {
	OldColumnName string `@Ident ("AS")?`
	NewColumName  string `@Ident`
}

type GenericSelectExpression struct {
	Select        string           `(@Ident".")? @"*"`
	ExceptColumns []*Column        `("EXCEPT" "(" @@ ("," @@)* ")")?`
	Replace       []*SelectReplace `("REPLACE" "(" @@ ("," @@)* ")")?`
}

type AliasedSelectExpression struct {
	Expression *Value `@@`
	Alias      *Value `( "AS"? @@ )?`
}

type RepeatableSelectExpression struct {
	GenericSelectExpression *GenericSelectExpression `@@`
	AliasedSelectExpression *AliasedSelectExpression `| @@`
}

// type SelectExpression struct {
// 	// Tokens []lexer.Token
// 	// TypeCast                   *TypeCast                     `@@?`
// 	// Selection                  string                        `( "ALL" | "DISTINCT" )?`
// 	// RepeatableSelectExpression []*RepeatableSelectExpression `@@ ( "," @@ )*`
// }

type SelectStatement struct {
	// SelectExpression SelectExpression `"SELECT" @@`
	// Tokens []lexer.Token
	TypeCast                   *TypeCast                     `"SELECT" @@?`
	Selection                  string                        `( @"ALL" | @"DISTINCT" )?`
	RepeatableSelectExpression []*RepeatableSelectExpression `@@ ( "," @@ )*`
	// From             []*FromStatement `("FROM" @@ ("," @@ )* )?`
	// Where            string             `("WHERE" @Ident)?`
	// GroupBy          *GroupByExpression `("GROUP BY" @@ )?`
	// Having           string             `("HAVING" @Ident)?`
	// Qualify          string             `("QUALIFY" @Ident)?`
	// Window           string             `("WINDOW" @Ident)?`
}
