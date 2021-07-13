package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	jsoniter "github.com/json-iterator/go"
)

// PanelCoverage contains all gene coverages
type PanelCoverage struct {
	Identifier string
	Genes      map[string]Coverage
}

// Coverage data
type Coverage struct {
	ENSG               string
	Symbol             string
	TotalTargetedBases int
	BasesCovered5x     int
	BasesCovered10x    int
	BasesCovered20x    int
	BasesCovered30x    int
}

// outputToJSON outputs coverage to JSON file
func outputToJSON(sample string, geneCoverageMap map[string]Coverage) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	var b bytes.Buffer
	w := gzip.NewWriter(&b)

	outputFilePath := path.Join(os.Getenv("HOME"), "coverage", (fmt.Sprintf("%s_coverage.json.gz", sample)))

	export := PanelCoverage{
		Identifier: sample,
		Genes:      geneCoverageMap,
	}

	toWriteJSON, err := json.Marshal(export)
	if err != nil {
		log.Fatalf("failed convert struct FileRepresentation of sample %s into json; got %v", sample, err)
	}

	w.Write(toWriteJSON)
	w.Close() // You must close this first to flush the bytes to the buffer.

	ioutil.WriteFile(outputFilePath, b.Bytes(), 0644)
}
