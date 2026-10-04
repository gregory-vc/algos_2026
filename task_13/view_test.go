package task13

import (
	"bytes"
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestConsoleViewBindsEvents(t *testing.T) {
	events := &Events{}
	var output bytes.Buffer
	view := NewConsoleView(events, strings.NewReader("raw input\n\nlast line"), &output)
	var inputs []string
	events.InputChanged.Subscribe(func(input string) { inputs = append(inputs, input) })
	events.StateChanged.Publish(ViewState{Text: "ready"})
	if err := view.Run(); err != nil {
		t.Fatal(err)
	}
	if want := []string{"raw input", "", "last line"}; !reflect.DeepEqual(inputs, want) {
		t.Fatalf("inputs = %q, want %q", inputs, want)
	}
	if got := output.String(); got != "ready\n" {
		t.Fatalf("output = %q, want ready followed by newline", got)
	}
}

type failingReader struct{ err error }

func (r failingReader) Read([]byte) (int, error) { return 0, r.err }

type failingWriter struct{ err error }

func (w failingWriter) Write([]byte) (int, error) { return 0, w.err }

func TestConsoleViewReturnsReadError(t *testing.T) {
	readError := errors.New("input unavailable")
	view := NewConsoleView(&Events{}, failingReader{readError}, &bytes.Buffer{})
	if err := view.Run(); !errors.Is(err, readError) {
		t.Fatalf("Run error = %v, want %v", err, readError)
	}
}

func TestConsoleViewStopsAfterWriteError(t *testing.T) {
	events := &Events{}
	NewCalculator(events)
	NewCalculatorViewModel(events)
	writeError := errors.New("output unavailable")
	view := NewConsoleView(events, strings.NewReader("+ 7 5\n- 7 5\n"), failingWriter{writeError})
	requests := 0
	events.CalculationRequested.Subscribe(func(Request) { requests++ })
	if err := view.Run(); !errors.Is(err, writeError) {
		t.Fatalf("Run error = %v, want %v", err, writeError)
	}
	if requests != 1 {
		t.Fatalf("handled %d requests after output failed, want 1", requests)
	}
}
