package gotesttmp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInc(t *testing.T) {
	tests := []struct {
		name string
		i    int
		want int
	}{
		{"inc 1", 1, 2},
		{"inc 2", 2, 3},
		{"inc 3", 3, 4},
		{"inc -1", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Inc(tt.i)
			assert.Equal(t, tt.want, got)
		})
	}
}
