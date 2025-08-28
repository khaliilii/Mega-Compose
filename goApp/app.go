package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port string = ":8080"
	handler := http.HandlerFunc(Helloweb)
	log.Printf("is adtress recive %s", port)
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatalf("could not listen port %s %v", port, err)

	}
}
func Helloweb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world! %s", r.URL.Path)
	log.Printf("recieve requre for path %s", r.URL.Path)
}
