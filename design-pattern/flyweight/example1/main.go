package main

import (
	"fmt"
	"strings"
	"unicode"
)

type FormattedText struct {
	plainText  string
	capitalize []bool
}

func (f *FormattedText) String() string {
	sb := strings.Builder{}
	for i, c := range f.plainText {
		if f.capitalize[i] {
			sb.WriteRune(unicode.ToUpper(rune(c)))
		} else {
			sb.WriteRune(rune(c))
		}
	}
	return sb.String()
}

func (f *FormattedText) Capitalize(start, end int) {
	for i := start; i <= end; i++ {
		f.capitalize[i] = true
	}
}

func NewFormattedText(plainText string) *FormattedText {
	return &FormattedText{
		plainText:  plainText,
		capitalize: make([]bool, len(plainText)),
	}
}

type TextRange struct {
	Start, End               int
	Capitalize, Bold, Italic bool
}

func (t *TextRange) Covers(position int) bool {
	return position >= t.Start && position <= t.End
}

//Sharing Text
type BetterFormattedText struct {
	plainText  string
	formatting []*TextRange
}

func NewBetterFormattedText(plainText string) *BetterFormattedText {
	return &BetterFormattedText{
		plainText: plainText,
	}
}

func (b *BetterFormattedText) Range(start, end int) *TextRange {
	r := &TextRange{start, end, false, false, false}
	b.formatting = append(b.formatting, r)
	return r
}

func (b *BetterFormattedText) String() string {
	sb := strings.Builder{}

	for i, c := range b.plainText {
		for _, r := range b.formatting {
			if r.Covers(i) && r.Capitalize {
				sb.WriteRune(unicode.ToUpper(rune(c)))
			} else {
				sb.WriteRune(rune(c))
			}
		}
	}
	return sb.String()
}

func main() {
	text := "This is a beautiful world"
	f := NewFormattedText(text)
	f.Capitalize(10, 18)
	fmt.Println(f.String())

	f2 := NewBetterFormattedText(text)
	f2.Range(10, 18).Capitalize = true
	fmt.Println(f2.String())
}
