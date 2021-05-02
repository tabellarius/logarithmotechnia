package vector

import (
	"reflect"
	"testing"
)

func TestEmpty_ByBool(t *testing.T) {
	vec := Empty()
	newVec := vec.Clone()
	if _, ok := vec.(EmptyVector); !ok {
		t.Error("Returned vector is not empty")
	}
	if reflect.ValueOf(vec).Pointer() == reflect.ValueOf(newVec).Pointer() {
		t.Error("Vector is not new")
	}
}

func TestEmpty_ByFromTo(t *testing.T) {
	vec := Empty()
	newVec := vec.Clone()
	if _, ok := vec.(EmptyVector); !ok {
		t.Error("Returned vector is not empty")
	}
	if reflect.ValueOf(vec).Pointer() == reflect.ValueOf(newVec).Pointer() {
		t.Error("Vector is not new")
	}
}

func TestEmpty_ByIndices(t *testing.T) {
	vec := Empty()
	newVec := vec.Clone()
	if _, ok := vec.(EmptyVector); !ok {
		t.Error("Returned vector is not empty")
	}
	if reflect.ValueOf(vec).Pointer() == reflect.ValueOf(newVec).Pointer() {
		t.Error("Vector is not new")
	}
}

func TestEmpty_Clone(t *testing.T) {
	vec := Empty()
	newVec := vec.Clone()
	if _, ok := vec.(EmptyVector); !ok {
		t.Error("Returned vector is not empty")
	}
	if reflect.ValueOf(vec).Pointer() == reflect.ValueOf(newVec).Pointer() {
		t.Error("Vector is not new")
	}
}

func TestEmpty_IsEmpty(t *testing.T) {
	vec := Empty()
	if !vec.IsEmpty() {
		t.Error("IsEmpty returns non-empty for empty vector")
	}
}

func TestEmpty_Length(t *testing.T) {
	vec := Empty()
	if vec.Length() != 0 {
		t.Error("Non-zero length for empty vector")
	}
}

func TestEmpty_Report(t *testing.T) {
	vec := Empty()
	if len(vec.Report().Warnings) > 0 || len(vec.Report().Errors) > 0 {
		t.Error("Non-empty warnings or errors for new empty vector")
	}
}
