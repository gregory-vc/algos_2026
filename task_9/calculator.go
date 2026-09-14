package task9

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Model struct {
	X, Y, Result int
	Operation    string
}

type DataReader interface {
	Read() (Model, error)
}

type ResultPrinter interface {
	Print(Model) error
}

type Operation interface {
	Symbol() string
	Apply(int, int) int
}

type Container struct {
	services []any
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) Add(service any) {
	c.services = append(c.services, service)
}

type PlusOperation struct{}

func (PlusOperation) Symbol() string            { return "+" }
func (PlusOperation) Apply(left, right int) int { return left + right }

type MinusOperation struct{}

func (MinusOperation) Symbol() string            { return "-" }
func (MinusOperation) Apply(left, right int) int { return left - right }

type OperationMaker struct {
	reader     DataReader
	printer    ResultPrinter
	operations map[string]Operation
}

func (m *OperationMaker) Init(container *Container) error {
	m.reader = nil
	m.printer = nil
	m.operations = make(map[string]Operation)

	for _, service := range container.services {
		if reader, ok := service.(DataReader); ok {
			m.reader = reader
		}
		if printer, ok := service.(ResultPrinter); ok {
			m.printer = printer
		}
		if operation, ok := service.(Operation); ok {
			if _, exists := m.operations[operation.Symbol()]; exists {
				return fmt.Errorf("initialize operation maker: duplicate operation %q", operation.Symbol())
			}
			m.operations[operation.Symbol()] = operation
		}
	}

	if m.reader == nil {
		return errors.New("initialize operation maker: data reader is not registered")
	}
	if m.printer == nil {
		return errors.New("initialize operation maker: result printer is not registered")
	}
	if len(m.operations) == 0 {
		return errors.New("initialize operation maker: no operations registered")
	}
	return nil
}

func (m *OperationMaker) Make() (Model, error) {
	model, err := m.reader.Read()
	if err != nil {
		return Model{}, fmt.Errorf("read model: %w", err)
	}
	operation, ok := m.operations[model.Operation]
	if !ok {
		return Model{}, fmt.Errorf("unknown operation %q", model.Operation)
	}
	model.Result = operation.Apply(model.X, model.Y)
	if err := m.printer.Print(model); err != nil {
		return Model{}, fmt.Errorf("print result: %w", err)
	}
	return model, nil
}

type TextDataReader struct {
	scanner *bufio.Scanner
}

func NewTextDataReader(reader io.Reader) *TextDataReader {
	return &TextDataReader{scanner: bufio.NewScanner(reader)}
}

func (r *TextDataReader) Read() (Model, error) {
	if !r.scanner.Scan() {
		if err := r.scanner.Err(); err != nil {
			return Model{}, err
		}
		return Model{}, io.EOF
	}
	parts := strings.Fields(r.scanner.Text())
	if len(parts) != 3 {
		return Model{}, fmt.Errorf("expected: operation left right")
	}
	left, err := strconv.Atoi(parts[1])
	if err != nil {
		return Model{}, fmt.Errorf("parse left operand: %w", err)
	}
	right, err := strconv.Atoi(parts[2])
	if err != nil {
		return Model{}, fmt.Errorf("parse right operand: %w", err)
	}
	return Model{X: left, Y: right, Operation: parts[0]}, nil
}

type TextResultPrinter struct {
	writer io.Writer
}

func NewTextResultPrinter(writer io.Writer) *TextResultPrinter {
	return &TextResultPrinter{writer: writer}
}

func (p *TextResultPrinter) Print(model Model) error {
	_, err := fmt.Fprintf(p.writer, "%d%s%d=%d\n", model.X, model.Operation, model.Y, model.Result)
	return err
}
