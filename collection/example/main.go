package main

import (
	"fmt"

	"github.com/shandialamp/pippi/collection"
)

type User struct {
	Name   string
	Age    int
	Active bool
	Score  float64
}

func main() {
	fmt.Println("=== Collection 示例 ===\n")

	// 示例 1: 基础操作
	fmt.Println("1. 基础操作:")
	numbers := collection.New(1, 2, 3, 4, 5)
	fmt.Printf("   原始集合: %v\n", numbers.All())
	fmt.Printf("   数量: %d\n", numbers.Count())
	first, _ := numbers.First()
	fmt.Printf("   第一个元素: %d\n", first)
	last, _ := numbers.Last()
	fmt.Printf("   最后一个元素: %d\n\n", last)

	// 示例 2: 过滤和映射
	fmt.Println("2. 过滤和映射:")
	evens := numbers.Filter(func(n int) bool { return n%2 == 0 })
	fmt.Printf("   偶数: %v\n", evens.All())
	doubled := collection.Map(evens, func(n int) int { return n * 2 })
	fmt.Printf("   偶数翻倍: %v\n\n", doubled.All())

	// 示例 3: 链式调用
	fmt.Println("3. 链式调用:")
	result := collection.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).
		Filter(func(n int) bool { return n%2 == 0 }).
		Skip(1).
		Take(3)
	fmt.Printf("   结果: %v\n\n", result.All())

	// 示例 4: 聚合操作
	fmt.Println("4. 聚合操作:")
	nums := collection.New(10, 20, 30, 40, 50)
	sum := collection.Sum(nums, func(n int) float64 { return float64(n) })
	avg := collection.Avg(nums, func(n int) float64 { return float64(n) })
	fmt.Printf("   总和: %.0f\n", sum)
	fmt.Printf("   平均值: %.0f\n\n", avg)

	// 示例 5: 用户数据处理
	fmt.Println("5. 用户数据处理:")
	users := collection.New(
		User{"Alice", 25, true, 85.5},
		User{"Bob", 30, false, 70.0},
		User{"Charlie", 25, true, 92.0},
		User{"David", 35, true, 78.5},
		User{"Eve", 30, false, 88.0},
	)

	// 活跃用户
	activeUsers := users.Filter(func(u User) bool { return u.Active })
	fmt.Printf("   活跃用户数量: %d\n", activeUsers.Count())

	// 平均分数
	avgScore := collection.Avg(activeUsers, func(u User) float64 { return u.Score })
	fmt.Printf("   活跃用户平均分数: %.2f\n", avgScore)

	// 按年龄分组
	groups := collection.GroupBy(users, func(u User) int { return u.Age })
	fmt.Printf("   按年龄分组: %d 个组\n", len(groups))
	for age, group := range groups {
		fmt.Printf("     年龄 %d: %d 人\n", age, group.Count())
	}

	// 最高分用户
	topUser, _ := collection.Max(users, func(u User) float64 { return u.Score })
	fmt.Printf("   最高分用户: %s (%.2f)\n\n", topUser.Name, topUser.Score)

	// 示例 6: 集合运算
	fmt.Println("6. 集合运算:")
	c1 := collection.New(1, 2, 3, 4, 5)
	c2 := collection.New(4, 5, 6, 7, 8)
	diff := collection.Diff(c1, c2)
	intersect := collection.Intersect(c1, c2)
	union := collection.Union(c1, c2)
	fmt.Printf("   差集: %v\n", diff.All())
	fmt.Printf("   交集: %v\n", intersect.All())
	fmt.Printf("   并集: %v\n\n", union.All())

	// 示例 7: 去重
	fmt.Println("7. 去重:")
	duplicates := collection.New(1, 2, 2, 3, 3, 3, 4, 4, 4, 4)
	unique := collection.Unique(duplicates)
	fmt.Printf("   原始: %v\n", duplicates.All())
	fmt.Printf("   去重后: %v\n\n", unique.All())

	// 示例 8: 分块
	fmt.Println("8. 分块:")
	items := collection.New(1, 2, 3, 4, 5, 6, 7, 8, 9)
	chunks := collection.Chunk(items, 3)
	fmt.Printf("   分块结果:\n")
	for i, chunk := range chunks {
		fmt.Printf("     块 %d: %v\n", i+1, chunk)
	}

	// 示例 9: 分区
	fmt.Println("\n9. 分区:")
	allNumbers := collection.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	evenNums, oddNums := allNumbers.Partition(func(n int) bool { return n%2 == 0 })
	fmt.Printf("   偶数: %v\n", evenNums.All())
	fmt.Printf("   奇数: %v\n\n", oddNums.All())

	// 示例 10: JSON 序列化
	fmt.Println("10. JSON 序列化:")
	data := collection.New(1, 2, 3, 4, 5)
	jsonStr, _ := data.ToJSON()
	fmt.Printf("   JSON: %s\n", jsonStr)
	restored, _ := collection.FromJSON[int](jsonStr)
	fmt.Printf("   从 JSON 恢复: %v\n", restored.All())
}
