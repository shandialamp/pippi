# Pippi

一个功能强大的 Go 工具库集合。

## 模块

### Collection

类似 Laravel Collection 的 Go 泛型集合库，提供丰富的数据处理方法。

**特性：**
- 🚀 完全使用 Go 泛型实现
- 📦 70+ 丰富的集合操作方法
- 🔗 支持链式调用
- 🎯 类型安全
- ⚡ 高性能

**安装：**
```bash
go get github.com/shandialamp/pippi/collection
```

**快速开始：**
```go
import "github.com/shandialamp/pippi/collection"

c := collection.New(1, 2, 3, 4, 5)

// 过滤、映射、链式调用
result := c.Filter(func(n int) bool {
    return n%2 == 0
}).Skip(1).Take(2)

fmt.Println(result.All()) // [4, 6]
```

[查看完整文档](collection/README.md)

### StructX

结构体工具库（待补充）

### Ptr

指针工具库，为基础数据类型提供方便的指针转换函数。

**特性：**
- 🎯 简洁的函数设计 - 只保留必要的转换
- 🔄 int64/uint64 为源头 - 转换为其他 int 类型的指针
- 🛡️ 安全处理 nil 指针，返回零值
- 📦 轻量设计，精简 API
- ⚡ 高效的指针操作

**快速开始：**
```go
import "github.com/shandialamp/pippi/ptr"

// int64 作为源头转换
int8Ptr := ptr.Int64ToInt8Ptr(100)       // int64 → *int8
int16Ptr := ptr.Int64ToInt16Ptr(1000)    // int64 → *int16
int32Ptr := ptr.Int64ToInt32Ptr(9999)    // int64 → *int32

// uint64 作为源头转换
uint8Ptr := ptr.Uint64ToUint8Ptr(200)    // uint64 → *uint8
uint16Ptr := ptr.Uint64ToUint16Ptr(5000) // uint64 → *uint16
uint32Ptr := ptr.Uint64ToUint32Ptr(50000) // uint64 → *uint32

// int64 指针转值
int64Ptr := ptr.Int64Ptr(999)
val := ptr.PtrToInt64(int64Ptr) // *int64 → int64

// 字符串转数字
numPtr := ptr.StringToInt64Ptr("9999")
if numPtr != nil {
	fmt.Println(*numPtr) // 9999
}

// 其他类型
strPtr := ptr.StringPtr("hello")
boolPtr := ptr.BoolPtr(true)
floatPtr := ptr.Float64Ptr(3.14)
```

[查看完整文档](ptr/README.md)

## 许可证

MIT
