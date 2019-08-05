package mssqlx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHostConnection(t *testing.T) {
	assert := assert.New(t)
	// assert equality
	c1 := Connection{Computer: "D30"}
	cs, err := c1.String()
	assert.NoError(err)
	assert.Equal("sqlserver://D30", cs)
}
