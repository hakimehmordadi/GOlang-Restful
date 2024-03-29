package controller

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

// The root directory for uploading & downloading porpuse
var base = "/home/mordadi/sample"

func Upload(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	n := fmt.Sprintf("%d - %s", time.Now().UTC().Unix(), file.Filename)
	dst := fmt.Sprintf("%s/%s", base, n)
	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)

	return n, err
}

func Download(n string) (string, []byte, error) {
	dst := fmt.Sprintf("%s/%s", base, n)
	b, err := ioutil.ReadFile(dst)
	if err != nil {
		return "", nil, err
	}
	// detect the content type of file
	m := http.DetectContentType(b[:512])

	return m, b, nil
}
