package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/matryer/is"
)

func TestRootCommand(t *testing.T) {
	is := is.New(t)

	inputReader, _, _ := createBufferedReaderAndWriter()
	outputReader, outputWriter, _ := createBufferedReaderAndWriter()

	rootCommand := GenerateRootCommand(inputReader, outputWriter, outputWriter)
	err := rootCommand.Execute()
	is.NoErr(err)

	err = outputWriter.Flush()
	is.NoErr(err)

	outputLines := readAllLines(outputReader)

	is.Equal(
		outputLines[0],
		"Split and combine secrets using Shamir's Secret Sharing algorithm.",
	)
}

func TestSplitCommand(t *testing.T) {
	is := is.New(t)

	secret := "SayHelloToMyLittleFriend\n"
	sharesCount := 3
	thresholdCount := 2

	inputReader, _, inputBuffer := createBufferedReaderAndWriter()
	outputReader, outputWriter, _ := createBufferedReaderAndWriter()

	_, err := inputBuffer.WriteString(secret)
	is.NoErr(err)

	// https://github.com/manifoldco/promptui/issues/63
	_, err = padBuffer(4096-len(secret), inputBuffer)
	is.NoErr(err)

	rootCommand := GenerateRootCommand(inputReader, outputWriter, outputWriter)
	rootCommand.SetArgs(
		[]string{
			"split",
			"-n",
			fmt.Sprint(sharesCount),
			"-t",
			fmt.Sprint(thresholdCount),
		},
	)

	err = rootCommand.Execute()
	is.NoErr(err)

	err = outputWriter.Flush()
	is.NoErr(err)

	outputLines := readAllLines(outputReader)
	shares := outputLines[len(outputLines)-sharesCount:]

	for i := range shares {
		shares[i] = cleanAnsiEscapeSequence(shares[i])
	}

	shareLength := len(shares[0])
	for _, share := range shares {
		is.Equal(len(share), shareLength)
	}
}

func TestCombineCommand(t *testing.T) {
	is := is.New(t)

	shares := []string{
		"67442ef838a57cbc3063a487d7ca861cf490b9026f5f3a41be\n",
		"9ef082cd4f3456dc4bf161460a7cd5f580ed1fd426fa3ff5d7\n",
	}
	thresholdCount := len(shares)

	inputReader, _, inputBuffer := createBufferedReaderAndWriter()
	outputReader, outputWriter, _ := createBufferedReaderAndWriter()

	for _, share := range shares {
		_, err := inputBuffer.WriteString(share)
		is.NoErr(err)

		// https://github.com/manifoldco/promptui/issues/63
		_, err = padBuffer(4096-len(share), inputBuffer)
		is.NoErr(err)
	}

	rootCommand := GenerateRootCommand(inputReader, outputWriter, outputWriter)
	rootCommand.SetArgs(
		[]string{
			"combine",
			"-t",
			fmt.Sprint(thresholdCount),
		},
	)

	err := rootCommand.Execute()
	is.NoErr(err)

	err = outputWriter.Flush()
	is.NoErr(err)

	outputLines := readAllLines(outputReader)
	lastLine := outputLines[len(outputLines)-1]
	cleanedLastLine := cleanAnsiEscapeSequence(lastLine)

	is.Equal(cleanedLastLine, "SayHelloToMyLittleFriend")
}

func createBufferedReaderAndWriter() (*bufio.Reader, *bufio.Writer, *bytes.Buffer) {
	buffer := bytes.NewBuffer(nil)
	reader := bufio.NewReader(buffer)
	writer := bufio.NewWriter(buffer)

	return reader, writer, buffer
}

func readAllLines(reader *bufio.Reader) []string {
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func padBuffer(byteCount int, buffer *bytes.Buffer) (n int, err error) {
	padding := make([]byte, byteCount)

	for i := 0; i < byteCount; i++ {
		padding[i] = 97 // a
	}

	return buffer.Write(padding)
}

func cleanAnsiEscapeSequence(text string) string {
	return strings.ReplaceAll(text, "\x1b[?25h", "")
}
