package main

type LegacyStandard struct {
	legacyOutput string
}

func (ls *LegacyStandard) getLegacyOutput() string {
	return ls.legacyOutput
}
