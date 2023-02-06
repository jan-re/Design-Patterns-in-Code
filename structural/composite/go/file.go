package main

import "fmt"

type File struct {
	name string
}

func (f *File) search(s string) {
	fmt.Printf("Searching file %s for string %s...\n", f.name, s)
}
