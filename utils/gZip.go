package utils

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//GzipAttributes Saves the original name TODO add more attribs
type GzipAttributes struct {
	name string
}

//GzipFile a file
func GzipFile(src string, dest string, attributes *GzipAttributes) string {

	var extension = filepath.Ext(dest)

	fmt.Printf("Gzipping source %v to %v \n", src, strings.Replace(dest, ".zip", ".gz", -1))
	// Open file on disk to read from.
	f, e1 := os.Open(src)

	if e1 != nil {
		fmt.Printf("Error on Gzip File %v \n", e1)
		return ""
	}

	// Create a Reader
	reader := bufio.NewReader(f)

	// Replace .zip extension with gz extension, actually replaces all extentions with .gz.
	name := dest[0:len(dest)-len(extension)] + ".gz"

	// Open file for writing.
	f1, _ := os.Create(name)

	// Write compressed data.
	w := gzip.NewWriter(f1)
	w.Name = attributes.name

	byteCount, err := io.Copy(w, reader)
	if err != nil {
		fmt.Println(err)
	}

	//close the writer first!
	w.Close()
	f.Close()
	f1.Close()

	fmt.Printf("\n")

	// Done.
	fmt.Printf("Done for %v wrote %v bytes \n", name, byteCount)

	return name
}

//UnGZip consumes source as path extracts to dest Path
func UnGZip(src string, dest string, password string) string {
	fmt.Println("Not yet implemented")
	return ""
}
