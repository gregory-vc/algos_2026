package task10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/gregory-vc/algos_2026/task_10/functional"
)

type person struct {
	name   string
	number int
	err    error
}

func Process(path string) (map[int][]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return GroupPeople(strings.Split(string(data), "\n"))
}

func GroupPeople(lines []string) (map[int][]string, error) {
	fields := functional.Map(lines, strings.Fields)
	numbered := functional.Filter(fields, func(parts []string) bool {
		return len(parts) == 2
	})
	people := functional.ParallelMap(numbered, parsePerson)

	invalid := functional.Filter(people, func(p person) bool {
		return p.err != nil
	})
	if len(invalid) > 0 {
		return nil, fmt.Errorf("invalid number for %q: %w", invalid[0].name, invalid[0].err)
	}

	return functional.GroupBy(people, func(p person) (int, string) {
		return p.number, p.name
	}), nil
}

func parsePerson(parts []string) person {
	number, err := strconv.Atoi(parts[1])
	return person{name: normalizeName(parts[0]), number: number, err: err}
}

func normalizeName(name string) string {
	lower := strings.ToLower(name)
	first, size := utf8.DecodeRuneInString(lower)
	return string(unicode.ToUpper(first)) + lower[size:]
}
