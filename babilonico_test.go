package main

import (
	"bytes"
	"errors"
	"math"
	"strings"
	"testing"
)

func TestParseSexagesimal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr error
	}{
		{name: "single digit", input: "5", want: 5},
		{name: "two digits", input: "1;30", want: 90},
		{name: "three digits", input: "1;2;3", want: 3723},
		{name: "fraction", input: "1.30", want: 1.5},
		{name: "integer and fraction", input: "1;30.15", want: 90.25},
		{name: "two fractional digits", input: "1.30;30", want: 1.508333333333333},
		{name: "surrounding spaces", input: "  2;30 ", want: 150},
		{name: "empty", input: "", wantErr: ErrEmptyInput},
		{name: "only spaces", input: "   ", wantErr: ErrEmptyInput},
		{name: "digit too high", input: "60", wantErr: ErrDigitOutOfRange},
		{name: "digit zero", input: "0", wantErr: ErrDigitOutOfRange},
		{name: "not a number", input: "1;a", wantErr: ErrInvalidNotation},
		{name: "two points", input: "1.2.3", wantErr: ErrInvalidNotation},
		{name: "empty fraction", input: "1.", wantErr: ErrInvalidNotation},
		{name: "empty digit", input: "1;;2", wantErr: ErrInvalidNotation},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ParseSexagesimal(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("error = %v, want %v", err, tt.wantErr)
			}

			if tt.wantErr == nil && math.Abs(got-tt.want) > 1e-9 {
				t.Fatalf("value = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRunPromptsAgainAfterAnInvalidEntry(t *testing.T) {
	t.Parallel()

	var out bytes.Buffer

	if err := run(strings.NewReader("99\n1;30\nsalir\n"), &out); err != nil {
		t.Fatalf("run: %v", err)
	}

	text := out.String()
	if !strings.Contains(text, "Alerta: integer part: digit out of range") {
		t.Errorf("missing alert in output: %q", text)
	}

	if !strings.Contains(text, "Valor decimal: 90") {
		t.Errorf("missing result in output: %q", text)
	}

	if got := strings.Count(text, "Número babilónico"); got != 3 {
		t.Errorf("prompts = %d, want 3", got)
	}
}
