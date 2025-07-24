package validator

import "reflect"

// ตรวจสอบ struct ว่าง
func IsEmptyStruct[T any](data, typeStruct T) bool {
	return reflect.DeepEqual(data, typeStruct)
}

// ตรวจสอบค่าว่าง
func IsEmpty(val any) bool {
	return IsEmptyPtr(val)
}

func IsEmptyPtr(val any) bool {
	if val == nil {
		return true
	}

	v := reflect.ValueOf(val)

	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Array, reflect.Slice, reflect.Map:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	}

	return false
}
