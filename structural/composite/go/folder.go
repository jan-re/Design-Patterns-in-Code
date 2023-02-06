package main

import "fmt"

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) search(s string) {
	fmt.Printf("Searching folder %s for string %s...\n", f.name, s)
	for _, f := range f.components {
		f.search(s)
	}
}

func (f *Folder) addComponent(c Component) {
	f.components = append(f.components, c)
}
