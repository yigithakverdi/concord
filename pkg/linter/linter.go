package linter

type Linter struct {
	config   map[string]string
	cwd      string
	flags    []string
	services []Service
}
