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
	singleDomain := flag.String("d", "single", "a domain of target website")
	listDomain := flag.String("df", "multiple", "a list of domains of target websites")
	protocolType := flag.String("p", "protocol", "protocol which will be used to connect to the website")
	numDomainsThread := flag.Int("t", 10, "number of domains handled by a single thread")
	flag.Parse()

	dFlagStatus := isFlagPassed("d")
	dfFlagStatus := isFlagPassed("df")
	pFlagStatus := isFlagPassed("p")

	if !pFlagStatus {
		*protocolType = "https"
	}

	var file *os.File
	var errFile error

	if dFlagStatus || dfFlagStatus {
		now := time.Now()
		fileName := "httprr_" + now.Format("20060102150405") + ".txt"
		file, errFile = os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		if errFile != nil {
			log.Fatalf("Failed creating file: %s", errFile)
		}
	} else {
		log.Fatalf("Domain not provided")
	}

	if dFlagStatus {
		domain := []string{*singleDomain}
		GetHeadersMultithreading(domain, *protocolType, file, 1)
	}

	if dfFlagStatus {
		readListDomains, err := os.Open(*listDomain)

		if err != nil {
			log.Fatalf("Failed reading file with domains: %s", err)
		}

		fileScanner := bufio.NewScanner(readListDomains)
		fileScanner.Split(bufio.ScanLines)

		domains := make([]string, 0)

		for fileScanner.Scan() {
			domains = append(domains, fileScanner.Text())
		}

		GetHeadersMultithreading(domains, *protocolType, file, *numDomainsThread)
		readListDomains.Close()
	}

	file.Close()

}
