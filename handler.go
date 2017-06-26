package cms

import (
	"net/http"
	"time"
)

func ServeIndex(w http.ResponseWriter, req *http.Request) {
	p := &Page{
		Title:   "TK Go CMS",
		Content: "Welcome to the home page!",
		Posts: []*Post{
			&Post{
				Title:         "Hello, Earth",
				Content:       "Hello, Thank you for being a G.",
				DatePublished: time.Now(),
			},
			&Post{
				Title:         "A Post w/ Comments",
				Content:       "This post will attract comments of course because it is phenomenal.",
				DatePublished: time.Now().Add(-time.Hour),
				Comments: []*Comment{
					&Comment{
						Author:        "Trevor Knott",
						Comment:       "What a terrible short post ehh -____-",
						DatePublished: time.Now().Add(-time.Hour / 2),
					},
				},
			},
		},
	}

	Tmpl.ExecuteTemplate(w, "page", p)
}

func HandleNew(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		Tmpl.ExecuteTemplate(w, "new", nil)

	case "POST":
		title := req.FormValue("post-title")
		content := req.FormValue("post-content")
		contentType := req.FormValue("post-content-type")
		req.ParseForm()

		if contentType == "page" {
			Tmpl.ExecuteTemplate(w, "page", &Page{
				Title:   title,
				Content: content,
			})
			return
		}
	default:
		http.Error(w, "method not supported: "*req.Method, http.StatusMethodNotAllowed)
	}
}

// TK
