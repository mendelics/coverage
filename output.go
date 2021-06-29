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

// outputToJSON outputs coverage to JSON file
func outputToJSON(sampleNames []string, geneCoverageMap map[string]Coverage) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	for _, identifier := range sampleNames {
		var b bytes.Buffer
		w := gzip.NewWriter(&b)

		outputFilePath := path.Join(os.Getenv("HOME"), "coverage", (fmt.Sprintf("%s_coverage.json.gz", identifier)))

		export := PanelCoverage{
			Identifier: identifier,
			Genes:      geneCoverageMap,
		}

		toWriteJSON, err := json.Marshal(export)
		if err != nil {
			log.Fatalf("failed convert struct FileRepresentation of sample %s into json; got %v", identifier, err)
		}

		w.Write(toWriteJSON)
		w.Close() // You must close this first to flush the bytes to the buffer.

		ioutil.WriteFile(outputFilePath, b.Bytes(), 0644)
	}
}
