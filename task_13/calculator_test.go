package task13

import (
	"errors"
	"testing"
)

func TestCalculatorPublishesResults(t *testing.T) {
	tests := []struct {
		name    string
		request Request
		want    int
	}{
		{"addition", Request{X: 7, Y: 5, Operation: "+"}, 12},
		{"subtraction", Request{X: 7, Y: 5, Operation: "-"}, 2},
		{"negative operands", Request{X: -7, Y: -5, Operation: "+"}, -12},
		{"zero result", Request{X: 7, Y: 7, Operation: "-"}, 0},
		{"custom operation", Request{X: 7, Y: 5, Operation: "*"}, 35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			events := &Events{}
			model := NewCalculator(events)
			if err := model.AddOperation("*", func(x, y int) (int, error) { return x * y, nil }); err != nil {
				t.Fatal(err)
			}
			var results []CalculationResult
			events.CalculationCompleted.Subscribe(func(result CalculationResult) {
				results = append(results, result)
			})

			events.CalculationRequested.Publish(tt.request)

			if len(results) != 1 {
				t.Fatalf("received %d results, want 1", len(results))
			}
			result := results[0]
			if result.Request != tt.request || result.Value != tt.want || result.Err != nil {
				t.Fatalf("result = %+v, want request %+v, value %d, no error", result, tt.request, tt.want)
			}
		})
	}
}

func TestCalculatorPublishesErrors(t *testing.T) {
	events := &Events{}
	model := NewCalculator(events)
	operationError := errors.New("division by zero")
	if err := model.AddOperation("/", func(x, y int) (int, error) {
		return 0, operationError
	}); err != nil {
		t.Fatal(err)
	}
	var results []CalculationResult
	events.CalculationCompleted.Subscribe(func(result CalculationResult) {
		results = append(results, result)
	})
	for _, symbol := range []string{"?", "/", "+"} {
		events.CalculationRequested.Publish(Request{X: 7, Y: 0, Operation: symbol})
	}
	if len(results) != 3 {
		t.Fatalf("received %d results, want 3", len(results))
	}
	if results[0].Err == nil || results[0].Err.Error() != `unknown operation "?"` {
		t.Fatalf("unknown operation error = %v", results[0].Err)
	}
	if !errors.Is(results[1].Err, operationError) {
		t.Fatalf("operation error = %v, want %v", results[1].Err, operationError)
	}
	if results[2].Err != nil || results[2].Value != 7 {
		t.Fatalf("calculation after errors = %+v, want value 7, no error", results[2])
	}
}

func TestCalculatorRejectsInvalidOperations(t *testing.T) {
	model := NewCalculator(&Events{})
	if err := model.AddOperation("", func(x, y int) (int, error) { return 0, nil }); err == nil {
		t.Fatal("empty operation symbol was accepted")
	}
	if err := model.AddOperation("*", nil); err == nil {
		t.Fatal("nil operation was accepted")
	}
}

func TestEventNotifiesAllSubscribers(t *testing.T) {
	var event Event[int]
	var received []int
	event.Subscribe(nil)
	event.Subscribe(func(value int) { received = append(received, value) })
	event.Subscribe(func(value int) { received = append(received, value*2) })
	event.Publish(7)
	if len(received) != 2 || received[0] != 7 || received[1] != 14 {
		t.Fatalf("received = %v, want [7 14]", received)
	}
}
