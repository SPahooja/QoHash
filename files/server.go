package main

import (
	"fmt"
	"net/http"

	"github.com/SPahooja/QoHash/files"
	"github.com/gorilla/mux"
)

var rfiles = []files.MyFile


func putAddr(w http.ResponseWriter, r *http.Request) {

}

func getFiles(w http.ResponseWriter, r *http.Request) {

}

func main() {
	rfiles, _, _ := files.FindFiles("C:")

	fmt.Println(rfiles)

	r := mux.NewRouter()

	r.HandleFunc("/", putAddr).Methods("PUT")
	r.HandleFunc("/", getFiles).Methods("GET")

	http.ListenAndServe(":8000", r)

}
