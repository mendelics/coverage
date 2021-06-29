package main

import (
	"flag"
	"log"
)

func main() {
	// Coverage vcf
	vcf := flag.String("vcf", "", "VCF file with coverage (GATK-format).")
	flag.Parse()

	if *vcf == "" {
		log.Fatalln("No vcf informed. Run `./coverage --vcf <COVERAGE_VCF>")
	}

	// Intersect coverage with transcripts
	geneCoverageMap, sampleNames := getCoverage(*vcf)

	// Output coverage to JSON file
	outputToJSON(sampleNames, geneCoverageMap)

}
