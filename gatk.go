package main

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseGatkVcf(coverageVcf string) (map[string][]TargetCoverage, int, int) {
	coverageFile, err := os.Open(coverageVcf)
	if err != nil {
		log.Fatal("unable to open coverage file")
	}
	defer coverageFile.Close()

	coverageRdr, err := gzip.NewReader(coverageFile)
	if err != nil {
		log.Fatal(err)
	}

	// read gencode file as TSV
	tsvReader := csv.NewReader(coverageRdr)
	tsvReader.Comma = '\t'
	tsvReader.Comment = '#'

	targets := make(map[string][]TargetCoverage)
	var globalTotalBases, globalCoveredBases10x int

	record, err := tsvReader.Read()
	for err == nil {
		if len(record) == 0 {
			break
		}

		position, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatalf("error parsing position, %v", err)
		}

		// END=153773270;IDP=137.83;IGC=0.538
		// END=153776164;IDP=56.41;IGC=0.704
		info := strings.Split(record[7], ";")
		endStr := strings.Split(info[0], "=")
		end, err := strconv.Atoi(endStr[1])
		if err != nil {
			log.Fatalf("error parsing end, %v", err)
		}

		lowCoverage, err := getLowCoverage(record[9])
		if err != nil {
			log.Fatalf("error parsing x10, %v", err)
		}

		if err != nil {
			log.Fatal("unable to parse vcf file", err, record)
		}

		globalTotalBases += end - position + 1
		globalCoveredBases10x += end - position + 1 - lowCoverage

		vcf := TargetCoverage{
			chr:   record[0],
			start: position - 1,
			end:   end,
			x10:   end - position + 1 - lowCoverage,
		}

		if beds, exists := targets[vcf.chr]; exists {
			beds = append(beds, vcf)
			targets[vcf.chr] = beds
		} else {
			targets[vcf.chr] = []TargetCoverage{vcf}
		}

		record, err = tsvReader.Read()
	}

	if err != nil && err != io.EOF {
		log.Fatal("error parsing coverage bed file", err, record)
	}

	return targets, globalTotalBases, globalCoveredBases10x
}

func getLowCoverage(coverageField string) (int, error) {
	// Coverage column 9
	// 137.83:0:0
	// LOW_COVERAGE:56.41:199:0
	coverageInfo := strings.Split(coverageField, ":")

	ll, err := strconv.Atoi(coverageInfo[len(coverageInfo)-2])
	if err != nil {
		return 0, fmt.Errorf("error converting LL to int: %s", coverageInfo)
	}

	zl, err := strconv.Atoi(coverageInfo[len(coverageInfo)-1])
	if err != nil {
		return 0, fmt.Errorf("error converting ZL to int: %s", coverageInfo)
	}

	return ll + zl, nil
}
