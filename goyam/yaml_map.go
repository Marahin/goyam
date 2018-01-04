package goyam

import (
	"log"
	"reflect"
	"sort"
	"strconv"

	"github.com/sergi/go-diff/diffmatchpatch"
	yaml "gopkg.in/yaml.v2"
)

/*
YAMLMap is just a generic parsed YAML structure
It can contain any form of YAML structure
*/
type YAMLMap = map[interface{}]interface{}

/*
ParseYAML takes stringified data and unmarshals it onto YAMLMap
(which then it returns, unless error happens).
*/
func ParseYAML(stringifiedData string) (YAMLMap, error) {
	m := YAMLMap{}
	err := yaml.Unmarshal([]byte(stringifiedData), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}
	return m, nil
}

/*
GrabTopLevelKeys returns array of stringified YAML keys
*/
func GrabTopLevelKeys(ym YAMLMap) []string {
	keys := []string{}
	for key := range ym {
		if str, isString := key.(string); isString {
			keys = append(keys, str)
		}
	}

	return keys
}

/*
stringInSlice returns a boolean value indicating
whether the passed target string can be found
in passed []string slice (using binary search)
*/
func stringInSlice(target string, slice []string) bool {
	sort.Strings(slice)
	i := sort.Search(len(slice),
		func(i int) bool { return slice[i] >= target })
	if i < len(slice) && slice[i] == target {
		return true
	}
	return false
}

// YAMLFindMutualKeys does the opposite of YAMLFindMissingKeys
func YAMLFindMutualKeys(first YAMLMap, second YAMLMap) []string {
	keys := []string{}
	firstKeys := GrabTopLevelKeys(first)
	secondKeys := GrabTopLevelKeys(second)

	for _, key := range firstKeys {
		if stringInSlice(key, secondKeys) {
			keys = append(keys, key)
		}
	}
	return keys
}

// YAMLFindMissingKeys returns keys that can be found
// in first, but not in the second key array
func YAMLFindMissingKeys(first YAMLMap, second YAMLMap) []string {
	firstKeys := GrabTopLevelKeys(first)
	secondKeys := GrabTopLevelKeys(second)
	missing := []string{}
	for _, key := range firstKeys {
		if !(stringInSlice(key, secondKeys)) {
			missing = append(missing, key)
		}
	}

	return missing
}

/*
YAMLCompare will print out every differing or missing key
between two input files. It also handles counting warnings and
errors.
*/
func YAMLCompare(first YAMLMap, second YAMLMap, prefix string) Summary {

	summary := NewSummary([]string{}, []string{})

	// Find missing keys
	for _, missingKey := range YAMLFindMissingKeys(first, second) {
		summary.AddError(Green(prefix) + " Key `" + Cyan(missingKey) + "` can be found in `" + Magenta(firstFilepath()) + "` but not in `" + Magenta(secondFilepath()) + "`.")
	}
	for _, missingKey := range YAMLFindMissingKeys(second, first) {
		summary.AddError(Green(prefix) + " Key `" + Cyan(missingKey) + "` can be found in `" + Magenta(secondFilepath()) + "` but not in `" + Magenta(firstFilepath()) + "`.")
	}

	// Compare types of values for mutual keys
	keys := YAMLFindMutualKeys(first, second)
	for _, key := range keys {
		dmp := diffmatchpatch.New()
		v1 := first[key]
		v2 := second[key]
		// Check if both sides have the key value set to string
		if reflect.TypeOf(v1) == reflect.TypeOf(v2) {
			switch v1.(type) {
			case string:
				if v1.(string) != v2.(string) {
					str1 := v1.(string)
					str2 := v2.(string)
					diffs := dmp.DiffMain(str1, str2, false)
					summary.AddWarning(Green(prefix) + " Key `" + Cyan(key) + "` (type `" + reflect.TypeOf(v1).String() + "`) has a different value (" + str1 + ") in `" + Magenta(firstFilepath()) + "` than `" + Magenta(secondFilepath()) + "` (" + str2 + "). [" + dmp.DiffPrettyText(diffs) + "]")
				}
			case int:
				if v1.(int) != v2.(int) {
					str1 := strconv.Itoa(v1.(int))
					str2 := strconv.Itoa(v2.(int))
					diffs := dmp.DiffMain(str1, str2, false)
					summary.AddWarning(Green(prefix) + " Key `" + Cyan(key) + "` (type `" + reflect.TypeOf(v1).String() + "`) has a different value (" + str1 + ") in `" + Magenta(firstFilepath()) + "` than `" + Magenta(secondFilepath()) + "` (" + str2 + "). [" + dmp.DiffPrettyText(diffs) + "]")
				}
			case YAMLMap:
				summary.Merge(YAMLCompare(v1.(YAMLMap), v2.(YAMLMap), prefix+" ("+key+") >"))
			}
		} else {
			summary.AddWarning(Green(prefix) + " Key `" + Cyan(key) + "` has a different type (" + reflect.TypeOf(v1).String() + ") in `" + Magenta(firstFilepath()) + "` than `" + Magenta(secondFilepath()) + "` (" + reflect.TypeOf(v2).String() + ").")
		}
	}

	return *summary
}
