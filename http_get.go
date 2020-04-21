package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// Save user inputted website url into webSite
	var ws string
	fmt.Println("Enter website")
	_, err := fmt.Scan(&ws)

	if err != nil {
		log.Fatal(err)
	}
	// Save user inputted file name into fileName

	var fileName string
	fmt.Println("Enter file to save to")
	_, err = fmt.Scan(&fileName)

	if err != nil {
		log.Fatal(err)
	}

	// search up website with GET method, read out response body and store in
	// variable body.
	resp, _ := http.Get(ws)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println("Source code saved to new file: ", ws)

	// create new file

	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// write source code to file.

	l, err := f.WriteString(string(body))
	if err != nil {
		fmt.Println(err)
		return
	}

	// print info and close file

	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}
