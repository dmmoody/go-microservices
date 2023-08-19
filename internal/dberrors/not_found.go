package dberrors

import "fmt"

type NotFoundError struct {
	Entity string
	Id     string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("unable to find %s with id %s", e.Entity, e.Id)
}
