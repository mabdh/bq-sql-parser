package main

type Number struct {
	Number *float64 `@Number`
}

type QuotedString struct {
}

// type String struct {
// 	Quoted       *QuotedString `( @QuotedString`
// 	TripleQuoted *string       `| @TripleQuotedString `
// 	Raw          *string       `| @RawString `
// 	Byte         *string       `| @ByteString `
// 	RawByte      *string       `| @RawByteString `
// 	Special      *string       `| @SpecialString )`
// }

type Name struct {
	Name *string `@String | """ @String """ | "(" @String ")"  | "'" @String "'"` // | "`" @String "`" //TODO
}

type CTEName struct {
	Name *Name `@@`
}

type ProjectName struct {
	Name *Name `@@`
}

type DatasetName struct {
	Name *Name `@@`
}

type TableName struct {
	Name *Name `@@`
}

type TableExpression struct {
	ProjectName     *ProjectName     `(( @@ "." )?`
	DatasetName     *DatasetName     `@@ "." )?`
	TableName       *TableName       `@@`
	TableExpression *TableExpression `| @@ ` // '`' table_expr '`'
}

// type Expression struct {
// 	Number *Number `( @@`
// 	String *String `| @@ )`
// }
