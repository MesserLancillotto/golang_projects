package main

import (
	"os"
	"fmt"
	"bytes"
)

type function func(data []byte) int

func countCharacters(data []byte) int {
	return len(data)
}

func countLines(data []byte) int {
	return bytes.Count(data, []byte("\n"))
}

func countWords(data []byte) int {
	return len(bytes.Fields(data))
}

func main() {
	var ops = map[string]function{
		"-c" : countCharacters,
		"-w" : countWords,
		"-l" : countLines,
	}

	var functionStack []function = []function{}
	var fileStack []string = []string{}
	var totalCount []int = []int{}

	for _, element := range os.Args[1:] {
		funct, exists := ops[element]
		if !exists {
			_, err := os.Stat(element)
			if err == nil {
				fileStack = append(fileStack, element)
			} else {
				fmt.Fprintf(os.Stderr, "Error: \"%s\" does not exist\n", element)
			}
		} else {
			functionStack = append(functionStack, funct)
			totalCount = append(totalCount, 0)
		}
	}
	if len(functionStack) == 0 {
		functionStack = append(functionStack, countLines)
	}
	for _, filename := range fileStack {
		data, error := os.ReadFile(filename)
		if error != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", error)
		} 
		for index, fun := range functionStack {
			var count int = fun(data) 
			totalCount[index] += count
			fmt.Print(count)
			fmt.Print("\t")

		}
		fmt.Println(filename)
	}

	if len(fileStack) == 1 {
		return
	}
	for _, e := range totalCount {
		fmt.Print(e)
		fmt.Print("\t")
	}
	fmt.Println("total")
}