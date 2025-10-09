package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectory(t *testing.T) {
	err := MakeDirectory("./a/b")
	assert.Nil(t, err)
	dirs, err := Directorys("./a")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(dirs))
	err = DeleteDirectory("./a/b")
	assert.Nil(t, err)
	dirs, err = Directorys("./a")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(dirs))
	err = DeleteDirectory("./a")
	assert.Nil(t, err)
}
