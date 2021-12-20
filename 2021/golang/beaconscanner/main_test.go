package beaconscanner_test

import (
	. "github.com/pjaneiro/advent-of-code-2021/beaconscanner"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		data     []Scanner
		expected int
		error    bool
	}{
		{
			name: "Example 1",
			data: []Scanner{
				Scanner{ID: 0, Beacons: []Beacon{Beacon{Pos: [3]int{404, -588, -901}}, Beacon{Pos: [3]int{528, -643, 409}}, Beacon{Pos: [3]int{-838, 591, 734}}, Beacon{Pos: [3]int{390, -675, -793}}, Beacon{Pos: [3]int{-537, -823, -458}}, Beacon{Pos: [3]int{-485, -357, 347}}, Beacon{Pos: [3]int{-345, -311, 381}}, Beacon{Pos: [3]int{-661, -816, -575}}, Beacon{Pos: [3]int{-876, 649, 763}}, Beacon{Pos: [3]int{-618, -824, -621}}, Beacon{Pos: [3]int{553, 345, -567}}, Beacon{Pos: [3]int{474, 580, 667}}, Beacon{Pos: [3]int{-447, -329, 318}}, Beacon{Pos: [3]int{-584, 868, -557}}, Beacon{Pos: [3]int{544, -627, -890}}, Beacon{Pos: [3]int{564, 392, -477}}, Beacon{Pos: [3]int{455, 729, 728}}, Beacon{Pos: [3]int{-892, 524, 684}}, Beacon{Pos: [3]int{-689, 845, -530}}, Beacon{Pos: [3]int{423, -701, 434}}, Beacon{Pos: [3]int{7, -33, -71}}, Beacon{Pos: [3]int{630, 319, -379}}, Beacon{Pos: [3]int{443, 580, 662}}, Beacon{Pos: [3]int{-789, 900, -551}}, Beacon{Pos: [3]int{459, -707, 401}}}},
				Scanner{ID: 1, Beacons: []Beacon{Beacon{Pos: [3]int{686, 422, 578}}, Beacon{Pos: [3]int{605, 423, 415}}, Beacon{Pos: [3]int{515, 917, -361}}, Beacon{Pos: [3]int{-336, 658, 858}}, Beacon{Pos: [3]int{95, 138, 22}}, Beacon{Pos: [3]int{-476, 619, 847}}, Beacon{Pos: [3]int{-340, -569, -846}}, Beacon{Pos: [3]int{567, -361, 727}}, Beacon{Pos: [3]int{-460, 603, -452}}, Beacon{Pos: [3]int{669, -402, 600}}, Beacon{Pos: [3]int{729, 430, 532}}, Beacon{Pos: [3]int{-500, -761, 534}}, Beacon{Pos: [3]int{-322, 571, 750}}, Beacon{Pos: [3]int{-466, -666, -811}}, Beacon{Pos: [3]int{-429, -592, 574}}, Beacon{Pos: [3]int{-355, 545, -477}}, Beacon{Pos: [3]int{703, -491, -529}}, Beacon{Pos: [3]int{-328, -685, 520}}, Beacon{Pos: [3]int{413, 935, -424}}, Beacon{Pos: [3]int{-391, 539, -444}}, Beacon{Pos: [3]int{586, -435, 557}}, Beacon{Pos: [3]int{-364, -763, -893}}, Beacon{Pos: [3]int{807, -499, -711}}, Beacon{Pos: [3]int{755, -354, -619}}, Beacon{Pos: [3]int{553, 889, -390}}}},
				Scanner{ID: 2, Beacons: []Beacon{Beacon{Pos: [3]int{649, 640, 665}}, Beacon{Pos: [3]int{682, -795, 504}}, Beacon{Pos: [3]int{-784, 533, -524}}, Beacon{Pos: [3]int{-644, 584, -595}}, Beacon{Pos: [3]int{-588, -843, 648}}, Beacon{Pos: [3]int{-30, 6, 44}}, Beacon{Pos: [3]int{-674, 560, 763}}, Beacon{Pos: [3]int{500, 723, -460}}, Beacon{Pos: [3]int{609, 671, -379}}, Beacon{Pos: [3]int{-555, -800, 653}}, Beacon{Pos: [3]int{-675, -892, -343}}, Beacon{Pos: [3]int{697, -426, -610}}, Beacon{Pos: [3]int{578, 704, 681}}, Beacon{Pos: [3]int{493, 664, -388}}, Beacon{Pos: [3]int{-671, -858, 530}}, Beacon{Pos: [3]int{-667, 343, 800}}, Beacon{Pos: [3]int{571, -461, -707}}, Beacon{Pos: [3]int{-138, -166, 112}}, Beacon{Pos: [3]int{-889, 563, -600}}, Beacon{Pos: [3]int{646, -828, 498}}, Beacon{Pos: [3]int{640, 759, 510}}, Beacon{Pos: [3]int{-630, 509, 768}}, Beacon{Pos: [3]int{-681, -892, -333}}, Beacon{Pos: [3]int{673, -379, -804}}, Beacon{Pos: [3]int{-742, -814, -386}}, Beacon{Pos: [3]int{577, -820, 562}}}},
				Scanner{ID: 3, Beacons: []Beacon{Beacon{Pos: [3]int{-589, 542, 597}}, Beacon{Pos: [3]int{605, -692, 669}}, Beacon{Pos: [3]int{-500, 565, -823}}, Beacon{Pos: [3]int{-660, 373, 557}}, Beacon{Pos: [3]int{-458, -679, -417}}, Beacon{Pos: [3]int{-488, 449, 543}}, Beacon{Pos: [3]int{-626, 468, -788}}, Beacon{Pos: [3]int{338, -750, -386}}, Beacon{Pos: [3]int{528, -832, -391}}, Beacon{Pos: [3]int{562, -778, 733}}, Beacon{Pos: [3]int{-938, -730, 414}}, Beacon{Pos: [3]int{543, 643, -506}}, Beacon{Pos: [3]int{-524, 371, -870}}, Beacon{Pos: [3]int{407, 773, 750}}, Beacon{Pos: [3]int{-104, 29, 83}}, Beacon{Pos: [3]int{378, -903, -323}}, Beacon{Pos: [3]int{-778, -728, 485}}, Beacon{Pos: [3]int{426, 699, 580}}, Beacon{Pos: [3]int{-438, -605, -362}}, Beacon{Pos: [3]int{-469, -447, -387}}, Beacon{Pos: [3]int{509, 732, 623}}, Beacon{Pos: [3]int{647, 635, -688}}, Beacon{Pos: [3]int{-868, -804, 481}}, Beacon{Pos: [3]int{614, -800, 639}}, Beacon{Pos: [3]int{595, 780, -596}}}},
				Scanner{ID: 4, Beacons: []Beacon{Beacon{Pos: [3]int{727, 592, 562}}, Beacon{Pos: [3]int{-293, -554, 779}}, Beacon{Pos: [3]int{441, 611, -461}}, Beacon{Pos: [3]int{-714, 465, -776}}, Beacon{Pos: [3]int{-743, 427, -804}}, Beacon{Pos: [3]int{-660, -479, -426}}, Beacon{Pos: [3]int{832, -632, 460}}, Beacon{Pos: [3]int{927, -485, -438}}, Beacon{Pos: [3]int{408, 393, -506}}, Beacon{Pos: [3]int{466, 436, -512}}, Beacon{Pos: [3]int{110, 16, 151}}, Beacon{Pos: [3]int{-258, -428, 682}}, Beacon{Pos: [3]int{-393, 719, 612}}, Beacon{Pos: [3]int{-211, -452, 876}}, Beacon{Pos: [3]int{808, -476, -593}}, Beacon{Pos: [3]int{-575, 615, 604}}, Beacon{Pos: [3]int{-485, 667, 467}}, Beacon{Pos: [3]int{-680, 325, -822}}, Beacon{Pos: [3]int{-627, -443, -432}}, Beacon{Pos: [3]int{872, -547, -609}}, Beacon{Pos: [3]int{833, 512, 582}}, Beacon{Pos: [3]int{807, 604, 487}}, Beacon{Pos: [3]int{839, -516, 451}}, Beacon{Pos: [3]int{891, -625, 532}}, Beacon{Pos: [3]int{-652, -548, -490}}, Beacon{Pos: [3]int{30, -46, -14}}}},
			},
			expected: 79,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge1(tc.data)
			if err != nil && tc.error == false {
				t.Errorf("Challenge1(%v) threw '%v', want %d", tc.data, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge1(%v) = %d, want to throw", tc.data, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge1(%v) = %d, want %d", tc.data, actual, tc.expected)
			}
		})
	}
}

