package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
)

func main() {
	directory := "./words.txt"
	countAndSortAZWordsFromAFileData(directory)
}
func countAndSortAZWordsFromAFileData(directory string) {
	dataFile, eventError := ioutil.ReadFile(directory)
	if eventError != nil {
		log.Printf("Error al abrir el archivo: %v", eventError)
		return
	}
	fileContent := strings.ToLower(string(dataFile))
	wordList := strings.Fields((regexp.MustCompile("[^a-z]+")).ReplaceAllString(fileContent, " "))
	wordFrequencyMap := map[string]int32{}
	for _, word := range wordList {
		wordFrequencyMap[word] += 1
	}
	wordAZList := make([]string, 0, len(wordFrequencyMap))
	for wordAZ := range wordFrequencyMap {
		wordAZList = append(wordAZList, wordAZ)
	}
	sort.Strings(wordAZList)
	for _, wordAZ := range wordAZList {
		fmt.Println(wordAZ, wordFrequencyMap[wordAZ])
	}
}
