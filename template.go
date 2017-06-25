package cms

import (
	"html/template"
)

// Capitol allows to be exported
var Tmpl = template.Must(template.ParseGlob("../templates/*"))

type Page struct {
	Title   string
	Content string
}
