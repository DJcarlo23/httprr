package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup

func GetHeaders(domains []string, protocolType string, file *os.File) {
	defer wg.Done()
	for i := 0; i < len(domains); i++ {
		resp, err := http.Get(protocolType + "://" + domains[i])
		fmt.Println(domains[i])
		datawriter := bufio.NewWriter(file)

		if err != nil {
			datawriter.WriteString("Domain: " + domains[i] + "\n")
			datawriter.WriteString("Failed connection to the website" + "\n")
		} else {
			datawriter.WriteString("Domain: " + domains[i] + "\n")

			for key, element := range resp.Header {
				datawriter.WriteString(key + ": " + element[0] + "\n")
			}
		}

		datawriter.WriteString("\n")
		datawriter.Flush()
	}
}

func GetHeadersMultithreading(domains []string, protocolType string, file *os.File, batchSize int) {
	for i := 0; i < len(domains); i += batchSize {
		wg.Add(1)
		end := i + batchSize
		if end > len(domains) {
			end = len(domains)
		}
		go GetHeaders(domains[i:end], protocolType, file)
	}
	wg.Wait()
}
