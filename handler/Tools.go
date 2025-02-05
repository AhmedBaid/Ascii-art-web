package ascii

import (
	"text/template"
)

var Tp *template.Template

type ErrorPage struct {
	Code         int
	ErrorMessage string
}
