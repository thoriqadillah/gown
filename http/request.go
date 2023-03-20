package http

import (
	"log"
	"net/http"
	"strings"
)

type response struct {
	url       string
	filename  string
	size      int64
	ranged    bool
	totalpart int
}

func Fetch(url string) (*response, error) {
	// get the redirected url
	res, err := http.Head(url)
	if err != nil {
		log.Printf("Error fetching url %v", err)
		return nil, err
	}
	url = res.Request.URL.String()

	// get content-length (size in bytes) of a file
	res, err = http.Head(url)
	if err != nil {
		log.Printf("Error fetching file url %v", err)
		return nil, err
	}
	size := res.ContentLength

	// get file name
	split := strings.Split(url, "/")
	filename := split[len(split)-1]

	// check if the file support ranged download
	ranged := res.Header.Get("Accept-Ranges") == "bytes"

	// total part is 8 if ranged is true
	totalpart := 1
	if ranged {
		totalpart = 8
	}

	response := &response{
		url:       url,
		filename:  filename,
		size:      size,
		ranged:    ranged,
		totalpart: totalpart,
	}

	return response, nil
}

func (r *response) Parts() int {
	return r.totalpart
}
