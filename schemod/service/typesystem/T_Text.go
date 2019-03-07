package typesystem

import (
	"unicode/utf8"
)

// T_Text implements the type: Text
type T_Text struct {
	v []rune
}

// ID returns the type identifier
func (o T_Text) ID() IDType {
	return PT__Text
}

// NewText constructs a new instance of type Text
func NewText(s []byte) T_Text {
	return T_Text{}
}

func (txt T_Text) Decode(b []byte) error {
	return nil
}

func (txt T_Text) Encode() []byte {
	size := 0
	for _, r := range txt.v {
		size += utf8.RuneLen(r)
	}

	bs := make([]byte, size)

	count := 0
	for _, r := range txt.v {
		count += utf8.EncodeRune(bs[count:], r)
	}

	return bs
}
