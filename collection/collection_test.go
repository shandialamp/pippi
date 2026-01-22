package collection

import (
	"encoding/json"
	"testing"
)

func TestNew(t *testing.T) {
	c := New(1, 2, 3, 4, 5)
	if c.Count() != 5 {
		t.Errorf("Expected count 5, got %d", c.Count())
	}
}

func TestFromSlice(t *testing.T) {
	slice := []int{1, 2, 3}
	c := FromSlice(slice)
	if c.Count() != 3 {
		t.Errorf("Expected count 3, got %d", c.Count())
	}
}

func TestAll(t *testing.T) {
	c := New(1, 2, 3)
	all := c.All()
	if len(all) != 3 || all[0] != 1 || all[1] != 2 || all[2] != 3 {
		t.Errorf("All() failed")
	}
}

func TestIsEmpty(t *testing.T) {
	c := New[int]()
	if !c.IsEmpty() {
		t.Error("Expected collection to be empty")
	}

	c.Push(1)
	if c.IsEmpty() {
		t.Error("Expected collection not to be empty")
	}
}

func TestFirst(t *testing.T) {
	c := New(1, 2, 3)
	first, ok := c.First()
	if !ok || first != 1 {
		t.Errorf("Expected first element to be 1, got %d", first)
	}

	empty := New[int]()
	_, ok = empty.First()
	if ok {
		t.Error("Expected no element in empty collection")
	}
}

func TestLast(t *testing.T) {
	c := New(1, 2, 3)
	last, ok := c.Last()
	if !ok || last != 3 {
		t.Errorf("Expected last element to be 3, got %d", last)
	}
}

func TestPushPop(t *testing.T) {
	c := New(1, 2, 3)
	c.Push(4, 5)
	if c.Count() != 5 {
		t.Errorf("Expected count 5, got %d", c.Count())
	}

	last, ok := c.Pop()
	if !ok || last != 5 {
		t.Errorf("Expected popped element to be 5, got %d", last)
	}
	if c.Count() != 4 {
		t.Errorf("Expected count 4, got %d", c.Count())
	}
}

func TestFilter(t *testing.T) {
	c := New(1, 2, 3, 4, 5)
	filtered := c.Filter(func(n int) bool {
		return n%2 == 0
	})
	if filtered.Count() != 2 {
		t.Errorf("Expected 2 even numbers, got %d", filtered.Count())
	}
}

func TestMap(t *testing.T) {
	c := New(1, 2, 3)
	mapped := Map(c, func(n int) int {
		return n * 2
	})
	all := mapped.All()
	if len(all) != 3 || all[0] != 2 || all[1] != 4 || all[2] != 6 {
		t.Error("Map failed")
	}
}

func TestReduce(t *testing.T) {
	c := New(1, 2, 3, 4, 5)
	sum := Reduce(c, func(acc int, n int) int {
		return acc + n
	}, 0)
	if sum != 15 {
		t.Errorf("Expected sum 15, got %d", sum)
	}
}

func TestUnique(t *testing.T) {
	c := New(1, 2, 2, 3, 3, 3, 4)
	unique := Unique(c)
	if unique.Count() != 4 {
		t.Errorf("Expected 4 unique elements, got %d", unique.Count())
	}
}

func TestContains(t *testing.T) {
	c := New(1, 2, 3, 4, 5)
	if !Contains(c, 3) {
		t.Error("Expected collection to contain 3")
	}
	if Contains(c, 6) {
		t.Error("Expected collection not to contain 6")
	}
}

func TestPartition(t *testing.T) {
	c := New(1, 2, 3, 4, 5)
	evens, odds := c.Partition(func(n int) bool { return n%2 == 0 })
	if evens.Count() != 2 {
		t.Errorf("Expected 2 even numbers, got %d", evens.Count())
	}
	if odds.Count() != 3 {
		t.Errorf("Expected 3 odd numbers, got %d", odds.Count())
	}
}

func TestGroupBy(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	c := New(
		Person{"Alice", 25},
		Person{"Bob", 30},
		Person{"Charlie", 25},
		Person{"David", 30},
	)
	groups := GroupBy(c, func(p Person) int { return p.Age })
	if len(groups) != 2 {
		t.Errorf("Expected 2 groups, got %d", len(groups))
	}
	if groups[25].Count() != 2 {
		t.Errorf("Expected 2 people aged 25, got %d", groups[25].Count())
	}
}

func TestSort(t *testing.T) {
	c := New(5, 2, 8, 1, 9)
	sorted := Sort(c, func(a, b int) bool { return a < b })
	all := sorted.All()
	if len(all) != 5 || all[0] != 1 || all[4] != 9 {
		t.Error("Sort failed")
	}
}

func TestSum(t *testing.T) {
	c := New(1, 2, 3, 4, 5)
	sum := Sum(c, func(n int) float64 { return float64(n) })
	if sum != 15.0 {
		t.Errorf("Expected sum 15.0, got %f", sum)
	}
}

func TestJSON(t *testing.T) {
	c := New(1, 2, 3, 4, 5)
	jsonStr, err := c.ToJSON()
	if err != nil {
		t.Errorf("ToJSON failed: %v", err)
	}

	c2, err := FromJSON[int](jsonStr)
	if err != nil {
		t.Errorf("FromJSON failed: %v", err)
	}
	if c2.Count() != 5 {
		t.Errorf("Expected 5 elements, got %d", c2.Count())
	}
}

func TestChainedOperations(t *testing.T) {
	result := New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).
		Filter(func(n int) bool { return n%2 == 0 }).
		Skip(1).
		Take(2)

	all := result.All()
	if len(all) != 2 || all[0] != 4 || all[1] != 6 {
		t.Errorf("Chained operations failed, got %v", all)
	}
}

func TestStructSerialization(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	c := New(
		Person{"Alice", 25},
		Person{"Bob", 30},
	)

	jsonStr, err := c.ToJSON()
	if err != nil {
		t.Fatalf("ToJSON failed: %v", err)
	}

	var persons []Person
	err = json.Unmarshal([]byte(jsonStr), &persons)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if len(persons) != 2 || persons[0].Name != "Alice" {
		t.Error("Struct serialization failed")
	}
}
