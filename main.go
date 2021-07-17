package main

import (
	"flag"
	"log"
	"time"
)

func main() {
	t0 := time.Now()

	// Args
	vcf := flag.String("vcf", "", "VCF file with coverage (GATK-format).")
	bed := flag.String("bed", "", "BED file with coverage (Mosdepth-format).")
	sample := flag.String("sample", "no_sample_id", "Sample ID")
	flag.Parse()

	targets := make(map[string][]TargetCoverage)
	var globalTotalBases, globalCoveredBases5x, globalCoveredBases10x, globalCoveredBases20x, globalCoveredBases30x int

	switch {
	// No coverage file
	case *vcf == "" && *bed == "":
		log.Fatalln("No vcf or bed informed. Run `./coverage --vcf <COVERAGE_VCF> or ./coverage --bed <COVERAGE_BED>")

	// Too many coverage files
	case *vcf != "" && *bed != "":
		log.Fatalln("vcf and bed are mutually exclusive. Run `./coverage --vcf <COVERAGE_VCF> or ./coverage --bed <COVERAGE_BED>")

	// GATK-derived coverage
	case *vcf != "":
		targets, globalTotalBases, globalCoveredBases10x = parseGatkVcf(*vcf)

	// Mosdepth-derived coverage
	case *bed != "":
		targets, globalTotalBases, globalCoveredBases5x, globalCoveredBases10x, globalCoveredBases20x, globalCoveredBases30x = parseMosdepthBed(*bed)

	}

	// Intersect coverage with transcripts
	geneCoverageMap := getGeneCoverage(targets)

	// Output coverage to JSON file
	outputToJSON(*sample, geneCoverageMap, globalTotalBases, globalCoveredBases5x, globalCoveredBases10x, globalCoveredBases20x, globalCoveredBases30x)

	// Time
	log.Printf("Executed in %.2f seconds\n", time.Since(t0).Seconds())
}
