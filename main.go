package main

import (
	"fmt"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/alecthomas/repr"
)

type Expression struct {
	Exp string `@Ident`
}

type ReplaceExpression struct {
	Expression  Expression   `@@ ("AS")?`
	ColumnNames []Expression `@@+`
}

type AliasExpression struct {
	Expression Expression `@@ `
	Alias      string     `( ( "AS" )? @Ident )?`
}

type InternalSelectExpression struct {
	Expression        Expression         `( (@@".")? "*"`
	ExceptColumnNames []string           `("EXCEPT" "(" @Ident+ ")")?`
	Replace           *ReplaceExpression `("REPLACE" "(" @@ ")")?`
	Alias             *AliasExpression   `| @@ )`
}

type Group struct {
	Name string `@Ident`
}

type RollUp struct {
	Name string `@Ident`
}

type GroupByExpression struct {
	Groups           []*Group  `@@ ("," @@)*`
	RollUpExpression []*RollUp `| "ROLLUP" "(" @@ ("," @@)* ")"`
}

type WindowClause struct {
	NamedWindow         string `@Ident "AS"`
	Alias               string `( @Ident `
	WindowSpecification string `| "(" ( @Ident )? ")")`
}

type CTE struct {
	Name            string          `@Ident "AS" `
	QueryExpression QueryExpression `"(" @@ ")"`
}

type WithStatement struct {
	CTE []*CTE `("WITH" @@ ("," @@)*)?`
}

type SelectQuerySetOperation struct {
	SelectStatement *SelectStatement `@@`
	QueryExpression *QueryExpression `| "(" @@ ")"`
	SetOperation    *SetOperation    `| @@`
}

type OrderByItem struct {
	Name       string `@Ident`
	Ascending  bool   `(@"ASC"`
	Descending bool   `|@"DESC")?`
}

type OrderByStatement struct {
	OrderBy []OrderByItem `("ORDER BY" @@ ("," @@)*)?`
}

type LimitStatement struct {
	Count    int `("LIMIT" @Ident`
	SkipRows int `("OFFSET" @Ident)?)?`
}

type QueryExpression struct {
	WithStatement           *WithStatement           `@@`
	SelectQuerySetOperation *SelectQuerySetOperation `@@`
	// OrderByStatement *OrderByStatement `@@`
	// LimitStatement   *LimitStatement   `@@`
}

// QueryStatement is a root
type QueryStatement struct {
	QueryExpression *QueryExpression `@@`
}

var sqlLexer = lexer.Must(lexer.NewSimple([]lexer.Rule{
	{`Keyword`, `(?i)\b(SELECT|FROM|TOP|DISTINCT|ALL|WHERE|GROUP|BY|HAVING|UNION|MINUS|EXCEPT|INTERSECT|ORDER|LIMIT|OFFSET|TRUE|FALSE|NULL|IS|NOT|ANY|SOME|BETWEEN|AND|OR|LIKE|AS|IN)\b`, nil},
	{`Ident`, `[a-zA-Z_][a-zA-Z0-9_]*`, nil},
	{`Number`, `[-+]?\d*\.?\d+([eE][-+]?\d+)?`, nil},
	{`String`, `'[^']*'|"[^"]*"`, nil},
	{`Operators`, `<>|!=|<=|>=|[-+*/%,.()=<>]`, nil},
	{"whitespace", `\s+`, nil},
}))

var parser = participle.MustBuild(
	&QueryStatement{},
	participle.Lexer(sqlLexer),
	participle.Unquote("String"))

// participle.CaseInsensitive("Keyword"),
func main() {

	sql := &QueryStatement{}
	sqlString := `
	SELECT Adams as LastName, 50 as SchoolID FROM Roster
	`
	err := parser.ParseString("", sqlString, sql)
	repr.Println(sql, repr.Indent("  "), repr.OmitEmpty(true))
	fmt.Println(err)

}
