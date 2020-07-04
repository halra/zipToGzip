package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"io"
	"log"
	"os"
)

//Sha256 returns the hash Bse64 encoded
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

	b64String := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return b64String
}
