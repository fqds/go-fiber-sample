package exception

import "fmt"

type NotFoundError struct {
	ID string
}

func (err NotFoundError) Error() string {
	return fmt.Sprintf("object with id %s not found", err.ID)
}
