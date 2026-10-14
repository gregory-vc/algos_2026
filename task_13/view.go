package task13

import (
	"bufio"
	"fmt"
	"io"
)

type ConsoleView struct {
	events  *Events
	scanner *bufio.Scanner
	writer  io.Writer
	err     error
}

func NewConsoleView(events *Events, reader io.Reader, writer io.Writer) *ConsoleView {
	view := &ConsoleView{
		events:  events,
		scanner: bufio.NewScanner(reader),
		writer:  writer,
	}
	events.StateChanged.Subscribe(view.render)
	return view
}

func (v *ConsoleView) Run() error {
	if v.err != nil {
		return v.err
	}
	for v.scanner.Scan() {
		v.events.InputChanged.Publish(v.scanner.Text())
		if v.err != nil {
			return v.err
		}
	}
	if err := v.scanner.Err(); err != nil {
		return fmt.Errorf("read input: %w", err)
	}
	return nil
}

func (v *ConsoleView) render(state ViewState) {
	if v.err != nil {
		return
	}
	if _, err := fmt.Fprintln(v.writer, state.Text); err != nil {
		v.err = fmt.Errorf("write output: %w", err)
	}
}
