package main

import (
	"fmt"
	"sort"
)
import (
	"bufio"
	"os"
	"strings"
)

type PropertyFile interface {
	sortKeys() PropertyFile
	duplicatedKeys(p PropertyFile) PropertyFile
	print()
}

func (f PropFile) sortKeys() PropertyFile {

	props := make([]PropLine, len(f.lines))
	copy(props, f.lines)

	sort.Slice(props, keyComparator(props))
	return PropFile{lines: props}
}
func (f PropFile) duplicatedKeys(p PropertyFile) PropertyFile {

	props := make([]PropLine, 0)
	copy(props, f.lines)
	sort.Slice(props, keyComparator(props))
	return PropFile{lines: props}
}
func (f PropFile) print() {
	fmt.Println("###########################")
	fmt.Println("Properties:  ")
	fmt.Println("###########################")
	for index, value := range f.lines {
		fmt.Printf("%d %s = %s\n", index, value.key, value.value)
	}
	fmt.Println("###########################")
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
		return array[i].key < array[j].key
	}
}

func main() {
	filePath := "first.properties"
	firstFile := readPropertiesFile(filePath)
	//filePath2 := "sec.properties"
	//secondFile := readPropertiesFile(filePath2)
	// Print the read firstFile
	firstFile.print()
	firstFile.sortKeys().print()
	//secondFile.print()
}
