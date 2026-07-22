package exception

import (
	"errors"
	"fmt"

	"go.uber.org/zap"
)

// Error 错误封装
type Error struct {
	err    error
	trace  string
	fields []zap.Field
}

// Error 实现 error 接口
func (e *Error) Error() string {
	return e.err.Error()
}

// Unwrap 保留原始错误链
func (e *Error) Unwrap() error {
	return e.err
}

// Stack 获取错误堆栈
func (e *Error) Trace() string {
	return e.trace
}

// Fields 获取附加日志字段
func (e *Error) Fields() []zap.Field {
	return append([]zap.Field(nil), e.fields...)
}

func (e *Error) TraceAndFields() []zap.Field {
	fileds := []zap.Field{
		zap.String("trace", e.trace),
	}
	return append(fileds, e.fields...)
}

// New 创建错误
//
// err 建议传入 pkg/errors.Wrap / WithStack 创建的错误
func New(err error, fields ...zap.Field) error {
	if err == nil {
		return nil
	}

	return &Error{
		err:    err,
		trace:  fmt.Sprintf("%+v", err),
		fields: fields,
	}
}

// As 判断并获取 exception.Error
func As(err error) (*Error, bool) {
	var ex *Error

	if errors.As(err, &ex) {
		return ex, true
	}

	return nil, false
}
