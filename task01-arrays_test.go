package homework

import (
    "fmt"
    "testing"
)

func TestAverage(t *testing.T) {
    type args struct {
        arr [15]float32
    }
    tests := []struct {
        name string
        args args
        want float64
    }{
        {"t1", args{[15]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}}, 8},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Average(tt.args.arr); fmt.Sprintf("%.5f", got) != fmt.Sprintf("%.5f", tt.want) {
                t.Errorf("average() = %v, want %v", got, tt.want)
            }
        })
    }
}
