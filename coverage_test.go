package main

import "testing"

func TestCoveragePerGeneFromBed(t *testing.T) {
	tt := []struct {
		ENSG               string
		Symbol             string
		TotalTargetedBases int
		BasesCovered5x     int
		BasesCovered10x    int
		BasesCovered20x    int
		BasesCovered30x    int
	}{
		{"ENSG00000001626", "CFTR", 4443, 4443, 4443, 4301, 3970},
		{"ENSG00000012048", "BRCA1", 5797, 5797, 5797, 5797, 5416},
		{"ENSG00000277027", "RMRP", 0, 0, 0, 0, 0},
		{"ENSG00000172062", "SMN1", 900, 900, 900, 890, 771},
		{"ENSG00000198947", "DMD", 11621, 11621, 11466, 10891, 9093},
	}

	for i, test := range tt {
		targets, _, _, _, _, _ := parseMosdepthBed("example_mosdepth_coverage.bed.gz")

		resultMap := getGeneCoverage(targets)

		gene := resultMap[test.ENSG]

		if gene.Symbol != test.Symbol {
			t.Errorf("expected %s symbol, got %s on gene %s, test number %d", test.Symbol, gene.Symbol, gene.ENSG, i)
		}

		if gene.TotalBases != test.TotalTargetedBases {
			t.Errorf("expected %d total bases for gene %s, got %d on test number %d", test.TotalTargetedBases, test.Symbol, gene.TotalBases, i)
		}

		if gene.BasesCovered5x != test.BasesCovered5x {
			t.Errorf("expected %d bases covered at 5x for gene %s, got %d on test number %d", test.BasesCovered5x, test.Symbol, gene.BasesCovered5x, i)
		}

		if gene.BasesCovered10x != test.BasesCovered10x {
			t.Errorf("expected %d bases covered at 10x for gene %s, got %d on test number %d", test.BasesCovered10x, test.Symbol, gene.BasesCovered10x, i)
		}

		if gene.BasesCovered20x != test.BasesCovered20x {
			t.Errorf("expected %d bases covered at 20x for gene %s, got %d on test number %d", test.BasesCovered20x, test.Symbol, gene.BasesCovered20x, i)
		}

		if gene.BasesCovered30x != test.BasesCovered30x {
			t.Errorf("expected %d bases covered at 30x for gene %s, got %d on test number %d", test.BasesCovered30x, test.Symbol, gene.BasesCovered30x, i)
		}

	}
}

func TestCoveragePerGeneFromVcf(t *testing.T) {
	tt := []struct {
		ENSG               string
		Symbol             string
		TotalTargetedBases int
		BasesCovered5x     int
		BasesCovered10x    int
		BasesCovered20x    int
		BasesCovered30x    int
	}{
		{"ENSG00000001626", "CFTR", 4443, 0, 4443, 0, 0},
		{"ENSG00000012048", "BRCA1", 5658, 0, 5658, 0, 0},
		{"ENSG00000277027", "RMRP", 0, 0, 0, 0, 0},
		{"ENSG00000172062", "SMN1", 900, 0, 900, 0, 0},
		{"ENSG00000198947", "DMD", 11213, 0, 11213, 0, 0},
	}

	for i, test := range tt {
		targets, _, _ := parseGatkVcf("example_gatk_coverage.vcf.gz")

		resultMap := getGeneCoverage(targets)

		gene := resultMap[test.ENSG]

		if gene.Symbol != test.Symbol {
			t.Errorf("expected %s symbol, got %s on gene %s, test number %d", test.Symbol, gene.Symbol, gene.ENSG, i)
		}

		if gene.TotalBases != test.TotalTargetedBases {
			t.Errorf("expected %d total bases for gene %s, got %d on test number %d", test.TotalTargetedBases, test.Symbol, gene.TotalBases, i)
		}

		if gene.BasesCovered5x != test.BasesCovered5x {
			t.Errorf("expected %d bases covered at 5x for gene %s, got %d on test number %d", test.BasesCovered5x, test.Symbol, gene.BasesCovered5x, i)
		}

		if gene.BasesCovered10x != test.BasesCovered10x {
			t.Errorf("expected %d bases covered at 10x for gene %s, got %d on test number %d", test.BasesCovered10x, test.Symbol, gene.BasesCovered10x, i)
		}

		if gene.BasesCovered20x != test.BasesCovered20x {
			t.Errorf("expected %d bases covered at 20x for gene %s, got %d on test number %d", test.BasesCovered20x, test.Symbol, gene.BasesCovered20x, i)
		}

		if gene.BasesCovered30x != test.BasesCovered30x {
			t.Errorf("expected %d bases covered at 30x for gene %s, got %d on test number %d", test.BasesCovered30x, test.Symbol, gene.BasesCovered30x, i)
		}

	}
}
