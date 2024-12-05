package common

import (
	"bufio"
	"os"
	"strings"
)

func GetLines() []string {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		panic(err)
	}

	return lines
}

func GetInput() string {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	return string(data)
}

func Split(input, sep string) (left, right string) {
	split := strings.SplitN(input, sep, 2)
	return split[0], split[1]
}
