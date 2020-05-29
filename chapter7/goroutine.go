package main

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	var urls = []string{
		"https://media.vlpt.us/post-images/veloss/4a03cfa0-c605-11e8-ab77-93833eaf1fa8/837200ed5211.jpg",
		"https://image.slidesharecdn.com/letsgo-150315195509-conversion-gate01/95/lets-go-golang-1-638.jpg",
		"https://img-a.udemycdn.com/course/750x422/2790842_bba8_2.jpg",
	}

	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			if _, err := download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}
	wg.Wait()

	filenames, err := filepath.Glob("./chapter7/download/*.jpg")
	if err != nil {
		log.Fatal(err)
	}
	err = writeZip("./chapter7/download/images.zip", filenames)
	if err != nil {
		log.Fatal(err)
	}
}

func writeZip(outputFilename string, filenames []string) error {
	outf, err := os.Create(outputFilename)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(outf)

	for _, filename := range filenames {
		w, err := zw.Create(filename)
		if err != nil {
			return err
		}

		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(w, f)
		if err != nil {
			return err
		}
	}
	return zw.Close()
}

func download(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	filename, err := urlToFilename(url)
	if err != nil {
		return "", nil
	}

	filenamePrefix := "./chapter7/download/"
	f, err := os.Create(filenamePrefix + filename)
	if err != nil {
		return "", nil
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)

	return filename, err
}

func urlToFilename(rawurl string) (string, error) {
	_url, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}

	return filepath.Base(_url.Path), nil
}
