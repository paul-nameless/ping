package main

import (
	"fmt"
	"flag"
	"net/http"
)

var response = flag.String("r", "pong", "response of /ping endpoint")
var bind = flag.String("b", ":8080", "bind address")

func pongHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, *response)
}



func main() {
	flag.Parse()
	http.HandleFunc("/ping", pongHandler)
	fmt.Printf("Listening %s: /ping -> %s\n", *bind, *response)
	http.ListenAndServe(*bind, nil)
}
