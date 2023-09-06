package main

import (
	"bufio"
	"net/http"
	"os"
)

func GetHeaders(url, protocolType string, file *os.File) {
	resp, err := http.Get(protocolType + "://" + url)
	datawriter := bufio.NewWriter(file)

	if err != nil {
		datawriter.WriteString("URL: " + url + "\n")
		datawriter.WriteString("Failed connection to the website")
	} else {
		datawriter.WriteString("URL: " + url + "\n")

		for key, element := range resp.Header {
			datawriter.WriteString(key + ": " + element[0] + "\n")
		}

	}

	datawriter.Flush()
}
