package quicksort

import "testing"

func Test_qsort(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "quick-1",
			args: args{
				data: []int{-100, -99, -98, 0, 98, 99, 100, 0, 1, -100, 98},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("before:%+v", tt.args.data)
			qsort(tt.args.data)
			t.Logf("after:%+v", tt.args.data)
		})
	}
}
