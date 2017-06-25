package cms


import (
	"net/http"
	"time"
)

func ServeIndex(w http.ResponseWriter, req *http.Request) {
	p := &Page{
		Title: "",
		Content: "",
		Posts: []*Post{
			&Post{
				
			}
		}
	}
}