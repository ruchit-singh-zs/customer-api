package errors

import "fmt"

type EntityNotFound struct {
	ID     string
	Entity string
}

func (e EntityNotFound) Error() string {
	return fmt.Sprintf("entity %s with ID %s not found", e.Entity, e.ID)
}
