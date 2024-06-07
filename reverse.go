package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

func XXDReverse(r io.Reader, w io.Writer) error {
	var (
		cols int
		octs int
		char = make([]byte, 1)
	)

	if opts.Columns != -1 {
		cols = opts.Columns
	}

	switch dumpType {
	case dumpBinary:
		octs = 8
	case dumpCformat:
		octs = 4
	default:
		octs = 2
	}

	if opts.Len != -1 && opts.Len < int64(cols) {
		cols = int(opts.Len)
	}

	if octs < 1 {
		octs = cols
	}

	c := int64(0)
	rd := bufio.NewReader(r)

	for {
		line, err := rd.ReadBytes('\n')
		if err != nil && !errors.Is(err, io.EOF) {
			return fmt.Errorf("hexxy: %v", err)
		}

		if len(line) == 0 {
			return nil
		}

		n := len(line)
		i := 0

		for i < n {
			if i+octs > n {
				break
			}

			switch dumpType {
			case dumpHex, dumpCformat:
				if rv, _ := hexDecode(char, line[i:i+octs]); rv == 0 {
					w.Write(char)
					i += octs
					c++
				} else if rv == -1 {
					i++
				} else {
					i += 2
				}
			case dumpBinary:
				if binaryDecode(char, line[i:i+octs]) == -1 {
					i++
				} else {
					w.Write(char)
					i += octs
					c++
				}
			case dumpPlain:
				if rv, _ := hexDecode(char, line[i:i+octs]); rv != 0 {
					w.Write(char)
					c++
				}
				i++
			}
		}

		if c >= int64(cols) && cols > 0 {
			return nil
		}
	}
}
