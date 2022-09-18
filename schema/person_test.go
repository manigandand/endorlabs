package schema

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPerson(t *testing.T) {
	t.Run("test person object interface methods", func(t *testing.T) {
		bt, err := time.Parse("2006-01-02T15:04:05.000Z", "1993-07-22T05:04:05.000Z")
		assert.Equal(t, nil, err)

		person := &Person{
			Name:      "Manigandan Dharmalingam",
			ID:        "8534ee57-2560-469c-baba-1584279d0b4e",
			LastName:  "Dharmalingam",
			Birthday:  "22/07/1993",
			BirthDate: bt,
		}

		kind := person.GetKind()
		assert.Equal(t, "*schema.Person", kind)

		id := person.GetID()
		assert.Equal(t, "8534ee57-2560-469c-baba-1584279d0b4e", id)

		name := person.GetName()
		assert.Equal(t, "Manigandan Dharmalingam", name)

		person.SetID("dummy-id")
		assert.Equal(t, "dummy-id", person.ID)

		person.SetName("dummy-name")
		assert.Equal(t, "dummy-name", person.Name)
	})
}
