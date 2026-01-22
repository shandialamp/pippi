# Collection

一个功能强大的 Go 泛型集合库，灵感来自 Laravel Collection。

## 特性

- 🚀 完全使用 Go 泛型实现
- 📦 丰富的集合操作方法（70+ 方法）
- 🔗 支持链式调用
- 🎯 类型安全
- ⚡ 高性能
- 📝 完整的测试覆盖

## 安装

```bash
go get github.com/shandialamp/pippi/collection
```

## 快速开始

```go
package main

import (
    "fmt"
    "github.com/shandialamp/pippi/collection"
)

func main() {
    // 创建集合
    c := collection.New(1, 2, 3, 4, 5)
    
    // 过滤偶数
    evens := c.Filter(func(n int) bool {
        return n%2 == 0
    })
    
    // 映射转换
    doubled := collection.Map(evens, func(n int) int {
        return n * 2
    })
    
    fmt.Println(doubled.All()) // [4, 8]
}
```

## 核心方法

### 创建集合

```go
// 直接创建
c := collection.New(1, 2, 3, 4, 5)

// 从切片创建
slice := []int{1, 2, 3}
c := collection.FromSlice(slice)

// 从 JSON 创建
c, err := collection.FromJSON[int](`[1, 2, 3]`)
```

### 基础操作

```go
c := collection.New(1, 2, 3, 4, 5)

// 获取所有元素
all := c.All()                    // [1, 2, 3, 4, 5]

// 获取数量
count := c.Count()                // 5

// 检查是否为空
isEmpty := c.IsEmpty()            // false

// 获取第一个元素
first, ok := c.First()            // 1, true

// 获取最后一个元素
last, ok := c.Last()              // 5, true

// 根据索引获取
item, ok := c.Get(2)              // 3, true
```

### 添加和移除

```go
c := collection.New(1, 2, 3)

// 向末尾添加
c.Push(4, 5)                      // [1, 2, 3, 4, 5]

// 向开头添加
c.Prepend(0)                      // [0, 1, 2, 3, 4, 5]

// 移除最后一个
last, ok := c.Pop()               // 5, true

// 移除第一个
first, ok := c.Shift()            // 0, true
```

### 过滤和映射

```go
c := collection.New(1, 2, 3, 4, 5)

// 过滤
evens := c.Filter(func(n int) bool {
    return n%2 == 0
})                                // [2, 4]

// 排除
odds := c.Reject(func(n int) bool {
    return n%2 == 0
})                                // [1, 3, 5]

// 映射
doubled := collection.Map(c, func(n int) int {
    return n * 2
})                                // [2, 4, 6, 8, 10]

// 遍历
c.Each(func(n int) {
    fmt.Println(n)
})
```

### 聚合操作

```go
c := collection.New(1, 2, 3, 4, 5)

// 求和
sum := collection.Sum(c, func(n int) float64 {
    return float64(n)
})                                // 15.0

// 平均值
avg := collection.Avg(c, func(n int) float64 {
    return float64(n)
})                                // 3.0

// 最小值
min, ok := collection.Min(c, func(n int) float64 {
    return float64(n)
})                                // 1, true

// 最大值
max, ok := collection.Max(c, func(n int) float64 {
    return float64(n)
})                                // 5, true

// 归约
product := collection.Reduce(c, func(acc, n int) int {
    return acc * n
}, 1)                             // 120
```

### 切片操作

```go
c := collection.New(1, 2, 3, 4, 5)

// 切片
sliced := c.Slice(1, 3)           // [2, 3]

// 取前 n 个
taken := c.Take(3)                // [1, 2, 3]

// 取后 n 个
takenFromEnd := c.Take(-2)        // [4, 5]

// 跳过前 n 个
skipped := c.Skip(2)              // [3, 4, 5]

// 分块
chunks := collection.Chunk(c, 2)  // [[1, 2], [3, 4], [5]]
```

### 排序和随机

```go
c := collection.New(5, 2, 8, 1, 9)

// 排序
sorted := collection.Sort(c, func(a, b int) bool {
    return a < b
})                                // [1, 2, 5, 8, 9]

// 降序排序
desc := collection.SortDesc(c, func(a, b int) bool {
    return a < b
})                                // [9, 8, 5, 2, 1]

// 反转
reversed := c.Reverse()           // [9, 1, 8, 2, 5]

// 打乱
shuffled := c.Shuffle()           // 随机顺序

// 随机取一个
random, ok := c.Random()          // 随机元素
```

### 去重和查找

```go
c := collection.New(1, 2, 2, 3, 3, 3)

// 去重
unique := collection.Unique(c)    // [1, 2, 3]

// 包含检查
contains := collection.Contains(c, 2)  // true

// 条件包含
hasEven := c.ContainsFunc(func(n int) bool {
    return n%2 == 0
})                                // true

// 所有元素满足条件
allEven := c.Every(func(n int) bool {
    return n%2 == 0
})                                // false

// 至少一个元素满足条件
someEven := c.Some(func(n int) bool {
    return n%2 == 0
})                                // true
```

