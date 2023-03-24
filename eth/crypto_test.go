package eth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyHexAddress(t *testing.T) {
	tests := []struct {
		name    string
		address string
		want    bool
	}{
		{"empty", "", false},
		{"short", "0x123456789012345678901234567890123456789", false},
		{"long", "0x12345678901234567890123456789012345678901", false},
		{"invalid", "0x1234567890123456789012345678901234567890g", false},
		{"valid", "0x1234567890123456789012345678901234567890", true},
		{"valid(no prefix)", "1234567890123456789012345678901234567890", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, verifyHexAddress(tt.address))
		})
	}
}
