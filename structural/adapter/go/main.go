package main

import "fmt"

func main() {
	ms := ModernStandard{}
	fmt.Println(ms.getOutput())

	ls := LegacyStandard{}
	fmt.Printf("Legacy output not understandable: %s\n", ls.getLegacyOutput())

	als := Adapter{
		LegacyStandard: &ls,
	}
	fmt.Printf("Legacy output translated: %s\n", als.getOutput())
}
