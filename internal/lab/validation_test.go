package lab

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func TestNormalizeOperatorKeepsMultibyteNamesValid(t *testing.T) {
	name := strings.Repeat("李", 81)
	got := NormalizeOperator(name)
	if !utf8.ValidString(got) {
		t.Fatalf("NormalizeOperator() returned invalid UTF-8: %q", got)
	}
	if got != strings.Repeat("李", 80) {
		t.Fatalf("NormalizeOperator() returned %d runes, want 80", utf8.RuneCountInString(got))
	}
}
