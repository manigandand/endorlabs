package schema

import "reflect"

// Animal implements Object interface
type Animal struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	OwnerID string `json:"owner_id"`
}

// GetKind returns the object kind
func (p *Animal) GetKind() string {
	return reflect.TypeOf(p).String()
}

// GetID returns the id of the object
func (p *Animal) GetID() string {
	return p.ID
}

// GetName returns the name of the object
func (p *Animal) GetName() string {
	return p.Name
}

// SetID replace the object id
func (p *Animal) SetID(s string) {
	p.ID = s
}

// SetName replace the object name
func (p *Animal) SetName(s string) {
	p.Name = s
}
