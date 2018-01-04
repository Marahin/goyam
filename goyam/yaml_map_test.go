package goyam

import (
	"reflect"
	"testing"
)

func TestGrabTopLevelKeys(t *testing.T) {
	type args struct {
		ym YAMLMap
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"returns top level keys", args{YAMLMap{"blah": "something", "something_else": "else"}}, []string{"blah", "something_else"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GrabTopLevelKeys(tt.args.ym); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GrabTopLevelKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
