package http

import (
	"changeme/gown/setting"
	"log"
	"math"
	"mime"
	"net/http"
	"regexp"
)

type Response struct {
	Url         string            `json:"url"`
	Size        int64             `json:"size"`
	ContentType string            `json:"contentType"`
	Cansplit    bool              `json:"cansplit"`
	Totalpart   int               `json:"totalpart"`
	Filename    string            `json:"filename"`
	Settings    *setting.Settings `json:"settings"`
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

	filename := filename(res)
	regex := regexp.MustCompile("^.*.(zip|tar)$")
	contentType := regex.ReplaceAllString(filename, "zip")
	log.Println(contentType)

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
		Filename:    filename,
		Settings:    setting,
	}

	return response, nil
}

func dynamicPartition(size int64, defaultParitionSize int64) int {
	num := math.Log10(float64(size / (1024 * 1024)))
	partsize := defaultParitionSize

	// dampening the total partition
	for i := 0; i < int(num); i++ {
		partsize *= 2 // 2 is just author's self configured number
	}

	return int(size / partsize)
}

func filename(res *http.Response) string {
	disposition := res.Header.Get("Content-Disposition")
	_, params, _ := mime.ParseMediaType(disposition)
	return params["filename"]
}
