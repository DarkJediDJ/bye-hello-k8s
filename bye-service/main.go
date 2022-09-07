package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func bye(w http.ResponseWriter, r *http.Request) {
	log.Println("recieve request on /bye")

	fmt.Fprintf(w, "Bye from bye-service: %v", time.Now())
}

func byeArtyom(w http.ResponseWriter, r *http.Request) {
	log.Println("recieve request on /bye/person")

	fmt.Fprint(w, "Bye Person!")
}

func main() {
	http.HandleFunc("/bye", bye)
	http.HandleFunc("/bye/person", byeArtyom)

	log.Fatalln(http.ListenAndServe(":3003", nil))
}
