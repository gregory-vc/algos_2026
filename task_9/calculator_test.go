package task9

import (
	"bytes"
	"strings"
	"testing"
)

func TestOperationMakerPullsDependencies(t *testing.T) {
	input := strings.NewReader("+ 7 5\n- 7 5\n")
	var output bytes.Buffer
	container := NewContainer()
	services := []any{
		PlusOperation{},
		MinusOperation{},
		NewTextDataReader(input),
		NewTextResultPrinter(&output),
	}
	for _, service := range services {
		container.Add(service)
	}

	maker := &OperationMaker{}
	if err := maker.Init(container); err != nil {
		t.Fatalf("initialize operation maker: %v", err)
	}
	container.Add(maker)
	for range 2 {
		if _, err := maker.Make(); err != nil {
			t.Fatalf("make operation: %v", err)
		}
	}

	if got, want := output.String(), "7+5=12\n7-5=2\n"; got != want {
		t.Fatalf("output = %q, want %q", got, want)
	}
	if stored := container.services[len(container.services)-1]; stored != maker {
		t.Fatalf("container did not retain initialized maker: got %T", stored)
	}
}

func TestOperationMakerRejectsMissingDependencies(t *testing.T) {
	container := NewContainer()
	container.Add(PlusOperation{})
	if err := (&OperationMaker{}).Init(container); err == nil {
		t.Fatal("initialization succeeded without reader and printer")
	}
}

func TestOperationMakerRejectsUnknownOperation(t *testing.T) {
	container := NewContainer()
	services := []any{
		PlusOperation{},
		NewTextDataReader(strings.NewReader("* 2 3\n")),
		NewTextResultPrinter(&bytes.Buffer{}),
	}
	for _, service := range services {
		container.Add(service)
	}
	maker := &OperationMaker{}
	if err := maker.Init(container); err != nil {
		t.Fatalf("initialize operation maker: %v", err)
	}
	container.Add(maker)
	if _, err := maker.Make(); err == nil {
		t.Fatal("unknown operation was accepted")
	}
}
