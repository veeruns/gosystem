package gosystem

import (
	"fmt"
	"io"
	"os"
)

const (
	MAXDIRS = 2000
)

func Walkfilesystem(dir string) {
	f, err := os.Open(dir)
	if err != nil {
		panic("Invalid open file " + err.Error())
	}
	files, err := f.Readdir(MAXDIRS)
	if err != io.EOF {
		if err != nil {
			panic("Inside maxdirs " + err.Error())
		}
		for _, fi := range files {
			fmt.Printf("%s/%s\n", dir, fi.Name())
			if fi.IsDir() {
				Walkfilesystem(fmt.Sprintf("%s/%s", dir, fi.Name()))
			}
		}
	}
}