func TestChallenge2(t *testing.T) {
	var testCases = []struct {
		name     string
		data     []Scanner
		expected int
		error    bool
	}{
		{
			name: "Example 1",
			data: []Scanner{
				Scanner{ID: 0, Beacons: []Beacon{Beacon{Pos: [3]int{404, -588, -901}}, Beacon{Pos: [3]int{528, -643, 409}}, Beacon{Pos: [3]int{-838, 591, 734}}, Beacon{Pos: [3]int{390, -675, -793}}, Beacon{Pos: [3]int{-537, -823, -458}}, Beacon{Pos: [3]int{-485, -357, 347}}, Beacon{Pos: [3]int{-345, -311, 381}}, Beacon{Pos: [3]int{-661, -816, -575}}, Beacon{Pos: [3]int{-876, 649, 763}}, Beacon{Pos: [3]int{-618, -824, -621}}, Beacon{Pos: [3]int{553, 345, -567}}, Beacon{Pos: [3]int{474, 580, 667}}, Beacon{Pos: [3]int{-447, -329, 318}}, Beacon{Pos: [3]int{-584, 868, -557}}, Beacon{Pos: [3]int{544, -627, -890}}, Beacon{Pos: [3]int{564, 392, -477}}, Beacon{Pos: [3]int{455, 729, 728}}, Beacon{Pos: [3]int{-892, 524, 684}}, Beacon{Pos: [3]int{-689, 845, -530}}, Beacon{Pos: [3]int{423, -701, 434}}, Beacon{Pos: [3]int{7, -33, -71}}, Beacon{Pos: [3]int{630, 319, -379}}, Beacon{Pos: [3]int{443, 580, 662}}, Beacon{Pos: [3]int{-789, 900, -551}}, Beacon{Pos: [3]int{459, -707, 401}}}},
				Scanner{ID: 1, Beacons: []Beacon{Beacon{Pos: [3]int{686, 422, 578}}, Beacon{Pos: [3]int{605, 423, 415}}, Beacon{Pos: [3]int{515, 917, -361}}, Beacon{Pos: [3]int{-336, 658, 858}}, Beacon{Pos: [3]int{95, 138, 22}}, Beacon{Pos: [3]int{-476, 619, 847}}, Beacon{Pos: [3]int{-340, -569, -846}}, Beacon{Pos: [3]int{567, -361, 727}}, Beacon{Pos: [3]int{-460, 603, -452}}, Beacon{Pos: [3]int{669, -402, 600}}, Beacon{Pos: [3]int{729, 430, 532}}, Beacon{Pos: [3]int{-500, -761, 534}}, Beacon{Pos: [3]int{-322, 571, 750}}, Beacon{Pos: [3]int{-466, -666, -811}}, Beacon{Pos: [3]int{-429, -592, 574}}, Beacon{Pos: [3]int{-355, 545, -477}}, Beacon{Pos: [3]int{703, -491, -529}}, Beacon{Pos: [3]int{-328, -685, 520}}, Beacon{Pos: [3]int{413, 935, -424}}, Beacon{Pos: [3]int{-391, 539, -444}}, Beacon{Pos: [3]int{586, -435, 557}}, Beacon{Pos: [3]int{-364, -763, -893}}, Beacon{Pos: [3]int{807, -499, -711}}, Beacon{Pos: [3]int{755, -354, -619}}, Beacon{Pos: [3]int{553, 889, -390}}}},
				Scanner{ID: 2, Beacons: []Beacon{Beacon{Pos: [3]int{649, 640, 665}}, Beacon{Pos: [3]int{682, -795, 504}}, Beacon{Pos: [3]int{-784, 533, -524}}, Beacon{Pos: [3]int{-644, 584, -595}}, Beacon{Pos: [3]int{-588, -843, 648}}, Beacon{Pos: [3]int{-30, 6, 44}}, Beacon{Pos: [3]int{-674, 560, 763}}, Beacon{Pos: [3]int{500, 723, -460}}, Beacon{Pos: [3]int{609, 671, -379}}, Beacon{Pos: [3]int{-555, -800, 653}}, Beacon{Pos: [3]int{-675, -892, -343}}, Beacon{Pos: [3]int{697, -426, -610}}, Beacon{Pos: [3]int{578, 704, 681}}, Beacon{Pos: [3]int{493, 664, -388}}, Beacon{Pos: [3]int{-671, -858, 530}}, Beacon{Pos: [3]int{-667, 343, 800}}, Beacon{Pos: [3]int{571, -461, -707}}, Beacon{Pos: [3]int{-138, -166, 112}}, Beacon{Pos: [3]int{-889, 563, -600}}, Beacon{Pos: [3]int{646, -828, 498}}, Beacon{Pos: [3]int{640, 759, 510}}, Beacon{Pos: [3]int{-630, 509, 768}}, Beacon{Pos: [3]int{-681, -892, -333}}, Beacon{Pos: [3]int{673, -379, -804}}, Beacon{Pos: [3]int{-742, -814, -386}}, Beacon{Pos: [3]int{577, -820, 562}}}},
				Scanner{ID: 3, Beacons: []Beacon{Beacon{Pos: [3]int{-589, 542, 597}}, Beacon{Pos: [3]int{605, -692, 669}}, Beacon{Pos: [3]int{-500, 565, -823}}, Beacon{Pos: [3]int{-660, 373, 557}}, Beacon{Pos: [3]int{-458, -679, -417}}, Beacon{Pos: [3]int{-488, 449, 543}}, Beacon{Pos: [3]int{-626, 468, -788}}, Beacon{Pos: [3]int{338, -750, -386}}, Beacon{Pos: [3]int{528, -832, -391}}, Beacon{Pos: [3]int{562, -778, 733}}, Beacon{Pos: [3]int{-938, -730, 414}}, Beacon{Pos: [3]int{543, 643, -506}}, Beacon{Pos: [3]int{-524, 371, -870}}, Beacon{Pos: [3]int{407, 773, 750}}, Beacon{Pos: [3]int{-104, 29, 83}}, Beacon{Pos: [3]int{378, -903, -323}}, Beacon{Pos: [3]int{-778, -728, 485}}, Beacon{Pos: [3]int{426, 699, 580}}, Beacon{Pos: [3]int{-438, -605, -362}}, Beacon{Pos: [3]int{-469, -447, -387}}, Beacon{Pos: [3]int{509, 732, 623}}, Beacon{Pos: [3]int{647, 635, -688}}, Beacon{Pos: [3]int{-868, -804, 481}}, Beacon{Pos: [3]int{614, -800, 639}}, Beacon{Pos: [3]int{595, 780, -596}}}},
				Scanner{ID: 4, Beacons: []Beacon{Beacon{Pos: [3]int{727, 592, 562}}, Beacon{Pos: [3]int{-293, -554, 779}}, Beacon{Pos: [3]int{441, 611, -461}}, Beacon{Pos: [3]int{-714, 465, -776}}, Beacon{Pos: [3]int{-743, 427, -804}}, Beacon{Pos: [3]int{-660, -479, -426}}, Beacon{Pos: [3]int{832, -632, 460}}, Beacon{Pos: [3]int{927, -485, -438}}, Beacon{Pos: [3]int{408, 393, -506}}, Beacon{Pos: [3]int{466, 436, -512}}, Beacon{Pos: [3]int{110, 16, 151}}, Beacon{Pos: [3]int{-258, -428, 682}}, Beacon{Pos: [3]int{-393, 719, 612}}, Beacon{Pos: [3]int{-211, -452, 876}}, Beacon{Pos: [3]int{808, -476, -593}}, Beacon{Pos: [3]int{-575, 615, 604}}, Beacon{Pos: [3]int{-485, 667, 467}}, Beacon{Pos: [3]int{-680, 325, -822}}, Beacon{Pos: [3]int{-627, -443, -432}}, Beacon{Pos: [3]int{872, -547, -609}}, Beacon{Pos: [3]int{833, 512, 582}}, Beacon{Pos: [3]int{807, 604, 487}}, Beacon{Pos: [3]int{839, -516, 451}}, Beacon{Pos: [3]int{891, -625, 532}}, Beacon{Pos: [3]int{-652, -548, -490}}, Beacon{Pos: [3]int{30, -46, -14}}}},
			},
			expected: 3621,
			error:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Challenge2(tc.data)
			if err != nil && tc.error == false {
				t.Errorf("Challenge2(%v) threw '%v', want %d", tc.data, err, tc.expected)
			} else if err == nil && tc.error == true {
				t.Errorf("Challenge2(%v) = %d, want to throw", tc.data, actual)
			} else if actual != tc.expected {
				t.Errorf("Challenge2(%v) = %d, want %d", tc.data, actual, tc.expected)
			}
		})
	}
}
