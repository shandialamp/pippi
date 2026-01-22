package collection

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
)

// Collection 泛型集合结构
type Collection[T any] struct {
	items []T
}

// New 创建一个新的集合
func New[T any](items ...T) *Collection[T] {
	return &Collection[T]{items: items}
}

// FromSlice 从切片创建集合
func FromSlice[T any](slice []T) *Collection[T] {
	return &Collection[T]{items: slice}
}

// All 获取集合中的所有项
func (c *Collection[T]) All() []T {
	return c.items
}

// Count 返回集合中的项数
func (c *Collection[T]) Count() int {
	return len(c.items)
}

// IsEmpty 检查集合是否为空
func (c *Collection[T]) IsEmpty() bool {
	return len(c.items) == 0
}

// IsNotEmpty 检查集合是否不为空
func (c *Collection[T]) IsNotEmpty() bool {
	return len(c.items) > 0
}

// First 获取集合的第一个元素
func (c *Collection[T]) First() (T, bool) {
	if len(c.items) == 0 {
		var zero T
		return zero, false
	}
	return c.items[0], true
}

// Last 获取集合的最后一个元素
func (c *Collection[T]) Last() (T, bool) {
	if len(c.items) == 0 {
		var zero T
		return zero, false
	}
	return c.items[len(c.items)-1], true
}

// Get 根据索引获取元素
func (c *Collection[T]) Get(index int) (T, bool) {
	if index < 0 || index >= len(c.items) {
		var zero T
		return zero, false
	}
	return c.items[index], true
}

// Push 向集合末尾添加元素
func (c *Collection[T]) Push(items ...T) *Collection[T] {
	c.items = append(c.items, items...)
	return c
}

// Pop 移除并返回集合的最后一个元素
func (c *Collection[T]) Pop() (T, bool) {
	if len(c.items) == 0 {
		var zero T
		return zero, false
	}
	last := c.items[len(c.items)-1]
	c.items = c.items[:len(c.items)-1]
	return last, true
}

// Shift 移除并返回集合的第一个元素
func (c *Collection[T]) Shift() (T, bool) {
	if len(c.items) == 0 {
		var zero T
		return zero, false
	}
	first := c.items[0]
	c.items = c.items[1:]
	return first, true
}

// Prepend 向集合开头添加元素
func (c *Collection[T]) Prepend(items ...T) *Collection[T] {
	c.items = append(items, c.items...)
	return c
}

// Filter 根据给定的回调函数过滤集合
func (c *Collection[T]) Filter(fn func(T) bool) *Collection[T] {
	filtered := make([]T, 0)
	for _, item := range c.items {
		if fn(item) {
			filtered = append(filtered, item)
		}
	}
	return &Collection[T]{items: filtered}
}

// Reject 根据给定的回调函数排除集合中的元素
func (c *Collection[T]) Reject(fn func(T) bool) *Collection[T] {
	return c.Filter(func(item T) bool {
		return !fn(item)
	})
}

// Map 对集合中的每个元素应用回调函数
func Map[T, U any](c *Collection[T], fn func(T) U) *Collection[U] {
	mapped := make([]U, len(c.items))
	for i, item := range c.items {
		mapped[i] = fn(item)
	}
	return &Collection[U]{items: mapped}
}

// Each 遍历集合中的每个元素
func (c *Collection[T]) Each(fn func(T)) *Collection[T] {
	for _, item := range c.items {
		fn(item)
	}
	return c
}

// EachWithIndex 遍历集合中的每个元素（带索引）
func (c *Collection[T]) EachWithIndex(fn func(int, T)) *Collection[T] {
	for i, item := range c.items {
		fn(i, item)
	}
	return c
}

// Reduce 将集合缩减为单个值
func Reduce[T, U any](c *Collection[T], fn func(U, T) U, initial U) U {
	result := initial
	for _, item := range c.items {
		result = fn(result, item)
	}
	return result
}

// Chunk 将集合分割成指定大小的块
func Chunk[T any](c *Collection[T], size int) [][]T {
	if size <= 0 {
		return [][]T{}
	}

	chunks := make([][]T, 0)
	for i := 0; i < len(c.items); i += size {
		end := i + size
		if end > len(c.items) {
			end = len(c.items)
		}
		chunks = append(chunks, c.items[i:end])
	}
	return chunks
}

// Slice 获取集合的切片
func (c *Collection[T]) Slice(start, end int) *Collection[T] {
	if start < 0 {
		start = 0
	}
	if end > len(c.items) {
		end = len(c.items)
	}
	if start > end {
		start = end
	}
	return &Collection[T]{items: c.items[start:end]}
}

// Take 获取集合的前n个元素
func (c *Collection[T]) Take(n int) *Collection[T] {
	if n > len(c.items) {
		n = len(c.items)
	}
	if n < 0 {
		// 取后n个元素
		if -n > len(c.items) {
			return &Collection[T]{items: c.items}
		}
		return &Collection[T]{items: c.items[len(c.items)+n:]}
	}
	return &Collection[T]{items: c.items[:n]}
}

