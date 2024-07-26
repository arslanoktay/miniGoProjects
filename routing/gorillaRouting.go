package gorillaRouting

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func MainRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/book/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r) // Vars take request as paramete, return map of segments
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})
	http.ListenAndServe(":8000", r) // r yerine nil olsa default olarak net/http kullan demek olur kendi routerımız için request verdik
}
