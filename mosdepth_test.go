package main

import "testing"

func TestParseCoverageBed(t *testing.T) {
	tt := []struct {
		chr     string
		targets int
	}{
		// {"ENSG00000001626", 0},
		// {"ENSG00000012048", 0},
		// {"ENSG00000277027", 0},
		// {"ENSG00000172062", 0},
		// {"ENSG00000198947", 0},

		{"chr1", 22061},
		{"chr2", 16644},
		{"chrX", 7315},
		{"chrY", 624},
		{"chrM", 13},
	}

	for i, test := range tt {
		resultMap := parseMosdepthBed("example_mosdepth_coverage.bed.gz")

		if len(resultMap[test.chr]) != test.targets {
			t.Errorf("expected %d targets for chr %s, got %d on test number %d", test.targets, test.chr, len(resultMap[test.chr]), i)
		}
	}
}
