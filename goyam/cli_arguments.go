package goyam

import "os"

func arguments() []string {
	return os.Args[1:]
}

func firstFilepath() string {
	return arguments()[0]
}

func secondFilepath() string {
	return arguments()[1]
}
