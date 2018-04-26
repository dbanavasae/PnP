package common

import (
	"io/ioutil"
	"encoding/json"
)

type Package struct {
	Name                  string        `json:"name"`
	Version               string        `json:"version"`
	CheckInstalledCmd     []string      `json:"checkInstalledCmd"`
	CheckInstalledVersion []string      `json:"checkInstalledVersion"`
	UnInstallInstructions []string      `json:"unInstallInstructions"`
	InstallFromFile       string        `json:"installFromFile"`
	InstallInstructions   []string      `json:"installInstructions"`
	UpdateRepo            []string      `json:"updateRepo"`
}

type PackageInfo struct {
	Packages []Package `json:"packages"`
}

func FromFile(file string) (*PackageInfo, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	p := PackageInfo{}
	if err = json.Unmarshal(b, &p); err != nil {
		return nil, err
	}
	return &p, nil
}