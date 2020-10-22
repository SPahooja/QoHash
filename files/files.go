package files

import (
	"io/ioutil"
	"sort"
	"strconv"
)

// MyFile is just a smaller version of FileInfo with only the Name, Size in bytes and ModTime
type MyFile struct {
	Name    string `json:"name" header:"Name"`
	Size    string `json:"size" header:"Size(bytes)"`
	ModTime string `json:"modtime" header:"Modtime"`
	IsDir   bool   `json:"isDir"`
}

// Address is a wrapper for address string
type Address struct {
	Addr string `json:"addr"`
}

// var files []MyFile
// var size int
// var addr string = "../"

// func main() {
// 	r := mux.NewRouter()

// 	r.HandleFunc("/", getFiles).Methods("GET")
// 	r.HandleFunc("/", getAddr).Methods("POST")

// 	ch := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))

// 	log.Fatal(http.ListenAndServe(":8000", ch(r)))
// }

// func getFiles(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	run()
// 	json.NewEncoder(w).Encode(files)
// }

// func getAddr(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 	r.ParseForm()
// 	addr = r.FormValue("addr")
// 	fmt.Println("New addr: ", addr)
// }

// func run() error {

// 	//err = new error("something")
// 	files, size, _ = FindFiles(addr)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	return nil
// }

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
		s := ""
		if f.Size() != 0 {
			s = strconv.FormatInt(f.Size(), 10)
		}
		myf := MyFile{Name: f.Name(), Size: s, ModTime: f.ModTime().Format("15:04 2006-01-02"), IsDir: f.IsDir()}
		myfiles = append(myfiles, myf)

	}

	return myfiles, size, err
}
