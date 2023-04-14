package http

import (
	"changeme/gown/setting"
	"log"
	"math"
	"net/http"
	"strings"
)

type Response struct {
	Url         string           `json:"url"`
	Size        int64            `json:"size"`
	ContentType string           `json:"contentType"`
	Cansplit    bool             `json:"cansplit"`
	Totalpart   int              `json:"totalpart"`
	Filename    string           `json:"filename"`
	Settings    setting.Settings `json:"settings"`
}

func Fetch(url string, setting *setting.Settings) (*Response, error) {
	log.Printf("Fetching the URL\n")

	res, err := http.Head(url)
	if err != nil {
		log.Printf("Error fetching url %v", err)
		return nil, err
	}

	// get the redirected URL
	newurl := res.Request.URL.String()
	if url != newurl {
		log.Printf("Following link to %s", newurl[:50]+"...")
	}

	url = newurl

	// get content-length (size in bytes) of a file
	res, err = http.Head(url)
	if err != nil {
		log.Printf("Error fetching file url %v", err)
		return nil, err
	}

	contentType := res.Header.Get("Content-Type")

	// check if the file support cansplit download
	cansplit := res.Header.Get("Accept-Ranges") == "bytes"
	size := res.ContentLength

	totalpart := dynamicPartition(size, setting.Partsize)
	if size == -1 || !cansplit {
		totalpart = 1
		log.Println("File does not support download in chunks. Downloading the file entirely")
	}

	setting.Concurrency = totalpart

	response := &Response{
		Url:         url,
		Size:        size,
		ContentType: contentType,
		Cansplit:    cansplit,
		Totalpart:   totalpart,
		Filename:    filename(contentType, url),
		Settings:    *setting,
	}

	return response, nil
}

func dynamicPartition(size int64, defaultParitionSize int64) int {
	num := math.Log10(float64(size / (1024 * 1024)))
	partsize := defaultParitionSize
	for i := 0; i < int(num); i++ {
		partsize *= 3 // 3 is just author's self configured number
	}

	return int(size / partsize)
}

func filename(contentType string, url string) string {
	split := strings.Split(contentType, "/")
	type_ := "." + split[len(split)-1]

	split = strings.Split(url, "/")
	filename := split[len(split)-1]
	filename = strings.Split(filename, "?")[0]
	if filename == "" {
		return "file" + type_
	}

	split = strings.Split(filename, ".")
	if len(split) != 0 {
		return filename + type_
	}

	return filename
}
