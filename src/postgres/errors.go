package postgres

import "fmt"

func NewNonExistentError(model string, method string, ID string) error {
	return fmt.Errorf("failed to %s a %s... there are no %ss with the ID %s", method, model, model, ID)
}
