package ulidx

import (
	"math/rand"

	"github.com/oklog/ulid"
)

var entropy = ulid.Monotonic(rand.New(rand.NewSource(int64(ulid.Now()))), 0)

func New() string {
	return ulid.MustNew(ulid.Now(), entropy).String()
}
