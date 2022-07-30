package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		n, err := fmt.Fprintf(w, "<h1>Hello Adiros 🤓</h1>")

		if err != nil {

			fmt.Println("Error:", err)
		}

		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))

	})

	_ = http.ListenAndServe(":3000", nil)

}
