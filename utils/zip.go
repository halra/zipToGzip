package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

//Unzip the file returns an array with the extracted filenames
func Unzip(src string, dest string, password *string) ([]string, error) {
	var filenames []string

	//consider using this reader as input for gzip ... :)
	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	//iterrate over all files in the archive
	for _, f := range r.File {

		if f.FileInfo().IsDir() {
			// Make Folder
			//os.MkdirAll(storedPath, os.ModePerm)
			continue
		}
		// Store filename/path for returning and using later on
		storedPath := filepath.Join(dest, f.Name)

		filenames = append(filenames, storedPath)

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
			return nil, err
		}
	}

	fmt.Printf("Filenames extracted: %v \n", filenames)
	return filenames, nil
}

//Zip compresses one or many files into a single zip archive file.
// Param 1: filename is the output zip file's name.
// Param 2: files is a list of files to add to the zip.
func Zip(src []string, dest string, password string) error {
	newZipFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range src {
		if err = addFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func addFileToZip(zipWriter *zip.Writer, filename string) error {

	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Using FileInfoHeader() above only uses the basename of the file. If we want
	// to preserve the folder structure we can overwrite this with the full path.
	header.Name = filename

	// Change to deflate to gain better compression
	// see http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
