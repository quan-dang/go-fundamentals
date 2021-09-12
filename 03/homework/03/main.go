package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// print current file or current position (.)
func printListing(entry string, depth int) {
	indent := strings.Repeat("|   ", depth)
	fmt.Printf("%s|-- %s\n", indent, entry)
}

// print directory
func printDirectory(path string, depth int) {
	// retrieve all entries in current path (folder or file)
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("error reading %s: %s\n", path, err.Error())
		return
	}

	printListing(path, depth)
	for _, entry := range entries {
		// check if entry is a symbolic link
		if (entry.Mode() & os.ModeSymlink) == os.ModeSymlink {
			// retrieve the original path
			full_path, err := os.Readlink(filepath.Join(path, entry.Name()))
			if err != nil {
				fmt.Printf("error reading link: %s\n", err.Error())
			} else {
				printListing(entry.Name()+" -> "+full_path, depth+1)
			}
		} else if entry.IsDir() {
			// call recursively if this entry is a directory
			printDirectory(filepath.Join(path, entry.Name()), depth+1)
		} else {
			// else print list
			printListing(entry.Name(), depth+1)
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run main.go <path>")
		os.Exit(1)
	}

	printDirectory(os.Args[1], 0)
}
