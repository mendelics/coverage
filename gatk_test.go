package main

import "testing"

func TestParseCoverageVcf(t *testing.T) {
	tt := []struct {
		chr     string
		targets int
	}{
		// {"ENSG00000001626", 0},
		// {"ENSG00000012048", 0},
		// {"ENSG00000277027", 0},
		// {"ENSG00000172062", 0},
		// {"ENSG00000198947", 0},

		{"chr1", 19712},
		{"chr2", 14730},
		{"chrX", 6769},
		{"chrY", 539},
		{"chrM", 0},
	}

	for i, test := range tt {
		resultMap := parseGatkVcf("example_gatk_coverage.vcf.gz")

		if len(resultMap[test.chr]) != test.targets {
			t.Errorf("expected %d targets for chr %s, got %d on test number %d", test.targets, test.chr, len(resultMap[test.chr]), i)
		}
	}
}
