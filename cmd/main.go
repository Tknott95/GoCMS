package main

import (
	"os"

	"github.com/tknott95/cms"
)

func main() {
	p := &cms.Page{
		Title:   "Hello, world!",
		Content: "this is the body of our webpage",
	}

	cms.Tmpl.ExecuteTemplate(os.Stdout, "index", p)
}
