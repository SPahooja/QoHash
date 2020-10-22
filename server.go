package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SPahooja/QoHash/files"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var fileslist []files.MyFile
var size int
var addr string = "../"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", getFiles).Methods("GET")
	r.HandleFunc("/", getAddr).Methods("POST")

	ch := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))

	log.Fatal(http.ListenAndServe(":8000", ch(r)))
}

func getFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := run()
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewEncoder(w).Encode(fileslist)
	if err != nil {
		log.Fatal(err)
	}
}

func getAddr(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}
	addr = r.FormValue("addr")
}

func run() error {

	f, s, err := files.FindFiles(addr)
	if err != nil {
		return err
	}
	fileslist = f
	size = s
	return nil
}
