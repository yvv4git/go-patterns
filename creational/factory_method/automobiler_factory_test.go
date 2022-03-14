package factory_method

import "testing"

func TestAutomobilerFactory(t *testing.T) {
	type args struct {
		avtoType AutomobileType
	}

	tests := []struct {
		name string
		args args
		want Automobiler
	}{
		{
			name: "CASE-1",
			args: args{
				avtoType: TypeRenaultLogan,
			},
			want: NewRenoLogan(),
		},
		{
			name: "CASE-2",
			args: args{
				avtoType: TypeToyotaCamry,
			},
			want: NewToyotaCamry(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AutomobilerFactory(tt.args.avtoType); got != tt.want {
				t.Fatalf("We want %v, but got %v", tt.want, got)
			}
		})
	}
}
