package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Gene struct {
	chr    string
	start  int
	end    int
	ensg   string
	symbol string
}

//go:embed genes.bed
var genesFile string

func getGeneCoverage(targets map[string][]TargetCoverage) map[string]Coverage {
	geneLines := strings.Split(string(genesFile), "\n")
	geneCoverageMap := make(map[string]Coverage)
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

		geneCoverageMap[gene.ensg] = Coverage{
			ENSG:            gene.ensg,
			Symbol:          gene.symbol,
			TotalBases:      totalTargets,
			CoveredBases:    totalCovered10x,
			BasesCovered5x:  totalCovered5x,
			BasesCovered10x: totalCovered10x,
			BasesCovered20x: totalCovered20x,
			BasesCovered30x: totalCovered30x,
		}

		fmt.Printf("%s\t%s\tCount: %d\tTotal Bases: %d\tCoverage 5x: %d\t10x: %d\t20x: %d\t30x: %d\n", gene.ensg, gene.symbol, overlapCount, totalTargets, totalCovered5x, totalCovered10x, totalCovered20x, totalCovered30x)
	}
	return geneCoverageMap
}
