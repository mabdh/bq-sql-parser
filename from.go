package main

type AliasName struct {
	Name string `@Ident`
}
type AsAlias struct {
	Alias *AliasName `"AS"? @@`
}

type TableInfo struct {
	TableExpression *Value  `@@`
	AsAlias         AsAlias `@@?`
	Timestamp       string  `("FOR" "SYSTEM_TIME" "AS" "OF" @Ident)?`
}

type AliasableQueryExpression struct {
	QueryExpression *QueryExpression `"(" @@ ")"`
	AsAlias         AsAlias          ` @@?`
}

type FromItem struct {
	TableInfo *TableInfo `( @@`
	// JoinOperation            []*JoinOperation          `| ( @@ | "(" @@ ")" )`
	AliasableQueryExpression *AliasableQueryExpression `| @@`
	JoinOperation            *JoinOperation            `| @@`
	FieldPath                *Value                    `| @@ )`
	// UnnestOperator  UnnestOperator  `| @@ `
	// CTEAlias AsAlias `@@?`
}

type FromStatement struct {
	FromItems []*FromItem `"FROM" @@ ( "," @@ )*`
}
