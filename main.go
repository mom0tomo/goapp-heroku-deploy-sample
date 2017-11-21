package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "ʕ ◔ϖ◔ʔ")
	})

	n := negroni.Classic()
	n.UseHandler(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3333"
	}
	http.ListenAndServe(":"+port, n)
}
