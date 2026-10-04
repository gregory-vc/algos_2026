package task13

import (
	"errors"
	"strings"
	"testing"
)

func TestViewModelBindsInputAndResult(t *testing.T) {
	events := &Events{}
	viewModel := NewCalculatorViewModel(events)
	var requests []Request
	events.CalculationRequested.Subscribe(func(request Request) {
		requests = append(requests, request)
	})
	var states []ViewState
	events.StateChanged.Subscribe(func(state ViewState) {
		if viewModel.State() != state {
			t.Fatal("state notification arrived before ViewModel state was updated")
		}
		states = append(states, state)
	})

	events.InputChanged.Publish("  -\t-7   5  ")
	request := Request{X: -7, Y: 5, Operation: "-"}
	if len(requests) != 1 || requests[0] != request {
		t.Fatalf("requests = %+v, want [%+v]", requests, request)
	}
	if len(states) != 0 {
		t.Fatalf("received UI state before calculation completed: %+v", states)
	}

	events.CalculationCompleted.Publish(CalculationResult{Request: request, Value: -12})
	if len(states) != 1 || states[0].Text != "-7-5=-12" || states[0].Err != nil {
		t.Fatalf("states = %+v, want successful state -7-5=-12", states)
	}
}

func TestViewModelRejectsInvalidInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty input", "", "expected: operation left right"},
		{"missing operand", "+ 7", "expected: operation left right"},
		{"extra operand", "+ 7 5 2", "expected: operation left right"},
		{"invalid left operand", "+ text 5", "parse left operand"},
		{"invalid right operand", "+ 7 text", "parse right operand"},
		{"operand out of range", "+ " + strings.Repeat("9", 100) + " 5", "parse left operand"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			events := &Events{}
			viewModel := NewCalculatorViewModel(events)
			events.CalculationRequested.Subscribe(func(Request) {
				t.Fatal("invalid input was sent to the Model")
			})
			var states []ViewState
			events.StateChanged.Subscribe(func(state ViewState) { states = append(states, state) })

			events.InputChanged.Publish(tt.input)

			if len(states) != 1 || states[0].Err == nil || !strings.Contains(states[0].Text, tt.want) {
				t.Fatalf("states = %+v, want one error containing %q", states, tt.want)
			}
			if viewModel.State() != states[0] {
				t.Fatal("ViewModel state differs from published state")
			}
		})
	}
}

func TestViewModelReplacesErrorWithSuccessfulState(t *testing.T) {
	events := &Events{}
	viewModel := NewCalculatorViewModel(events)
	operationError := errors.New("division by zero")
	events.CalculationCompleted.Publish(CalculationResult{Err: operationError})
	state := viewModel.State()
	if !errors.Is(state.Err, operationError) || state.Text != "error: division by zero" {
		t.Fatalf("error state = %+v", state)
	}
	events.CalculationCompleted.Publish(CalculationResult{
		Request: Request{X: 2, Y: 3, Operation: "+"}, Value: 5,
	})
	state = viewModel.State()
	if state.Err != nil || state.Text != "2+3=5" {
		t.Fatalf("state after recovery = %+v, want 2+3=5, no error", state)
	}
}
