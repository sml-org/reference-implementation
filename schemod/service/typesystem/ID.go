package typesystem

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"regexp"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

var identifierPattern = regexp.MustCompile("^[0-9a-f]{32}$")

// ID implements the type: ID
// it represents a 16 byte universally unique identifier
type ID struct {
	v [16]byte
}

// NewID creates a new universally unique identifier
func NewID() (id ID) {
	copy(id.v[:], uuid.NewV4().Bytes())
	return id
}

// String returns the textual representation of the identifier
func (id ID) String() string {
	return hex.EncodeToString(id.v[:])
}

// FromString parses the identifier from a hex encoded 32 char string
func (id ID) FromString(str string) (err error) {
	if !identifierPattern.MatchString(str) {
		return fmt.Errorf(
			"the given identifier (%s) doesn't match the string pattern",
			str,
		)
	}
	bytes, err := hex.DecodeString(str)
	if err != nil {
		return fmt.Errorf("couldn't decode hex string to bytes: %s", err)
	}
	copy(id.v[:], bytes)
	return nil
}

// MarshalJSON implements the Go JSON interface
func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(id.v[:]))
}

// UnmarshalJSON implements the Go JSON interface
func (id ID) UnmarshalJSON(bytes []byte) (err error) {
	if len(bytes) != 34 {
		return errors.Errorf(
			"invalid identifier (wrong length: %d)",
			len(bytes),
		)
	}
	return id.FromString(string(bytes[1:33]))
}

// Serialize serializes the value to the given byte stream
func (o ID) Serialize(
	byteOrder binary.ByteOrder,
	stream io.Writer,
) error {
	_, err := stream.Write(o.v[:])
	return err
}
