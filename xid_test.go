package xid

import (
	"bufio"
	"compress/gzip"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"unicode"
)

func readMatches(file string) ([]bool, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	rdr, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer rdr.Close()
	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanBytes)
	var matches []bool
	for scanner.Scan() {
		b, err := strconv.ParseBool(scanner.Text())
		if err != nil {
			return nil, err
		}
		matches = append(matches, b)
	}
	return matches, nil
}

func TestExhaustive(t *testing.T) {
	cases := []struct {
		class string
		f     func(rune) bool
	}{
		{"xid_start", Start},
		{"xid_continue", Continue},
	}
	for _, c := range cases {
		t.Run(c.class, func(t *testing.T) {
			matches, err := readMatches(filepath.Join("testdata", c.class+unicodeTestVersion+".txt.gz"))
			if err != nil {
				t.Fatal(err)
			}
			for r := rune(0); r <= unicode.MaxRune; r++ {
				want := matches[r]
				if got := c.f(r); got != want {
					t.Fatalf("%s(%s)=%v, got=%v", c.class, strconv.QuoteRuneToASCII(r), want, got)
				}
			}
		})
	}
}
