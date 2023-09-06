package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"time"
)

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func main() {
	singleURL := flag.String("u", "single url", "a url of the target website")
	listURL := flag.String("uf", "multiple urls", "a list of urls of target websites")
	flag.Parse()

	uFlagStatus := isFlagPassed("u")
	ufFlagStatus := isFlagPassed("uf")

	var file *os.File
	var errFile error

	if uFlagStatus || ufFlagStatus {
		now := time.Now()
		fileName := "httprr_" + now.Format("20060102150405") + ".txt"
		file, errFile = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if errFile != nil {
			log.Fatalf("Failed creating file: %s", errFile)
		}
	} else {
		log.Fatalf("URL not provided")
	}

	if uFlagStatus {
		GetHeaders(*singleURL, file)
	}

	if ufFlagStatus {
		readListURL, err := os.Open(*listURL)

		if err != nil {
			log.Fatalf("Failed reading file with urls: %s", err)
		}

		fileScanner := bufio.NewScanner(readListURL)
		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			GetHeaders(fileScanner.Text(), file)
		}

		readListURL.Close()
	}

	file.Close()

}
