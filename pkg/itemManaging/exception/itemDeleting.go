package exception

import "fmt"

type ItemDeleting struct {
	ItemID uint64
}

func (e *ItemDeleting) Error() string {
	return fmt.Sprintf("deleting item id: %d failed", e.ItemID)
}
