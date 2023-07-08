package entity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewID(t *testing.T) {
	id := NewID()

	assert.NotEmpty(t, id)
	assert.NotNil(t, id)
}

func TestParseID(t *testing.T) {
	id, err := ParseID("00000000-0000-0000-0000-000000000000")

	assert.NoError(t, err)
	assert.Equal(t, "00000000-0000-0000-0000-000000000000", id.String())

	// invalid id
	id, err = ParseID("invalid")
	fmt.Println(id)
	assert.Error(t, err)
	assert.Equal(t, "00000000-0000-0000-0000-000000000000", id.String())
}
