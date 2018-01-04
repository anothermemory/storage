// Package storage contains common interfaces and implementation for all supported storage implementations.
package storage

import (
	"github.com/anothermemory/unit"
)

// Storage represents some storage which can be used to store units
type Storage interface {
	// SaveUnit saves given unit to the storage.
	SaveUnit(u unit.Unit) error

	// RemoveUnit removes given unit from the storage.
	RemoveUnit(u unit.Unit) error

	// LoadUnit loads unit with given id from the storage
	LoadUnit(id string) (unit.Unit, error)

	// IsCreated checks whether storage was already created or not
	IsCreated() bool

	// Create creates storage
	Create() error

	// Remove removes storage
	Remove() error
}
