package main

import (
	"fmt"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/alecthomas/repr"
)

// type ReplaceExpression struct {
// 	Expression  Expression   `@@ ("AS")?`
// 	ColumnNames []Expression `@@+`
// }

// type AliasExpression struct {
// 	Expression Expression `@@ `
// 	Alias      string     `( ( "AS" )? @Ident )?`
// }

// type InternalSelectExpression struct {
// 	Expression        Expression         `( (@@".")? "*"`
// 	ExceptColumnNames []string           `("EXCEPT" "(" @Ident+ ")")?`
// 	Replace           *ReplaceExpression `("REPLACE" "(" @@ ")")?`
// 	Alias             *AliasExpression   `| @@ )`
// }

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

type OrderByItem struct {
	// Expression *Expression `@@`
	Expression *Value `@@`
	OrderType  string `(@"ASC" |@"DESC")?`
}

type OrderClause struct {
	OrderBy []OrderByItem `"ORDER" "BY" @@ ("," @@)*`
}

type SkipRows struct {
	Number *Number `@@`
}

type LimitClause struct {
	Count    *Value `"LIMIT" @@`
	SkipRows *Value `("OFFSET" @@)?`
}

type FirstFormQueryExpression struct {
	SelectStatement     *SelectStatement `@@`
	SetOperation        *SetOperation    `@@?`
	QueryExpressionPost *QueryExpression `@@?`
	OrderByStatement    *OrderClause     `@@?`
	LimitStatement      *LimitClause     `@@?`
}

type SecondFormQueryExpression struct {
	QueryExpression     *QueryExpression `"(" @@ ")"`
	SetOperation        *SetOperation    `@@?`
	QueryExpressionPost *QueryExpression `@@?`
	OrderByStatement    *OrderClause     `@@?`
	LimitStatement      *LimitClause     `@@?`
}

type SetOperation struct {
	Operator string `( @"UNION" ( @"ALL" | @"DISTINCT" )? | @"INTERSECT DISTINCT" | @"EXCEPT DISTINCT" )`
}

// type QueryExpression struct {
// 	SelectStatement     *SelectStatement `( @@`
// 	QueryExpressionPre  *QueryExpression `| "(" @@ ")" )`
// 	SetOperation        *SetOperation    `@@?`
// 	QueryExpressionPost *QueryExpression `@@?`
// 	OrderByStatement    *OrderClause     `@@?`
// 	LimitStatement      *LimitClause     `@@?`
// }

type QueryExpression struct {
	SelectStatement     *SelectStatement `( @@`
	QueryExpressionPre  *QueryExpression `| "(" @@ ")" )`
	SetOperation        *SetOperation    `@@?`
	QueryExpressionPost *QueryExpression `@@?`
	OrderByStatement    *OrderClause     `@@?`
	LimitStatement      *LimitClause     `@@?`
}

// type QueryExpression struct {
// 	QueryExpressionPre  *QueryExpression `(@@`
// 	SetOperation        *SetOperation    ` @@`
// 	QueryExpressionPost *QueryExpression ` @@ `
// 	SelectStatement     *SelectStatement `| @@`
// 	QueryExpression     *QueryExpression `| "(" @@ ")" )`
// 	OrderByStatement    *OrderClause     `@@?`
// 	LimitStatement      *LimitClause     `@@?`
// }

// QueryStatement is a root
type QueryStatement struct {
	Tokens          []lexer.Token
	WithStatement   *WithStatement   `@@? ";"?`
	QueryExpression *QueryExpression `@@ ";"?`
}

type Value struct {
	Wildcard   bool     `(  @"*"`
	Number     *float64 ` | @Number`
	Identifier string   ` | @Ident`
	String     *string  ` | @String`
	Boolean    *Boolean ` | @("TRUE" | "FALSE")`
	Null       bool     ` | @"NULL" )`
}

type Boolean bool

func (b *Boolean) Capture(values []string) error {
	*b = values[0] == "TRUE"
	return nil
}

var sqlLexer = lexer.Must(lexer.NewSimple([]lexer.Rule{
	{`Keyword`, `(?i)\b(SELECT|FROM|TOP|DISTINCT|ALL|WHERE|GROUP|BY|HAVING|UNION|MINUS|EXCEPT|INTERSECT|ORDER|LIMIT|OFFSET|TRUE|FALSE|NULL|IS|NOT|ANY|SOME|BETWEEN|AND|OR|LIKE|AS|IN)\b`, nil},
	{`Ident`, `[a-zA-Z_][a-zA-Z0-9_]*`, nil},
	{`Number`, `[-+]?\d*\.?\d+([eE][-+]?\d+)?`, nil},
	{`String`, `'[^']*'|"[^"]*"`, nil},
	{`Operators`, `<>|!=|<=|>=|[-+*/%,.()=<>]`, nil},
	{"Punct", `,`, nil},
	{"whitespace", `\s+`, nil},
}))

var parser = participle.MustBuild(
	&QueryStatement{},
	participle.Lexer(sqlLexer),
	participle.Unquote("String"),
	participle.CaseInsensitive("Keyword"),
	participle.UseLookahead(2))

func main() {

	sql := &QueryStatement{}
	// sqlString := `
	// WITH PlayerStats AS
	// (SELECT 'Adams' as LastName, 51 as OpponentID, 3 as PointsScored UNION ALL
	// SELECT 'Buchanan', 77, 0 UNION ALL
	// SELECT 'Coolidge', 77, 1 UNION ALL
	// SELECT 'Adams', 52, 4 UNION ALL
	// SELECT 'Buchanan', 50, 13)
	// SELECT * FROM PlayerStats
	// `

	sqlString := `SELECT * FROM (SELECT "apple" AS fruit, "carrot" AS vegetable);`

	err := parser.ParseString("", sqlString, sql)
	repr.Println(sql, repr.Indent("  "), repr.OmitEmpty(true))
	fmt.Println(err)

}
