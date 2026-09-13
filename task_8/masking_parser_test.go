package task8

import (
	"bufio"
	"os"
	"testing"
)

func TestMaskingParser(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "short and long logins",
			input: "login=romanow login=testadmin",
			want:  "login=r****** login=te***n",
		},
		{
			name:  "two and three part names",
			input: "name=Гогия Владимир; name=Романов Алексей Сергеевич",
			want:  "name=Г**** Владимир; name=Р****** Алексей Се***ч",
		},
		{
			name:  "email and card",
			input: "email=admin@example.com card=4111 1111 1111 1234",
			want:  "email=a***@example.com card=4111 11** **** 1234",
		},
		{
			name:  "card separators",
			input: "card=4111-1111-1111-1234 card=4111111111111234",
			want:  "card=4111-11**-****-1234 card=411111******1234",
		},
		{
			name:  "arbitrary order and repeated fields",
			input: "card=5555 4444 3333 2222 login=user01 email=user01@example.com login=user02",
			want:  "card=5555 44** **** 2222 login=u***** email=u***@example.com login=u*****",
		},
		{
			name:  "unrelated text",
			input: "2026-03-05 INFO request completed",
			want:  "2026-03-05 INFO request completed",
		},
	}

	var parser MaskingParser
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := parser.Mask(test.input); got != test.want {
				t.Fatalf("Mask(%q) = %q, want %q", test.input, got, test.want)
			}
		})
	}
}

func TestMaskingParserGoldenFile(t *testing.T) {
	input := openTestFile(t, "testdata/input.log")
	defer input.Close()
	expected := openTestFile(t, "testdata/result.log")
	defer expected.Close()

	parser := NewMaskingParser()
	inputScanner := bufio.NewScanner(input)
	expectedScanner := bufio.NewScanner(expected)
	lineNumber := 0
	for inputScanner.Scan() {
		lineNumber++
		if !expectedScanner.Scan() {
			t.Fatalf("result.log has no line %d", lineNumber)
		}
		if got, want := parser.Mask(inputScanner.Text()), expectedScanner.Text(); got != want {
			t.Fatalf("line %d: got %q, want %q", lineNumber, got, want)
		}
	}
	if err := inputScanner.Err(); err != nil {
		t.Fatalf("read input.log: %v", err)
	}
	if expectedScanner.Scan() {
		t.Fatalf("result.log contains extra lines after line %d", lineNumber)
	}
	if err := expectedScanner.Err(); err != nil {
		t.Fatalf("read result.log: %v", err)
	}
}

func openTestFile(t *testing.T, name string) *os.File {
	t.Helper()
	file, err := os.Open(name)
	if err != nil {
		t.Fatalf("open %s: %v", name, err)
	}
	return file
}
