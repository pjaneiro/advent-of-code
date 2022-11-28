package reactorreboot_test

import (
	. "github.com/pjaneiro/advent-of-code/2021/golang/reactorreboot"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []Cuboid
		expected int64
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Cuboid{Cuboid{Type: true, XMin: 10, XMax: 12, YMin: 10, YMax: 12, ZMin: 10, ZMax: 12}, Cuboid{Type: true, XMin: 11, XMax: 13, YMin: 11, YMax: 13, ZMin: 11, ZMax: 13}, Cuboid{Type: false, XMin: 9, XMax: 11, YMin: 9, YMax: 11, ZMin: 9, ZMax: 11}, Cuboid{Type: true, XMin: 10, XMax: 10, YMin: 10, YMax: 10, ZMin: 10, ZMax: 10}},
			expected: 39,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    []Cuboid{Cuboid{Type: true, XMin: -20, XMax: 26, YMin: -36, YMax: 17, ZMin: -47, ZMax: 7}, Cuboid{Type: true, XMin: -20, XMax: 33, YMin: -21, YMax: 23, ZMin: -26, ZMax: 28}, Cuboid{Type: true, XMin: -22, XMax: 28, YMin: -29, YMax: 23, ZMin: -38, ZMax: 16}, Cuboid{Type: true, XMin: -46, XMax: 7, YMin: -6, YMax: 46, ZMin: -50, ZMax: -1}, Cuboid{Type: true, XMin: -49, XMax: 1, YMin: -3, YMax: 46, ZMin: -24, ZMax: 28}, Cuboid{Type: true, XMin: 2, XMax: 47, YMin: -22, YMax: 22, ZMin: -23, ZMax: 27}, Cuboid{Type: true, XMin: -27, XMax: 23, YMin: -28, YMax: 26, ZMin: -21, ZMax: 29}, Cuboid{Type: true, XMin: -39, XMax: 5, YMin: -6, YMax: 47, ZMin: -3, ZMax: 44}, Cuboid{Type: true, XMin: -30, XMax: 21, YMin: -8, YMax: 43, ZMin: -13, ZMax: 34}, Cuboid{Type: true, XMin: -22, XMax: 26, YMin: -27, YMax: 20, ZMin: -29, ZMax: 19}, Cuboid{Type: false, XMin: -48, XMax: -32, YMin: 26, YMax: 41, ZMin: -47, ZMax: -37}, Cuboid{Type: true, XMin: -12, XMax: 35, YMin: 6, YMax: 50, ZMin: -50, ZMax: -2}, Cuboid{Type: false, XMin: -48, XMax: -32, YMin: -32, YMax: -16, ZMin: -15, ZMax: -5}, Cuboid{Type: true, XMin: -18, XMax: 26, YMin: -33, YMax: 15, ZMin: -7, ZMax: 46}, Cuboid{Type: false, XMin: -40, XMax: -22, YMin: -38, YMax: -28, ZMin: 23, ZMax: 41}, Cuboid{Type: true, XMin: -16, XMax: 35, YMin: -41, YMax: 10, ZMin: -47, ZMax: 6}, Cuboid{Type: false, XMin: -32, XMax: -23, YMin: 11, YMax: 30, ZMin: -14, ZMax: 3}, Cuboid{Type: true, XMin: -49, XMax: -5, YMin: -3, YMax: 45, ZMin: -29, ZMax: 18}, Cuboid{Type: false, XMin: 18, XMax: 30, YMin: -20, YMax: -8, ZMin: -3, ZMax: 13}, Cuboid{Type: true, XMin: -41, XMax: 9, YMin: -7, YMax: 43, ZMin: -33, ZMax: 15}, Cuboid{Type: true, XMin: -54112, XMax: -39298, YMin: -85059, YMax: -49293, ZMin: -27449, ZMax: 7877}, Cuboid{Type: true, XMin: 967, XMax: 23432, YMin: 45373, YMax: 81175, ZMin: 27513, ZMax: 53682}},
			expected: 590784,
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
		input    []Cuboid
		expected int64
		error    bool
	}{
		{
			name:     "Example 1",
			input:    []Cuboid{Cuboid{Type: true, XMin: -5, XMax: 47, YMin: -31, YMax: 22, ZMin: -19, ZMax: 33}, Cuboid{Type: true, XMin: -44, XMax: 5, YMin: -27, YMax: 21, ZMin: -14, ZMax: 35}, Cuboid{Type: true, XMin: -49, XMax: -1, YMin: -11, YMax: 42, ZMin: -10, ZMax: 38}, Cuboid{Type: true, XMin: -20, XMax: 34, YMin: -40, YMax: 6, ZMin: -44, ZMax: 1}, Cuboid{Type: false, XMin: 26, XMax: 39, YMin: 40, YMax: 50, ZMin: -2, ZMax: 11}, Cuboid{Type: true, XMin: -41, XMax: 5, YMin: -41, YMax: 6, ZMin: -36, ZMax: 8}, Cuboid{Type: false, XMin: -43, XMax: -33, YMin: -45, YMax: -28, ZMin: 7, ZMax: 25}, Cuboid{Type: true, XMin: -33, XMax: 15, YMin: -32, YMax: 19, ZMin: -34, ZMax: 11}, Cuboid{Type: false, XMin: 35, XMax: 47, YMin: -46, YMax: -34, ZMin: -11, ZMax: 5}, Cuboid{Type: true, XMin: -14, XMax: 36, YMin: -6, YMax: 44, ZMin: -16, ZMax: 29}, Cuboid{Type: true, XMin: -57795, XMax: -6158, YMin: 29564, YMax: 72030, ZMin: 20435, ZMax: 90618}, Cuboid{Type: true, XMin: 36731, XMax: 105352, YMin: -21140, YMax: 28532, ZMin: 16094, ZMax: 90401}, Cuboid{Type: true, XMin: 30999, XMax: 107136, YMin: -53464, YMax: 15513, ZMin: 8553, ZMax: 71215}, Cuboid{Type: true, XMin: 13528, XMax: 83982, YMin: -99403, YMax: -27377, ZMin: -24141, ZMax: 23996}, Cuboid{Type: true, XMin: -72682, XMax: -12347, YMin: 18159, YMax: 111354, ZMin: 7391, ZMax: 80950}, Cuboid{Type: true, XMin: -1060, XMax: 80757, YMin: -65301, YMax: -20884, ZMin: -103788, ZMax: -16709}, Cuboid{Type: true, XMin: -83015, XMax: -9461, YMin: -72160, YMax: -8347, ZMin: -81239, ZMax: -26856}, Cuboid{Type: true, XMin: -52752, XMax: 22273, YMin: -49450, YMax: 9096, ZMin: 54442, ZMax: 119054}, Cuboid{Type: true, XMin: -29982, XMax: 40483, YMin: -108474, YMax: -28371, ZMin: -24328, ZMax: 38471}, Cuboid{Type: true, XMin: -4958, XMax: 62750, YMin: 40422, YMax: 118853, ZMin: -7672, ZMax: 65583}, Cuboid{Type: true, XMin: 55694, XMax: 108686, YMin: -43367, YMax: 46958, ZMin: -26781, ZMax: 48729}, Cuboid{Type: true, XMin: -98497, XMax: -18186, YMin: -63569, YMax: 3412, ZMin: 1232, ZMax: 88485}, Cuboid{Type: true, XMin: -726, XMax: 56291, YMin: -62629, YMax: 13224, ZMin: 18033, ZMax: 85226}, Cuboid{Type: true, XMin: -110886, XMax: -34664, YMin: -81338, YMax: -8658, ZMin: 8914, ZMax: 63723}, Cuboid{Type: true, XMin: -55829, XMax: 24974, YMin: -16897, YMax: 54165, ZMin: -121762, ZMax: -28058}, Cuboid{Type: true, XMin: -65152, XMax: -11147, YMin: 22489, YMax: 91432, ZMin: -58782, ZMax: 1780}, Cuboid{Type: true, XMin: -120100, XMax: -32970, YMin: -46592, YMax: 27473, ZMin: -11695, ZMax: 61039}, Cuboid{Type: true, XMin: -18631, XMax: 37533, YMin: -124565, YMax: -50804, ZMin: -35667, ZMax: 28308}, Cuboid{Type: true, XMin: -57817, XMax: 18248, YMin: 49321, YMax: 117703, ZMin: 5745, ZMax: 55881}, Cuboid{Type: true, XMin: 14781, XMax: 98692, YMin: -1341, YMax: 70827, ZMin: 15753, ZMax: 70151}, Cuboid{Type: true, XMin: -34419, XMax: 55919, YMin: -19626, YMax: 40991, ZMin: 39015, ZMax: 114138}, Cuboid{Type: true, XMin: -60785, XMax: 11593, YMin: -56135, YMax: 2999, ZMin: -95368, ZMax: -26915}, Cuboid{Type: true, XMin: -32178, XMax: 58085, YMin: 17647, YMax: 101866, ZMin: -91405, ZMax: -8878}, Cuboid{Type: true, XMin: -53655, XMax: 12091, YMin: 50097, YMax: 105568, ZMin: -75335, ZMax: -4862}, Cuboid{Type: true, XMin: -111166, XMax: -40997, YMin: -71714, YMax: 2688, ZMin: 5609, ZMax: 50954}, Cuboid{Type: true, XMin: -16602, XMax: 70118, YMin: -98693, YMax: -44401, ZMin: 5197, ZMax: 76897}, Cuboid{Type: true, XMin: 16383, XMax: 101554, YMin: 4615, YMax: 83635, ZMin: -44907, ZMax: 18747}, Cuboid{Type: false, XMin: -95822, XMax: -15171, YMin: -19987, YMax: 48940, ZMin: 10804, ZMax: 104439}, Cuboid{Type: true, XMin: -89813, XMax: -14614, YMin: 16069, YMax: 88491, ZMin: -3297, ZMax: 45228}, Cuboid{Type: true, XMin: 41075, XMax: 99376, YMin: -20427, YMax: 49978, ZMin: -52012, ZMax: 13762}, Cuboid{Type: true, XMin: -21330, XMax: 50085, YMin: -17944, YMax: 62733, ZMin: -112280, ZMax: -30197}, Cuboid{Type: true, XMin: -16478, XMax: 35915, YMin: 36008, YMax: 118594, ZMin: -7885, ZMax: 47086}, Cuboid{Type: false, XMin: -98156, XMax: -27851, YMin: -49952, YMax: 43171, ZMin: -99005, ZMax: -8456}, Cuboid{Type: false, XMin: 2032, XMax: 69770, YMin: -71013, YMax: 4824, ZMin: 7471, ZMax: 94418}, Cuboid{Type: true, XMin: 43670, XMax: 120875, YMin: -42068, YMax: 12382, ZMin: -24787, ZMax: 38892}, Cuboid{Type: false, XMin: 37514, XMax: 111226, YMin: -45862, YMax: 25743, ZMin: -16714, ZMax: 54663}, Cuboid{Type: false, XMin: 25699, XMax: 97951, YMin: -30668, YMax: 59918, ZMin: -15349, ZMax: 69697}, Cuboid{Type: false, XMin: -44271, XMax: 17935, YMin: -9516, YMax: 60759, ZMin: 49131, ZMax: 112598}, Cuboid{Type: true, XMin: -61695, XMax: -5813, YMin: 40978, YMax: 94975, ZMin: 8655, ZMax: 80240}, Cuboid{Type: false, XMin: -101086, XMax: -9439, YMin: -7088, YMax: 67543, ZMin: 33935, ZMax: 83858}, Cuboid{Type: false, XMin: 18020, XMax: 114017, YMin: -48931, YMax: 32606, ZMin: 21474, ZMax: 89843}, Cuboid{Type: false, XMin: -77139, XMax: 10506, YMin: -89994, YMax: -18797, ZMin: -80, ZMax: 59318}, Cuboid{Type: false, XMin: 8476, XMax: 79288, YMin: -75520, YMax: 11602, ZMin: -96624, ZMax: -24783}, Cuboid{Type: true, XMin: -47488, XMax: -1262, YMin: 24338, YMax: 100707, ZMin: 16292, ZMax: 72967}, Cuboid{Type: false, XMin: -84341, XMax: 13987, YMin: 2429, YMax: 92914, ZMin: -90671, ZMax: -1318}, Cuboid{Type: false, XMin: -37810, XMax: 49457, YMin: -71013, YMax: -7894, ZMin: -105357, ZMax: -13188}, Cuboid{Type: false, XMin: -27365, XMax: 46395, YMin: 31009, YMax: 98017, ZMin: 15428, ZMax: 76570}, Cuboid{Type: false, XMin: -70369, XMax: -16548, YMin: 22648, YMax: 78696, ZMin: -1892, ZMax: 86821}, Cuboid{Type: true, XMin: -53470, XMax: 21291, YMin: -120233, YMax: -33476, ZMin: -44150, ZMax: 38147}, Cuboid{Type: false, XMin: -93533, XMax: -4276, YMin: -16170, YMax: 68771, ZMin: -104985, ZMax: -24507}},
			expected: 2758514936282235,
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
