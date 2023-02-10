package main

import (
	"fmt"
)

type SkinFactory struct {
	skinMap map[string]Skin
}

func (d *SkinFactory) getSkinByType(skinType string) (Skin, error) {
	if d.skinMap[skinType] != nil {
		return d.skinMap[skinType], nil
	} else {
		return nil, fmt.Errorf("Wrong skinType passed. Wrong value: %s\n", skinType)
	}
}

func getSkinFactorySingleInstance() *SkinFactory {
	return &skinFactorySingleInstance
}

var (
	skinFactorySingleInstance = SkinFactory{
		skinMap: map[string]Skin{
			"tSkin":  newTSkin(),
			"ctSkin": newCTSkin(),
		}}
)
