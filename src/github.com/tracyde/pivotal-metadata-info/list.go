package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list collection of resources from Pivotal metadata",
	Long:  "list a collection of resources from Pivotal metadata files downloaded from pivnet",
	Run:   nil,
}

var listStemcellsCmd = &cobra.Command{
	Use:   "stemcells",
	Short: "list all stemcells",
	Long:  "list all stemcells referenced in Pivotal metadata",
	Run:   listStemcells,
}

var displayUrl bool

func init() {
	listCmd.AddCommand(listStemcellsCmd)

	listStemcellsCmd.Flags().BoolVarP(&displayUrl, "url", "u", false, "Display URL used to download stemcell")
}

func listStemcells(cmd *cobra.Command, args []string) {
	filename := fileName
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var meta PivotalMetadata
	if err := meta.Parse(data); err != nil {
		log.Fatal(err)
	}
	v := meta.Stemcell_Criteria.Version

	if displayUrl {
		// fmt.Printf("https://bosh.io/d/stemcells/bosh-aws-xen-ubuntu-trusty-go_agent?v=%v\n", v) // bosh.io has no param to request full stemcells
		fmt.Printf("https://d26ekeud912fhb.cloudfront.net/bosh-stemcell/aws/bosh-stemcell-%v-aws-xen-ubuntu-trusty-go_agent.tgz\n", v)
	} else {
		fmt.Printf("bosh-stemcell-%v-aws-xen-ubuntu-trusty-go_agent.tgz\n", v)
	}
}
