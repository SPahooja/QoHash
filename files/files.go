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

//FindFiles takes in an address and return the files in the folder as an array of MyFile, size of the array and error
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

	return myfiles, size, nil
}
