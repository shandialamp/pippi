package structx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name *string
	Age  *int
}

func TestWithDefaults(t *testing.T) {
	person1 := &Person{
		Name: PtrString("张三"),
		Age:  nil,
	}
	person2 := &Person{
		Name: PtrString("李四"),
		Age:  PtrInt(18),
	}
	WithDefaults(person1, person2)

	assert.Equal(t, "张三", *person1.Name)
	assert.Equal(t, 18, *person1.Age)
}
