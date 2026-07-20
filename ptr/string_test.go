package ptr

import (
	"testing"
)

// ============ String 类型测试 ============

func TestStringPtr(t *testing.T) {
	val := "hello"
	ptr := StringPtr(val)
	if ptr == nil || *ptr != "hello" {
		t.Errorf("StringPtr(hello) failed, got %v", ptr)
	}
}

func TestPtrToString(t *testing.T) {
	val := "hello"
	ptr := &val
	result := PtrToString(ptr)
	if result != "hello" {
		t.Errorf("PtrToString(&hello) failed, got %s", result)
	}

	result = PtrToString(nil)
	if result != "" {
		t.Errorf("PtrToString(nil) failed, expected empty string, got %s", result)
	}
}

func TestStringToIntPtr(t *testing.T) {
	result := StringToIntPtr("42")
	if result == nil || *result != 42 {
		t.Errorf("StringToIntPtr(\"42\") failed, got %v", result)
	}

	result = StringToIntPtr("invalid")
	if result != nil {
		t.Errorf("StringToIntPtr(\"invalid\") should return nil, got %v", result)
	}
}

func TestStringToInt64Ptr(t *testing.T) {
	result := StringToInt64Ptr("9999")
	if result == nil || *result != 9999 {
		t.Errorf("StringToInt64Ptr(\"9999\") failed, got %v", result)
	}

	result = StringToInt64Ptr("invalid")
	if result != nil {
		t.Errorf("StringToInt64Ptr(\"invalid\") should return nil, got %v", result)
	}
}

func TestStringToFloat64Ptr(t *testing.T) {
	result := StringToFloat64Ptr("3.14")
	if result == nil || *result != 3.14 {
		t.Errorf("StringToFloat64Ptr(\"3.14\") failed, got %v", result)
	}

	result = StringToFloat64Ptr("invalid")
	if result != nil {
		t.Errorf("StringToFloat64Ptr(\"invalid\") should return nil, got %v", result)
	}
}
