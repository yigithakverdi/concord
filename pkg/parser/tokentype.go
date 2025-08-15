package parser

type TokenType struct {
	label      string
	keyword    string
	beforeExpr string
	startsExpr string
	isLoop     string
	isAssigng  string
	prefix     string
}
