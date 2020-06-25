package utils

import (
	"archive/zip"
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//Unzip the file
func Unzip(src string, dest string) ([]string, error) {
	var filenames []string

	//consider using this reader as input for gzip ... :)
	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	//this should only return a single filename
	for _, f := range r.File {

		// Store filename/path for returning and using later on
		storedPath := filepath.Join(dest, f.Name)

		filenames = append(filenames, storedPath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(storedPath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(storedPath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(storedPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}

	fmt.Printf("Filenames extracted: %v \n", filenames)
	return filenames, nil
}

//GzipFile a file
func GzipFile(uncompressedName string, compressedString string) {

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
	w.Name = uncompressedName

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
