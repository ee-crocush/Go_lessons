package pkg

import (
	"reflect"
	"testing"
)

func Test_someHelperFunction(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "basic test 1",
			args: args{s: "a,b,c"},
			want: []string{"a", "b", "c"},
		},
		{
			name: "basic test 2",
			args: args{s: "a,b;c"},
			want: []string{"a", "bc"},
		},
		{
			name: "basic test 3",
			args: args{s: "a,"},
			want: []string{"a", ""},
		},
		{
			name: "basic test 4",
			args: args{s: ","},
			want: []string{"", ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := someHelperFunction(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("someHelperFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}
