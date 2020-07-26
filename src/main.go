package main

import(
	"fmt"
	"net/http"
	"time"
	"log"
	"html/template"
	"os"
	"github.com/gorilla/mux"
)


var templates = template.Must(template.ParseGlob("templates/*"))

func main(){
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/",welcomeHandler).Methods("GET")
	r.HandleFunc("/about",aboutHandler).Methods("GET")
//	r.HandleFunc("/{name}",nameHandler).Methods("GET")
	r.HandleFunc("/home",homeHandler).Methods("GET")

	//manage PORT
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}


	server := &http.Server{
		Addr: ":"+PORT,
		Handler :r,
	}
	log.Printf("Listening on port %s...",PORT)
	server.ListenAndServe()
}

func welcomeHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Welcome Mr Edwin")
}

func aboutHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"<h1 align=\"center\">MY SITE</h1>")
}

func nameHandler(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w,"<h1 align=\"center\">WELCOME MR. %s TO YOUR SITE.<br> YOU VISITED THIS SITE ON %s</h1>",name,time.Now())
}

func homeHandler(w http.ResponseWriter,r *http.Request){
	err := templates.ExecuteTemplate(w,"base",nil)
	if err != nil{
		log.Println(err)
	}
}
