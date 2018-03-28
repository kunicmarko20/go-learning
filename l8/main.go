package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	response, error := http.Get("https://www.gutenberg.org/files/2701/old/moby10b.txt")

	if error != nil {
		log.Fatal(error)
	}

	scanner := bufio.NewScanner(response.Body)

	defer response.Body.Close()

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
