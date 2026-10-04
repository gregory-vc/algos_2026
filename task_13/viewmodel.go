package task13

import (
	"fmt"
	"strconv"
	"strings"
)

type CalculatorViewModel struct {
	events *Events
	state  ViewState
}

func NewCalculatorViewModel(events *Events) *CalculatorViewModel {
	viewModel := &CalculatorViewModel{events: events}
	events.InputChanged.Subscribe(viewModel.onInput)
	events.CalculationCompleted.Subscribe(viewModel.onResult)
	return viewModel
}

func (vm *CalculatorViewModel) State() ViewState {
	return vm.state
}

func (vm *CalculatorViewModel) onInput(input string) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		vm.showError(fmt.Errorf("expected: operation left right"))
		return
	}
	x, err := strconv.Atoi(parts[1])
	if err != nil {
		vm.showError(fmt.Errorf("parse left operand: %w", err))
		return
	}
	y, err := strconv.Atoi(parts[2])
	if err != nil {
		vm.showError(fmt.Errorf("parse right operand: %w", err))
		return
	}
	vm.events.CalculationRequested.Publish(Request{X: x, Y: y, Operation: parts[0]})
}

func (vm *CalculatorViewModel) onResult(result CalculationResult) {
	if result.Err != nil {
		vm.showError(result.Err)
		return
	}
	request := result.Request
	vm.updateState(ViewState{
		Text: fmt.Sprintf("%d%s%d=%d", request.X, request.Operation, request.Y, result.Value),
	})
}

func (vm *CalculatorViewModel) showError(err error) {
	vm.updateState(ViewState{Text: "error: " + err.Error(), Err: err})
}

func (vm *CalculatorViewModel) updateState(state ViewState) {
	vm.state = state
	vm.events.StateChanged.Publish(state)
}