// Skip 跳过集合的前n个元素
func (c *Collection[T]) Skip(n int) *Collection[T] {
	if n >= len(c.items) {
		return &Collection[T]{items: []T{}}
	}
	if n < 0 {
		n = 0
	}
	return &Collection[T]{items: c.items[n:]}
}

// Reverse 反转集合
func (c *Collection[T]) Reverse() *Collection[T] {
	reversed := make([]T, len(c.items))
	for i, item := range c.items {
		reversed[len(c.items)-1-i] = item
	}
	return &Collection[T]{items: reversed}
}

// Shuffle 随机打乱集合
func (c *Collection[T]) Shuffle() *Collection[T] {
	shuffled := make([]T, len(c.items))
	copy(shuffled, c.items)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return &Collection[T]{items: shuffled}
}

// Random 随机获取一个元素
func (c *Collection[T]) Random() (T, bool) {
	if len(c.items) == 0 {
		var zero T
		return zero, false
	}
	return c.items[rand.Intn(len(c.items))], true
}

// Unique 移除集合中的重复元素（需要元素类型可比较）
func Unique[T comparable](c *Collection[T]) *Collection[T] {
	seen := make(map[T]bool)
	unique := make([]T, 0)
	for _, item := range c.items {
		if !seen[item] {
			seen[item] = true
			unique = append(unique, item)
		}
	}
	return &Collection[T]{items: unique}
}

// Contains 检查集合是否包含给定的元素
func Contains[T comparable](c *Collection[T], value T) bool {
	for _, item := range c.items {
		if item == value {
			return true
		}
	}
	return false
}

// ContainsFunc 使用回调函数检查集合是否包含满足条件的元素
func (c *Collection[T]) ContainsFunc(fn func(T) bool) bool {
	for _, item := range c.items {
		if fn(item) {
			return true
		}
	}
	return false
}

// Every 检查集合中的所有元素是否都满足条件
func (c *Collection[T]) Every(fn func(T) bool) bool {
	for _, item := range c.items {
		if !fn(item) {
			return false
		}
	}
	return true
}

// Some 检查集合中是否至少有一个元素满足条件
func (c *Collection[T]) Some(fn func(T) bool) bool {
	return c.ContainsFunc(fn)
}

// Partition 将集合分为两个集合，一个包含满足条件的元素，另一个包含不满足条件的元素
func (c *Collection[T]) Partition(fn func(T) bool) (*Collection[T], *Collection[T]) {
	passed := make([]T, 0)
	failed := make([]T, 0)
	for _, item := range c.items {
		if fn(item) {
			passed = append(passed, item)
		} else {
			failed = append(failed, item)
		}
	}
	return &Collection[T]{items: passed}, &Collection[T]{items: failed}
}

// GroupBy 根据给定的键对集合进行分组
func GroupBy[T any, K comparable](c *Collection[T], fn func(T) K) map[K]*Collection[T] {
	groups := make(map[K]*Collection[T])
	for _, item := range c.items {
		key := fn(item)
		if _, ok := groups[key]; !ok {
			groups[key] = &Collection[T]{items: make([]T, 0)}
		}
		groups[key].items = append(groups[key].items, item)
	}
	return groups
}

// Sort 对集合进行排序
func Sort[T any](c *Collection[T], less func(T, T) bool) *Collection[T] {
	sorted := make([]T, len(c.items))
	copy(sorted, c.items)
	sort.Slice(sorted, func(i, j int) bool {
		return less(sorted[i], sorted[j])
	})
	return &Collection[T]{items: sorted}
}

// SortDesc 对集合进行降序排序
func SortDesc[T any](c *Collection[T], less func(T, T) bool) *Collection[T] {
	return Sort(c, func(a, b T) bool {
		return less(b, a)
	})
}

// Pluck 从集合中提取给定键的所有值
func Pluck[T any, U any](c *Collection[T], fn func(T) U) *Collection[U] {
	return Map(c, fn)
}

// Sum 计算集合元素的总和
func Sum[T any](c *Collection[T], fn func(T) float64) float64 {
	sum := 0.0
	for _, item := range c.items {
		sum += fn(item)
	}
	return sum
}

// Avg 计算集合元素的平均值
func Avg[T any](c *Collection[T], fn func(T) float64) float64 {
	if len(c.items) == 0 {
		return 0
	}
	return Sum(c, fn) / float64(len(c.items))
}

// Min 获取集合中的最小值
func Min[T any](c *Collection[T], fn func(T) float64) (T, bool) {
	if len(c.items) == 0 {
		var zero T
		return zero, false
	}
	minItem := c.items[0]
	minValue := fn(minItem)
	for _, item := range c.items[1:] {
		value := fn(item)
		if value < minValue {
			minValue = value
			minItem = item
		}
	}
	return minItem, true
}

