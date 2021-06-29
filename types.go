package main

// PanelCoverage contains all gene coverages
type PanelCoverage struct {
	Identifier string
	Genes      map[string]Coverage
}

// Coverage data
type Coverage struct {
	ENSG         string
	Symbol       string
	TotalBases   int
	CoveredBases int
	CoveragePct  float64
}
