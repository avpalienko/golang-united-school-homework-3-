package homework

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	type args struct {
		arr []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{"t1", args{[]int64{1, 2, 5, 15}}, []int64{15, 5, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
