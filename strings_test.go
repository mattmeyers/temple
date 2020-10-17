package temple

import (
	"testing"
)

func Test_numCommas(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{s: "1"},
			want: 0,
		},
		{
			name: "1,234",
			args: args{s: "1234"},
			want: 1,
		},
		{
			name: "1,234,567",
			args: args{s: "1234567"},
			want: 2,
		},
		{
			name: "-1",
			args: args{s: "-1"},
			want: 0,
		},
		{
			name: "-123",
			args: args{s: "-123"},
			want: 0,
		},
		{
			name: "-1,234",
			args: args{s: "-1234"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numCommas(tt.args.s); got != tt.want {
				t.Errorf("numCommas() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommas(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "1",
			args:    args{s: "1"},
			want:    "1",
			wantErr: false,
		},
		{
			name:    "1,234",
			args:    args{s: "1234"},
			want:    "1,234",
			wantErr: false,
		},
		{
			name:    "1,234,567",
			args:    args{s: "1234567"},
			want:    "1,234,567",
			wantErr: false,
		},
		{
			name:    "-1",
			args:    args{s: "-1"},
			want:    "-1",
			wantErr: false,
		},
		{
			name:    "-123",
			args:    args{s: "-123"},
			want:    "-123",
			wantErr: false,
		},
		{
			name:    "-1,234",
			args:    args{s: "-1234"},
			want:    "-1,234",
			wantErr: false,
		},
		{
			name:    "1,234.56",
			args:    args{s: "1234.56"},
			want:    "1,234.56",
			wantErr: false,
		},
		{
			name:    ".56",
			args:    args{s: ".56"},
			want:    ".56",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Commas(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Commas() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Commas() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumeric(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "123",
			args: args{s: "123"},
			want: true,
		},
		{
			name: "-123",
			args: args{s: "-123"},
			want: true,
		},
		{
			name: "-",
			args: args{s: "-"},
			want: false,
		},
		{
			name: "123.5",
			args: args{s: "123.5"},
			want: true,
		},
		{
			name: "-123.5",
			args: args{s: "-123.5"},
			want: true,
		},
		{
			name: "-1-23.5",
			args: args{s: "-1-23.5"},
			want: false,
		},
		{
			name: "1.23.5",
			args: args{s: "1.23.5"},
			want: false,
		},
		{
			name: "-1,234.5",
			args: args{s: "-1,234.5"},
			want: true,
		},
		{
			name: "1,234,567",
			args: args{s: "1,234,567"},
			want: true,
		},
		{
			name: "1234567.56",
			args: args{s: "1234567.56"},
			want: true,
		},
		{
			name: ".567",
			args: args{s: ".567"},
			want: true,
		},
		{
			name: "empty",
			args: args{s: ""},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumeric(tt.args.s); got != tt.want {
				t.Errorf("isNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatMask(t *testing.T) {
	type args struct {
		mask string
		s    string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "US phone number",
			args:    args{mask: "(###) ###-####", s: "5551234567"},
			want:    "(555) 123-4567",
			wantErr: false,
		},
		{
			name:    "Invalid mask length",
			args:    args{mask: "#", s: "12"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Too few string characters",
			args:    args{mask: "##", s: "1"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Too many string characters",
			args:    args{mask: "##", s: "123"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Emoji support",
			args:    args{mask: "🙁 # #", s: "🙂🙃"},
			want:    "🙁 🙂 🙃",
			wantErr: false,
		},
		{
			name:    "Contains #",
			args:    args{mask: `\##\##abc\#`, s: "12"},
			want:    "#1#2abc#",
			wantErr: false,
		},
		{
			name:    "Empty string",
			args:    args{mask: "mask", s: ""},
			want:    "mask",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FormatMask(tt.args.mask, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("FormatMask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FormatMask() = %v, want %v", got, tt.want)
			}
		})
	}
}
