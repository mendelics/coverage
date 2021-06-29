package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getTotalBases(startField, endField string) (int, error) {
	start, err := strconv.Atoi(startField)
	if err != nil {
		return 0, fmt.Errorf("ERROR CONVERTING START TO INT: %s", startField)
	}
	start-- // converts variant to 0-based

	// Info column 7
	// END=153773270;IDP=137.83;IGC=0.538
	// END=153776164;IDP=56.41;IGC=0.704
	info := strings.Split(endField, ";")
	endStr := strings.Split(info[0], "=")
	end, err := strconv.Atoi(endStr[1])
	if err != nil {
		return 0, fmt.Errorf("ERROR CONVERTING END TO INT: %s", endStr)
	}

	return end - start, nil
}

func getLowCoverage(coverageField string) (int, error) {
	// Coverage column 9
	// 137.83:0:0
	// LOW_COVERAGE:56.41:199:0
	coverageInfo := strings.Split(coverageField, ":")

	ll, err := strconv.Atoi(coverageInfo[len(coverageInfo)-2])
	if err != nil {
		return 0, fmt.Errorf("ERROR CONVERTING LL TO INT: %s", coverageInfo)
	}

	zl, err := strconv.Atoi(coverageInfo[len(coverageInfo)-1])
	if err != nil {
		return 0, fmt.Errorf("ERROR CONVERTING ZL TO INT: %s", coverageInfo)
	}

	return ll + zl, nil
}

func parseCoverageVcfLine(line string) ([]int, error) {
	// chr1    153773193       .       A       <DT>    .       PASS    END=153773270;IDP=137.83;IGC=0.538      IDP:LL:ZL       137.83:0:0
	// chr1    153775357       .       A       <DT>    .       LOW_COVERAGE    END=153776164;IDP=56.41;IGC=0.704       FT:IDP:LL:ZL    LOW_COVERAGE:56.41:199:0
	line = strings.TrimSpace(line)
	fields := strings.Split(line, "\t")
	if len(fields) < 10 {
		return nil, fmt.Errorf("wrong amount of columns: %d", len(fields))
	}

	totalBases, err := getTotalBases(fields[1], fields[7])
	if err != nil {
		return nil, err
	}

	lowCoverage, err := getLowCoverage(fields[9])
	if err != nil {
		return nil, err
	}

	return []int{totalBases, lowCoverage}, nil
}
