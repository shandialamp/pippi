package collection

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c1 := []int{1, 2, 3}

	collect := New(c1)
	c2 := collect.ToSlice()
	assert.Equal(t, c1, c2)
}

func TestMap(t *testing.T) {
	c1 := []int{1, 2, 3}
	collect1 := New(c1)
	collect2 := Map(collect1, func(item int) string {
		return strconv.Itoa(item)
	})
	c2 := collect2.ToSlice()
	assert.Equal(t, len(c1), len(c2))
	assert.Equal(t, "1", c2[0])
	assert.Equal(t, "2", c2[1])
	assert.Equal(t, "3", c2[2])
}

func TestPush(t *testing.T) {
	c1 := []int{1, 2, 3}
	collect := New(c1)
	collect.Push(4)
	c2 := collect.ToSlice()
	assert.Equal(t, 4, len(c2))
	assert.Equal(t, 4, c2[3])
}

func TestPull(t *testing.T) {
	c1 := []int{1, 2, 3, 4}
	collect := New(c1)
	v := collect.Pull()
	c2 := collect.ToSlice()
	assert.Equal(t, 3, len(c2))
	assert.Equal(t, 4, v)
}

func TestEach(t *testing.T) {
	c1 := []int{1, 2, 3, 4}
	collect := New(c1)
	sum := 0
	collect.Each(func(i int) {
		sum += i
	})
	assert.Equal(t, 10, sum)
}

func TestFilter(t *testing.T) {
	c1 := []int{1, 2, 3, 4}
	collect1 := New(c1)
	collect2 := collect1.Filter(func(i int) bool {
		return i == 4
	})
	assert.Equal(t, 1, collect2.Size())
}
