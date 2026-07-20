package ptr

// BoolPtr 将 bool 转换为指针
func BoolPtr(v bool) *bool {
	return &v
}

// PtrToBool 将 bool 指针转换为值
func PtrToBool(p *bool) bool {
	if p == nil {
		return false
	}
	return *p
}
