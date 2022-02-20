package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type StdinMsg string

type ErrorMsg struct{ err error }

func (e ErrorMsg) Error() string {
	return e.err.Error()
}

func ReadFromStdin() tea.Msg {
	hasStdin, err := HasStdin()
	if err != nil {
		return ErrorMsg{err: err}
	} else if !hasStdin {
		return StdinMsg("")
	}

	scanner := bufio.NewScanner(os.Stdin)
	var output []string

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if scanner.Err() != nil {
		return ErrorMsg{fmt.Errorf("Failed to read data from stdin: %v\n", err)}
	}

	return StdinMsg(strings.Join(output, "\n"))
}

func HasStdin() (bool, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}
	if info.Mode()&os.ModeCharDevice == os.ModeCharDevice {
		// No stdin available - not an error state thought
		return false, nil
	}
	return true, nil
}
