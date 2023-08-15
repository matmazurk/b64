package b64

import (
	"fmt"
	"testing"
)

func Test_bits_addLeft(t *testing.T) {
	tests := []struct {
		input    *bits
		toAdd    *bits
		expected *bits
	}{
		{
			input:    newEmptyBits(),
			toAdd:    newEmptyBits(),
			expected: newEmptyBits(),
		},
		{
			input:    newBits(5, 0b10000),
			toAdd:    newEmptyBits(),
			expected: newBits(5, 0b10000),
		},
		{
			input:    newEmptyBits(),
			toAdd:    newBits(5, 0b10000),
			expected: newBits(5, 0b10000),
		},
		{
			input:    newBits(5, 0b11111),
			toAdd:    newBits(5, 0b10000),
			expected: newBits(10, 0b1000011111),
		},
	}
	for _, tc := range tests {
		t.Run(
			fmt.Sprintf("%s.addLeft(%s) should result in %s", tc.input, tc.toAdd, tc.expected),
			func(t *testing.T) {
				in := tc.input
				in.addLeft(tc.toAdd)

				if !in.equals(tc.expected) {
					t.Errorf("actual different than expected: \n%s != \n%s", in, tc.expected)
				}
			},
		)
	}
}

func Test_bits_addRight(t *testing.T) {
	tests := []struct {
		input    *bits
		toAdd    *bits
		expected *bits
	}{
		{
			input:    newEmptyBits(),
			toAdd:    newEmptyBits(),
			expected: newEmptyBits(),
		},
		{
			input:    newBits(5, 0b10000),
			toAdd:    newEmptyBits(),
			expected: newBits(5, 0b10000),
		},
		{
			input:    newEmptyBits(),
			toAdd:    newBits(5, 0b10000),
			expected: newBits(5, 0b10000),
		},
		{
			input:    newBits(5, 0b11111),
			toAdd:    newBits(5, 0b10000),
			expected: newBits(10, 0b1111110000),
		},
	}
	for _, tc := range tests {
		t.Run(
			fmt.Sprintf("%s.addRight(%s) should result in %s", tc.input, tc.toAdd, tc.expected),
			func(t *testing.T) {
				in := tc.input
				in.addRight(tc.toAdd)

				if !in.equals(tc.expected) {
					t.Errorf("actual different than expected: \n%s != \n%s", in, tc.expected)
				}
			},
		)
	}
}

func Test_bits_cutSignificantBits(t *testing.T) {
	tests := []struct {
		in          *bits
		toCut       int
		expectedCut *bits
		expectedIn  *bits
	}{
		{
			in:          newEmptyBits(),
			toCut:       0,
			expectedCut: newEmptyBits(),
			expectedIn:  newEmptyBits(),
		},
		{
			in:          newBits(6, 0b110011),
			toCut:       3,
			expectedCut: newBits(3, 0b110),
			expectedIn:  newBits(3, 0b011),
		},
		{
			in:          newBits(8, 0b00110011),
			toCut:       5,
			expectedCut: newBits(5, 0b00110),
			expectedIn:  newBits(3, 0b011),
		},
	}
	for _, tc := range tests {
		t.Run(
			fmt.Sprintf("%s.cutSignificantBits(%d) should result in %s", tc.in, tc.toCut, tc.expectedCut),
			func(t *testing.T) {
				in := tc.in
				cut := in.cutSignificantBits(tc.toCut)

				if !cut.equals(tc.expectedCut) {
					t.Errorf("actual cut different than expected: \n%s != \n%s", cut, tc.expectedCut)
				}

				if !in.equals(tc.expectedIn) {
					t.Errorf("actual in different than expected: \n%s != \n%s", in, tc.expectedIn)
				}
			},
		)
	}
}
