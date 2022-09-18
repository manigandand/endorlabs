package schema

import (
	"reflect"
	"time"
)

// Person implements Object interface
type Person struct {
	Name      string    `json:"name"`
	ID        string    `json:"id"`
	LastName  string    `json:"last_name"`
	Birthday  string    `json:"birth_day"`
	BirthDate time.Time `json:"birth_date"`
}

// GetKind returns the object kind
func (p *Person) GetKind() string {
	return reflect.TypeOf(p).String()
}

// GetID returns the id of the object
func (p *Person) GetID() string {
	return p.ID
}

// GetName returns the name of the object
func (p *Person) GetName() string {
	return p.Name
}

// SetID replace the object id
func (p *Person) SetID(s string) {
	p.ID = s
}

// SetName replace the object name
func (p *Person) SetName(s string) {
	p.Name = s
}
