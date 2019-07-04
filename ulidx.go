package ulidx

import (
	"math/rand"
	"sync"

	"github.com/oklog/ulid"
)

var entropy = ulid.Monotonic(rand.New(rand.NewSource(int64(ulid.Now()))), 0)
var mtx sync.Mutex

func New() string {
	mtx.Lock()
	defer mtx.Unlock()
	return ulid.MustNew(ulid.Now(), entropy).String()
}
