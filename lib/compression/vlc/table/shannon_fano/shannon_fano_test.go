package shannon_fano

import (
	"reflect"
	"testing"
)

func Test_bestDividerPosition(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  int
	}{
		{
			name: "one element",
			codes: []code{
				{Quantity: 2},
			},
			want: 0,
		},
		{
			name: "two elements",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
			},
			want: 1,
		},
		{
			name: "three elements",
			codes: []code{
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
		{
			name: "many elements",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 2,
		},
		{
			name: "uncertainty (need rightmost)",
			codes: []code{
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
		{
			name: "uncertainty (need rightmost)",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestDividerPosition(tt.codes); got != tt.want {
				t.Errorf("bestDividerPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assignCodes(t *testing.T) {
	tests := []struct {
		name  string
		codes []code
		want  []code
	}{
		{
			name: "two elements",
			codes: []code{
				{Quantity: 2},
				{Quantity: 2},
			},
			want: []code{
				{Quantity: 2, Bits: 0, Size: 1},
				{Quantity: 2, Bits: 1, Size: 1},
			},
		},
		{
			name: "three elements, certain position",
			codes: []code{
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: []code{
				{Quantity: 2, Bits: 0, Size: 1}, // 0
				{Quantity: 1, Bits: 2, Size: 2}, // 10
				{Quantity: 1, Bits: 3, Size: 2}, // 11
			},
		},
		{
			name: "three elements, uncertain position",
			codes: []code{
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: []code{
				{Quantity: 1, Bits: 0, Size: 1}, // 0
				{Quantity: 1, Bits: 2, Size: 2}, // 10
				{Quantity: 1, Bits: 3, Size: 2}, // 11
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assignCodes(tt.codes)

			if !reflect.DeepEqual(tt.codes, tt.want) {
				t.Errorf("got: %v, want: %v", tt.codes, tt.want)
			}
		})
	}
}

func Test_build(t *testing.T) {
	tests := []struct {
		name string
		text string
		want encodingTable
	}{
		{
			name: "base test",
			text: "abbbcc",
			want: encodingTable{
				'a': code{
					Char:     'a',
					Quantity: 1,
					Bits:     3,
					Size:     2,
				},
				'b': code{
					Char:     'b',
					Quantity: 3,
					Bits:     0,
					Size:     1,
				},
				'c': code{
					Char:     'c',
					Quantity: 2,
					Bits:     2,
					Size:     2,
				},
			},
		},
		{
			name: "base test",
			text: "aaccbb",
			want: encodingTable{
				'a': code{
					Char:     'a',
					Quantity: 2,
					Bits:     0,
					Size:     1,
				},
				'b': code{
					Char:     'b',
					Quantity: 2,
					Bits:     2,
					Size:     2,
				},
				'c': code{
					Char:     'c',
					Quantity: 2,
					Bits:     3,
					Size:     2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := build(newCharStat(tt.text)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("build() = %v, want %v", got, tt.want)
			}
		})
	}
}
