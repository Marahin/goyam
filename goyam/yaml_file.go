package goyam

import (
	"io/ioutil"
)

/*
LoadFile loads a file and returns content in form of a string
*/
func LoadFile(filepath string) (string, error) {
	data, err := ioutil.ReadFile(filepath) // just pass the file name
	if err != nil {
		return "", err
	}
	return string(data), nil
}

/*
LoadYAMLFile calls LoadFile and then ParseYAML
functions (as you could have guessed).
*/
func LoadYAMLFile(filepath string) (YAMLMap, error) {
	data, err := LoadFile(filepath)
	if err != nil {
		return nil, err
	}

	yamlData, err := ParseYAML(data)
	return yamlData, err
}
