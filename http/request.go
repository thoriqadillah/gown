package http

import (
	"log"
	"net/http"
	"strings"
)

type response struct {
	url         string
	filename    string
	size        int64
	contentType string
	cansplit    bool
	totalpart   int
}

func Fetch(url string, splitnum int) (*response, error) {
	// get the redirected url
	res, err := http.Head(url)
	if err != nil {
		log.Printf("Error fetching url %v", err)
		return nil, err
	}

	newurl := res.Request.URL.String()
	if url != newurl {
		log.Printf("Following link to %s", newurl)
	}

	url = newurl

	// get content-length (size in bytes) of a file
	res, err = http.Head(url)
	if err != nil {
		log.Printf("Error fetching file url %v", err)
		return nil, err
	}

	size := res.ContentLength

	contentType := res.Header.Get("Content-Type")

	split := strings.Split(url, "/")
	filename := split[len(split)-1]

	// check if the file support cansplit download
	cansplit := res.Header.Get("Accept-Ranges") == "bytes"
	if !cansplit {
		log.Println("Does not support split download. Downloading the file entirely")
	}

	// total part is 8 if cansplit is true
	totalpart := 1
	if cansplit {
		totalpart = splitnum
	}

	response := &response{
		url:         url,
		filename:    filename,
		size:        size,
		contentType: contentType,
		cansplit:    cansplit,
		totalpart:   totalpart,
	}

	return response, nil
}

func (r *response) Parts() int {
	return r.totalpart
}

func (r *response) Size() int64 {
	return r.size
}

func (r *response) Filename() string {
	return r.filename
}

func (r *response) Type() string {
	split := strings.Split(r.contentType, "/")
	contentType := split[len(split)-1]

	return "." + contentType
}
