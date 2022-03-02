package main

import (
	"github.com/c0llinn/prolog/internal/server"
	"log"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatalln(srv.ListenAndServe())
}
