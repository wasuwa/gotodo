package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	ts := All()
	assert.Contains(t, "hello", "hello")
}
