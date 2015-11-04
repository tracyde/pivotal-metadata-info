package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "print fields from Pivotal metadata files",
	Long:  "print the fields from Pivotal metadata files downloaded from pivnet",
}

func init() {
	printCmd.Run = print
}

func print(cmd *cobra.Command, args []string) {
	filename := fileName
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var meta PivotalMetadata
	if err := meta.Parse(data); err != nil {
		log.Fatal(err)
	}
	j, err := json.MarshalIndent(meta, "", "\t")
	if err != nil {
		log.Fatal("error:", err)
	}
	fmt.Printf("%+v\n", string(j))
}
