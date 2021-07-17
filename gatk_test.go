package main

import "testing"

func TestParseCoverageVcf(t *testing.T) {
	tt := []struct {
		chr     string
		targets int
	}{
		{"chr1", 19712},
		{"chr2", 14730},
		{"chrX", 6769},
		{"chrY", 539},
		{"chrM", 0},
	}

	var expectedGlobal = 33167129
	var expected10x = 32369680

	for i, test := range tt {
		resultMap, globalTotal, global10x := parseGatkVcf("example_gatk_coverage.vcf.gz")

		if globalTotal != expectedGlobal {
			t.Errorf("expected globalTotal %d, got %d on test number %d", expectedGlobal, globalTotal, i)
		}
		if global10x != expected10x {
			t.Errorf("expected global10x %d, got %d on test number %d", expected10x, global10x, i)
		}
		if len(resultMap[test.chr]) != test.targets {
			t.Errorf("expected %d targets for chr %s, got %d on test number %d", test.targets, test.chr, len(resultMap[test.chr]), i)
		}
	}
}
