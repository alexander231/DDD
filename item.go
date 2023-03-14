package tavern

import "github.com/google/uuid"

// Item represents an Item for all subdomains
type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}
