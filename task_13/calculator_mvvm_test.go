package task13

import (
	"bytes"
	"strings"
	"testing"
)

func TestCalculatorMVVM(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"arithmetic", "+ 7 5\n- 7 5\n+ -7 5\n- 0 0", "7+5=12\n7-5=2\n-7+5=-2\n0-0=0\n"},
		{"custom operation", "* 7 5\n", "7*5=35\n"},
		{"recovery", "+ 7\n? 7 5\n+ 2 3\n", "error: expected: operation left right\nerror: unknown operation \"?\"\n2+3=5\n"},
		{"empty stream", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			events := &Events{}
			model := NewCalculator(events)
			if err := model.AddOperation("*", func(x, y int) (int, error) { return x * y, nil }); err != nil {
				t.Fatal(err)
			}
			NewCalculatorViewModel(events)
			var output bytes.Buffer
			view := NewConsoleView(events, strings.NewReader(tt.input), &output)
			if err := view.Run(); err != nil {
				t.Fatal(err)
			}
			if got := output.String(); got != tt.want {
				t.Fatalf("output = %q, want %q", got, tt.want)
			}
		})
	}
}
