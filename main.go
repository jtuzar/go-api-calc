package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello World!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/hello", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}
