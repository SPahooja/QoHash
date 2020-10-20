package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/gorilla/mux"
)

// MyFile is just a smaller version of FileInfo with only the Name, Size in bytes and ModTime
type MyFile struct {
	Name    string    `json:"name"`
	Size    int64     `json:"size"`
	ModTime time.Time `json:"modtime"`
	IsDir   bool      `json:"isDir"`
}

// Addresss is a wrapper for address string
type Address struct {
	Addr string `json:"addr"`
}

var files []MyFile
var size int
var addr string = "../"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", getFiles).Methods("GET")
	r.HandleFunc("/", getAddr).Methods("POST")

	//files, size, err := FindFiles("C:")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Total Files: %d", size)
	//fmt.Println(files)

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	run()
	//fmt.Println(files)
	json.NewEncoder(w).Encode(files)
}

func getAddr(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)
	//addr = params["addr"]
	var ad Address
	_ = json.NewDecoder(r.Body).Decode(&ad)
	addr = ad.Addr
	fmt.Println("New addr: ", addr)
	w.Header().Set("Content-Type", "application/json")
	run()
	//fmt.Println(files)
	json.NewEncoder(w).Encode(files)
}

func run() error {

	//err = new error("something")
	files, size, _ = FindFiles(addr)
	// if err != nil {
	// 	return err
	// }
	return nil
}

//FindFiles takes in an address and return the files in the folder as an array of MyFile
func FindFiles(addr string) ([]MyFile, int, error) {
	files, err := ioutil.ReadDir(addr)
	if err != nil {
		return nil, 0, nil
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Size() < files[j].Size()
	})

	myfiles := []MyFile{}

	size := len(files)

	for _, f := range files {
		myf := MyFile{Name: f.Name(), Size: f.Size(), ModTime: f.ModTime(), IsDir: f.IsDir()}
		myfiles = append(myfiles, myf)

	}

	return myfiles, size, err
}
