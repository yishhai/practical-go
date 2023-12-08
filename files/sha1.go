package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	h, err := sha1sum("http.log.gz")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("Hash: %v\n", h)

	h, err = sha1sum("sha1.go")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("Hash 2: %v\n", h)
}

/*
	 if filename ends with .gz
		$ cat filename.gz| gunzip | sha1sum/shasum
	 else
	 	# cat filename.*| sha1sum/shasum

This is the flow of this function,
the file will be uncompressed with gunzip only if it ends with .gz
*/
func sha1sum(filename string) (string, error) {
	// idiom: acquire a resource, check for error, defer release
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer file.Close() // defer are called in LIFO order, in the case were there a multiple defers
	var ir io.Reader = file

	/*
		fileInfo, err := file.Stat()
		if err != nil {
			// Handle error
		}
		buf := make([]byte, fileInfo.Size())

		file.Read(buf)
	*/

	if strings.HasSuffix(filename, "gz") {
		gzCompressedFile, err := gzip.NewReader(file)
		if err != nil {
			return "", err
		}
		defer gzCompressedFile.Close()
		ir = gzCompressedFile
	}

	// io.CopyN(os.Stdout, r, 100)
	// fmt.Println()

	w := sha1.New()
	if _, err := io.Copy(w, ir); err != nil {
		return "", err
	}

	hashBytes := w.Sum(nil)

	// fmt.Printf("buf: %v\n", buf[0:32])
	return fmt.Sprintf("%x", hashBytes), nil
}
