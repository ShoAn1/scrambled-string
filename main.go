package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

func checkSubstring(checkDict, windowString []string) int {
	sort.Strings(windowString)
	sort.Strings(checkDict)
	if strings.Join(windowString, "") == strings.Join(checkDict, "") {
		return 1
	}
	return 0
}

func subStringsCheck(str1 string, dict []string) int {
	count := 0
	m2 := make(map[string]map[string]bool)
	for _, j := range dict {
		window := len(j)
		for i := 0; i < len(str1)-window; i++ {
			if j[0] == str1[i] && string(str1[i+window-1]) == string(j[len(j)-1]) {
				if _, ok := m2[j][str1[i:i+window-1]]; ok {
					continue
				} else {
					count = count + checkSubstring(strings.Split(j, ""), strings.Split(str1[i:i+window], ""))
					if m2[j] == nil {
						m2[j] = make(map[string]bool)
					}
					m2[j][str1[i:i+window-1]] = true
					i += window - 1
				}

			}
		}
	}
	return count
}
func GetListFromFile(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("failed to open")

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text
}

func main() {
	dictPath := flag.String("dictionary", "", "Path to Dictionary File")
	stringPath := flag.String("input", "", "Path to list of Grambled String")
	flag.Parse()
	dict := GetListFromFile(string(*dictPath))
	check_string := GetListFromFile(string(*stringPath))
	// fmt.Println("dictionary and inputfile:", dict, check_string)
	for i, j := range check_string {
		num_of_words_appear := subStringsCheck(j, dict)
		fmt.Printf("Case #%v : %v \n", i+1, num_of_words_appear)
	}
}
