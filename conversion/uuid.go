package conversion

import (
	"reflect"

	"github.com/google/uuid"
)

var UUID uuid.UUID

func Uuid(s interface{}) uuid.UUID {
	pc := processConvertUuidPtr(s)
	if pc == nil {
		return UUID
	}
	return *pc
}

func UuidPtr(s interface{}) *uuid.UUID {
	pc := processConvertUuidPtr(s)
	if pc == nil {
		return &UUID
	}

	return pc
}

func processConvertUuidPtr(s interface{}) *uuid.UUID {
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

	if uidx, ok := s.(string); ok {
		uid, _ := uuid.Parse(uidx)
		return &uid
	}

	if uidx, ok := s.(uuid.UUID); ok {
		return &uidx
	}

	if ui, ok := s.([]uint8); ok {
		uuidStr := string(ui)
		uid, _ := uuid.Parse(uuidStr)
		return &uid
	}

	return nil
}
