package blankendline

import (
	"testing"
)

func TestIsComment(t *testing.T) {
	tests := []struct {
		name string
		init map[int]struct{}
		arg  int
		want bool
	}{
		{
			name: "Exist",
			init: map[int]struct{}{
				2:  {},
				4:  {},
				15: {},
			},
			arg:  4,
			want: true,
		},
		{
			name: "Not Exist",
			init: map[int]struct{}{
				2:  {},
				4:  {},
				15: {},
			},
			arg:  3,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := checker{
				comments: tt.init,
			}

			if got := c.isComment(tt.arg); got != tt.want {
				t.Errorf("checker.isComment() = %v, want: %v", got, tt.want)
			}
		})
	}
}
