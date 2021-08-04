package main

import "testing"

func TestParseCoverageBed(t *testing.T) {
	tt := []struct {
		chr     string
		targets int
	}{
		{"chr1", 22061},
		{"chr2", 16644},
		{"chrX", 7315},
		{"chrY", 624},
		{"chrM", 13},
	}

	var expectedGlobal = 36726635
	var expected5x = 36462766
	var expected10x = 36290670
	var expected20x = 35708523
	var expected30x = 34512655

	for i, test := range tt {
		resultMap, globalTotal, global5x, global10x, global20x, global30x := parseMosdepthBed("samples/example_mosdepth_coverage.bed.gz")

		if globalTotal != expectedGlobal {
			t.Errorf("expected globalTotal %d, got %d on test number %d", expectedGlobal, globalTotal, i)
		}
		if global5x != expected5x {
			t.Errorf("expected global5x %d, got %d on test number %d", expected5x, global5x, i)
		}
		if global10x != expected10x {
			t.Errorf("expected global10x %d, got %d on test number %d", expected10x, global10x, i)
		}
		if global20x != expected20x {
			t.Errorf("expected global20x %d, got %d on test number %d", expected20x, global20x, i)
		}
		if global30x != expected30x {
			t.Errorf("expected global30x %d, got %d on test number %d", expected30x, global30x, i)
		}
		if len(resultMap[test.chr]) != test.targets {
			t.Errorf("expected %d targets for chr %s, got %d on test number %d", test.targets, test.chr, len(resultMap[test.chr]), i)
		}
	}
}
