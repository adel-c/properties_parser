package main

import (
	"testing"
)

func TestReadPropertiesFile(t *testing.T) {

	file := PropFile{lines: []PropLine{
		{key: "key1.sk1.ssk1", value: "f_value1"},
		{key: "key1.sk1.ssk2", value: "f_value2"},
		{key: "key1.sk2.ssk2", value: "f_value3"},
		{key: "key1.sk2.ssk1", value: "f_value4"},
	}}
	s := file.sortKeys().print(2, false)
	expected := `key1.sk1.ssk1 = f_value1
key1.sk1.ssk2 = f_value2

key1.sk2.ssk1 = f_value4
key1.sk2.ssk2 = f_value3
`
	if s != expected {
		t.Fatalf(`generated file
%s
doesn't match expected 
%s`, s, expected)
	}

}
