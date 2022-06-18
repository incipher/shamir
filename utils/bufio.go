package utils

import (
	"bufio"
	"io"
	"strings"
)

// Reads lines from the given Reader.
func ReadLines(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var lines []string
	var line string

	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}

	return lines
}

// Returns an io.ReadCloser with a no-op Close method wrapping the provided reader.
func NopReadCloser(reader io.Reader) io.ReadCloser {
	return io.NopCloser(reader)
}

// Returns an io.WriteCloser with a no-op Close method wrapping the provided writer.
func NopWriteCloser(writer io.Writer) io.WriteCloser {
	return &WriteCloser{Writer: writer}
}

func (writeCloser *WriteCloser) Close() error {
	// Noop
	return nil
}

type WriteCloser struct {
	io.Writer
}
