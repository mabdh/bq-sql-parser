package main

type PivotAggregateFunctionCall struct {
	Function string  `@Ident `
	Alias    AsAlias `( @@ )?`
}

type InPivotColumn struct {
	PivotColumn string  `@Ident`
	Alias       AsAlias `( @@ )?`
}

type PivotOperator struct {
	PivotAggregateFunctionCalls []PivotAggregateFunctionCall `"PIVOT ("( @@ )+`
	ForInputColumn              string                       `"FOR" @Ident`
	InPivotColumns              []InPivotColumn              `"IN (" ( @@ )+ ")"`
	Alias                       AsAlias                      `")" ( "AS" @Ident )?`
}
