package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
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
		Url:           "https://www.dropbox.com/s/f63i7s11ydm2cu6/542.Carousel1-potassium-1024w-1366h%402x~ipad.jpg?dl=1",
		TargetPath:    "wallpaper.jpg",
		TotalSections: 10,
	}

	err := d.Do()
	if err != nil {
		log.Fatalf("Error occured while downloading the file: %s\n", err)
	}
	fmt.Printf("Download completed in %v seconds\n", time.Now().Sub(startTime).Seconds())
	fmt.Printf("Welcome to the Download Manager\n")
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
	fmt.Printf("Status %v\n", resp.StatusCode)

	if resp.StatusCode > 299 {
		return errors.New(fmt.Sprintf("Can't process! Response status is: %v", resp.StatusCode))
	}

	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		return err
	}
	fmt.Printf("File size: %v bytes\n", size)
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
