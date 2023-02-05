package main

type Adapter struct {
	LegacyStandard *LegacyStandard
}

func (a Adapter) getOutput() string {
	lo := a.LegacyStandard.getLegacyOutput()

	rns := []rune(lo)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}
