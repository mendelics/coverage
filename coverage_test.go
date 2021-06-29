// +build integration

package coverage

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestCoverage001(t *testing.T) {
	expectedIdentifier := "TEST-00"
	expectedGlobalTotalBases := 102734451
	expectedGlobalCoverage := 0.981938
	expectedCoverageCFTRPct := 1.000000

	sampleFolder := path.Join(os.Getenv("HOME"), "vsa2", "data", "samples")

	// InputVcf
	samplePath := path.Join(sampleFolder, "GRCh38_sample_coverage.vcf.gz")

	// OutputJsonl
	outputJSON := path.Join(sampleFolder, "TEST-00_vsa_coverage.json")

	// Delete old OutputJson
	_ = os.Remove(outputJSON)

	GetCoverage(samplePath, sampleFolder, "")

	jsonFile, err := ioutil.ReadFile(outputJSON)
	if err != nil {
		t.Errorf("unexpected err %v on output JSON read", err)
	}

	var exported PanelCoverage
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	err = json.Unmarshal(jsonFile, &exported)
	if err != nil {
		t.Errorf("unexpected err %v on JSON unmarshall", err)
	}

	cftrTotalBases := exported.PerGeneCoverage["ENSG00000001626"].TotalBases
	cftrCoveredBases := exported.PerGeneCoverage["ENSG00000001626"].CoveredBases
	pctCFTR := float64(cftrCoveredBases) / float64(cftrTotalBases)

	if expectedIdentifier != exported.Identifier {
		t.Errorf("expected identifier %s, got %s", expectedIdentifier, exported.Identifier)
	}
	if expectedGlobalTotalBases != exported.GlobalTotalBases {
		t.Errorf("expected global total bases %d, got %d", expectedGlobalTotalBases, exported.GlobalTotalBases)
	}
	if expectedGlobalCoverage-exported.GlobalCoverage > 0.001 || expectedGlobalCoverage-exported.GlobalCoverage < -0.001 {
		t.Errorf("expected global coverage %f, got %f", expectedGlobalCoverage, exported.GlobalCoverage)
	}
	if expectedCoverageCFTRPct-pctCFTR > 0.001 || expectedCoverageCFTRPct-pctCFTR < -0.001 {
		t.Errorf("expected CFTR coverage %f, got %f", expectedCoverageCFTRPct, pctCFTR)
	}
}
