package main

type TypeCast struct {
	TypeName string `("AS" (@Ident|"STRUCT"|"VALUE"))?`
}

type Column struct {
	Name *Name `@@`
}

type SelectReplace struct {
	// OldColumnName *Expression `@@?`
	// NewColumName  *Expression `AsAlias?`
	OldColumnName *Value   `@@?`
	NewColumName  *AsAlias `@@?`
}

type ExceptStatement struct {
	ExceptColumns []*Column `"EXCEPT" "(" @@ ("," @@)* ")"`
}

type ReplaceStatement struct {
	Replaces []*SelectReplace `"REPLACE" "(" @@ ("," @@)* ")"`
}

type FirstFormSelectExpression struct {
	// Select        string           `(@Ident ".")? @"*"`
	// ExceptColumns []*Column        `("EXCEPT" "(" @@ ("," @@)* ")")?`
	// Replace       []*SelectReplace `("REPLACE" "(" @@ ("," @@)* ")")?`
	Expression *Value            `@@? "."? "*"`
	Except     *ExceptStatement  `@@?`
	Replace    *ReplaceStatement `@@?`
}

type SecondFormSelectExpression struct {
	Expression *Value   `@@`
	Alias      *AsAlias `@@?`
}

type RepeatableSelectExpression struct {
	FirstFormSelectExpression  *FirstFormSelectExpression  `@@`
	SecondFormSelectExpression *SecondFormSelectExpression `| @@`
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
	From                       []*FromStatement              `@@?`
	// Where            string             `("WHERE" @Ident)?`
	// GroupBy          *GroupByExpression `("GROUP BY" @@ )?`
	// Having           string             `("HAVING" @Ident)?`
	// Qualify          string             `("QUALIFY" @Ident)?`
	// Window           string             `("WINDOW" @Ident)?`
}
