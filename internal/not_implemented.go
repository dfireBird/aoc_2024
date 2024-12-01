package internal

import (
	"fmt"
)

type NotImplemented struct{}

func (e *NotImplemented) Error() string {
	return fmt.Sprintf("not implemented")
}
