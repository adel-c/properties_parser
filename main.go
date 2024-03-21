package main

import "fmt"
import (
	"bufio"
	"os"
	"strings"
)

type PropertyFile interface {
	sortKeys() PropertyFile
	duplicatedKeys(p PropertyFile) PropertyFile
}
type PropFile struct {
	lines []PropLine
}
type PropLine struct {
	key   string
	value string
}

func readPropertiesFile(filePath string) PropFile {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	properties := make(map[string]string)
	props := make([]PropLine, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Index(strings.TrimSpace(line), "#") == 0 {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			properties[key] = value
			l := PropLine{key: key, value: value}
			props = append(props, l)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return PropFile{lines: props}
}
func keyComparator(array []PropLine) func(i, j int) bool {
	return func(i, j int) bool {
		return array[i].key > array[j].key
	}
}

func main() {
	fmt.Println("Hello, World!")
	filePath := "first.properties"
	firstFile := readPropertiesFile(filePath)
	filePath2 := "sec.properties"
	secondFile := readPropertiesFile(filePath2)
	// Print the read firstFile
	fmt.Println("Properties:")
	for index, value := range firstFile.lines {
		fmt.Printf("%d %s = %s\n", index, value.key, value.value)
	}

	for index, value := range secondFile.lines {
		fmt.Printf("%d %s = %s\n", index, value.key, value.value)
	}
}
