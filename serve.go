package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.ListenAndServe("199.58.162.130:80", nil)
}