// Max 获取集合中的最大值
func Max[T any](c *Collection[T], fn func(T) float64) (T, bool) {
	if len(c.items) == 0 {
		var zero T
		return zero, false
	}
	maxItem := c.items[0]
	maxValue := fn(maxItem)
	for _, item := range c.items[1:] {
		value := fn(item)
		if value > maxValue {
			maxValue = value
			maxItem = item
		}
	}
	return maxItem, true
}

// Flatten 将多维集合扁平化为一维集合
func Flatten[T any](c *Collection[[]T]) *Collection[T] {
	flattened := make([]T, 0)
	for _, items := range c.items {
		flattened = append(flattened, items...)
	}
	return &Collection[T]{items: flattened}
}

// FlatMap 对集合应用映射函数，然后扁平化结果
func FlatMap[T, U any](c *Collection[T], fn func(T) []U) *Collection[U] {
	flattened := make([]U, 0)
	for _, item := range c.items {
		flattened = append(flattened, fn(item)...)
	}
	return &Collection[U]{items: flattened}
}

// Zip 将多个集合合并为一个集合
func Zip[T, U any](c1 *Collection[T], c2 *Collection[U]) *Collection[[2]any] {
	length := len(c1.items)
	if len(c2.items) < length {
		length = len(c2.items)
	}
	zipped := make([][2]any, length)
	for i := 0; i < length; i++ {
		zipped[i] = [2]any{c1.items[i], c2.items[i]}
	}
	return &Collection[[2]any]{items: zipped}
}

// Join 将集合元素连接成字符串
func Join[T any](c *Collection[T], separator string, fn func(T) string) string {
	if len(c.items) == 0 {
		return ""
	}
	result := fn(c.items[0])
	for _, item := range c.items[1:] {
		result += separator + fn(item)
	}
	return result
}

// ToJSON 将集合转换为JSON字符串
func (c *Collection[T]) ToJSON() (string, error) {
	data, err := json.Marshal(c.items)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// FromJSON 从JSON字符串创建集合
func FromJSON[T any](jsonStr string) (*Collection[T], error) {
	var items []T
	err := json.Unmarshal([]byte(jsonStr), &items)
	if err != nil {
		return nil, err
	}
	return &Collection[T]{items: items}, nil
}

// String 实现Stringer接口
func (c *Collection[T]) String() string {
	return fmt.Sprintf("%v", c.items)
}

// Clone 克隆集合
func (c *Collection[T]) Clone() *Collection[T] {
	cloned := make([]T, len(c.items))
	copy(cloned, c.items)
	return &Collection[T]{items: cloned}
}

// Merge 合并多个集合
func (c *Collection[T]) Merge(others ...*Collection[T]) *Collection[T] {
	merged := make([]T, len(c.items))
	copy(merged, c.items)
	for _, other := range others {
		merged = append(merged, other.items...)
	}
	return &Collection[T]{items: merged}
}

// Diff 返回集合中存在但不在给定集合中的元素
func Diff[T comparable](c *Collection[T], other *Collection[T]) *Collection[T] {
	otherMap := make(map[T]bool)
	for _, item := range other.items {
		otherMap[item] = true
	}
	diff := make([]T, 0)
	for _, item := range c.items {
		if !otherMap[item] {
			diff = append(diff, item)
		}
	}
	return &Collection[T]{items: diff}
}

// Intersect 返回两个集合的交集
func Intersect[T comparable](c *Collection[T], other *Collection[T]) *Collection[T] {
	otherMap := make(map[T]bool)
	for _, item := range other.items {
		otherMap[item] = true
	}
	intersect := make([]T, 0)
	seen := make(map[T]bool)
	for _, item := range c.items {
		if otherMap[item] && !seen[item] {
			intersect = append(intersect, item)
			seen[item] = true
		}
	}
	return &Collection[T]{items: intersect}
}

// Union 返回两个集合的并集
func Union[T comparable](c *Collection[T], other *Collection[T]) *Collection[T] {
	seen := make(map[T]bool)
	union := make([]T, 0)
	for _, item := range c.items {
		if !seen[item] {
			union = append(union, item)
			seen[item] = true
		}
	}
	for _, item := range other.items {
		if !seen[item] {
			union = append(union, item)
			seen[item] = true
		}
	}
	return &Collection[T]{items: union}
}

// Tap 执行给定的回调函数并返回集合本身
func (c *Collection[T]) Tap(fn func(*Collection[T])) *Collection[T] {
	fn(c)
	return c
}

// Pipe 通过给定的回调函数传递集合并返回结果
func Pipe[T, U any](c *Collection[T], fn func(*Collection[T]) U) U {
	return fn(c)
}

// When 当给定条件为真时执行回调函数
func (c *Collection[T]) When(condition bool, fn func(*Collection[T]) *Collection[T]) *Collection[T] {
	if condition {
		return fn(c)
	}
	return c
}

// Unless 当给定条件为假时执行回调函数
func (c *Collection[T]) Unless(condition bool, fn func(*Collection[T]) *Collection[T]) *Collection[T] {
	if !condition {
		return fn(c)
	}
	return c
}
