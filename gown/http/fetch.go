package http

import (
	"changeme/gown/setting"
	"log"
	"math"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
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

	filename := filename(res, setting)
	contentType := contentType(filename)

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

func contentType(filename string) string {
	if match, _ := regexp.MatchString(`^.*.(jpg|jpeg|png|gif|svg|bmp)$`, filename); match {
		return "image"
	}

	if match, _ := regexp.MatchString(`^.*\.(mp4|mov|avi|mkv|wmv|flv|webm|mpeg|mpg|3gp|m4v|m4a)$`, filename); match {
		return "video"
	}

	if match, _ := regexp.MatchString(`^.*.(mp3|wav|flac|aac|ogg|opus)$`, filename); match {
		return "audio"
	}

	if match, _ := regexp.MatchString(`^.*.(doc|docx|pdf|txt|ppt|pptx|xls|xlsx|odt|ods|odp|odg|odf|rtf|tex|texi|texinfo|wpd|wps|wpg|wks|wqd|wqx|w)$`, filename); match {
		return "document"
	}

	if match, _ := regexp.MatchString(`^.*.(zip|rar|7z|tar|gz|bz2|tgz|tbz2|xz|txz|zst|zstd)$`, filename); match {
		return "compressed"
	}

	return "other"
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

func filename(res *http.Response, s *setting.Settings) string {
	//TODO: there is no guarantee that this approach is always works. Add more implementation
	disposition := res.Header.Get("Content-Disposition")
	_, params, _ := mime.ParseMediaType(disposition)

	return handleDuplicate(params["filename"], s)
}

func handleDuplicate(name string, s *setting.Settings) string {
	filename := filepath.Join(s.SaveLocation, name)
	if _, err := os.Stat(filename); err != nil {
		return name
	}

	regex, err := regexp.Compile(`\((.*?)\)`)
	if err != nil { // if there is no number prefix
		return name
	}

	prefix := regex.FindStringSubmatch(name)
	if len(prefix) == 0 {
		split := strings.Split(name, ".")
		split[0] += " (1)"
		name = strings.Join(split, ".")
		name = handleDuplicate(name, s)
		return name
	}

	name = strings.ReplaceAll(name, " "+prefix[0], "")
	number, err := strconv.Atoi(prefix[1])
	if err != nil {
		return name
	}
	split := strings.Split(name, ".")
	split[0] += " (" + strconv.Itoa(number+1) + ")"
	name = strings.Join(split, ".")
	name = handleDuplicate(name, s)

	return name
}
