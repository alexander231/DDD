// Package aggregates holds aggregates that combines many entities into a full object
package aggregate

import (
	"github.com/alexander231/DDD/entity"
	"github.com/alexander231/DDD/valueobject"
)

// Customer is an aggregate that combines all entities needed to represent a customer
type Customer struct {
	// person is the root entity of a customer
	// which means the person.ID is the main identifier for this aggregate
	person *entity.Person
	// a customer can hold many products
	products []*entity.Item
	// a customer can perform many transactions
	transactions []valueobject.Transaction
}
