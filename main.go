package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "hello, world.")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":1324", router))
}