### 分组和分区

```go
type Person struct {
    Name string
    Age  int
}

people := collection.New(
    Person{"Alice", 25},
    Person{"Bob", 30},
    Person{"Charlie", 25},
)

// 分组
groups := collection.GroupBy(people, func(p Person) int {
    return p.Age
})
// map[25:[{Alice 25}, {Charlie 25}] 30:[{Bob 30}]]

// 分区
c := collection.New(1, 2, 3, 4, 5)
evens, odds := c.Partition(func(n int) bool {
    return n%2 == 0
})
// evens: [2, 4], odds: [1, 3, 5]
```

### 集合运算

```go
c1 := collection.New(1, 2, 3, 4)
c2 := collection.New(3, 4, 5, 6)

// 差集
diff := collection.Diff(c1, c2)   // [1, 2]

// 交集
intersect := collection.Intersect(c1, c2)  // [3, 4]

// 并集
union := collection.Union(c1, c2) // [1, 2, 3, 4, 5, 6]

// 合并
merged := c1.Merge(c2)            // [1, 2, 3, 4, 3, 4, 5, 6]
```

### 扁平化

```go
// 扁平化二维数组
nested := collection.New(
    []int{1, 2},
    []int{3, 4},
    []int{5},
)
flat := collection.Flatten(nested)  // [1, 2, 3, 4, 5]

// FlatMap
c := collection.New(1, 2, 3)
flatMapped := collection.FlatMap(c, func(n int) []int {
    return []int{n, n * 2}
})                                  // [1, 2, 2, 4, 3, 6]
```

### 条件执行

```go
c := collection.New(1, 2, 3)

// 当条件为真时执行
result := c.When(true, func(col *collection.Collection[int]) *collection.Collection[int] {
    return col.Push(4)
})                                  // [1, 2, 3, 4]

// 当条件为假时执行
result := c.Unless(false, func(col *collection.Collection[int]) *collection.Collection[int] {
    return col.Push(4)
})                                  // [1, 2, 3, 4]

// Tap（执行回调但返回原集合）
c.Tap(func(col *collection.Collection[int]) {
    fmt.Println("Count:", col.Count())
})                                  // 打印后返回原集合

// Pipe（传递到回调并返回结果）
count := collection.Pipe(c, func(col *collection.Collection[int]) int {
    return col.Count()
})                                  // 3
```

### JSON 序列化

```go
c := collection.New(1, 2, 3, 4, 5)

// 转为 JSON
jsonStr, err := c.ToJSON()        // "[1,2,3,4,5]"

// 从 JSON 创建
c2, err := collection.FromJSON[int](jsonStr)
```

### 链式调用

```go
result := collection.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).
    Filter(func(n int) bool { return n%2 == 0 }).  // 偶数
    Skip(1).                                        // 跳过第一个
    Take(3).                                        // 取前三个
    Reverse()                                       // 反转

fmt.Println(result.All())  // [8, 6, 4]
```

## 完整示例

### 用户数据处理

```go
package main

import (
    "fmt"
    "github.com/shandialamp/pippi/collection"
)

type User struct {
    ID     int
    Name   string
    Age    int
    Active bool
    Score  float64
}

func main() {
    users := collection.New(
        User{1, "Alice", 25, true, 85.5},
        User{2, "Bob", 30, false, 70.0},
        User{3, "Charlie", 25, true, 92.0},
        User{4, "David", 35, true, 78.5},
        User{5, "Eve", 30, false, 88.0},
    )

    // 获取所有活跃用户的平均分数
    activeUsers := users.Filter(func(u User) bool {
        return u.Active
    })
    
    avgScore := collection.Avg(activeUsers, func(u User) float64 {
        return u.Score
    })
    
    fmt.Printf("活跃用户平均分数: %.2f\n", avgScore)

    // 按年龄分组
    groups := collection.GroupBy(users, func(u User) int {
        return u.Age
    })
    
    for age, group := range groups {
        names := collection.Map(group, func(u User) string {
            return u.Name
        })
        fmt.Printf("年龄 %d: %v\n", age, names.All())
    }

    // 获取分数最高的用户
    topUser, ok := collection.Max(users, func(u User) float64 {
        return u.Score
    })
    
    if ok {
        fmt.Printf("最高分用户: %s (%.2f)\n", topUser.Name, topUser.Score)
    }

    // 获取活跃用户的名字列表
    activeNames := collection.Map(activeUsers, func(u User) string {
        return u.Name
    })
    
    nameList := collection.Join(activeNames, ", ", func(name string) string {
        return name
    })
    
    fmt.Printf("活跃用户: %s\n", nameList)
}
```

### 电商订单处理

