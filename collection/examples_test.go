package collection_test

import (
	"fmt"

	"github.com/shandialamp/pippi/collection"
)

// Example_basic 基础使用示例
func Example_basic() {
	c := collection.New(1, 2, 3, 4, 5)
	
	// 过滤偶数
	evens := c.Filter(func(n int) bool {
		return n%2 == 0
	})
	
	fmt.Println(evens.All())
	// Output: [2 4]
}

// Example_chaining 链式调用示例
func Example_chaining() {
	result := collection.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).
		Filter(func(n int) bool { return n%2 == 0 }). // 过滤偶数
		Skip(1).                                       // 跳过第一个
		Take(3)                                        // 取前三个
	
	fmt.Println(result.All())
	// Output: [4 6 8]
}

// Example_userProcessing 用户数据处理示例
func Example_userProcessing() {
	type User struct {
		Name   string
		Age    int
		Active bool
		Score  float64
	}
	
	users := collection.New(
		User{"Alice", 25, true, 85.5},
		User{"Bob", 30, false, 70.0},
		User{"Charlie", 25, true, 92.0},
		User{"David", 35, true, 78.5},
	)
	
	// 获取活跃用户的平均分数
	activeUsers := users.Filter(func(u User) bool {
		return u.Active
	})
	
	avgScore := collection.Avg(activeUsers, func(u User) float64 {
		return u.Score
	})
	
	fmt.Printf("%.2f\n", avgScore)
	// Output: 85.33
}

// Example_groupBy 分组示例
func Example_groupBy() {
	type Person struct {
		Name string
		Age  int
	}
	
	people := collection.New(
		Person{"Alice", 25},
		Person{"Bob", 30},
		Person{"Charlie", 25},
		Person{"David", 30},
	)
	
	// 按年龄分组
	groups := collection.GroupBy(people, func(p Person) int {
		return p.Age
	})
	
	// 按年龄排序输出
	ages := []int{25, 30}
	for _, age := range ages {
		if group, ok := groups[age]; ok {
			names := collection.Map(group, func(p Person) string {
				return p.Name
			})
			fmt.Printf("Age %d: %d people\n", age, names.Count())
		}
	}
	// Output:
	// Age 25: 2 people
	// Age 30: 2 people
}

// Example_aggregation 聚合操作示例
func Example_aggregation() {
	numbers := collection.New(10, 20, 30, 40, 50)
	
	sum := collection.Sum(numbers, func(n int) float64 {
		return float64(n)
	})
	
	avg := collection.Avg(numbers, func(n int) float64 {
		return float64(n)
	})
	
	min, _ := collection.Min(numbers, func(n int) float64 {
		return float64(n)
	})
	
	max, _ := collection.Max(numbers, func(n int) float64 {
		return float64(n)
	})
	
	fmt.Printf("Sum: %.0f\n", sum)
	fmt.Printf("Avg: %.0f\n", avg)
	fmt.Printf("Min: %d\n", min)
	fmt.Printf("Max: %d\n", max)
	// Output:
	// Sum: 150
	// Avg: 30
	// Min: 10
	// Max: 50
}

// Example_setOperations 集合运算示例
func Example_setOperations() {
	c1 := collection.New(1, 2, 3, 4, 5)
	c2 := collection.New(4, 5, 6, 7, 8)
	
	diff := collection.Diff(c1, c2)
	intersect := collection.Intersect(c1, c2)
	union := collection.Union(c1, c2)
	
	fmt.Println("Diff:", diff.All())
	fmt.Println("Intersect:", intersect.All())
	fmt.Println("Union:", union.All())
	// Output:
	// Diff: [1 2 3]
	// Intersect: [4 5]
	// Union: [1 2 3 4 5 6 7 8]
}

// Example_partition 分区示例
func Example_partition() {
	numbers := collection.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	
	evens, odds := numbers.Partition(func(n int) bool {
		return n%2 == 0
	})
	
	fmt.Println("Evens:", evens.All())
	fmt.Println("Odds:", odds.All())
	// Output:
	// Evens: [2 4 6 8 10]
	// Odds: [1 3 5 7 9]
}

