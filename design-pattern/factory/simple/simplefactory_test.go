package main

import (
	"reflect"
	"testing"
)

func TestCreateFood(t *testing.T) {
	type args struct {
		objName string
	}
	tests := []struct {
		name string
		args args
		want Food
	}{
		{
			name: "Meat",
			args: args{objName: "Meat"},
			want: &Meat{},
		},
		{
			name: "Hamberger",
			args: args{objName: "Hamberger"},
			want: &Hamberger{},
		},
	}

	f := FoodFactory{}
	for _, v := range tests {
		tt := v
		t.Run(tt.name, func(t *testing.T) {
			if got := f.CreateFood(tt.args.objName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateFood() = %v, want %v", got, tt.want)
			}
		})
	}
}
