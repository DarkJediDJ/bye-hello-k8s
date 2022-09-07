package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("recieve request on /hello")

	fmt.Fprintf(w, "Hello from hello-service: %v", time.Now())
}

func helloArtyom(w http.ResponseWriter, r *http.Request) {
	log.Println("recieve request on /hello/person")

	fmt.Fprint(w, "Hello Person!")

}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/hello/person", helloArtyom)

	log.Fatalln(http.ListenAndServe(":3001", nil))
}
