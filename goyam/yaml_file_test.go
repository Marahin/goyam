package goyam

import (
	"reflect"
	"testing"
)

func TestLoadFile(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// We don't need to test the loading, as it's the part of STD library
		{"load inexistent file", args{"./blah"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadFile(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LoadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadYAMLFile(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    YAMLMap
		wantErr bool
	}{
		{"load small YAML file", args{"../examples/yaml files/smaller_yaml_1.yml"}, YAMLMap{"name": "John", "height": "123", "data": YAMLMap{"age": 29, "profession": "Plumber"}}, false},
		{"load inexistent YAML file", args{"./blah"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadYAMLFile(tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadYAMLFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadYAMLFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
