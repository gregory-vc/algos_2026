package task7

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindLCSExample(t *testing.T) {
	got := Find("ABDEFADRFG", "DAFERG")
	expected := []string{"DAFG", "AFRG", "AERG", "DFRG", "DERG", "DARG"}

	sort.Strings(got)
	sort.Strings(expected)

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("got %v, want %v", got, expected)
	}
}

func TestFindLCSSingleChar(t *testing.T) {
	got := Find("ABC", "DEF")
	if len(got) != 1 || got[0] != "" {
		t.Fatalf("got %v, want [\"\"]", got)
	}
}

func TestFindLCSOverlapping(t *testing.T) {
	got := Find("AA", "AA")
	if len(got) != 1 || got[0] != "AA" {
		t.Fatalf("got %v, want [\"AA\"]", got)
	}
}
