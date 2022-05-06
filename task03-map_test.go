package homework

import (
	"reflect"
	"testing"
)

func TestSortMapValues(t *testing.T) {
	type args struct {
		arr map[int]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"t1",
			args{map[int]string{2: "a", 0: "b", 1: "c"}},
			[]string{"b", "c", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortMapValues(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortMapValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
