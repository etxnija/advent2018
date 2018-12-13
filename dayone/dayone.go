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
	sa := FrequencyFromFile(args[1])
	ia, err := stringAtoIntA(sa)
	if err != nil {
		panic(err.Error)
	}
	fmt.Println(Frequency(ia))
	fmt.Println("Collibate: ", Callibration(ia))
}

// Frequency is calculated
func Frequency(in []int) int {
	result := 0
	for _, d := range in {
		result += d
	}
	return result
}

func Callibration(in []int) int {
	result := 0
	l := len(in)
	m := make(map[int]int, l)
	m[0] = 1

	for i := 0; i < l*10000; i++ {
		result += in[i%l]
		if m[result] > 0 {
			return result
		}
		m[result] = m[result] + 1

	}
	return 0
}

// FrequencyString is calculated
func FrequencyString(in string) []int {
	ia, err := stringAtoIntA(strings.Split(in, ","))
	if err != nil {
		return []int{}
	}
	return ia
}

func stringAtoIntA(sa []string) ([]int, error) {
	ia := make([]int, len(sa))
	for i := 0; i < len(sa); i++ {
		d, err := strconv.Atoi(strings.TrimSpace(sa[i]))
		if err != nil {
			return []int{}, err
		}
		ia[i] = d
	}
	return ia, nil
}

func FrequencyFromFile(file string) []string {
	sa := []string{}
	fileHandle, _ := os.Open(file)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	for fileScanner.Scan() {
		sa = append(sa, fileScanner.Text())
	}
	return sa
}
