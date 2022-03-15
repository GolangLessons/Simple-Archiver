package shannon_fano

import (
	"reflect"
	"testing"
)

func Test_countCharacters(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want charStat
	}{
		{
			name: "base test",
			str:  "zaaza",
			want: charStat{
				'z': 2,
				'a': 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newCharStat(tt.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newCharStat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_abs(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "positive number",
			x:    4,
			want: 4,
		},
		{
			name: "negative number",
			x:    -4,
			want: 4,
		},
		{
			name: "zero",
			x:    0,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := abs(tt.x); got != tt.want {
				t.Errorf("abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_My(t *testing.T) {

}

func Test_bestDividePosition(t *testing.T) {
	tests := []struct {
		name  string
		codes []Code
		want  int
	}{
		{
			name: "one element",
			codes: []Code{
				{Quantity: 2},
			},
			want: 0,
		},
		{
			name: "one element",
			codes: []Code{
				{Quantity: 2},
			},
			want: 0,
		},
		{
			name: "two elements",
			codes: []Code{
				{Quantity: 2},
				{Quantity: 2},
			},
			want: 1,
		},
		{
			name: "certain position",
			codes: []Code{
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
		{
			name: "certain position",
			codes: []Code{
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
			name: "uncertainty (need leftmost)",
			codes: []Code{
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: 1,
		},
		{
			name: "uncertainty (need leftmost)",
			codes: []Code{
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
			if got := bestDividePosition(tt.codes); got != tt.want {
				t.Errorf("bestDividePosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assignCodes(t *testing.T) {
	tests := []struct {
		name  string
		codes []Code
		want  []Code
	}{
		//{
		//	name: "one element",
		//	codes: []Code{
		//		{Quantity: 2},
		//	},
		//	want: []Code{
		//		{Quantity: 2, Bits: 0, Size: 1},
		//	},
		//},
		{
			name: "two elements",
			codes: []Code{
				{Quantity: 2},
				{Quantity: 2},
			},
			want: []Code{
				{Quantity: 2, Bits: 0, Size: 1},
				{Quantity: 2, Bits: 1, Size: 1},
			},
		},
		{
			name: "certain position",
			codes: []Code{
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: []Code{
				{Quantity: 2, Bits: 0, Size: 1}, // 0
				{Quantity: 1, Bits: 2, Size: 2}, // 11
				{Quantity: 1, Bits: 3, Size: 2}, // 10
			},
		},
		{
			name: "uncertainty (need leftmost)",
			codes: []Code{
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: []Code{
				{Quantity: 1, Bits: 0, Size: 1}, // 0
				{Quantity: 1, Bits: 2, Size: 2}, // 11
				{Quantity: 1, Bits: 3, Size: 2}, // 10
			},
		},
		{
			name: "uncertainty (need leftmost)",
			codes: []Code{
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: []Code{
				{Quantity: 1, Bits: 0, Size: 2}, // 00
				{Quantity: 1, Bits: 1, Size: 2}, // 01
				{Quantity: 1, Bits: 2, Size: 2}, // 10
				{Quantity: 1, Bits: 3, Size: 2}, // 11
			},
		},
		{
			name: "uncertainty (need leftmost)",
			codes: []Code{
				{Quantity: 2},
				{Quantity: 2},
				{Quantity: 1},
				{Quantity: 1},
			},
			want: []Code{
				{Quantity: 2, Bits: 0, Size: 1}, // 0
				{Quantity: 2, Bits: 2, Size: 2}, // 10
				{Quantity: 1, Bits: 6, Size: 3}, // 110
				{Quantity: 1, Bits: 7, Size: 3}, // 111
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assignCodes(tt.codes)

			if !reflect.DeepEqual(tt.codes, tt.want) {
				t.Errorf("Encode() = %v, want %v", tt.codes, tt.want)
			}
		})
	}
}

func Test_build(t *testing.T) {
	tests := []struct {
		name string
		text string
		want Table
	}{
		//{
		//	name: "solution",
		//	text: ``,
		//},
		{
			name: "base test",
			text: "abbbcc",
			want: Table{
				'a': Code{
					Char:     'a',
					Quantity: 1,
					Bits:     3,
					Size:     2,
				},
				'b': Code{
					Char:     'b',
					Quantity: 3,
					Bits:     0,
					Size:     1,
				},
				'c': Code{
					Char:     'c',
					Quantity: 2,
					Bits:     2,
					Size:     2,
				},
			},
		},
		{
			name: "equal number of letters",
			text: "aaccbb",
			want: Table{
				'a': Code{
					Char:     'a',
					Quantity: 2,
					Bits:     0,
					Size:     1,
				},
				'b': Code{
					Char:     'b',
					Quantity: 2,
					Bits:     2,
					Size:     2,
				},
				'c': Code{
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
