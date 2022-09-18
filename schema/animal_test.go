package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnimal(t *testing.T) {
	t.Run("test animal object interface methods", func(t *testing.T) {
		animal := &Animal{
			Name:    "Scooby",
			ID:      "8534ee57-2560-469c-baba-1584279d0b4e",
			Type:    "dog",
			OwnerID: "2334ee57-2560-469c-baba-1584279d0b4e",
		}

		kind := animal.GetKind()
		assert.Equal(t, "*schema.Animal", kind)

		id := animal.GetID()
		assert.Equal(t, "8534ee57-2560-469c-baba-1584279d0b4e", id)

		name := animal.GetName()
		assert.Equal(t, "Scooby", name)

		animal.SetID("dummy-id")
		assert.Equal(t, "dummy-id", animal.ID)

		animal.SetName("dummy-name")
		assert.Equal(t, "dummy-name", animal.Name)
	})
}
