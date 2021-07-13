package main

import (
	"compress/gzip"
	_ "embed"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type TargetCoverage struct {
	chr   string
	start int
	end   int
	x5    int
	x10   int
	x20   int
	x30   int
}

func parseMosdepthBed(coverageBed string) map[string][]TargetCoverage {
	coverageFile, err := os.Open(coverageBed)
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

	record, err := tsvReader.Read()
	for err == nil {
		if len(record) == 0 {
			break
		}

		start, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatalf("error parsing start, %v", err)
		}

		end, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatalf("error parsing end, %v", err)
		}

		x5, err := strconv.Atoi(record[4])
		if err != nil {
			log.Fatalf("error parsing x5, %v", err)
		}

		x10, err := strconv.Atoi(record[5])
		if err != nil {
			log.Fatalf("error parsing x10, %v", err)
		}

		x20, err := strconv.Atoi(record[6])
		if err != nil {
			log.Fatalf("error parsing x20, %v", err)
		}

		x30, err := strconv.Atoi(record[7])
		if err != nil {
			log.Fatalf("error parsing x30, %v", err)
		}

		if err != nil {
			log.Fatal("unable to parse bed file", err, record)
		}

		bed := TargetCoverage{
			chr:   record[0],
			start: start,
			end:   end,
			x5:    x5,
			x10:   x10,
			x20:   x20,
			x30:   x30,
		}

		if beds, exists := targets[bed.chr]; exists {
			beds = append(beds, bed)
			targets[bed.chr] = beds
		} else {
			targets[bed.chr] = []TargetCoverage{bed}
		}

		record, err = tsvReader.Read()
	}

	if err != nil && err != io.EOF {
		log.Fatal("error parsing coverage bed file", err, record)
	}

	return targets
}
