package uuid

import (
	"github.com/google/uuid"
)

type UUID [16]byte

type UuidGenerator struct{}

func (*UuidGenerator) generate() UUID {
	return UUID(uuid.New())
}
