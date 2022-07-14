package chapter1

import "testing"

func TestIsUniqueRunes(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "Unique String 1",
			arg:  "ABCDEFGabcdefg",
			want: true,
		},
		{
			name: "Unique String 2",
			arg:  "ABCDEFGabcdefg",
			want: true,
		},
		{
			name: "Unique String 3",
			arg:  "ABCDEFGabcdefg",
			want: true,
		},
		{
			name: "Duplicate String 1",
			arg:  "AbcdefgA",
			want: false,
		},
		{
			name: "Duplicate String 2",
			arg:  "BcdefghB",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUniqueRunes(tt.arg); got != tt.want {
				t.Errorf("IsUniqueRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUniqueRunes2(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "Unique String 1",
			arg:  "ABCDEFGabcdefg",
			want: true,
		},
		{
			name: "Unique String 2",
			arg:  "ABCDEFGabcdefg",
			want: true,
		},
		{
			name: "Unique String 3",
			arg:  "ABCDEFGabcdefg",
			want: true,
		},
		{
			name: "Duplicate String 1",
			arg:  "AbcdefgA",
			want: false,
		},
		{
			name: "Duplicate String 2",
			arg:  "BcdefghB",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUniqueRunes2(tt.arg); got != tt.want {
				t.Errorf("IsUniqueRunes2() = %v, want %v", got, tt.want)
			}
		})
	}
}
