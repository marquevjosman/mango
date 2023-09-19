package main

import (
	"fmt"
	"os"
)

var args map[string]string

func isCommand(word string) bool {
	if word[0] == '-' {
		return true
	}
	return false
}

func getArgs() {
	args = make(map[string]string)
	ls := os.Args[1:]
	l := len(ls)
	for i := 0; i < l-1; i++ {
		word := ls[i]
		if isCommand(word) {
			i++
			args[word] = ls[i]
		}
	}
}

func main() {
	fmt.Println("Hello world")
	getArgs()
	printArgs()
}

func printArgs() {
	fmt.Println(args)
}
