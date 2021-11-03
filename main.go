package main

import (
	// "bufio"
	"fmt"
	"net/http"

	// "go/format"
	"log"
	"time"
	// "go/format"
	// "os"
)

type Download struct {
	//Download URL
	Url string

	//Path to download the file to
	TargetPath string

	//Number of connections to the server
	TotalSections int
}

func main() {
	startTime := time.Now()
	d := Download{
		Url:           "https://www.dropbox.com/s/f63i7s11ydm2cu6/542.Carousel1-potassium-1024w-1366h%402x~ipad.jpg?dl=0",
		TargetPath:    "ipadwallpaper.jpg",
		TotalSections: 10,
	}

	err := d.Do()
	if err != nil {
		log.Fatalf("Error occured while downloading the file: %s\n", err)
	}
	fmt.Printf("Download completed in %v seconds\n", time.Now().Sub(startTime).Seconds())
	fmt.Printf("Welcome to the Download Manager")
}

func (d Download) Do() error {
	fmt.Println("Making connection...")

	r, err := d.getNewRequest("HEAD")
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	fmt.Printf("Status %v", resp.StatusCode)

	return nil
}

//Get a new http request
func (d Download) getNewRequest(method string) (*http.Request, error) {
	r, err := http.NewRequest(
		method,
		d.Url,
		nil,
	)
	if err != nil {
		return nil, err
	}
	r.Header.Set("User-Agent", "Concurrent Download Manager") //key val pair
	return r, nil
}
