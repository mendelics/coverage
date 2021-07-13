package main

import (
	"compress/gzip"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/markbates/pkger"
)

type Gene struct {
	chr    string
	start  int
	end    int
	ensg   string
	symbol string
}

type Bed struct {
	chr   string
	start int
	end   int
	x5    int
	x10   int
	x20   int
	x30   int
}

func getMosdepth(coverageBed string) {
	targets := parseMosdepthBed(coverageBed)

	f, err := pkger.Open("genes.bed")
	if err != nil {
		log.Fatalln("couldn't find genes.bed")
	}
	defer f.Close()

	genesFile, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("unable to open genes file")
	}

	geneLines := strings.Split(string(genesFile), "\n")

	for _, line := range geneLines {
		fields := strings.Split(line, "\t")
		if len(fields) != 5 {
			continue
		}

		start, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalf("error parsing start, %v", err)
		}

		end, err := strconv.Atoi(fields[2])
		if err != nil {
			log.Fatalf("error parsing end, %v", err)
		}

		gene := Gene{
			chr:    fields[0],
			start:  start,
			end:    end,
			ensg:   fields[3],
			symbol: fields[4],
		}

		var overlapCount, totalTargets, totalCovered5x, totalCovered10x, totalCovered20x, totalCovered30x int

		for _, target := range targets[gene.chr] {
			if !(gene.end < target.start || gene.start > target.end) {
				overlapCount++
				totalTargets += target.end - target.start
				totalCovered5x += target.x5
				totalCovered10x += target.x10
				totalCovered20x += target.x20
				totalCovered30x += target.x30
			}
		}

		fmt.Printf("%s\t%s\tCount: %d\tTotal Bases: %d\tCoverage 5x: %d\t10x: %d\t20x: %d\t30x: %d", gene.ensg, gene.symbol, overlapCount, totalTargets, totalCovered5x, totalCovered10x, totalCovered20x, totalCovered30x)
	}
}

func parseMosdepthBed(coverageBed string) map[string][]Bed {
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

	targets := make(map[string][]Bed)

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

		bed := Bed{
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
			targets[bed.chr] = []Bed{bed}
		}

		record, err = tsvReader.Read()
	}

	if err != nil && err != io.EOF {
		log.Fatal("error parsing coverage bed file", err, record)
	}

	return targets
}
