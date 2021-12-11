package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start")

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/callback", lineHandler)

	fmt.Println("http:localhost:8080で起動中・・・")

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello World!"
	fmt.Fprintf(w, msg)
}
