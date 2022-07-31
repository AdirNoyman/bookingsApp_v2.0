package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {

	_, err := fmt.Fprintf(w, "<h1>This is the home page 🤓</h1>")
	if err != nil {
		println(err)

	}

}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("<h1>This is the about page 😎\nAnd 2 + 2 = %d</h1>", sum))

}

func addValues(x, y int) int {

	return x + y

}

func Divide(w http.ResponseWriter, r *http.Request) {

	x, y := 100.0, 20.0
	f, err := DivideValues(100.0, 20.0)
	if err != nil {

		fmt.Fprintf(w, "Can't devide by zero 😩")
		return
	}

	_, err = fmt.Fprintf(w, fmt.Sprintf("<h1>%f divided by %f is %f</h1>", x, y, f))
	if err != nil {
		println(err)
	}

}

func DivideValues(x, y float32) (float32, error) {

	if y <= 0 {
		err := errors.New("hey ASSHOLE 🤨 - can't devide by zero")
		return 0, err
	}
	result := x / y

	return result, nil
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/divide", Divide)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Server started. Listening on port %s 🤓🤟", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
