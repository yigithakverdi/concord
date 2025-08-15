package parser

type Type struct {
}

type Value struct {
}

type Token struct {
	_type  Type
	_value Value
	_start int
	_end   int
}
