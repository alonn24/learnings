package permutations

import (
	"testing"
)

func TestPerm(t *testing.T) {
	ans := Perm("abc")
	if len(ans) != 6 {
		t.Fatalf("Perm(\"abc\") = %v, wanted 6 results", ans)
	}
}
