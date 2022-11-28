package packetdecoder_test

import (
	. "github.com/pjaneiro/advent-of-code/2021/golang/packetdecoder"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected int64
		error    bool
	}{
		{
			name:     "Example 1",
			input:    "8A004A801A8002F478",
			expected: 16,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    "620080001611562C8802118E34",
			expected: 12,
			error:    false,
		},
		{
			name:     "Example 3",
			input:    "C0015000016115A2E0802F182340",
			expected: 23,
			error:    false,
		},
		{
			name:     "Example 4",
			input:    "A0016C880162017C3686B18A3D4780",
			expected: 31,
			error:    false,
		},
		{
			name:     "Example 5",
			input:    "D2FE28",
			expected: 6,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.input)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v) threw '%v', want %d", tc.input, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v) = %d, want to throw", tc.input, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v) = %d, want %d", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		input    string
		expected int64
		error    bool
	}{
		{
			name:     "Example 1",
			input:    "8A004A801A8002F478",
			expected: 15,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    "620080001611562C8802118E34",
			expected: 46,
			error:    false,
		},
		{
			name:     "Example 3",
			input:    "C0015000016115A2E0802F182340",
			expected: 46,
			error:    false,
		},
		{
			name:     "Example 4",
			input:    "A0016C880162017C3686B18A3D4780",
			expected: 54,
			error:    false,
		},
		{
			name:     "Example 5",
			input:    "D2FE28",
			expected: 2021,
			error:    false,
		},
		{
			name:     "Example 6",
			input:    "C200B40A82",
			expected: 3,
			error:    false,
		},
		{
			name:     "Example 7",
			input:    "04005AC33890",
			expected: 54,
			error:    false,
		},
		{
			name:     "Example 8",
			input:    "880086C3E88112",
			expected: 7,
			error:    false,
		},
		{
			name:     "Example 9",
			input:    "CE00C43D881120",
			expected: 9,
			error:    false,
		},
		{
			name:     "Example 10",
			input:    "D8005AC2A8F0",
			expected: 1,
			error:    false,
		},
		{
			name:     "Example 11",
			input:    "F600BC2D8F",
			expected: 0,
			error:    false,
		},
		{
			name:     "Example 12",
			input:    "9C005AC2F8F0",
			expected: 0,
			error:    false,
		},
		{
			name:     "Example 13",
			input:    "9C0141080250320F1802104A08",
			expected: 1,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.input)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v) threw '%v', want %d", tc.input, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v) = %d, want to throw", tc.input, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v) = %d, want %d", tc.input, actual, tc.expected)
			}
		})
	}
}
