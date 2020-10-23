// Package xid implements validation functions for unicode identifiers,
// as defined in UAX#31: https://unicode.org/reports/tr31/.
// The syntax for an identifier is:
//
//     <identifier> := <xid_start> <xid_continue>*
//
// where <xid_start> and <xid_continue> derive from <id_start> and
// <id_continue>, respectively, and check their NFKC normalized forms.
package xid

import (
	"unicode"

	"golang.org/x/text/unicode/norm"
)

type set func(rune) bool

func (a set) add(rt *unicode.RangeTable) set {
	b := in(rt)
	return func(r rune) bool { return a(r) || b(r) }
}

func (a set) sub(rt *unicode.RangeTable) set {
	b := in(rt)
	return func(r rune) bool { return a(r) && !b(r) }
}

func in(rt *unicode.RangeTable) set {
	return func(r rune) bool { return unicode.Is(rt, r) }
}

var id_start = set(unicode.IsLetter).
	add(unicode.Nl).
	add(unicode.Other_ID_Start).
	sub(unicode.Pattern_Syntax).
	sub(unicode.Pattern_White_Space)

var id_continue = id_start.
	add(unicode.Mn).
	add(unicode.Mc).
	add(unicode.Nd).
	add(unicode.Pc).
	add(unicode.Other_ID_Continue).
	sub(unicode.Pattern_Syntax).
	sub(unicode.Pattern_White_Space)

// Start checks that the rune begins an identifier.
func Start(r rune) bool {
	// id_start(r) && NFKC(r) in "id_start xid_continue*"
	if !id_start(r) {
		return false
	}
	s := norm.NFKC.String(string(r))
	if s == "" {
		return false
	}
	for i, r := range s {
		if i == 0 {
			if !id_start(r) {
				return false
			}
		} else {
			if !Continue(r) {
				return false
			}
		}
	}
	return true
}

// Continue checks that the rune continues an identifier.
func Continue(r rune) bool {
	// id_continue(r) && NFKC(r) in "id_continue*"
	if !id_continue(r) {
		return false
	}
	for _, r := range norm.NFKC.String(string(r)) {
		if !id_continue(r) {
			return false
		}
	}
	return true
}
