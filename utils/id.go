package utils

import (
	srand "crypto/rand"
	"time"

	ulid "github.com/oklog/ulid/v2"
)

var secureSource *ulid.MonotonicEntropy

// SecureID returns a Universally Unique Lexicographically Sortable Identifier
// obtained via crypto/rand entropy
func SecureID() ulid.ULID {
	return ulid.MustNew(ulid.Timestamp(time.Now()), secureSource)
}

// IsIDValid checks if a ULID is valid
func IsIDValid(id string) bool {
	if _, err := ulid.Parse(id); err != nil {
		return false
	}
	return true
}

func init() {
	secureSource = ulid.Monotonic(srand.Reader, 0)
}
