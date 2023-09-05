package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
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
	ulrFlag := flag.String("u", "url", "a url of the target website")
	flag.Parse()

	if flagStatus := isFlagPassed("u"); flagStatus == true {
		resp, err := http.Get(*ulrFlag)

		if err != nil {
			fmt.Print(errors.New("Invalid request"))
		}

		for key, element := range resp.Header {
			fmt.Printf("%s: %v \n", key, element[0])
		}

	}

}
