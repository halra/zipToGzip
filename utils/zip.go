package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/kjk/lzmadec"
)

//Unzip the file returns an array with the extracted filenames
func Unzip(src string, dest string, password *string) ([]string, error) {
	var filenames []string

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
			fmt.Printf("Detected directory, dir scanning not yet implemented, skipping entry %v\n", f.Name)
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

		bytesCount, err1 := io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()
		fmt.Printf("Filenames extracted: %v wrote %v bytes \n", filenames, bytesCount)
		if err1 != nil {
			return nil, err
		}
	}

	//fmt.Printf("Filenames extracted: %v \n", filenames)
	return filenames, nil
}

//Unzip7z Prototype .... ... this is very rudimentary ...
func Unzip7z(src string, dest string, password *string) ([]string, error) {
	var archive *lzmadec.Archive
	fmt.Println(src)
	fmt.Println(dest)
	archive, e := lzmadec.NewArchive("./" + src)

	if e != nil {
		fmt.Printf("ERROR: %v \n", e)
	}
	var filenames []string

	// list all files inside archive
	//firstFile := archive.Entries[0].Path
	for _, e := range archive.Entries {
		fmt.Printf("name: %s, size: %d\n", e.Path, e.Size)
	}

	//TODO get extension from right file, check if file size is big enough ... this is very rudimentary ...
	firstFile := archive.Entries[0].Path
	var extension = filepath.Ext(archive.Entries[1].Path)
	filenames = append(filenames, firstFile+extension)
	// extract to a file
	archive.ExtractToFile(firstFile+extension, firstFile)

	// decompress to in-memory buffer
	//r, _ := archive.GetFileReader(firstFile)
	//var buf bytes.Buffer
	//_, _ = io.Copy(&buf, r)
	// if not fully read, calling Close() ensures that sub-launched 7z executable
	// is terminated
	//r.Close()
	//fmt.Printf("size of file %s after decompression: %d\n", firstFile, len(buf.Bytes()))

	return filenames, nil
}

//Zip compresses one or more files into a single zip archive
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
	header.Name = filename

	//deflate for better compression
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
