package ptr

// ============ Int64 转换 ============

// Int64Ptr 将 int64 转换为指针
func Int64Ptr(v int64) *int64 {
	return &v
}

// Int8PtrToInt64 将 int8 指针转换为 int64
func Int8PtrToInt64(p *int8) int64 {
	if p == nil {
		return 0
	}
	return int64(*p)
}

// Int16PtrToInt64 将 int16 指针转换为 int64
func Int16PtrToInt64(p *int16) int64 {
	if p == nil {
		return 0
	}
	return int64(*p)
}

// Int32PtrToInt64 将 int32 指针转换为 int64
func Int32PtrToInt64(p *int32) int64 {
	if p == nil {
		return 0
	}
	return int64(*p)
}

// Int64PtrToInt64 将 int64 指针转换为 int64
func Int64PtrToInt64(p *int64) int64 {
	if p == nil {
		return 0
	}
	return *p
}

// IntPtrToInt64 将 int 指针转换为 int64
func IntPtrToInt64(p *int) int64 {
	if p == nil {
		return 0
	}
	return int64(*p)
}

// Int64ToInt8Ptr 将 int64 转换为 int8 指针
func Int64ToInt8Ptr(v int64) *int8 {
	i8 := int8(v)
	return &i8
}

// Int64ToInt16Ptr 将 int64 转换为 int16 指针
func Int64ToInt16Ptr(v int64) *int16 {
	i16 := int16(v)
	return &i16
}

// Int64ToInt32Ptr 将 int64 转换为 int32 指针
func Int64ToInt32Ptr(v int64) *int32 {
	i32 := int32(v)
	return &i32
}

// Int64ToInt64Ptr 将 int64 转换为 int64 指针
func Int64ToInt64Ptr(v int64) *int64 {
	return &v
}

// Int64ToIntPtr 将 int64 转换为 int 指针
func Int64ToIntPtr(v int64) *int {
	i := int(v)
	return &i
}

// ============ Uint64 转换 ============

// Uint64Ptr 将 uint64 转换为指针
func Uint64Ptr(v uint64) *uint64 {
	return &v
}

// Uint8PtrToUint64 将 uint8 指针转换为 uint64
func Uint8PtrToUint64(p *uint8) uint64 {
	if p == nil {
		return 0
	}
	return uint64(*p)
}

// Uint16PtrToUint64 将 uint16 指针转换为 uint64
func Uint16PtrToUint64(p *uint16) uint64 {
	if p == nil {
		return 0
	}
	return uint64(*p)
}

// Uint32PtrToUint64 将 uint32 指针转换为 uint64
func Uint32PtrToUint64(p *uint32) uint64 {
	if p == nil {
		return 0
	}
	return uint64(*p)
}

// Uint64PtrToUint64 将 uint64 指针转换为 uint64
func Uint64PtrToUint64(p *uint64) uint64 {
	if p == nil {
		return 0
	}
	return *p
}

// UintPtrToUint64 将 uint 指针转换为 uint64
func UintPtrToUint64(p *uint) uint64 {
	if p == nil {
		return 0
	}
	return uint64(*p)
}

// Uint64ToUint8Ptr 将 uint64 转换为 uint8 指针
func Uint64ToUint8Ptr(v uint64) *uint8 {
	u8 := uint8(v)
	return &u8
}

// Uint64ToUint16Ptr 将 uint64 转换为 uint16 指针
func Uint64ToUint16Ptr(v uint64) *uint16 {
	u16 := uint16(v)
	return &u16
}

// Uint64ToUint32Ptr 将 uint64 转换为 uint32 指针
func Uint64ToUint32Ptr(v uint64) *uint32 {
	u32 := uint32(v)
	return &u32
}

// Uint64ToUint64Ptr 将 uint64 转换为 uint64 指针
func Uint64ToUint64Ptr(v uint64) *uint64 {
	return &v
}

// Uint64ToUintPtr 将 uint64 转换为 uint 指针
func Uint64ToUintPtr(v uint64) *uint {
	u := uint(v)
	return &u
}
