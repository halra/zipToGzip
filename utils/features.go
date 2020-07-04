package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//ZipToGZip converts zip archives in a directory to gzip archives
func ZipToGZip() {
	files, err := ioutil.ReadDir("./") // read current dir
	if err != nil {
		log.Fatal(err)
	}

	var gzipAttribs *GzipAttributes
	gzipAttribs = new(GzipAttributes)
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".zip") {
			fmt.Printf("Decompressing: %v \n", f.Name())
			newFilename, _ := Unzip(f.Name(), "", nil)

			if len(newFilename) > 0 {
				//TODO iter over array and compress all files!
				//fmt.Printf("hash of [%v] : -> [%v]", newFilename[0], Sha256(newFilename[0]))
				gzipAttribs.name = newFilename[0]
				GzipFile(newFilename[0], f.Name(), gzipAttribs)
				//fmt.Printf("written %v with hash %v \n", gzippedName, Sha256(gzippedName))
				os.Remove(newFilename[0])
				fmt.Printf("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n\n")
			} else {
				fmt.Printf("No filename found is the archive empty?  returned array: %v \n", newFilename)
			}

		}

	}
}
