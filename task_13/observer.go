package task13

type Event[T any] struct {
	listeners []func(T)
}

func (e *Event[T]) Subscribe(listener func(T)) {
	if listener != nil {
		e.listeners = append(e.listeners, listener)
	}
}

func (e *Event[T]) Publish(value T) {
	for _, listener := range e.listeners {
		listener(value)
	}
}

type Request struct {
	X, Y      int
	Operation string
}

type CalculationResult struct {
	Request Request
	Value   int
	Err     error
}

type ViewState struct {
	Text string
	Err  error
}

type Events struct {
	InputChanged         Event[string]
	CalculationRequested Event[Request]
	CalculationCompleted Event[CalculationResult]
	StateChanged         Event[ViewState]
}
