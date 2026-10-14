package task13

import (
	"errors"
	"fmt"
)

type Operation func(int, int) (int, error)

type Calculator struct {
	events     *Events
	operations map[string]Operation
}

func NewCalculator(events *Events) *Calculator {
	calculator := &Calculator{
		events: events,
		operations: map[string]Operation{
			"+": func(x, y int) (int, error) { return x + y, nil },
			"-": func(x, y int) (int, error) { return x - y, nil },
		},
	}
	events.CalculationRequested.Subscribe(calculator.calculate)
	return calculator
}

func (c *Calculator) AddOperation(symbol string, operation Operation) error {
	if symbol == "" {
		return errors.New("operation symbol is empty")
	}
	if operation == nil {
		return errors.New("operation is nil")
	}
	c.operations[symbol] = operation
	return nil
}

func (c *Calculator) calculate(request Request) {
	result := CalculationResult{Request: request}
	operation, ok := c.operations[request.Operation]
	if !ok {
		result.Err = fmt.Errorf("unknown operation %q", request.Operation)
	} else {
		result.Value, result.Err = operation(request.X, request.Y)
	}
	c.events.CalculationCompleted.Publish(result)
}
