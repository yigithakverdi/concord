package parser

type Point struct {
	X int
	Y int
}

type Parser struct {
	file     string
	position Point
	ast      map[string]string
}
