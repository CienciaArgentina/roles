package adapter

import (
	"reflect"
	"testing"
)

func TestAdapt(t *testing.T) {
	type args struct {
		o interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "slice_adapter",
			args: args{
				o: []string{"1", "2", "3"},
			},
			want: map[string]interface{}{
				"total":   3,
				"results": []string{"1", "2", "3"},
			},
		},
		{
			name: "not_slice_adapter",
			args: args{
				o: "Hello There Traveller",
			},
			want: map[string]interface{}{
				"total":   1,
				"results": []interface{}{"Hello There Traveller"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Adapt(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Adapt() = %v, want %v", got, tt.want)
			}
		})
	}
}
