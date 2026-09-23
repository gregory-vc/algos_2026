package task10

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func TestProcessExample(t *testing.T) {
	got, err := Process("testdata/people.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := map[int][]string{5: {"Вася", "Аня"}, 3: {"Петя"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	t.Logf("Processed file: %v", got)
}

func TestGroupPeople(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  map[int][]string
	}{
		{
			name:  "duplicates are preserved",
			lines: []string{"вася 5", "ВАСЯ 5", "вася 5", "АНЯ 5"},
			want:  map[int][]string{5: {"Вася", "Вася", "Вася", "Аня"}},
		},
		{
			name:  "same name with different numbers",
			lines: []string{"вася 5", "ВАСЯ 3", "Петя 3"},
			want:  map[int][]string{5: {"Вася"}, 3: {"Вася", "Петя"}},
		},
		{
			name:  "Unicode and compound names",
			lines: []string{"ёЖ 1", "я 1", "ÉLODIE 1", "АННА-МАРИЯ 1", "о'КОННОР 1"},
			want:  map[int][]string{1: {"Ёж", "Я", "Élodie", "Анна-мария", "О'коннор"}},
		},
		{
			name:  "spaces tabs and CRLF",
			lines: []string{"  ваСЯ\t5\r", "\tАНЯ    5  ", "\u00a0ПЕТЯ\u00a03\u00a0", "  Тото  ", "\t\r"},
			want:  map[int][]string{5: {"Вася", "Аня"}, 3: {"Петя"}},
		},
		{
			name:  "lines without exactly two fields are skipped",
			lines: []string{"", "   ", "Тото", "5", "АНЯ 5 лишнее", "ПЕТЯ 3"},
			want:  map[int][]string{3: {"Петя"}},
		},
		{
			name:  "numbers are compared numerically",
			lines: []string{"вася 005", "аня +5", "петя 0", "тото -3"},
			want:  map[int][]string{5: {"Вася", "Аня"}, 0: {"Петя"}, -3: {"Тото"}},
		},
		{
			name:  "no people with numbers",
			lines: []string{"Вася", "Петя", "", "\t"},
			want:  map[int][]string{},
		},
		{
			name: "empty input",
			want: map[int][]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := append([]string(nil), tt.lines...)
			got, err := GroupPeople(tt.lines)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.lines, original) {
				t.Fatalf("input lines were modified: %v", tt.lines)
			}
		})
	}
}

func TestGroupPeopleInvalidNumber(t *testing.T) {
	tests := []struct {
		number string
		want   error
	}{
		{number: "abc", want: strconv.ErrSyntax},
		{number: "3.5", want: strconv.ErrSyntax},
		{number: strings.Repeat("9", 100), want: strconv.ErrRange},
	}

	for _, tt := range tests {
		t.Run(tt.number, func(t *testing.T) {
			got, err := GroupPeople([]string{"Вася 5", "АНЯ " + tt.number})
			if !errors.Is(err, tt.want) {
				t.Fatalf("got error %v, want %v", err, tt.want)
			}
			if got != nil {
				t.Fatalf("expected no partial result on error, got %v", got)
			}
		})
	}
}

func TestProcessFiles(t *testing.T) {
	tests := []struct {
		name string
		text string
		want map[int][]string
	}{
		{name: "empty file", want: map[int][]string{}},
		{
			name: "CRLF without trailing newline",
			text: "вася 5\r\nПетя 3\r\nАНЯ 5\r\nТото",
			want: map[int][]string{5: {"Вася", "Аня"}, 3: {"Петя"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "people.txt")
			if err := os.WriteFile(path, []byte(tt.text), 0600); err != nil {
				t.Fatal(err)
			}
			got, err := Process(path)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessMissingFile(t *testing.T) {
	got, err := Process(filepath.Join(t.TempDir(), "missing.txt"))
	if !errors.Is(err, os.ErrNotExist) || got != nil {
		t.Fatalf("got (%v, %v), expected file not found error", got, err)
	}
}

func TestProcessInvalidNumber(t *testing.T) {
	path := filepath.Join(t.TempDir(), "people.txt")
	if err := os.WriteFile(path, []byte("АНЯ abc"), 0600); err != nil {
		t.Fatal(err)
	}
	got, err := Process(path)
	if !errors.Is(err, strconv.ErrSyntax) || got != nil {
		t.Fatalf("got (%v, %v), expected number parsing error", got, err)
	}
}

func TestGroupPeopleParallelOrder(t *testing.T) {
	lines := make([]string, 1001)
	want := make(map[int][]string)
	for i := range lines {
		number := i % 7
		lines[i] = fmt.Sprintf("ИМЯ%d %d", i, number)
		want[number] = append(want[number], fmt.Sprintf("Имя%d", i))
	}

	original := runtime.GOMAXPROCS(0)
	t.Cleanup(func() { runtime.GOMAXPROCS(original) })
	for _, workers := range []int{1, 2, 4} {
		t.Run(fmt.Sprintf("GOMAXPROCS=%d", workers), func(t *testing.T) {
			runtime.GOMAXPROCS(workers)
			got, err := GroupPeople(lines)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatal("group contents or name order depend on worker count")
			}
		})
	}
}

func TestProcessOutput(t *testing.T) {
	groups, err := Process("testdata/people.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := "map[3:[Петя] 5:[Вася Аня]]"
	if got := fmt.Sprint(groups); got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}
