package ptr

import "strconv"

// StringPtr 将 string 转换为指针
func StringPtr(v string) *string {
	return &v
}

// PtrToString 将 string 指针转换为值
func PtrToString(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

// StringToIntPtr 将 string 转换为 int 指针，转换失败返回 nil
func StringToIntPtr(v string) *int {
	i, err := strconv.Atoi(v)
	if err != nil {
		return nil
	}
	return &i
}

// StringToInt64Ptr 将 string 转换为 int64 指针，转换失败返回 nil
func StringToInt64Ptr(v string) *int64 {
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return nil
	}
	return &i
}

// StringToFloat64Ptr 将 string 转换为 float64 指针，转换失败返回 nil
func StringToFloat64Ptr(v string) *float64 {
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return nil
	}
	return &f
}
