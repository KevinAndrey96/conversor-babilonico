package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const exitCommand = "salir"

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run asks for Babylonian numbers until the user types "salir" or the input
// ends. Invalid entries produce an alert and the number is requested again.
func run(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)

	for {
		if _, err := fmt.Fprint(out, "Número babilónico (ej. 1;30.15) o 'salir': "); err != nil {
			return fmt.Errorf("prompt: %w", err)
		}

		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())
		if strings.EqualFold(line, exitCommand) {
			break
		}

		value, err := ParseSexagesimal(line)
		if err != nil {
			if _, werr := fmt.Fprintf(out, "Alerta: %v\n", err); werr != nil {
				return fmt.Errorf("alert: %w", werr)
			}

			continue
		}

		if _, werr := fmt.Fprintf(out, "Valor decimal: %g\n", value); werr != nil {
			return fmt.Errorf("result: %w", werr)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("read input: %w", err)
	}

	return nil
}
