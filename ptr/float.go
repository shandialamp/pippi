package ptr

// Float32Ptr 将 float32 转换为指针
func Float32Ptr(v float32) *float32 {
	return &v
}

// PtrToFloat32 将 float32 指针转换为值
func PtrToFloat32(p *float32) float32 {
	if p == nil {
		return 0
	}
	return *p
}

// Float64Ptr 将 float64 转换为指针
func Float64Ptr(v float64) *float64 {
	return &v
}

// PtrToFloat64 将 float64 指针转换为值
func PtrToFloat64(p *float64) float64 {
	if p == nil {
		return 0
	}
	return *p
}
