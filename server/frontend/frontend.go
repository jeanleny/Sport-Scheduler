package front

import (
	"net/http"
	"fmt"
)


type S_frontEnd struct {}

func (f *S_frontEnd) Index (w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func (f *S_frontEnd) Clicked (w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "button clicked")	
}

