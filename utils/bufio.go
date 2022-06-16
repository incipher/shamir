package utils

import (
	"io"
)

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
