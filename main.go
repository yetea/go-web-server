package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if error := r.ParseForm(); error != nil {
		fmt.Print(w, "ParseForm() err: %v", error)
		return
	}
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	password := r.FormValue("password")
	fmt.Fprintf(w, "firstName = %s\n", firstName)
	fmt.Fprintf(w, "lastName = %s\n", lastName)
	fmt.Fprintf(w, "password = %s\n", password)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n")
	if error := http.ListenAndServe(":8080", nil); error != nil {
		log.Fatal(error)
	}
}
