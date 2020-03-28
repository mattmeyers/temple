package temple

import (
	"reflect"
	"testing"
)

func Test_parseIntArgs(t *testing.T) {
	type args struct {
		arg1 interface{}
		arg2 []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "ints",
			args:    args{arg1: 1, arg2: []interface{}{2, 3}},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "single int",
			args:    args{arg1: 1, arg2: nil},
			want:    []int{1},
			wantErr: false,
		},
		{
			name:    "interface slice",
			args:    args{arg1: []interface{}{1, 2, 3}, arg2: nil},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "Slice collection",
			args:    args{arg1: NewSlice(1, 2, 3), arg2: nil},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "int slice",
			args:    args{arg1: []int{1, 2, 3}, arg2: nil},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseIntArgs(tt.args.arg1, tt.args.arg2...)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseIntArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseIntArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
