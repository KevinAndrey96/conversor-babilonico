// Package main implements a converter from Babylonian (sexagesimal) numbers
// to decimal.
package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	base       = 60
	minDigit   = 1
	maxDigit   = 59
	digitSep   = ";"
	pointSep   = "."
	maxPartLen = 64
)

var (
	// ErrEmptyInput reports an input without any digit.
	ErrEmptyInput = errors.New("empty input")
	// ErrInvalidNotation reports an input that does not follow the notation.
	ErrInvalidNotation = errors.New("invalid notation")
	// ErrDigitOutOfRange reports a digit outside the range 1 to 59.
	ErrDigitOutOfRange = errors.New("digit out of range")
)

// ParseSexagesimal converts a Babylonian number written as digits separated by
// ";" and an optional fractional part after "." into its decimal value.
// Each digit must be between 1 and 59. Digits before the point are weighted by
// decreasing powers of 60 ending in 60^0, and digits after the point by
// 60^-1, 60^-2 and so on.
func ParseSexagesimal(input string) (float64, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return 0, ErrEmptyInput
	}

	parts := strings.Split(input, pointSep)
	if len(parts) > 2 {
		return 0, fmt.Errorf("point separator: %w", ErrInvalidNotation)
	}

	integer, err := parseDigits(parts[0])
	if err != nil {
		return 0, fmt.Errorf("integer part: %w", err)
	}

	var result float64

	for i, digit := range integer {
		result += float64(digit) * math.Pow(base, float64(len(integer)-1-i))
	}

	if len(parts) == 1 {
		return result, nil
	}

	fraction, err := parseDigits(parts[1])
	if err != nil {
		return 0, fmt.Errorf("fractional part: %w", err)
	}

	for i, digit := range fraction {
		result += float64(digit) * math.Pow(base, -float64(i+1))
	}

	return result, nil
}

// parseDigits splits a part by ";" and validates each digit.
func parseDigits(part string) ([]int, error) {
	if strings.TrimSpace(part) == "" || len(part) > maxPartLen {
		return nil, ErrInvalidNotation
	}

	fields := strings.Split(part, digitSep)
	digits := make([]int, 0, len(fields))

	for _, field := range fields {
		digit, err := strconv.Atoi(strings.TrimSpace(field))
		if err != nil {
			return nil, ErrInvalidNotation
		}

		if digit < minDigit || digit > maxDigit {
			return nil, ErrDigitOutOfRange
		}

		digits = append(digits, digit)
	}

	return digits, nil
}
