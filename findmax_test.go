package mcommon

import (
	"reflect"
	"testing"
	"errors"
)

func TestMax(t *testing.T) {
	type testStruct struct {
		some string
		value int
	}

	tests := []struct {
		name string
		data []interface{}
		want interface{}
		wantedErr error
	}{
		{"empty", []interface{}{}, nil, errors.New("Cannot find maximum for empty slice")},
		{"single value", []interface{}{testStruct{"asdf", 15}}, testStruct{"asdf", 15}, nil},
		{"multiple values", []interface{}{
			testStruct{"aaa", 32},
			testStruct{"bbb", 42},  // max
			testStruct{"ccc", 18},
			testStruct{"e", 8},
			testStruct{"xx", 1},
		}, testStruct{"bbb", 42}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Max(tt.data, func(i, j int) bool { 
				return tt.data[i].(testStruct).value > tt.data[j].(testStruct).value
			})
			if !reflect.DeepEqual(err, tt.wantedErr) {
				t.Fatalf("Max() got error = '%v', want '%v'", err, tt.wantedErr)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
