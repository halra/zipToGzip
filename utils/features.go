package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//Foo print foo
func Foo() {
	files, err := ioutil.ReadDir("./") // consider making this windows compatible
	if err != nil {
		log.Fatal(err)
	}

	var gzipAttribs *GzipAttributes

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".zip") {
			fmt.Printf("Decompressing: %v \n", f.Name())
			newFilename, _ := Unzip(f.Name(), "", nil)
			if len(newFilename) > 0 {
				//TODO iter over array and compress all files!
				gzipAttribs.name = newFilename[0]
				GzipFile(newFilename[0], f.Name(), gzipAttribs)
				os.Remove(newFilename[0])
			} else {
				fmt.Printf("No filename found array: %v \n", newFilename)
			}

		}

	}
}
