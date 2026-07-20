package ptr

import (
	"testing"
)

// ============ Bool 类型测试 ============

func TestBoolPtr(t *testing.T) {
	val := true
	ptr := BoolPtr(val)
	if ptr == nil || *ptr != true {
		t.Errorf("BoolPtr(true) failed, got %v", ptr)
	}
}

func TestPtrToBool(t *testing.T) {
	val := true
	ptr := &val
	result := PtrToBool(ptr)
	if result != true {
		t.Errorf("PtrToBool(&true) failed, got %v", result)
	}

	result = PtrToBool(nil)
	if result != false {
		t.Errorf("PtrToBool(nil) failed, expected false, got %v", result)
	}
}
