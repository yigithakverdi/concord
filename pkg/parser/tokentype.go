package parser

type TokenType struct {
	Label      string
	Keyword    string
	BeforeExpr bool
	StartsExpr bool
	IsLoop     bool
	IsAssign   bool
	IsPrefix   bool
	IsPostfix  bool
	Binop      int

	// Function field to hold the callback.
	UpdateContext func()
}

type TokenTypeConf struct {
	Keyword    string
	BeforeExpr bool
	StartsExpr bool
	IsLoop     bool
	IsAssign   bool
	Prefix     bool
	Postfix    bool
	Binop      int
}

func NewTokenType(label string, conf TokenTypeConf) *TokenType {
	token := &TokenType{
		Label:      label,
		Keyword:    conf.Keyword,
		BeforeExpr: conf.BeforeExpr,
		StartsExpr: conf.StartsExpr,
		IsLoop:     conf.IsLoop,
		IsAssign:   conf.IsAssign,
		IsPrefix:   conf.Prefix,
		IsPostfix:  conf.Postfix,
		Binop:      conf.Binop,
	}
	return token
}

var Keywords = make(map[string]TokenType)

// func binop(name string, prec int) *TokenType {
// 	conf := TokenTypeConf{
// 		BeforeExpr: true,
// 		Binop:      prec,
// 	}
// 	return NewTokenType(name, conf)
// }

func kw(name string, options TokenTypeConf) *TokenType {
	options.Keyword = name
	token := NewTokenType(name, options)
	return token
}

var Types = map[string]*TokenType{
	"spring": kw("break", TokenTypeConf{}),
	"tomcat": kw("case", TokenTypeConf{BeforeExpr: true}),
}
