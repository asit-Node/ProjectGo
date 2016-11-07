package textFrequency

import "strings"

func GetWordFrequency(text string) (result map[string]int) {
	var wordMap map[string]int
	wordMap = make(map[string]int)
	keySlice := []string{}
	splitText := strings.Split(text, " ")
	for index := range splitText {
		word := strings.ToLower(splitText[index])
		if val, exists := wordMap[word]; exists {
			wordMap[word] = val + 1
		} else {
			wordMap[word] = 1
			keySlice = append(keySlice, strings.ToLower(splitText[index]))
		}
		//fmt.Println(index, splitText[index])
	}

	result = wordMap
	return
}
