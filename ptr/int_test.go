package ptr

import (
	"testing"
)

// ============ Int64 转换测试 ============

func TestInt64Ptr(t *testing.T) {
	val := int64(999999)
	ptr := Int64Ptr(val)
	if ptr == nil || *ptr != 999999 {
		t.Errorf("Int64Ptr(999999) failed, got %v", ptr)
	}
}

func TestInt8PtrToInt64(t *testing.T) {
	val := int8(10)
	result := Int8PtrToInt64(&val)
	if result != 10 {
		t.Errorf("Int8PtrToInt64(&10) failed, got %d", result)
	}
	result = Int8PtrToInt64(nil)
	if result != 0 {
		t.Errorf("Int8PtrToInt64(nil) failed, expected 0, got %d", result)
	}
}

func TestInt16PtrToInt64(t *testing.T) {
	val := int16(200)
	result := Int16PtrToInt64(&val)
	if result != 200 {
		t.Errorf("Int16PtrToInt64(&200) failed, got %d", result)
	}
	result = Int16PtrToInt64(nil)
	if result != 0 {
		t.Errorf("Int16PtrToInt64(nil) failed, expected 0, got %d", result)
	}
}

func TestInt32PtrToInt64(t *testing.T) {
	val := int32(50000)
	result := Int32PtrToInt64(&val)
	if result != 50000 {
		t.Errorf("Int32PtrToInt64(&50000) failed, got %d", result)
	}
	result = Int32PtrToInt64(nil)
	if result != 0 {
		t.Errorf("Int32PtrToInt64(nil) failed, expected 0, got %d", result)
	}
}

func TestInt64PtrToInt64(t *testing.T) {
	val := int64(42)
	ptr := &val
	result := Int64PtrToInt64(ptr)
	if result != 42 {
		t.Errorf("Int64PtrToInt64(&42) failed, got %d", result)
	}

	result = Int64PtrToInt64(nil)
	if result != 0 {
		t.Errorf("Int64PtrToInt64(nil) failed, expected 0, got %d", result)
	}
}

func TestIntPtrToInt64(t *testing.T) {
	val := int(888)
	result := IntPtrToInt64(&val)
	if result != 888 {
		t.Errorf("IntPtrToInt64(&888) failed, got %d", result)
	}
	result = IntPtrToInt64(nil)
	if result != 0 {
		t.Errorf("IntPtrToInt64(nil) failed, expected 0, got %d", result)
	}
}

func TestInt64ToInt8Ptr(t *testing.T) {
	ptr := Int64ToInt8Ptr(100)
	if ptr == nil || *ptr != 100 {
		t.Errorf("Int64ToInt8Ptr(100) failed, got %v", ptr)
	}
}

func TestInt64ToInt16Ptr(t *testing.T) {
	ptr := Int64ToInt16Ptr(1000)
	if ptr == nil || *ptr != 1000 {
		t.Errorf("Int64ToInt16Ptr(1000) failed, got %v", ptr)
	}
}

func TestInt64ToInt32Ptr(t *testing.T) {
	ptr := Int64ToInt32Ptr(9999)
	if ptr == nil || *ptr != 9999 {
		t.Errorf("Int64ToInt32Ptr(9999) failed, got %v", ptr)
	}
}

func TestInt64ToInt64Ptr(t *testing.T) {
	ptr := Int64ToInt64Ptr(123456)
	if ptr == nil || *ptr != 123456 {
		t.Errorf("Int64ToInt64Ptr(123456) failed, got %v", ptr)
	}
}

func TestInt64ToIntPtr(t *testing.T) {
	ptr := Int64ToIntPtr(5555)
	if ptr == nil || *ptr != 5555 {
		t.Errorf("Int64ToIntPtr(5555) failed, got %v", ptr)
	}
}

// ============ Uint64 转换测试 ============

func TestUint64Ptr(t *testing.T) {
	val := uint64(9999)
	ptr := Uint64Ptr(val)
	if ptr == nil || *ptr != 9999 {
		t.Errorf("Uint64Ptr(9999) failed, got %v", ptr)
	}
}

func TestUint8PtrToUint64(t *testing.T) {
	val := uint8(20)
	result := Uint8PtrToUint64(&val)
	if result != 20 {
		t.Errorf("Uint8PtrToUint64(&20) failed, got %d", result)
	}
	result = Uint8PtrToUint64(nil)
	if result != 0 {
		t.Errorf("Uint8PtrToUint64(nil) failed, expected 0, got %d", result)
	}
}

func TestUint16PtrToUint64(t *testing.T) {
	val := uint16(5000)
	result := Uint16PtrToUint64(&val)
	if result != 5000 {
		t.Errorf("Uint16PtrToUint64(&5000) failed, got %d", result)
	}
	result = Uint16PtrToUint64(nil)
	if result != 0 {
		t.Errorf("Uint16PtrToUint64(nil) failed, expected 0, got %d", result)
	}
}

func TestUint32PtrToUint64(t *testing.T) {
	val := uint32(50000)
	result := Uint32PtrToUint64(&val)
	if result != 50000 {
		t.Errorf("Uint32PtrToUint64(&50000) failed, got %d", result)
	}
	result = Uint32PtrToUint64(nil)
	if result != 0 {
		t.Errorf("Uint32PtrToUint64(nil) failed, expected 0, got %d", result)
	}
}

func TestUint64PtrToUint64(t *testing.T) {
	val := uint64(100)
	ptr := &val
	result := Uint64PtrToUint64(ptr)
	if result != 100 {
		t.Errorf("Uint64PtrToUint64(&100) failed, got %d", result)
	}

	result = Uint64PtrToUint64(nil)
	if result != 0 {
		t.Errorf("Uint64PtrToUint64(nil) failed, expected 0, got %d", result)
	}
}

func TestUintPtrToUint64(t *testing.T) {
	val := uint(999)
	result := UintPtrToUint64(&val)
	if result != 999 {
		t.Errorf("UintPtrToUint64(&999) failed, got %d", result)
	}
	result = UintPtrToUint64(nil)
	if result != 0 {
		t.Errorf("UintPtrToUint64(nil) failed, expected 0, got %d", result)
	}
}

func TestUint64ToUint8Ptr(t *testing.T) {
	ptr := Uint64ToUint8Ptr(200)
	if ptr == nil || *ptr != 200 {
		t.Errorf("Uint64ToUint8Ptr(200) failed, got %v", ptr)
	}
}

func TestUint64ToUint16Ptr(t *testing.T) {
	ptr := Uint64ToUint16Ptr(5000)
	if ptr == nil || *ptr != 5000 {
		t.Errorf("Uint64ToUint16Ptr(5000) failed, got %v", ptr)
	}
}

func TestUint64ToUint32Ptr(t *testing.T) {
	ptr := Uint64ToUint32Ptr(50000)
	if ptr == nil || *ptr != 50000 {
		t.Errorf("Uint64ToUint32Ptr(50000) failed, got %v", ptr)
	}
}

func TestUint64ToUint64Ptr(t *testing.T) {
	ptr := Uint64ToUint64Ptr(999999)
	if ptr == nil || *ptr != 999999 {
		t.Errorf("Uint64ToUint64Ptr(999999) failed, got %v", ptr)
	}
}

func TestUint64ToUintPtr(t *testing.T) {
	ptr := Uint64ToUintPtr(7777)
	if ptr == nil || *ptr != 7777 {
		t.Errorf("Uint64ToUintPtr(7777) failed, got %v", ptr)
	}
}
