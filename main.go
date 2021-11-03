package main

import (
	// "bufio"
	"fmt"

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
		Url:           "https://drive.google.com/uc?export=download&id=1ShtjiHZWFeH--gJyMhfYBvklEG4Auu2S",
		TargetPath:    "rdfinale.pdf",
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
	return nil
}
