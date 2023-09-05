package main

import (
	"bufio"
	"errors"
	"log"
	"net/http"
	"os"
)

func GetHeaders(url string, file *os.File) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(errors.New("failed connection to the site"))
	}

	datawriter := bufio.NewWriter(file)
	datawriter.WriteString("URL: " + url + "\n")

	for key, element := range resp.Header {
		datawriter.WriteString(key + ": " + element[0] + "\n")
	}

	datawriter.Flush()
}
