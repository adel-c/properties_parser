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

func readPropertiesFile(filePath string) []PropLine {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	properties := make(map[string]string)
	props := make([]PropLine, 1)

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

	return properties
}
func main() {
	fmt.Println("Hello, World!")
	filePath := "first.properties"
	properties := readPropertiesFile(filePath)

	// Print the read properties
	fmt.Println("Properties:")
	for key, value := range properties {
		fmt.Printf("%s = %s\n", key, value)
	}
}
