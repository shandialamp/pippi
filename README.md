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

## 许可证

MIT
