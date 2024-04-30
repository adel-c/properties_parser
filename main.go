package main

import (
	"fmt"
)

func main() {
	filePath := "first.properties"
	firstFile := ReadPropertiesFile(filePath)
	//filePath2 := "sec.properties"
	//secondFile := readPropertiesFile(filePath2)
	// Print the read firstFile

	sortedFile := firstFile.sortKeys()

	fmt.Println("###########################")
	fmt.Println("Properties:  ")
	fmt.Println("###########################")
	print(sortedFile.print(2, true))
	fmt.Println("###########################")
	//secondFile.print()
	fmt.Println("###########################")
	fmt.Println("Properties:  ")
	fmt.Println("###########################")
	print(sortedFile.print(2, false))
	fmt.Println("###########################")
}
