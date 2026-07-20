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

// ============ Uint64 转换测试 ============

func TestUint64Ptr(t *testing.T) {
	val := uint64(9999)
	ptr := Uint64Ptr(val)
	if ptr == nil || *ptr != 9999 {
		t.Errorf("Uint64Ptr(9999) failed, got %v", ptr)
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
