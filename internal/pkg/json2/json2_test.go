package json2

import (
	"reflect"
	"testing"
)

type Order struct {
	TagName string `json2:"tag_name"`
}

func TestMarshal(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
		{"", args{Order{"abc"}}, []byte(`{"tag_name":"abc"}`), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Marshal(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshal(t *testing.T) {
	o := Order{}
	err := Unmarshal([]byte(`{"tag_name":"abc"}`), &o)
	if err != nil || o.TagName != "abc" {
		t.Error("unmarshal failed")
	}
}
