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
	print(groupLevel int) string
	properties() []PropLine
}

func (f PropFile) properties() []PropLine {
	return f.lines
}
func (f PropFile) sortKeys() PropertyFile {

	props := make([]PropLine, len(f.lines))
	copy(props, f.lines)

	sort.Slice(props, keyComparator(props))
	return PropFile{lines: props}
}
func (f PropFile) duplicatedKeys(p PropertyFile) PropertyFile {
	v := make(map[string]string)
	for _, value := range f.lines {
		v[value.key] = value.value
	}
	for _, sv := range p.sortKeys().properties() {
		if v[sv.key] == sv.value {

		}
	}
	props := make([]PropLine, 0)
	copy(props, f.lines)
	sort.Slice(props, keyComparator(props))
	return PropFile{lines: props}
}
func (l PropLine) levelKey(groupLevel int) string {
	values := strings.Split(l.key, ".")
	values = values[0:groupLevel]

	return strings.Join(values, ".")
}
func (f PropFile) print(groupLevel int) string {
	result := ""
	result += fmt.Sprintf("###########################\n")
	result += fmt.Sprintf("Properties:  \n")
	result += fmt.Sprintf("###########################\n")

	currentGroup := ""
	if len(f.lines) > 0 {
		currentGroup = f.lines[0].levelKey(groupLevel)
	}
	for _, value := range f.lines {

		key := value.levelKey(groupLevel)
		if key != currentGroup {
			result += "\n"
			currentGroup = key
		}
		println(key)
		result += fmt.Sprintf("%s = %s\n", value.key, value.value)
	}
	result += fmt.Sprintf("###########################\n")
	return result
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
	print(firstFile.print(2))

	print(firstFile.sortKeys().print(2))
	//secondFile.print()
}
