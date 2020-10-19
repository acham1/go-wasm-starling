package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Listening at :8080")
	http.ListenAndServe(":8080", http.FileServer(http.Dir("docs")))
}
