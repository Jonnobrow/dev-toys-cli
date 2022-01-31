package helpers

import (
	"github.com/atotto/clipboard"
)

func WriteToClipboard(data string) error {
	return clipboard.WriteAll(data)
}

func ReadFromClipboard() (string, error) {
	return clipboard.ReadAll()
}
