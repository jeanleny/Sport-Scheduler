package front

import (
	"net/http"
	"fmt"
	"html"
	"log"
)


type S_frontEnd struct {}

func (f *S_frontEnd) Index (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func (f *S_frontEnd) Clicked (w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "MAIS NAN LE CLICKAAA")
}

func (f *S_frontEnd) EnteredText (w http.ResponseWriter, r *http.Request){
	q := r.URL.Query().Get("q")
	log.Println("received:", q)
	fmt.Fprintf(w, "<p>You sent: %s</p>", html.EscapeString(q))
}
