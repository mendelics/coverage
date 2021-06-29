package main

import "testing"

func TestParseCoverageVcfLine(t *testing.T) {
	tt := []struct {
		line         string
		expectedInt0 int
		expectedInt1 int
		expectedErr  bool
	}{
		{"chr1\t153773193\t.\tA\t<DT>\t.\tPASS\tEND=153773270;IDP=137.83;IGC=0.538\tIDP:LL:ZL\t137.83:0:0", 78, 0, false},
		{"chr1\t153775357\t.\tA\t<DT>\t.\tLOW_COVERAGE\tEND=153776164;IDP=56.41;IGC=0.704\tFT:IDP:LL:ZL\tLOW_COVERAGE:56.41:199:0", 808, 199, false},
		{"chr1\t153775357\t.\tA\t<DT>\t.\tLOW_COVERAGE\tEND=153776164;IDP=56.41;IGC=0.704\tFT:IDP:LL:ZL\tLOW_COVERAGE:56.41:199:1", 808, 200, false},
		{"chr1\t153775357\t.\tA\t<DT>\t.\tLOW_COVERAGE\tEND=153776164;IDP=56.41;IGC=0.704\tFT:IDP:LL:ZL\tLOW_COVERAGE:56.41:0:1", 808, 1, false},
	}
	for i, test := range tt {
		result, err := parseCoverageVcfLine(test.line)
		if (err != nil && !test.expectedErr) || (err == nil && test.expectedErr) {
			if test.expectedErr {
				t.Errorf("expected err but got nil on test number %d", i)
			} else {
				t.Errorf("unexpected err %v on test number %d", err, i)
			}

		}
		if result[0] != test.expectedInt0 {
			t.Errorf("expected totalBases %d got %d on test number %d", test.expectedInt0, result[0], i)
		}
		if result[1] != test.expectedInt1 {
			t.Errorf("expected lowCoverage %d got %d on test number %d", test.expectedInt1, result[1], i)
		}
	}
}
