package ptr

import (
	"testing"
)

// ============ Float 类型测试 ============

func TestFloat32Ptr(t *testing.T) {
	val := float32(3.14)
	ptr := Float32Ptr(val)
	if ptr == nil || *ptr != 3.14 {
		t.Errorf("Float32Ptr(3.14) failed, got %v", ptr)
	}
}

func TestPtrToFloat32(t *testing.T) {
	val := float32(3.14)
	ptr := &val
	result := PtrToFloat32(ptr)
	if result != 3.14 {
		t.Errorf("PtrToFloat32(&3.14) failed, got %f", result)
	}

	result = PtrToFloat32(nil)
	if result != 0 {
		t.Errorf("PtrToFloat32(nil) failed, expected 0, got %f", result)
	}
}

func TestFloat64Ptr(t *testing.T) {
	val := 3.14159
	ptr := Float64Ptr(val)
	if ptr == nil || *ptr != 3.14159 {
		t.Errorf("Float64Ptr(3.14159) failed, got %v", ptr)
	}
}

func TestPtrToFloat64(t *testing.T) {
	val := 3.14159
	ptr := &val
	result := PtrToFloat64(ptr)
	if result != 3.14159 {
		t.Errorf("PtrToFloat64(&3.14159) failed, got %f", result)
	}

	result = PtrToFloat64(nil)
	if result != 0 {
		t.Errorf("PtrToFloat64(nil) failed, expected 0, got %f", result)
	}
}
