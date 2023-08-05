package snip

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/gertd/go-pluralize"
)

type Snipper struct {
	Head       int
	Tail       int
	SkipSnip   bool
	LineNumber bool
}

var plural = pluralize.NewClient()

func New() *Snipper {
	return &Snipper{}
}

func (s *Snipper) printLineNumber(out *bufio.Writer, count int) {
	if !s.LineNumber {
		return
	}
	if count < s.Head {
		out.WriteString(fmt.Sprintf("%6d  ", count+1))
	} else if s.Tail > 0 {
		out.WriteString(fmt.Sprintf("%6d  ", count+1))
	}
}

func (s *Snipper) isInvalid() error {
	if s.Head < 0 || s.Tail < 0 {
		return fmt.Errorf("head=%d, tail=%d: both of head and tail must be positive", s.Head, s.Tail)
	}
	return nil
}

func (s *Snipper) PerformEach(fileName string, writer io.Writer) error {
	if err := s.isInvalid(); err != nil {
		return err
	}
	f, err := openFile(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	lines := []string{}
	count := 0
	scanner := bufio.NewScanner(f)
	out := bufio.NewWriter(writer)
	for scanner.Scan() {
		text := scanner.Text()
		s.printHead(out, text, count)
		lines = s.updateLines(lines, text)
		count++
	}
	s.printSnip(out, count)
	s.printTail(out, lines, count)
	out.Flush()
	return scanner.Err()
}

func (s *Snipper) updateLines(lines []string, newLine string) []string {
	if s.Tail > 0 {
		lines = append(lines, newLine)
		if len(lines) > s.Tail {
			lines = lines[1:]
		}
	}
	return lines
}

func (s *Snipper) printSnip(out *bufio.Writer, count int) {
	if !s.SkipSnip {
		snippedCount := count - s.Head - s.Tail
		if snippedCount > 0 {
			out.WriteString(fmt.Sprintf("        ... snip ... (%s)\n", plural.Pluralize("line", snippedCount, true)))
		}
	}
}

func (s *Snipper) printHead(out *bufio.Writer, line string, count int) {
	if s.Head > 0 && count < s.Head {
		s.printLineNumber(out, count)
		out.WriteString(line)
		out.WriteByte('\n')
	}
}

func normalize(lines []string, tail int) []string {
	for i := len(lines); i < tail; i++ {
		// insert empty lines to index 0.
		lines, lines[0] = append(lines[:1], lines[0:]...), ""
	}
	return lines
}

func (s *Snipper) printTail(out *bufio.Writer, lines []string, count int) {
	index := 0
	lines = normalize(lines, s.Tail)
	if s.Head > count { // no more printing line.
		return
	} else if s.Head+s.Tail < count {
		count = count - s.Tail
	} else if s.Head+s.Tail >= count { // print last TAIL lines.
		index = len(lines) - count + s.Head
		count = s.Head
	}
	for i := index; i < len(lines); i++ {
		s.printLineNumber(out, count)
		out.WriteString(lines[i])
		out.WriteByte('\n')
		count++
	}
	out.Flush()
}

func openFile(fileName string) (io.ReadCloser, error) {
	if fileName == "" {
		return os.Stdin, nil
	}
	return os.Open(fileName)
}
