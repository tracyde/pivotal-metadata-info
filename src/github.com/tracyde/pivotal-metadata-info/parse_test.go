package main

import "testing"

var data = `
---
name: 'p-concourse'
product_version: '0.62.0'
metadata_version: "1.5"
target_installer_version: 1.4.0.0
stemcell_criteria:
  os: ubuntu-trusty
  requires_cpi: false
  version: '2989'
releases:
- name: 'concourse'
  file: 'concourse-0.62.0.tgz'
  version: '0.62.0'
  sha1: 0d38e6fadf323888477aec42844749a8d04a6c77
  md5: 163d38e205a2215eff80bf2dbb789d05
  url: https://bosh.io/d/github.com/cloudfoundry-incubator/garden-linux-release?v=0.308.0
- name: 'garden-linux'
  file: 'release.tgz'
  version: '0.303.0'

provides_product_versions:
  - name:    'p-concourse'
    version: '0.62.0'

label:       Concourse
description: CI that scales with your project
`

func ensureEquals(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected '%v', got %v", b, a)
	}
}

func TestParse(t *testing.T) {
	var meta PivotalMetadata
	if err := meta.Parse([]byte(data)); err != nil {
		t.Error("Failed to parse")
	}
	ensureEquals(t, meta.Name, "p-concourse")
	ensureEquals(t, meta.ProductVersion, "0.62.0")
	ensureEquals(t, meta.MetadataVersion, "1.5")
	ensureEquals(t, meta.TargetInstallerVersion, "1.4.0.0")
	ensureEquals(t, meta.Stemcell_Criteria.OS, "ubuntu-trusty")
	ensureEquals(t, meta.Stemcell_Criteria.RequiresCPI, false)
	ensureEquals(t, meta.Stemcell_Criteria.Version, "2989")
	ensureEquals(t, len(meta.Releases), 2)
	ensureEquals(t, meta.Releases[0].SHA1, "0d38e6fadf323888477aec42844749a8d04a6c77")
	ensureEquals(t, meta.Releases[0].MD5, "163d38e205a2215eff80bf2dbb789d05")
	ensureEquals(t, meta.Releases[0].URL, "https://bosh.io/d/github.com/cloudfoundry-incubator/garden-linux-release?v=0.308.0")
	ensureEquals(t, meta.Releases[1].Name, "garden-linux")
	ensureEquals(t, meta.Releases[1].File, "release.tgz")
	ensureEquals(t, meta.Releases[1].Version, "0.303.0")
	ensureEquals(t, len(meta.Provides_Product_Versions), 1)
	ensureEquals(t, meta.Provides_Product_Versions[0].Name, "p-concourse")
	ensureEquals(t, meta.Provides_Product_Versions[0].Version, "0.62.0")
	ensureEquals(t, meta.Label, "Concourse")
	ensureEquals(t, meta.Description, "CI that scales with your project")
}
