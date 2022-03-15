package table

import (
	"reflect"
	"testing"
)

func Test_encodingTable_DecodingTree(t *testing.T) {
	tests := []struct {
		name string
		et   EncodingTable
		want decodingTree
	}{
		{
			name: "base tree test",
			et: EncodingTable{
				'a': "11",
				'b': "1001",
				'z': "0101",
			},
			want: decodingTree{
				Zero: &decodingTree{
					One: &decodingTree{
						Zero: &decodingTree{
							One: &decodingTree{
								Value: "z",
							},
						},
					},
				},
				One: &decodingTree{
					Zero: &decodingTree{
						Zero: &decodingTree{
							One: &decodingTree{
								Value: "b",
							},
						},
					},
					One: &decodingTree{
						Value: "a",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.et.decodingTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodingTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
