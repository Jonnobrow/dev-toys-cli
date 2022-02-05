package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFromStdin() (string, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}

	if info.Mode()&os.ModeCharDevice == os.ModeCharDevice {
		return "", fmt.Errorf("Invalid input device for stdin")
	}

	scanner := bufio.NewScanner(os.Stdin)
	var output []string

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if scanner.Err() != nil {
		return "", fmt.Errorf("Failed to read data from stdin: %v\n", err)
	}

	return strings.Join(output, "\n"), nil
}

func WriteToStdout(data string) error {
	info, err := os.Stdout.Stat()
	if err != nil {
		return err
	}

	if info.Mode()&os.ModeCharDevice == os.ModeCharDevice {
		return fmt.Errorf("Invalid input device for stdin")
	}

	writer := bufio.NewWriter(os.Stdout)
	_, err = writer.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}
