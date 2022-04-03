package options

import (
	"reflect"
	"testing"
)

func TestNewHouse(t *testing.T) {
	type args struct {
		opts []HouseOption
	}

	tests := []struct {
		name string
		args args
		want *House
	}{
		{
			name: "CASE-1",
			args: args{
				opts: []HouseOption{},
			},
			want: &House{
				Material:     defaultMaterialWood,
				HasFireplace: defaultHasFireplace,
				Floors:       defaultFloors,
			},
		},
		{
			name: "CASE-2",
			args: args{
				opts: []HouseOption{
					WithMaterialKerpic(),
				},
			},
			want: &House{
				Material:     defaultMaterialKerpic,
				HasFireplace: defaultHasFireplace,
				Floors:       defaultFloors,
			},
		},
		{
			name: "CASE-3",
			args: args{
				opts: []HouseOption{
					WithMaterialKerpic(),
					WithoutFireplace(),
				},
			},
			want: &House{
				Material:     defaultMaterialKerpic,
				HasFireplace: defaultNoFireplace,
				Floors:       defaultFloors,
			},
		},
		{
			name: "CASE-4",
			args: args{
				opts: []HouseOption{
					WithMaterialKerpic(),
					WithoutFireplace(),
					WithFloors(3),
				},
			},
			want: &House{
				Material:     defaultMaterialKerpic,
				HasFireplace: defaultNoFireplace,
				Floors:       3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHouse(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHouse() = %v, want %v", got, tt.want)
			}
		})
	}
}
