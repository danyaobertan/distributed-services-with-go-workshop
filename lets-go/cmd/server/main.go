package main

import (
	"log"

	"github.com/danyaobertan/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Println("Server listening on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
