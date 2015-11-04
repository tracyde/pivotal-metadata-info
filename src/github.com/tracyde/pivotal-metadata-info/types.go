package main

import (
	"gopkg.in/yaml.v2"
)

// Pivotal metadata file
type PivotalMetadata struct {
	Name                      string
	ProductVersion            string       `yaml:"product_version"`
	MetadataVersion           string       `yaml:"metadata_version"`
	TargetInstallerVersion    string       `yaml:"target_installer_version"`
	Stemcell_Criteria         StemcellInfo //`yaml: "stemcell_criteria"` Can not rename this field otherwise will not work (https://github.com/go-yaml/yaml/issues/141)
	Releases                  []ReleaseInfo
	Provides_Product_Versions []ProductVersion //`yaml: "provides_product_versions"` Can not rename this field otherwise will not work (https://github.com/go-yaml/yaml/issues/141)
	Label                     string
	Description               string
}

type StemcellInfo struct {
	OS          string `yaml: "os"`
	RequiresCPI bool   `yaml: "requires_cpi"`
	Version     string `yaml: "version"`
}

type ReleaseInfo struct {
	Name    string
	File    string
	Version string
	SHA1    string
	MD5     string
	URL     string
}

type ProductVersion struct {
	Name    string
	Version string
}

func (c *PivotalMetadata) Parse(data []byte) error {
	if err := yaml.Unmarshal(data, c); err != nil {
		return err
	}
	return nil
}
