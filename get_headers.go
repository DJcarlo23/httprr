package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func GetHeaders(domain, protocolType string, file *os.File) {
	resp, err := http.Get(protocolType + "://" + domain)
	fmt.Println(domain)
	datawriter := bufio.NewWriter(file)

	if err != nil {
		datawriter.WriteString("Domain: " + domain + "\n")
		datawriter.WriteString("Failed connection to the website" + "\n")
	} else {
		datawriter.WriteString("Domain: " + domain + "\n")

		for key, element := range resp.Header {
			datawriter.WriteString(key + ": " + element[0] + "\n")
		}

	}

	datawriter.WriteString("\n")

	datawriter.Flush()
}
