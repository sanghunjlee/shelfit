package shelfit

import (
	"github.com/google/uuid"
)

func newUUID() string {
	return uuid.NewString()
}
