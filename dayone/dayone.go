package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("HEllo")
	args := os.Args
	fmt.Println("input: ", args[1])
	fmt.Println(FrequencyFromFile(args[1]))
}

// Frequency is calculated
func Frequency(in string) int {
	sa := strings.Split(in, ",")
	result := 0
	for _, s := range sa {
		i, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			panic("not at int")
		}
		result += i
	}
	return result
}

func FrequencyFromFile(file string) int {
	result := 0

	fileHandle, _ := os.Open(file)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {
		i, err := strconv.Atoi(strings.TrimSpace(fileScanner.Text()))
		if err != nil {
			panic("not at int")
		}
		result += i
	}
	return result
}
