package main

import (
	"GoLang/zipToGzip/utils"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("./") // consider making this windows compatible
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".zip") {
			fmt.Printf("Decompressing: %v \n", f.Name())
			newFilename, _ := utils.Unzip(f.Name(), "")
			if len(newFilename) > 0 {
				//TODO iter over array and compress all files!
				utils.GzipFile(newFilename[0], f.Name())
				os.Remove(newFilename[0])
			} else {
				fmt.Printf("No filename found array: %v \n", newFilename)
			}

		}

	}
}
