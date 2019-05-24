package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClone(t *testing.T) {
	err := Clone()
	assert.NoError(t, err)
}

func TestPlush(t *testing.T) {
	err := Push()
	assert.NoError(t, err)
}

func TestTags(t *testing.T) {
	err := Tags()
	assert.NoError(t, err)
}
