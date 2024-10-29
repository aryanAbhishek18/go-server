package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	
	fileServer := http.FileServer(http.Dir("./static"));
	http.Handle("/", fileServer)
	
	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
