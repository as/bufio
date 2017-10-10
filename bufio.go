package bufio

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// Scanner embeds a bufio scanner and implements
// additional functionality on top of it.
type Scanner struct {
	*bufio.Scanner
}

//  NewScanner returns a new Scanner to read from r.
// The split function defaults to ScanLines.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{bufio.NewScanner(r)}
}

//   Scan advances the Scanner to the next token, which will then be available
//   through the Bytes or Text method. It returns false when the scan stops,
//   either by reaching the end of the input or an error. After Scan returns
//   false, the Err method will return any error that occurred during scanning,
//   except that if it was io.EOF, Err will return nil. Scan panics if the split
//   function returns 100 empty tokens without advancing the input. This is a
//   common error mode for scanners.
func (sc *Scanner) Scan() bool {
	return sc.Scanner.Scan()
}

// Atoi returns the result of ParseInt(s, 10, 0) converted to type int
// for each integer in the scanner's Text(). All elements of the Text() must
// be parsable
func (sc *Scanner) Atoi() ([]int, error) {
	s := strings.Fields(sc.Text())
	ia := make([]int, len(s))
	for i, v := range s {
		w, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ia[i] = w
	}
	return ia, nil
}