```go
package main

import (
    "fmt"
    "github.com/shandialamp/pippi/collection"
)

type Product struct {
    Name     string
    Price    float64
    Category string
    Stock    int
}

func main() {
    products := collection.New(
        Product{"iPhone", 999.0, "Electronics", 10},
        Product{"iPad", 599.0, "Electronics", 5},
        Product{"Book", 29.9, "Books", 100},
        Product{"Pen", 2.5, "Stationery", 200},
        Product{"Laptop", 1299.0, "Electronics", 8},
    )

    // 按类别分组并计算总价值
    groups := collection.GroupBy(products, func(p Product) string {
        return p.Category
    })
    
    for category, items := range groups {
        totalValue := collection.Sum(items, func(p Product) float64 {
            return p.Price * float64(p.Stock)
        })
        fmt.Printf("%s 类别总价值: $%.2f\n", category, totalValue)
    }

    // 找出低库存产品（库存 < 10）
    lowStock := products.Filter(func(p Product) bool {
        return p.Stock < 10
    })
    
    fmt.Println("\n低库存产品:")
    lowStock.Each(func(p Product) {
        fmt.Printf("- %s: %d 件\n", p.Name, p.Stock)
    })

    // 计算所有产品的平均价格
    avgPrice := collection.Avg(products, func(p Product) float64 {
        return p.Price
    })
    
    fmt.Printf("\n平均价格: $%.2f\n", avgPrice)

    // 获取价格最高的 3 个产品
    topProducts := collection.Sort(products, func(a, b Product) bool {
        return a.Price > b.Price // 降序
    }).Take(3)
    
    fmt.Println("\n价格最高的 3 个产品:")
    topProducts.Each(func(p Product) {
        fmt.Printf("- %s: $%.2f\n", p.Name, p.Price)
    })
}
```

## API 参考

### 创建方法
- `New[T](items ...T)` - 创建新集合
- `FromSlice[T](slice []T)` - 从切片创建
- `FromJSON[T](jsonStr string)` - 从 JSON 创建

### 基础方法
- `All()` - 获取所有元素
- `Count()` - 元素数量
- `IsEmpty()` - 是否为空
- `IsNotEmpty()` - 是否不为空
- `First()` - 第一个元素
- `Last()` - 最后一个元素
- `Get(index)` - 根据索引获取

### 修改方法
- `Push(items...)` - 末尾添加
- `Pop()` - 移除最后一个
- `Shift()` - 移除第一个
- `Prepend(items...)` - 开头添加

### 转换方法
- `Filter(fn)` - 过滤
- `Reject(fn)` - 排除
- `Map(c, fn)` - 映射
- `Reduce(c, fn, initial)` - 归约
- `FlatMap(c, fn)` - 映射并扁平化
- `Pluck(c, fn)` - 提取字段

### 切片方法
- `Slice(start, end)` - 切片
- `Take(n)` - 取前 n 个
- `Skip(n)` - 跳过前 n 个
- `Chunk(size)` - 分块

### 排序方法
- `Sort(c, less)` - 排序
- `SortDesc(c, less)` - 降序排序
- `Reverse()` - 反转
- `Shuffle()` - 打乱

### 查找方法
- `Contains(c, value)` - 包含检查
- `ContainsFunc(fn)` - 条件包含
- `Every(fn)` - 所有满足
- `Some(fn)` - 至少一个满足
- `Random()` - 随机元素

### 集合运算
- `Unique(c)` - 去重
- `Diff(c1, c2)` - 差集
- `Intersect(c1, c2)` - 交集
- `Union(c1, c2)` - 并集
- `Merge(others...)` - 合并

### 分组方法
- `GroupBy(c, fn)` - 分组
- `Partition(fn)` - 分区

### 聚合方法
- `Sum(c, fn)` - 求和
- `Avg(c, fn)` - 平均值
- `Min(c, fn)` - 最小值
- `Max(c, fn)` - 最大值

### 工具方法
- `Each(fn)` - 遍历
- `EachWithIndex(fn)` - 带索引遍历
- `Tap(fn)` - 执行回调
- `Pipe(c, fn)` - 管道传递
- `When(condition, fn)` - 条件执行
- `Unless(condition, fn)` - 条件执行（反向）
- `Clone()` - 克隆
- `ToJSON()` - 转 JSON
- `String()` - 转字符串

## 性能建议

1. **避免不必要的复制**：大多数方法返回新集合，如果需要修改原集合，使用修改类方法（Push, Pop 等）

2. **链式调用**：链式调用可以减少中间变量，代码更简洁

3. **使用合适的方法**：
   - 需要查找：使用 `ContainsFunc` 而不是 `Filter` + `Count`
   - 需要条件判断：使用 `Every/Some` 而不是 `Filter` + `Count`

4. **大数据集**：对于大数据集，考虑使用流式处理或分批处理
