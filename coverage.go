package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/biogo/hts/bgzf"
	"github.com/brentp/bix"
	"github.com/brentp/irelate"
	"github.com/brentp/irelate/interfaces"
	"github.com/brentp/irelate/parsers"
)

// getCoverage parses the GATK-derived coverage VCF and returns a map
func getCoverage(coverageVcfFile string) (map[string]Coverage, []string) {
	transcriptsPath := path.Join(os.Getenv("HOME"), "coverage", "transcripts_GRCh38.bed.gz")

	// chr1	358856	366052	ENSG00000236601
	qCoverageBed, err := bix.New(transcriptsPath, 0)
	if err != nil {
		log.Fatalf("error generating indexing %v", err)
	}

	geneCoverageMap := make(map[string]Coverage)

	// covStream is the streaming of coverage intervals
	covStream, sampleNames := prepareVcfToStream(coverageVcfFile)

	// iterate over coverage intervals
	for interval := range irelate.PIRelate(4000, 40000, covStream, false, nil, qCoverageBed) {

		if len(interval.Related()) != 0 {

			// coverage for specific interval
			intervalCoverageMap, err := sumCoverage(interval)
			if err != nil {
				log.Println(err)
				continue
			}

			// update coverage map that will be used in exported JSON
			geneCoverageMap = updateCoverage(geneCoverageMap, intervalCoverageMap)
		}
	}

	geneCoverageWithPctMap := make(map[string]Coverage)

	for ensg, val := range geneCoverageMap {
		if val.TotalBases != 0 {
			val.CoveragePct = float64(val.CoveredBases) / float64(val.TotalBases)
		}
		geneCoverageWithPctMap[ensg] = val
	}

	return geneCoverageWithPctMap, sampleNames
}

func prepareVcfToStream(coverageVcfFile string) (interfaces.RelatableIterator, []string) {
	// Check coverage from GATK-generated coverage VCF
	rdr, err := ioutil.ReadFile(coverageVcfFile)
	if err != nil {
		log.Fatalf("error reading file %v", err)
	}

	byteRdr := bytes.NewReader(rdr)

	// Presumes VCF with denormalized variants ("bcftools norm -m -any"), bgzipped and tabix indexed
	qrdr, err := bgzf.NewReader(byteRdr, 0)
	if err != nil {
		log.Fatalf("error opening query file %s: %v", coverageVcfFile, err)
	}

	// Parse VCF file with vcfgo to use in irelate
	covStream, vcfRdr, err := parsers.VCFIterator(qrdr)
	if err != nil {
		log.Fatalf("error parsing VCF query file %s: %v", coverageVcfFile, err)
	}

	return covStream, vcfRdr.Header.SampleNames
}

func sumCoverage(interval interfaces.Relatable) (map[string]Coverage, error) {

	// Parse VCF line
	vcfLineStr := fmt.Sprintf("%v", interval)
	cov, err := parseCoverageVcfLine(vcfLineStr)
	if err != nil {
		log.Println("Failed to parse vcf line", vcfLineStr, err)
		return nil, err
	}

	intervalCoverageMap := make(map[string]Coverage)

	for _, relatedInterval := range interval.Related() {
		// chr1    358856  365704  ENSG00000223181;ENST00000411249;NONCODING;-;RNU6-1199P
		relatedStr := fmt.Sprintf("%v", relatedInterval)
		bedFields := strings.Split(relatedStr, "\t")
		info := strings.Split(bedFields[3], ";")

		// ENSG00000223181;ENST00000411249;NONCODING;-;RNU6-1199P
		intervalCoverageMap[info[0]] = Coverage{
			ENSG:         info[0],
			Symbol:       info[4],
			TotalBases:   cov[0],
			CoveredBases: (cov[0] - cov[1]),
		}

	}

	return intervalCoverageMap, nil
}

func updateCoverage(geneCoverageMap map[string]Coverage, intervalCoverageMap map[string]Coverage) map[string]Coverage {

	for ensg, intervalCov := range intervalCoverageMap {
		if geneCoverage, exists := geneCoverageMap[ensg]; exists {
			geneCoverage.ENSG = ensg
			geneCoverage.Symbol = intervalCov.Symbol
			geneCoverage.TotalBases += intervalCov.TotalBases
			geneCoverage.CoveredBases += intervalCov.CoveredBases
			geneCoverageMap[ensg] = geneCoverage

		} else {
			newCov := Coverage{
				ENSG:         ensg,
				Symbol:       intervalCov.Symbol,
				TotalBases:   intervalCov.TotalBases,
				CoveredBases: intervalCov.CoveredBases,
			}
			geneCoverageMap[ensg] = newCov

		}
	}
	return geneCoverageMap
}
