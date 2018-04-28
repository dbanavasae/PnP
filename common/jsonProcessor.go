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

func GetConfigFromJson(file string, configStruct interface{}) (err error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, configStruct); err != nil {
		return err
	}

	return nil
}