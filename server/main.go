package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, azd!")
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":3100", nil)
}
