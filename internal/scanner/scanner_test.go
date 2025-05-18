package scanner

import "testing"

func TestScanFindsPublicResources(t *testing.T) {
	res := []Resource{{Name: "a", Public: true}, {Name: "b", Public: false}}
	findings := Scan(res)
	if len(findings) != 1 {
		t.Fatalf("expected 1 finding, got %d", len(findings))
	}
}
