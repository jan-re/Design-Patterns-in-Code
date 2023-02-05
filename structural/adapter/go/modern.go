package main

type ModernStandard struct {
	output string
}

func (ms *ModernStandard) getOutput() string {
	ms.output = "Output format produced by the modern code standard."

	return ms.output
}
