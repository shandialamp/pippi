# Ptr - 指针工具库

Ptr 是一个轻量级的 Go 指针工具库，提供了方便的指针转换函数。

## 特性

- 🎯 **简洁 API** - 类型特定的指针转换函数
- 🔄 **双向转换** - `TypePtr()` 值转指针、`PtrToType()` 指针转值
- 🛡️ **安全操作** - 自动处理 nil 指针，返回零值
- 📦 **轻量设计** - 精简 API，专注核心功能
- ⚡ **高性能** - 直接的指针操作

## 安装

```bash
go get github.com/shandialamp/pippi/ptr
```

## 快速开始

### 整数类型（int64/uint64 为源头）

```go
package main

import (
	"fmt"
	"github.com/shandialamp/pippi/ptr"
)

func main() {
	// int64 作为源头，转换为其他 int 类型的指针
	int8Ptr := ptr.Int64ToInt8Ptr(100)      // int64 → *int8
	int16Ptr := ptr.Int64ToInt16Ptr(1000)   // int64 → *int16
	int32Ptr := ptr.Int64ToInt32Ptr(9999)   // int64 → *int32
	fmt.Println(*int8Ptr, *int16Ptr, *int32Ptr)

	// uint64 作为源头，转换为其他 uint 类型的指针
	uint8Ptr := ptr.Uint64ToUint8Ptr(200)   // uint64 → *uint8
	uint16Ptr := ptr.Uint64ToUint16Ptr(5000) // uint64 → *uint16
	uint32Ptr := ptr.Uint64ToUint32Ptr(50000) // uint64 → *uint32
	fmt.Println(*uint8Ptr, *uint16Ptr, *uint32Ptr)

	// int64 指针转换
	int64Ptr := ptr.Int64Ptr(999999)
	val := ptr.PtrToInt64(int64Ptr) // *int64 → int64
	fmt.Println(val)

	// *int64 → uint64
	uval := ptr.Int64PtrToUint64(int64Ptr)
	fmt.Println(uval)
}
```

### 字符串类型

```go
// 字符串指针转换
strPtr := ptr.StringPtr("hello")
fmt.Println(*strPtr) // "hello"

// 字符串转整数指针
intPtr := ptr.StringToIntPtr("42")
fmt.Println(*intPtr) // 42

// 转换失败返回 nil
intPtr = ptr.StringToIntPtr("invalid")
fmt.Println(intPtr) // nil
```

### 浮点数类型

```go
// float64 指针转换
float64Ptr := ptr.Float64Ptr(3.14)
fmt.Println(*float64Ptr) // 3.14

// float32 指针转换
float32Ptr := ptr.Float32Ptr(3.14)
fmt.Println(*float32Ptr) // 3.14
```

### 布尔类型

```go
// bool 指针转换
boolPtr := ptr.BoolPtr(true)
fmt.Println(*boolPtr) // true
```

## API 文档

### 整数类型函数

**Int64 - 有符号整数源头**
- `Int64Ptr(v int64) *int64` - int64 值 → 指针
- `PtrToInt64(p *int64) int64` - int64 指针 → 值（nil 返回 0）
- `Int64ToInt8Ptr(v int64) *int8` - int64 → *int8 指针
- `Int64ToInt16Ptr(v int64) *int16` - int64 → *int16 指针
- `Int64ToInt32Ptr(v int64) *int32` - int64 → *int32 指针
- `Int64PtrToInt64(p *int64) int64` - *int64 指针 → int64 值（nil 返回 0）
- `Int64PtrToUint64(p *int64) uint64` - *int64 指针 → uint64 值（nil 返回 0）

**Uint64 - 无符号整数源头**
- `Uint64Ptr(v uint64) *uint64` - uint64 值 → 指针
- `PtrToUint64(p *uint64) uint64` - uint64 指针 → 值（nil 返回 0）
- `Uint64ToUint8Ptr(v uint64) *uint8` - uint64 → *uint8 指针
- `Uint64ToUint16Ptr(v uint64) *uint16` - uint64 → *uint16 指针
- `Uint64ToUint32Ptr(v uint64) *uint32` - uint64 → *uint32 指针
- `Uint64PtrToUint64(p *uint64) uint64` - *uint64 指针 → uint64 值（nil 返回 0）

### 字符串类型函数

- `StringPtr(v string) *string` - 值转指针
- `PtrToString(p *string) string` - 指针转值（nil 返回空字符串）
- `StringToIntPtr(v string) *int` - 字符串转 int 指针（失败返回 nil）
- `StringToInt64Ptr(v string) *int64` - 字符串转 int64 指针（失败返回 nil）
- `StringToFloat64Ptr(v string) *float64` - 字符串转 float64 指针（失败返回 nil）

### 浮点数类型函数

- `Float32Ptr(v float32) *float32` / `PtrToFloat32(p *float32) float32`
- `Float64Ptr(v float64) *float64` / `PtrToFloat64(p *float64) float64`

### 布尔类型函数

- `BoolPtr(v bool) *bool` - 值转指针
- `PtrToBool(p *bool) bool` - 指针转值（nil 返回 false）

## 错误处理

- Int64/Uint64 指针：nil 指针返回 `0`
- 字符串类型：nil 指针返回 `""`
- 浮点数类型：nil 指针返回 `0.0`
- 布尔类型：nil 指针返回 `false`
- 字符串转换失败：返回 `nil` 指针

## 常见用例

### 处理可选字段

```go
type User struct {
	ID    int64      // 使用 int64
	Name  *string    // 可选字段
	Email *string
	Age   *int64     // 使用 int64 指针
}

user := User{
	ID:    1,
	Name:  ptr.StringPtr("Alice"),
	Email: ptr.StringPtr("alice@example.com"),
	Age:   ptr.Int64Ptr(30),
}

name := ptr.PtrToString(user.Name)
if name == "" {
	name = "Unknown"
}

age := ptr.PtrToInt64(user.Age)
```

### 数据库模型

```go
type Product struct {
	ID       int64
	Name     string
	Price    *float64  // 可选价格
	Discount *float64  // 可选折扣
}

product := Product{
	ID:       1,
	Name:     "Widget",
	Price:    ptr.Float64Ptr(99.99),
	Discount: ptr.Float64Ptr(0.1),
}

price := ptr.PtrToFloat64(product.Price)
discount := ptr.PtrToFloat64(product.Discount)
```

### 类型转换

```go
// int64 转换为其他 int 类型的指针
val := int64(42)
int8Ptr := ptr.Int64ToInt8Ptr(val)     // int64 → *int8
int16Ptr := ptr.Int64ToInt16Ptr(val)   // int64 → *int16
int32Ptr := ptr.Int64ToInt32Ptr(val)   // int64 → *int32
fmt.Println(*int8Ptr, *int16Ptr, *int32Ptr)

// uint64 转换为其他 uint 类型的指针
uval := uint64(100)
uint8Ptr := ptr.Uint64ToUint8Ptr(uval)    // uint64 → *uint8
uint16Ptr := ptr.Uint64ToUint16Ptr(uval)  // uint64 → *uint16
uint32Ptr := ptr.Uint64ToUint32Ptr(uval)  // uint64 → *uint32
fmt.Println(*uint8Ptr, *uint16Ptr, *uint32Ptr)

// int64 指针转值
int64Ptr := ptr.Int64Ptr(999)
val = ptr.PtrToInt64(int64Ptr)      // *int64 → int64
uval2 := ptr.Int64PtrToUint64(int64Ptr) // *int64 → uint64
```

## 许可证

MIT

