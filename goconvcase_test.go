package goconvcase

import (
	"testing"
)

func TestConvertCase(t *testing.T) {
	testTbl := []struct {
		comment string
		src     string
		expect  string
	}{
		{"test convert 1",
			`package hoge
var HOGE_VAR int`,
			`package hoge
var HogeVar int`,
		},
	}

	for _, tt := range testTbl {
		got, err := ConvertCase(tt.src, UpperSnake, UpperCamel)
		if err != nil {
			t.Error(err)
		}
		if got != tt.expect {
			t.Errorf("got=%v, expect=%v", got, tt.expect)
		}
	}
}
