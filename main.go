package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = The name is %s\n", name)
	fmt.Fprintf(w, "Address = The address is %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Starting server at port 3006\n")
	// if err := http.ListenAndServe(":3006", nil); err != nil {
	//     log.Fatal(err)
	// }
	if err := http.ListenAndServeTLS(":3008", "/etc/letsencrypt/live/parapera.gr/fullchain.pem", "/etc/letsencrypt/live/parapera.gr/privkey.pem", nil); err != nil {
		log.Fatal(err)
	}
}
