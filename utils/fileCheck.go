package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"log"
	"os"
)

func Sha256(file string) string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	uEnc := base64.StdEncoding.EncodeToString(h.Sum(nil))
	//fmt.Println(uEnc)

	//fmt.Printf("%v", uEnc)
	return uEnc
}