// Example_flatMap FlatMap 示例
func Example_flatMap() {
	numbers := collection.New(1, 2, 3)
	
	result := collection.FlatMap(numbers, func(n int) []int {
		return []int{n, n * 10}
	})
	
	fmt.Println(result.All())
	// Output: [1 10 2 20 3 30]
}

// Example_unique 去重示例
func Example_unique() {
	numbers := collection.New(1, 2, 2, 3, 3, 3, 4, 4, 4, 4)
	
	unique := collection.Unique(numbers)
	
	fmt.Println(unique.All())
	// Output: [1 2 3 4]
}

// Example_sort 排序示例
func Example_sort() {
	numbers := collection.New(5, 2, 8, 1, 9, 3)
	
	sorted := collection.Sort(numbers, func(a, b int) bool {
		return a < b
	})
	
	fmt.Println(sorted.All())
	// Output: [1 2 3 5 8 9]
}

// Example_chunk 分块示例
func Example_chunk() {
	numbers := collection.New(1, 2, 3, 4, 5, 6, 7)
	
	chunks := collection.Chunk(numbers, 3)

	for _, chunk := range chunks {
		fmt.Println(chunk)
	}
}

// Example_reduce 归约示例
func Example_reduce() {
	numbers := collection.New(1, 2, 3, 4, 5)
	
	// 计算乘积
	product := collection.Reduce(numbers, func(acc, n int) int {
		return acc * n
	}, 1)
	
	fmt.Println(product)
	// Output: 120
}

// Example_whenUnless 条件执行示例
func Example_whenUnless() {
	numbers := collection.New(1, 2, 3)

	// 当条件为真时添加元素
	result1 := numbers.Clone().When(true, func(c *collection.Collection[int]) *collection.Collection[int] {
		return c.Push(4)
	})

	// 当条件为假时添加元素
	result2 := numbers.Clone().Unless(false, func(c *collection.Collection[int]) *collection.Collection[int] {
		return c.Push(5)
	})

	fmt.Println(result1.Count())
	fmt.Println(result2.Count())
	// Output:
	// 4
	// 4
}

// Example_ecommerce 电商场景示例
func Example_ecommerce() {
	type Product struct {
		Name     string
		Price    float64
		Category string
		InStock  bool
	}
	
	products := collection.New(
		Product{"iPhone", 999, "Electronics", true},
		Product{"iPad", 599, "Electronics", false},
		Product{"Book", 29.9, "Books", true},
		Product{"Laptop", 1299, "Electronics", true},
	)
	
	// 获取有货的电子产品
	inStockElectronics := products.
		Filter(func(p Product) bool { return p.InStock }).
		Filter(func(p Product) bool { return p.Category == "Electronics" })
	
	// 计算总价值
	totalValue := collection.Sum(inStockElectronics, func(p Product) float64 {
		return p.Price
	})
	
	fmt.Printf("In-stock electronics: %d\n", inStockElectronics.Count())
	fmt.Printf("Total value: $%.0f\n", totalValue)
	// Output:
	// In-stock electronics: 2
	// Total value: $2298
}

// Example_dataTransformation 数据转换示例
func Example_dataTransformation() {
	// 原始数据
	rawData := collection.New("1", "2", "3", "4", "5")
	
	// 转换为整数并过滤
	numbers := collection.Map(rawData, func(s string) int {
		n := 0
		fmt.Sscanf(s, "%d", &n)
		return n
	})
	
	// 只保留偶数
	evens := numbers.Filter(func(n int) bool {
		return n%2 == 0
	})
	
	// 转回字符串
	result := collection.Map(evens, func(n int) string {
		return fmt.Sprintf("%d", n)
	})
	
	// 连接
	joined := collection.Join(result, ", ", func(s string) string {
		return s
	})
	
	fmt.Println(joined)
	// Output: 2, 4
}
