package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}
		fmt.Fprintf(w, "Hello there %s", "visitor")
	})

	srv := http.Server{
		Addr: ":8080",
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
