//TODO html with http serve templates templating
//TODO ajax
package main

import (
	"fmt"
	"html/template"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//tmpl, err := template.ParseFiles("/var/www/apps/go/index.html")

		myvar := map[string]interface{}{"User": "Go website", "Name": "A (modular, highly tweakable) responsive one-page template designed by HTML5 UP and released for free under the Creative Commons."}
		outputHTML(w, "/static/index.html", myvar)
		//  http.HandleFunc("/greet/", func(w http.ResponseWriter, r *http.Request) {
		//    name := r.URL.Path[len("/greet/"):]
		//  fmt.Fprintf(w, "Hello %s\n", name)
		//})
	})

	//http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Starting server at port 3008\n")
	// if err := http.ListenAndServe(":3006", nil); err != nil {
	//     log.Fatal(err)
	// }
	//connect maria
	mariacon()
	//connect postgresql
	pgcon()
	//connect mongodb
	moncon()
	//get data from api
	getapi()
	if err := http.ListenAndServeTLS(":3008", "/etc/letsencrypt/live/parapera.gr/fullchain.pem", "/etc/letsencrypt/live/parapera.gr/privkey.pem", nil); err != nil {
		log.Fatal(err)
	}
}

func outputHTML(w http.ResponseWriter, filename string, data interface{}) {
	//fileServer := http.FileServer(http.Dir("./static"))
	t, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
