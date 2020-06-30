package utils

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

//GzipAttributes foobar
type GzipAttributes struct {
	name string
}

//GzipFile a file
func GzipFile(uncompressedName string, compressedString string, attributes *GzipAttributes) {

	fmt.Printf("Gzipping source %v to %v \n", uncompressedName, strings.Replace(compressedString, ".zip", ".gz", -1))
	// Open file on disk.
	name := uncompressedName
	f, _ := os.Open(name)

	// Create a Reader and use ReadAll to get all the bytes from the file.
	reader := bufio.NewReader(f)
	//content, _ := ioutil.ReadAll(reader)

	// Replace .zip extension with gz extension.
	name = strings.Replace(compressedString, ".zip", ".gz", -1)

	// Open file for writing.
	f1, _ := os.Create(name)

	// Write compressed data.
	w := gzip.NewWriter(f1)
	w.Name = attributes.name

	if _, err := io.Copy(w, reader); err != nil {
		fmt.Println(err)
	}

	f.Close()
	f1.Close()
	fmt.Printf("\n")

	w.Close()

	// Done.
	fmt.Printf("Done for %v \n", name)
}

//UnGZip consumes source as path extracts to dest Path
func UnGZip(src string, dest string, password string) string {
	return ""
}
