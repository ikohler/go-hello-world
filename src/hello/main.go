package main

import (
	"log"
	"net/http"
)

func main() {
	<brokemblabla>
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8000", router))
}
