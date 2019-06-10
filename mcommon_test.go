package mcommon

import (
	"reflect"
	"testing"
	"strconv"
)

func TestMostCommon10(t *testing.T) {
	tests := []struct {
		text string
		want []string
	}{
		{"", []string{}},
		{"xx bb xx bb xx cc", []string{"xx", "bb", "cc"}},
		{`xx xx xx xx xx xx xx xx xx xx xx
			ii ii ii
			ee ee ee ee ee ee ee
			gg gg gg gg gg
			cc cc cc cc cc cc cc cc cc
			hh hh hh hh
			ff ff ff ff ff ff
			dd dd dd dd dd dd dd dd
			jj jj
			bb bb bb bb bb bb bb bb bb bb
		  kk
		  ll`, []string{"xx", "bb", "cc", "dd", "ee", "ff", "gg", "hh", "ii", "jj"}},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := MostCommon10(tt.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MostCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}
