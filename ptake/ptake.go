package ptake

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func performLine(taker Taker, reader io.Reader, out *bufio.Writer) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		result := taker.TakeLine(line)
		if result {
			fmt.Fprintln(out, line)
		} else {
			break
		}
	}
	return nil
}

func performByte(taker Taker, reader io.Reader, out *bufio.Writer) error {
	buffer := make([]byte, 256)
	printBuffer := bytes.NewBuffer([]byte{})
	for {
		length, err := reader.Read(buffer)
		if length == 0 && err == io.EOF {
			return nil
		}
		for i := 0; i < length; i++ {
			if taker.TakeByte(buffer[i]) {
				printBuffer.Write([]byte{buffer[i]})
			} else {
				out.Write(printBuffer.Bytes())
				return nil
			}
		}
		if err != nil && err != io.EOF {
			return err
		}
		out.Write(printBuffer.Bytes())
		printBuffer.Reset()
	}
}

func performStream(taker Taker, reader io.Reader, out *bufio.Writer) error {
	var f func(Taker, io.Reader, *bufio.Writer) error
	if taker.IsLine() {
		f = performLine
	} else {
		f = performByte
	}
	err := f(taker, reader, out)
	out.Flush()
	return err
}

func PerformEach(taker Taker, arg string, out *bufio.Writer) error {
	if arg == "" {
		return performStream(taker, os.Stdin, out)
	}
	file, err := os.Open(arg)
	if err != nil {
		return err
	}
	defer file.Close()
	return performStream(taker, file, out)
}
