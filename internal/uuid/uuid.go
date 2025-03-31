package uuid

import (
	"github.com/google/uuid"
)

type UUID [16]byte

type UuidStrategy interface {
	generate() UUID
}

type uuidStrategy struct{}

func NewUuidStrategy() UuidStrategy {
	return &uuidStrategy{}
}

func (*uuidStrategy) generate() UUID {
	return UUID(uuid.New())
}
