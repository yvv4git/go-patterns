package prototype

import (
	"reflect"
	"testing"
)

func TestSheep_Clone(t *testing.T) {
	type fields struct {
		name string
	}

	tests := []struct {
		name   string
		fields fields
		want   Mutable
	}{
		{
			name: "CASE-1",
			fields: fields{
				name: "Eva",
			},
			want: &Sheep{
				name: "Eva_Dolly",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSheep(tt.fields.name)

			if got := s.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sheep.Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}
