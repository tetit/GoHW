package taskTests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	tests := map[string]struct {
		inputI int
		inputS string
		inputD bool
		expRez string
		expErr string
	}{
		"valid I first case valid S D":  {inputI: 3, inputS: "d", inputD: true, expRez: "[3D]", expErr: ""},
		"valid I second case valid S D": {inputI: 21, inputS: "d", inputD: true, expRez: "D", expErr: ""},
		"invalid I valid S D":           {inputI: 10, inputS: "d", inputD: true, expRez: "", expErr: "invalid i"},
		"invalid S valid I D":           {inputI: 3, inputS: "e", inputD: true, expRez: "", expErr: "invalid s"},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			actual, err := Do(tt.inputS, tt.inputI, tt.inputD)
			assert.Equal(t, tt.expRez, actual)

			if tt.expErr != "" {
				assert.EqualError(t, err, tt.expErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
