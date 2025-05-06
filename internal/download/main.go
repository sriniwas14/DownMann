package download

import (
	"downmann/internal/configloader"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"mime"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/lucsky/cuid"
)

type DownloadPart struct {
	From      int64 `json:"from"`
	To        int64 `json:"to"`
	Completed bool  `json:"completed"`
}

type Download struct {
	Id          string          `json:"id"`
	Url         string          `json:"url"`
	Size        int64           `json:"size"`
	Parts       []*DownloadPart `json:"parts"`
	Cursor      int
	Range       bool   `json:"range"`
	Status      int    `json:"status"`
	Destination string `json:"destination"`
}

var client = &http.Client{
	Timeout: 10 * time.Second,
}

func NewDownload(url string) (*Download, error) {
	res, err := http.Head(url)

	if err != nil {
		return nil, err
	}

	rangeSupported := false
	rangeHeader := res.Header.Get("Accept-Ranges")
	contentType := strings.Split(res.Header.Get("Content-Type"), ";")[0]
	contentLength, err := strconv.Atoi(res.Header.Get("Content-Length"))
	contentDisposition := strings.Split(res.Header.Get("Content-Disposition"), ";")[0]

	if err != nil {
		return nil, err
	}

	dest, err := getDestFileName(url, contentType, contentDisposition)
	log.Println(dest)

	parts := []*DownloadPart{}
	if rangeHeader == "bytes" {
		rangeSupported = true

		log.Println("Range is supported, splitting chunks")

		chunkSize := math.Ceil(float64(contentLength) / float64(configloader.MaxThreads))
		for i := range configloader.MaxThreads {
			from := int64(chunkSize) * i
			to := int64(chunkSize) * (i + 1)
			if i > 0 {
				from++
			}
			parts = append(parts, &DownloadPart{
				From:      from,
				To:        to,
				Completed: false,
			})
		}
	} else {
		log.Println("Range not supported!")
		parts = append(parts, &DownloadPart{
			From:      0,
			To:        int64(contentLength),
			Completed: false,
		})
	}

	if err != nil {
		return nil, err
	}

	return &Download{
		Id:          cuid.Slug(),
		Url:         url,
		Parts:       parts,
		Size:        int64(contentLength),
		Range:       rangeSupported,
		Destination: dest,
	}, nil
}

func (d *Download) Start() error {
	jobs := make(chan *DownloadPart, len(d.Parts))
	errors := make(chan error, len(d.Parts))

	log.Printf("Download has %d parts\n", len(d.Parts))

	for w := 1; w <= 2; w++ {
		go d.DownloadPart(jobs, errors)
	}

	for j := range len(d.Parts) {
		jobs <- d.Parts[j]
	}
	close(jobs)

	for range len(d.Parts) {
		err := <-errors
		log.Println("Error ", err)
	}

	return nil
}

func (d *Download) DownloadPart(jobs <-chan *DownloadPart, errors chan<- error) {
	for part := range jobs {
		cursor := d.Cursor
		var m sync.Mutex
		m.Lock()
		d.Cursor++
		m.Unlock()
		log.Printf("Downloading part %d\n", cursor)
		req, err := http.NewRequest("GET", d.Url, nil)
		if err != nil {
			errors <- err
			return
		}

		rangeValue := fmt.Sprintf("bytes=%d-%d", part.From, part.To)
		req.Header.Add("Range", rangeValue)

		dest := fmt.Sprintf("%s/%s/%s_%d", configloader.DestFolder, "/Temp/", d.Id, cursor)
		out, err := os.Create(dest)
		writer := &ProgressWriter{
			part:   d.Parts[0],
			writer: out,
		}

		if err != nil {
			errors <- err
			return
		}
		defer out.Close()

		res, err := client.Do(req)
		if err != nil {
			errors <- err
			return
		}

		io.Copy(writer, res.Body)

		errors <- nil
	}
}

func (d *Download) Debug() {
	data, _ := json.Marshal(d)
	log.Println(string(data))
}

func (d *Download) Pause(cb func()) {

}

func getDestFileName(url, contentType, contentDisposition string) (string, error) {

	filename := ""

	if contentDisposition == "attachment" {
		nameParts := strings.Split(url, "/")
		filename = nameParts[len(nameParts)-1]
	} else {
		ext, err := mime.ExtensionsByType(contentType)
		if err != nil {
			return "", err
		}

		if len(ext) > 0 {
			filename = cuid.Slug() + ext[0]
		}
	}

	return configloader.DestFolder + "/Downloads/" + filename, nil
}
