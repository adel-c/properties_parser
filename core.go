package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type PropertyFile interface {
	keepOverLoadedKeys(child PropertyFile) PropertyFile
	print(groupLevel int, displayHeaders bool) string
	properties() []PropLine
}

func (f PropFile) properties() []PropLine {
	return f.lines
}

func (f PropFile) keepOverLoadedKeys(child PropertyFile) PropertyFile {
	v := make(map[string]string)
	props := make([]PropLine, 0)
	for _, value := range f.properties() {
		v[value.key] = value.value
	}
	for _, sv := range child.properties() {
		if v[sv.key] != sv.value {
			props = append(props, sv)
		}
	}

	return PropFile{lines: props}
}
func (l PropLine) levelKey(groupLevel int) string {
	values := strings.Split(l.key, ".")
	values = values[0:min(groupLevel, len(values))]

	return strings.Join(values, ".")
}
func (f PropFile) print(groupLevel int, displayHeader bool) string {
	result := ""

	currentGroup := ""
	if len(f.lines) > 0 && !displayHeader {
		currentGroup = f.lines[0].levelKey(groupLevel)
	}
	for _, value := range f.lines {

		key := value.levelKey(groupLevel)
		if key != currentGroup {
			if displayHeader {
				result += fmt.Sprintf("####################\n")
				result += fmt.Sprintf("### %s\n", key)
				result += fmt.Sprintf("####################\n")
			}
			result += "\n"
			currentGroup = key
		}
		result += fmt.Sprintf("%s = %s\n", value.key, value.value)
	}

	return result
}

type PropFile struct {
	lines []PropLine
}
type PropLine struct {
	key   string
	value string
}

func ReadPropertiesFile(filePath string) PropFile {
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
	sort.Slice(props, keyComparator(props))
	return PropFile{lines: props}
}
func keyComparator(array []PropLine) func(i, j int) bool {
	return func(i, j int) bool {
		return array[i].key < array[j].key
	}
}
