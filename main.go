package main

import (
	"github.com/JeremiasArriondo/http-server/server"
)

func main() {

	srv := server.New(":8080")

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
