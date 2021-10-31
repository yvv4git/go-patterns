package sharding

import "testing"

func TestCalculatorMD5_CalculateShardIndex(t *testing.T) {
	type args struct {
		key        string
		shardCount int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				key:        "red",
				shardCount: 5,
			},
			want: 2,
		},
		{
			name: "CASE-2",
			args: args{
				key:        "orange",
				shardCount: 5,
			},
			want: 0,
		},
		{
			name: "CASE-3",
			args: args{
				key:        "green",
				shardCount: 5,
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := CalculatorMD5{}
			if got := i.CalculateShardIndex(tt.args.key, tt.args.shardCount); got != tt.want {
				t.Errorf("CalculateShardIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
