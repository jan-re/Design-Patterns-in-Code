package main

import "fmt"

func main() {
	ms := ModernStandard{
		output: "Output format produced by the modern code standard.",
	}
	fmt.Println(ms.getOutput())

	ls := LegacyStandard{
		legacyOutput: ".dradnatSycageL eht fo roivaheb laicepS",
	}
	fmt.Printf("Legacy output not understandable: %s\n", ls.getLegacyOutput())

	als := Adapter{
		LegacyStandard: &ls,
	}
	fmt.Printf("Legacy output translated: %s\n", als.getOutput())
}
