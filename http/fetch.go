package http

import (
	"log"
	"net/http"
	"strings"

	"github.com/thoriqadillah/gown/config"
)

type response struct {
	url         string
	Filename    string
	size        int64
	contentType string
	cansplit    bool
	partition   int
	*config.Config
}

func Fetch(url string, conf ...*config.Config) (*response, error) {
	// get the redirected url
	res, err := http.Head(url)
	if err != nil {
		log.Printf("Error fetching url %v", err)
		return nil, err
	}

	newurl := res.Request.URL.String()
	if url != newurl {
		log.Printf("Following link to %s", newurl[:len(newurl)/2]+"...")
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
	split := strings.Split(contentType, "/")
	contentType = "." + split[len(split)-1]

	split = strings.Split(url, "/")
	filename := split[len(split)-1]
	if len(filename) > 256 {
		filename = "file" + contentType
	}

	// check if the file support cansplit download
	cansplit := res.Header.Get("Accept-Ranges") == "bytes"

	config := config.Default()
	if conf != nil {
		config = conf[0]
	}

	partition := size / config.Partsize

	response := &response{
		url:         url,
		Filename:    filename,
		size:        size,
		contentType: contentType,
		cansplit:    cansplit,
		partition:   int(partition),
		Config:      config,
	}

	return response, nil
}

func (r *response) Size() int64 {
	return r.size
}
