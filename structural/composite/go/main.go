package main

func main() {
	file1 := File{
		name: "File1",
	}
	file2 := File{
		name: "File2",
	}
	file3 := File{
		name: "File3",
	}

	folder1 := Folder{
		name: "Folder1",
	}

	folder1.addComponent(&file1)

	folder2 := Folder{
		name: "Folder2",
	}
	folder2.addComponent(&file2)
	folder2.addComponent(&file3)
	folder2.addComponent(&folder1)

	folder2.search("rose")
}

// Example coded according to https://refactoring.guru/design-patterns/composite/go/example#example-0--file-go
