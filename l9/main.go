package main

import (
	"bufio"
	"log"
	"net/http"
)

const number_of_buckets = 12

func main() {
	response, error := http.Get("https://www.gutenberg.org/files/2701/old/moby10b.txt")

	if error != nil {
		log.Fatal(error)
	}

	scanner := bufio.NewScanner(response.Body)

	defer response.Body.Close()

	scanner.Split(bufio.ScanWords)

	var buckets = make([][]string, number_of_buckets)

	for scanner.Scan() {
		word := scanner.Text()

		position := HashBucket(word)

		buckets[position] = append(buckets[position], word)
	}
}

func HashBucket(word string) int {
	var sum int

	for _, value := range word {
		sum += int(value)
	}

	return sum % number_of_buckets
}
