package mssqlx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMSSQLXObject(t *testing.T) {
	assert := assert.New(t)
	mssqlx, err := New("mssql")
	assert.NoError(err)
	assert.NotNil(mssqlx)
}

func TestMSSQLXOpen(t *testing.T) {
	assert := assert.New(t)
	mssqlx, err := New("mssql")
}
