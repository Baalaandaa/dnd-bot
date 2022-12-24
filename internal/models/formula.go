package models

type FormulaState struct {
	Entered bool
}

var DefaultFormulaState = FormulaState{
	Entered: false,
}
