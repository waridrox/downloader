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

	//Splitting the file into sections of equal length
	// Example 50MB =>
	// section 0 = from 0MB to 5MB
	// section 1 = from 6MB to 11MB
	// section 2 = from 12MB to 17MB... till section 10

	//if file size is 100 bytes, section would be: [[0 10] [11 21] [22 32] [33 43] [44 54] [55 65] [66 76] [77 87] [88 98] [99 99]]
	var sections = make([][2]int, d.TotalSections)
	sectionSize := size / d.TotalSections
	fmt.Printf("Each sub section of file: %v bytes\n", sectionSize)

	for i := range sections {
		if i == 0 {
			//starting byte of other sections
			sections[i][0] = 0
		} else {
			//starting byte of remaining sections
			sections[i][0] = sections[i-1][0] + 1 //start of the previous section + 1
		}

		if i < d.TotalSections-1 {
			//ending byte of other sections
			sections[i][1] = sections[i][0] + sectionSize
		} else {
			//ending byte of other sections
			sections[i][1] = size - 1
		}
	}

	fmt.Println(sections)

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
