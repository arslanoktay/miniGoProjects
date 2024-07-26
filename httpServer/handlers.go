package handlers

import (
	"fmt"
	"net/http"
)

func HttpServerHandlers() {

	// register a new handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Wellcome")
	})

	// serves static files like css or js or image
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
