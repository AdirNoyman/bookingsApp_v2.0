package main

import (
	"bookingsApp_2.0/pkg/handlers"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Server started. Listening on port %s 🤓🤟", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
