package conversion

import (
	"reflect"
	"time"
)

var TIME time.Time

func Time(s interface{}) time.Time {

	pc := processTimePtr(s)
	if pc == nil {
		return TIME
	}

	return *pc
}

func TimePtr(s interface{}) *time.Time {

	pc := processTimePtr(s)
	if pc == nil {
		return nil
	}

	return nil
}

func processTimePtr(s interface{}) *time.Time {
	if s == nil {
		return nil
	}

	// Is pointer?
	if reflect.TypeOf(s).Kind() == reflect.Pointer {
		vs := reflect.ValueOf(s)
		// Has value?
		if vs.IsNil() {
			return nil
		}
		s = reflect.Indirect(vs).Interface()
	}

	if val, ok := s.(time.Time); ok {
		return &val
	}

	if val, ok := s.(string); ok {

		t, err := time.Parse(time.Layout, val)
		if err == nil {
			return &t
		}

		t, err = time.Parse(time.ANSIC, val)
		if err == nil {
			return &t
		}

		t, err = time.Parse(time.UnixDate, val)
		if err == nil {
			return &t
		}

		t, err = time.Parse(time.RubyDate, val)
		if err == nil {
			return &t
		}

		t, err = time.Parse(time.RFC822, val)
		if err == nil {
			return &t
		}

		t, err = time.Parse(time.RFC822Z, val)
		if err == nil {
			return &t
		}

		t, err = time.Parse(time.RFC850, val)
		if err == nil {
			return &t
		}

		t, err = time.Parse(time.RFC1123, val)
		if err == nil {
			return &t
		}

		t, err = time.Parse(time.RFC1123Z, val)
		if err == nil {
			return &t
		}

		t, err = time.Parse(time.RFC3339, val)
		if err == nil {
			return &t
		}

		t, err = time.Parse(time.RFC3339Nano, val)
		if err == nil {
			return &t
		}

		t, err = time.Parse(time.Kitchen, val)
		if err == nil {
			return &t
		}

		t, err = time.Parse(time.Stamp, val)
		if err == nil {
			return &t
		}

		t, err = time.Parse(time.StampMilli, val)
		if err == nil {
			return &t
		}

		t, err = time.Parse(time.StampMicro, val)
		if err == nil {
			return &t
		}

		t, err = time.Parse(time.StampNano, val)
		if err == nil {
			return &t
		}

		t, err = time.Parse(time.DateTime, val)
		if err == nil {
			return &t
		}

		t, err = time.Parse(time.DateOnly, val)
		if err == nil {
			return &t
		}

		t, err = time.Parse(time.TimeOnly, val)
		if err == nil {
			return &t
		}
	}

	return nil
}
