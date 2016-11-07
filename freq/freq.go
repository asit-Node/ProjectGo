package main

import (
	"bufio"
	"fmt"
	"numberFrequency"
	"os"

	"./textFrequency"
)

func main() {

	var input string
	fmt.Println("Enter a number")
	fmt.Scanf("%s", &input)
	result := numberFrequency.GetNumberFrequency(string(input))
	for contentIndex := 0; contentIndex < len(result); contentIndex++ {
		fmt.Print(result[contentIndex], " ")
	}
	fmt.Println()
	fmt.Println("Enter a sentence")
	bio := bufio.NewReader(os.Stdin)
	line, hasMoreInLine, err := bio.ReadLine()
	//fmt.Println(line)
	if err == nil && hasMoreInLine == false {
		wordMap := textFrequency.GetWordFrequency(string(line))
		for key, value := range wordMap {
			fmt.Println(key, ": ", value)
		}
	}
}
