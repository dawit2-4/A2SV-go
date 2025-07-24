package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func WordFrequency(text string) map[string]int {
	text = strings.ToLower(text)
	reg, _ := regexp.Compile(`[^\w\s]`)
	text = reg.ReplaceAllString(text, "")
	words := strings.Fields(text)

	freq := make(map[string]int)
	for _, word := range words {
		freq[word]++
	}
	return freq
}

func main() {
	fmt.Println("Enter text to analyze (press Enter when done):")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	result := WordFrequency(input)

	fmt.Println("Word Frequencies:")
	for word, count := range result {
		fmt.Printf("%s: %d\n", word, count)
	}
}
