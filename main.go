package main

import "fmt"
import (
	"bufio"
	"os"
	"strings"
)

func readPropertiesFile(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	properties := make(map[string]string)

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
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return properties, nil
}
func main() {
	fmt.Println("Hello, World!")
	filePath := "first.properties"
	properties, err := readPropertiesFile(filePath)
	if err != nil {
		fmt.Println("Error reading properties file:", err)
		return
	}

	// Print the read properties
	fmt.Println("Properties:")
	for key, value := range properties {
		fmt.Printf("%s = %s\n", key, value)
	}
}
