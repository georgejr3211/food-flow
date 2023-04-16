package valueobject

import (
	"github.com/google/uuid"
)

type ProductID struct {
	value string
}

func NewProductId(id string) *ProductID {
	if id == "" {
		id = uuid.New().String()
	}

	return &ProductID{
		value: id,
	}
}

func (p ProductID) Value() string {
	return p.value
}

func (p ProductID) Equals(other ProductID) bool {
	return p.value == other.value
}
