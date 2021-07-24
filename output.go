package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"

	jsoniter "github.com/json-iterator/go"
)

// PanelCoverage contains all gene coverages
type PanelCoverage struct {
	Identifier            string
	GitVersion            string
	GlobalTotalBases      int
	GlobalCoveredBases5x  int
	GlobalCoveredBases10x int
	GlobalCoveredBases20x int
	GlobalCoveredBases30x int
	PerGeneCoverage       map[string]Coverage
}

// Coverage data
type Coverage struct {
	ENSG            string
	Symbol          string
	TotalBases      int // total targeted bases
	CoveredBases    int // 	copy of BasesCovered10x for backwards compatibility
	BasesCovered5x  int
	BasesCovered10x int
	BasesCovered20x int
	BasesCovered30x int
}

// outputToJSON outputs coverage to JSON file
func outputToJSON(sample string, geneCoverageMap map[string]Coverage, globalTotalBases, globalCoveredBases5x, globalCoveredBases10x, globalCoveredBases20x, globalCoveredBases30x int) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	var b bytes.Buffer
	w := gzip.NewWriter(&b)

	outputFilePath := fmt.Sprintf("%s_coverage.json.gz", sample)

	export := PanelCoverage{
		Identifier:            sample,
		PerGeneCoverage:       geneCoverageMap,
		GlobalTotalBases:      globalTotalBases,
		GlobalCoveredBases5x:  globalCoveredBases5x,
		GlobalCoveredBases10x: globalCoveredBases10x,
		GlobalCoveredBases20x: globalCoveredBases20x,
		GlobalCoveredBases30x: globalCoveredBases30x,
	}

	toWriteJSON, err := json.Marshal(export)
	if err != nil {
		log.Fatalf("failed convert struct FileRepresentation of sample %s into json; got %v", sample, err)
	}

	w.Write(toWriteJSON)
	w.Close() // You must close this first to flush the bytes to the buffer.

	err = ioutil.WriteFile(outputFilePath, b.Bytes(), 0644)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
