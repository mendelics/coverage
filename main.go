package main

import (
	"flag"
	"log"
)

func main() {
	// Coverage vcf
	vcf := flag.String("vcf", "", "VCF file with coverage (GATK-format).")
	bed := flag.String("bed", "", "BED file with coverage (Mosdepth-format).")
	flag.Parse()

	geneCoverageMap := make(map[string]Coverage)
	sampleNames := make([]string, 0)

	switch {
	case *vcf == "" && *bed == "":
		log.Fatalln("No vcf or bed informed. Run `./coverage --vcf <COVERAGE_VCF> or ./coverage --bed <COVERAGE_BED>")
	case *vcf != "" && *bed != "":
		log.Fatalln("vcf and bed are mutually exclusive. Run `./coverage --vcf <COVERAGE_VCF> or ./coverage --bed <COVERAGE_BED>")
	case *vcf != "":
		// Intersect coverage with transcripts
		geneCoverageMap, sampleNames = getCoverage(*vcf)
	case *bed != "":
		// Intersect coverage with transcripts
		getMosdepth(*bed)
	}

	// Output coverage to JSON file
	outputToJSON(sampleNames, geneCoverageMap)
}
