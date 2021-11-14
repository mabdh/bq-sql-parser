package main

type AsAlias struct {
	Alias string `("AS")? @Ident`
}

type TableInfo struct {
	TableName string  `@Ident`
	AsAlias   AsAlias `(@@)?`
	Timestamp string  `("FOR SYSTEM_TIME AS OF" @Ident)?`
}

type AliasableQueryExpression struct {
	QueryExpression *QueryExpression `"(" @@ ")"`
	AsAlias         AsAlias          ` @@?`
}

type FromItem struct {
	TableInfo                *TableInfo                `( @@`
	JoinOperation            []*JoinOperation          `| ( @@ | "(" @@ ")" )`
	AliasableQueryExpression *AliasableQueryExpression `| @@`
	FieldPath                *Value                    `| @@ )`
	// UnnestOperator  UnnestOperator  `| @@ `
	// CTEAlias AsAlias `@@?`
}

type FromStatement struct {
	FromItem *FromItem `@@`
	// PivotOperator       PivotOperator       `( @@`
	// UnpivotOperator     UnpivotOperator     `| @@)?`
	// TableSampleOperator TableSampleOperator `( @@ )?`
}
