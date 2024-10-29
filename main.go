package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(res, "method not supported", http.StatusNotFound)
	}
	fmt.Fprint(res, "Hello user!");
}

func main() {
	
	fileServer := http.FileServer(http.Dir("./static"));
	http.Handle("/", fileServer)
	
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
