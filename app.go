package main

import (
	"fmt"
	"net/http"
)

func rootRout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {

	fmt.Println("Hello world")

	http.HandleFunc("/", rootRout)
	http.ListenAndServe(":80", nil)

}
