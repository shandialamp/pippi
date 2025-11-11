package structx

import "reflect"

func WithDefaults[T any](dst *T, def *T) {
	dstVal := reflect.ValueOf(dst).Elem()
	defVal := reflect.ValueOf(def).Elem()

	for i := 0; i < dstVal.NumField(); i++ {
		field := dstVal.Field(i)
		defField := defVal.Field(i)

		if field.Kind() == reflect.Ptr && field.IsNil() && !defField.IsNil() {
			field.Set(defField)
		}
	}
}

func PtrString(s string) *string { return &s }
func PtrInt(i int) *int          { return &i }
