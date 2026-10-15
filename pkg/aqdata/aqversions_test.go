package aqdata

import (
	"slices"
	"testing"
)

// tests type conversions for AquesTalk versions
func TestAquesVersion(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		n    int
		ver  AquesVersion
		text string
	}{
		{1, AqVer1, "AquesTalk1"},
		{2, AqVer2, "AquesTalk2"},
		{10, AqVer10, "AquesTalk10"},
	} {
		if got := ToAquesVersion(tc.n); got != tc.ver {
			t.Errorf("ToAquesVersion(%d) = %v, expected %v", tc.n, got, tc.ver)
		}
		if got := tc.ver.ToInt(); got != tc.n {
			t.Errorf("ToInt(%v) = %d, expected %d", tc.ver, got, tc.n)
		}
		if got := tc.ver.ToString(); got != tc.text {
			t.Errorf("ToString(%v) = %q, expected %q", tc.ver, got, tc.text)
		}
	}
	for _, n := range []int{-1, 0, 3, 114514} {
		if got := ToAquesVersion(n); got != AqVer2 {
			t.Errorf("ToAquesVersion(%d) = %v, expected %v", n, got, AqVer2)
		}
	}
	for _, v := range []AquesVersion{0, 3, 114514} {
		if got := v.ToInt(); got != 2 {
			t.Errorf("ToInt(%v) = %d, expected 2", v, got)
		}
		if got := v.ToString(); got != "Not Defined" {
			t.Errorf("ToString(%v) = %q, expected %q", v, got, "Not Defined")
		}
	}
	got, expected := GetAllVersions(), []AquesVersion{AqVer1, AqVer2, AqVer10}
	if !slices.Equal(got, expected) {
		t.Errorf("GetAllVersions() = %v, expected %v", got, expected)
	}
}
