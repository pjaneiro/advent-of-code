package spacestoichiometry_test

import (
	. "github.com/pjaneiro/advent-of-code/2019/spacestoichiometry"
	"testing"
)

func TestChallenge1(t *testing.T) {
	var testCases = []struct {
		name     string
		input    ReactionList
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    ReactionList{Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 10}}, Output: Chemical{Name: "A", Quantity: 10}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 1}}, Output: Chemical{Name: "B", Quantity: 1}}, Reaction{Input: []Chemical{Chemical{Name: "A", Quantity: 7}, Chemical{Name: "B", Quantity: 1}}, Output: Chemical{Name: "C", Quantity: 1}}, Reaction{Input: []Chemical{Chemical{Name: "A", Quantity: 7}, Chemical{Name: "C", Quantity: 1}}, Output: Chemical{Name: "D", Quantity: 1}}, Reaction{Input: []Chemical{Chemical{Name: "A", Quantity: 7}, Chemical{Name: "D", Quantity: 1}}, Output: Chemical{Name: "E", Quantity: 1}}, Reaction{Input: []Chemical{Chemical{Name: "A", Quantity: 7}, Chemical{Name: "E", Quantity: 1}}, Output: Chemical{Name: "FUEL", Quantity: 1}}},
			expected: 31,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    ReactionList{Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 9}}, Output: Chemical{Name: "A", Quantity: 2}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 8}}, Output: Chemical{Name: "B", Quantity: 3}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 7}}, Output: Chemical{Name: "C", Quantity: 5}}, Reaction{Input: []Chemical{Chemical{Name: "A", Quantity: 3}, Chemical{Name: "B", Quantity: 4}}, Output: Chemical{Name: "AB", Quantity: 1}}, Reaction{Input: []Chemical{Chemical{Name: "B", Quantity: 5}, Chemical{Name: "C", Quantity: 7}}, Output: Chemical{Name: "BC", Quantity: 1}}, Reaction{Input: []Chemical{Chemical{Name: "C", Quantity: 4}, Chemical{Name: "A", Quantity: 1}}, Output: Chemical{Name: "CA", Quantity: 1}}, Reaction{Input: []Chemical{Chemical{Name: "AB", Quantity: 2}, Chemical{Name: "BC", Quantity: 3}, Chemical{Name: "CA", Quantity: 4}}, Output: Chemical{Name: "FUEL", Quantity: 1}}},
			expected: 165,
			error:    false,
		},
		{
			name:     "Example 3",
			input:    ReactionList{Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 157}}, Output: Chemical{Name: "NZVS", Quantity: 5}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 165}}, Output: Chemical{Name: "DCFZ", Quantity: 6}}, Reaction{Input: []Chemical{Chemical{Name: "XJWVT", Quantity: 44}, Chemical{Name: "KHKGT", Quantity: 5}, Chemical{Name: "QDVJ", Quantity: 1}, Chemical{Name: "NZVS", Quantity: 29}, Chemical{Name: "GPVTF", Quantity: 9}, Chemical{Name: "HKGWZ", Quantity: 48}}, Output: Chemical{Name: "FUEL", Quantity: 1}}, Reaction{Input: []Chemical{Chemical{Name: "HKGWZ", Quantity: 12}, Chemical{Name: "GPVTF", Quantity: 1}, Chemical{Name: "PSHF", Quantity: 8}}, Output: Chemical{Name: "QDVJ", Quantity: 9}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 179}}, Output: Chemical{Name: "PSHF", Quantity: 7}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 177}}, Output: Chemical{Name: "HKGWZ", Quantity: 5}}, Reaction{Input: []Chemical{Chemical{Name: "DCFZ", Quantity: 7}, Chemical{Name: "PSHF", Quantity: 7}}, Output: Chemical{Name: "XJWVT", Quantity: 2}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 165}}, Output: Chemical{Name: "GPVTF", Quantity: 2}}, Reaction{Input: []Chemical{Chemical{Name: "DCFZ", Quantity: 3}, Chemical{Name: "NZVS", Quantity: 7}, Chemical{Name: "HKGWZ", Quantity: 5}, Chemical{Name: "PSHF", Quantity: 10}}, Output: Chemical{Name: "KHKGT", Quantity: 8}}},
			expected: 13312,
			error:    false,
		},
		{
			name:     "Example 4",
			input:    ReactionList{Reaction{Input: []Chemical{Chemical{Name: "VPVL", Quantity: 2}, Chemical{Name: "FWMGM", Quantity: 7}, Chemical{Name: "CXFTF", Quantity: 2}, Chemical{Name: "MNCFX", Quantity: 11}}, Output: Chemical{Name: "STKFG", Quantity: 1}}, Reaction{Input: []Chemical{Chemical{Name: "NVRVD", Quantity: 17}, Chemical{Name: "JNWZP", Quantity: 3}}, Output: Chemical{Name: "VPVL", Quantity: 8}}, Reaction{Input: []Chemical{Chemical{Name: "STKFG", Quantity: 53}, Chemical{Name: "MNCFX", Quantity: 6}, Chemical{Name: "VJHF", Quantity: 46}, Chemical{Name: "HVMC", Quantity: 81}, Chemical{Name: "CXFTF", Quantity: 68}, Chemical{Name: "GNMV", Quantity: 25}}, Output: Chemical{Name: "FUEL", Quantity: 1}}, Reaction{Input: []Chemical{Chemical{Name: "VJHF", Quantity: 22}, Chemical{Name: "MNCFX", Quantity: 37}}, Output: Chemical{Name: "FWMGM", Quantity: 5}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 139}}, Output: Chemical{Name: "NVRVD", Quantity: 4}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 144}}, Output: Chemical{Name: "JNWZP", Quantity: 7}}, Reaction{Input: []Chemical{Chemical{Name: "MNCFX", Quantity: 5}, Chemical{Name: "RFSQX", Quantity: 7}, Chemical{Name: "FWMGM", Quantity: 2}, Chemical{Name: "VPVL", Quantity: 2}, Chemical{Name: "CXFTF", Quantity: 19}}, Output: Chemical{Name: "HVMC", Quantity: 3}}, Reaction{Input: []Chemical{Chemical{Name: "VJHF", Quantity: 5}, Chemical{Name: "MNCFX", Quantity: 7}, Chemical{Name: "VPVL", Quantity: 9}, Chemical{Name: "CXFTF", Quantity: 37}}, Output: Chemical{Name: "GNMV", Quantity: 6}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 145}}, Output: Chemical{Name: "MNCFX", Quantity: 6}}, Reaction{Input: []Chemical{Chemical{Name: "NVRVD", Quantity: 1}}, Output: Chemical{Name: "CXFTF", Quantity: 8}}, Reaction{Input: []Chemical{Chemical{Name: "VJHF", Quantity: 1}, Chemical{Name: "MNCFX", Quantity: 6}}, Output: Chemical{Name: "RFSQX", Quantity: 4}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 176}}, Output: Chemical{Name: "VJHF", Quantity: 6}}},
			expected: 180697,
			error:    false,
		},
		{
			name:     "Example 5",
			input:    ReactionList{Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 171}}, Output: Chemical{Name: "CNZTR", Quantity: 8}}, Reaction{Input: []Chemical{Chemical{Name: "ZLQW", Quantity: 7}, Chemical{Name: "BMBT", Quantity: 3}, Chemical{Name: "XCVML", Quantity: 9}, Chemical{Name: "XMNCP", Quantity: 26}, Chemical{Name: "WPTQ", Quantity: 1}, Chemical{Name: "MZWV", Quantity: 2}, Chemical{Name: "RJRHP", Quantity: 1}}, Output: Chemical{Name: "PLWSL", Quantity: 4}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 114}}, Output: Chemical{Name: "BHXH", Quantity: 4}}, Reaction{Input: []Chemical{Chemical{Name: "VRPVC", Quantity: 14}}, Output: Chemical{Name: "BMBT", Quantity: 6}}, Reaction{Input: []Chemical{Chemical{Name: "BHXH", Quantity: 6}, Chemical{Name: "KTJDG", Quantity: 18}, Chemical{Name: "WPTQ", Quantity: 12}, Chemical{Name: "PLWSL", Quantity: 7}, Chemical{Name: "FHTLT", Quantity: 31}, Chemical{Name: "ZDVW", Quantity: 37}}, Output: Chemical{Name: "FUEL", Quantity: 1}}, Reaction{Input: []Chemical{Chemical{Name: "WPTQ", Quantity: 6}, Chemical{Name: "BMBT", Quantity: 2}, Chemical{Name: "ZLQW", Quantity: 8}, Chemical{Name: "KTJDG", Quantity: 18}, Chemical{Name: "XMNCP", Quantity: 1}, Chemical{Name: "MZWV", Quantity: 6}, Chemical{Name: "RJRHP", Quantity: 1}}, Output: Chemical{Name: "FHTLT", Quantity: 6}}, Reaction{Input: []Chemical{Chemical{Name: "XDBXC", Quantity: 15}, Chemical{Name: "LTCX", Quantity: 2}, Chemical{Name: "VRPVC", Quantity: 1}}, Output: Chemical{Name: "ZLQW", Quantity: 6}}, Reaction{Input: []Chemical{Chemical{Name: "WPTQ", Quantity: 13}, Chemical{Name: "LTCX", Quantity: 10}, Chemical{Name: "RJRHP", Quantity: 3}, Chemical{Name: "XMNCP", Quantity: 14}, Chemical{Name: "MZWV", Quantity: 2}, Chemical{Name: "ZLQW", Quantity: 1}}, Output: Chemical{Name: "ZDVW", Quantity: 1}}, Reaction{Input: []Chemical{Chemical{Name: "BMBT", Quantity: 5}}, Output: Chemical{Name: "WPTQ", Quantity: 4}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 189}}, Output: Chemical{Name: "KTJDG", Quantity: 9}}, Reaction{Input: []Chemical{Chemical{Name: "MZWV", Quantity: 1}, Chemical{Name: "XDBXC", Quantity: 17}, Chemical{Name: "XCVML", Quantity: 3}}, Output: Chemical{Name: "XMNCP", Quantity: 2}}, Reaction{Input: []Chemical{Chemical{Name: "VRPVC", Quantity: 12}, Chemical{Name: "CNZTR", Quantity: 27}}, Output: Chemical{Name: "XDBXC", Quantity: 2}}, Reaction{Input: []Chemical{Chemical{Name: "KTJDG", Quantity: 15}, Chemical{Name: "BHXH", Quantity: 12}}, Output: Chemical{Name: "XCVML", Quantity: 5}}, Reaction{Input: []Chemical{Chemical{Name: "BHXH", Quantity: 3}, Chemical{Name: "VRPVC", Quantity: 2}}, Output: Chemical{Name: "MZWV", Quantity: 7}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 121}}, Output: Chemical{Name: "VRPVC", Quantity: 7}}, Reaction{Input: []Chemical{Chemical{Name: "XCVML", Quantity: 7}}, Output: Chemical{Name: "RJRHP", Quantity: 6}}, Reaction{Input: []Chemical{Chemical{Name: "BHXH", Quantity: 5}, Chemical{Name: "VRPVC", Quantity: 4}}, Output: Chemical{Name: "LTCX", Quantity: 5}}},
			expected: 2210736,
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
		input    ReactionList
		expected int
		error    bool
	}{
		{
			name:     "Example 1",
			input:    ReactionList{Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 157}}, Output: Chemical{Name: "NZVS", Quantity: 5}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 165}}, Output: Chemical{Name: "DCFZ", Quantity: 6}}, Reaction{Input: []Chemical{Chemical{Name: "XJWVT", Quantity: 44}, Chemical{Name: "KHKGT", Quantity: 5}, Chemical{Name: "QDVJ", Quantity: 1}, Chemical{Name: "NZVS", Quantity: 29}, Chemical{Name: "GPVTF", Quantity: 9}, Chemical{Name: "HKGWZ", Quantity: 48}}, Output: Chemical{Name: "FUEL", Quantity: 1}}, Reaction{Input: []Chemical{Chemical{Name: "HKGWZ", Quantity: 12}, Chemical{Name: "GPVTF", Quantity: 1}, Chemical{Name: "PSHF", Quantity: 8}}, Output: Chemical{Name: "QDVJ", Quantity: 9}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 179}}, Output: Chemical{Name: "PSHF", Quantity: 7}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 177}}, Output: Chemical{Name: "HKGWZ", Quantity: 5}}, Reaction{Input: []Chemical{Chemical{Name: "DCFZ", Quantity: 7}, Chemical{Name: "PSHF", Quantity: 7}}, Output: Chemical{Name: "XJWVT", Quantity: 2}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 165}}, Output: Chemical{Name: "GPVTF", Quantity: 2}}, Reaction{Input: []Chemical{Chemical{Name: "DCFZ", Quantity: 3}, Chemical{Name: "NZVS", Quantity: 7}, Chemical{Name: "HKGWZ", Quantity: 5}, Chemical{Name: "PSHF", Quantity: 10}}, Output: Chemical{Name: "KHKGT", Quantity: 8}}},
			expected: 82892753,
			error:    false,
		},
		{
			name:     "Example 2",
			input:    ReactionList{Reaction{Input: []Chemical{Chemical{Name: "VPVL", Quantity: 2}, Chemical{Name: "FWMGM", Quantity: 7}, Chemical{Name: "CXFTF", Quantity: 2}, Chemical{Name: "MNCFX", Quantity: 11}}, Output: Chemical{Name: "STKFG", Quantity: 1}}, Reaction{Input: []Chemical{Chemical{Name: "NVRVD", Quantity: 17}, Chemical{Name: "JNWZP", Quantity: 3}}, Output: Chemical{Name: "VPVL", Quantity: 8}}, Reaction{Input: []Chemical{Chemical{Name: "STKFG", Quantity: 53}, Chemical{Name: "MNCFX", Quantity: 6}, Chemical{Name: "VJHF", Quantity: 46}, Chemical{Name: "HVMC", Quantity: 81}, Chemical{Name: "CXFTF", Quantity: 68}, Chemical{Name: "GNMV", Quantity: 25}}, Output: Chemical{Name: "FUEL", Quantity: 1}}, Reaction{Input: []Chemical{Chemical{Name: "VJHF", Quantity: 22}, Chemical{Name: "MNCFX", Quantity: 37}}, Output: Chemical{Name: "FWMGM", Quantity: 5}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 139}}, Output: Chemical{Name: "NVRVD", Quantity: 4}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 144}}, Output: Chemical{Name: "JNWZP", Quantity: 7}}, Reaction{Input: []Chemical{Chemical{Name: "MNCFX", Quantity: 5}, Chemical{Name: "RFSQX", Quantity: 7}, Chemical{Name: "FWMGM", Quantity: 2}, Chemical{Name: "VPVL", Quantity: 2}, Chemical{Name: "CXFTF", Quantity: 19}}, Output: Chemical{Name: "HVMC", Quantity: 3}}, Reaction{Input: []Chemical{Chemical{Name: "VJHF", Quantity: 5}, Chemical{Name: "MNCFX", Quantity: 7}, Chemical{Name: "VPVL", Quantity: 9}, Chemical{Name: "CXFTF", Quantity: 37}}, Output: Chemical{Name: "GNMV", Quantity: 6}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 145}}, Output: Chemical{Name: "MNCFX", Quantity: 6}}, Reaction{Input: []Chemical{Chemical{Name: "NVRVD", Quantity: 1}}, Output: Chemical{Name: "CXFTF", Quantity: 8}}, Reaction{Input: []Chemical{Chemical{Name: "VJHF", Quantity: 1}, Chemical{Name: "MNCFX", Quantity: 6}}, Output: Chemical{Name: "RFSQX", Quantity: 4}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 176}}, Output: Chemical{Name: "VJHF", Quantity: 6}}},
			expected: 5586022,
			error:    false,
		},
		{
			name:     "Example 3",
			input:    ReactionList{Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 171}}, Output: Chemical{Name: "CNZTR", Quantity: 8}}, Reaction{Input: []Chemical{Chemical{Name: "ZLQW", Quantity: 7}, Chemical{Name: "BMBT", Quantity: 3}, Chemical{Name: "XCVML", Quantity: 9}, Chemical{Name: "XMNCP", Quantity: 26}, Chemical{Name: "WPTQ", Quantity: 1}, Chemical{Name: "MZWV", Quantity: 2}, Chemical{Name: "RJRHP", Quantity: 1}}, Output: Chemical{Name: "PLWSL", Quantity: 4}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 114}}, Output: Chemical{Name: "BHXH", Quantity: 4}}, Reaction{Input: []Chemical{Chemical{Name: "VRPVC", Quantity: 14}}, Output: Chemical{Name: "BMBT", Quantity: 6}}, Reaction{Input: []Chemical{Chemical{Name: "BHXH", Quantity: 6}, Chemical{Name: "KTJDG", Quantity: 18}, Chemical{Name: "WPTQ", Quantity: 12}, Chemical{Name: "PLWSL", Quantity: 7}, Chemical{Name: "FHTLT", Quantity: 31}, Chemical{Name: "ZDVW", Quantity: 37}}, Output: Chemical{Name: "FUEL", Quantity: 1}}, Reaction{Input: []Chemical{Chemical{Name: "WPTQ", Quantity: 6}, Chemical{Name: "BMBT", Quantity: 2}, Chemical{Name: "ZLQW", Quantity: 8}, Chemical{Name: "KTJDG", Quantity: 18}, Chemical{Name: "XMNCP", Quantity: 1}, Chemical{Name: "MZWV", Quantity: 6}, Chemical{Name: "RJRHP", Quantity: 1}}, Output: Chemical{Name: "FHTLT", Quantity: 6}}, Reaction{Input: []Chemical{Chemical{Name: "XDBXC", Quantity: 15}, Chemical{Name: "LTCX", Quantity: 2}, Chemical{Name: "VRPVC", Quantity: 1}}, Output: Chemical{Name: "ZLQW", Quantity: 6}}, Reaction{Input: []Chemical{Chemical{Name: "WPTQ", Quantity: 13}, Chemical{Name: "LTCX", Quantity: 10}, Chemical{Name: "RJRHP", Quantity: 3}, Chemical{Name: "XMNCP", Quantity: 14}, Chemical{Name: "MZWV", Quantity: 2}, Chemical{Name: "ZLQW", Quantity: 1}}, Output: Chemical{Name: "ZDVW", Quantity: 1}}, Reaction{Input: []Chemical{Chemical{Name: "BMBT", Quantity: 5}}, Output: Chemical{Name: "WPTQ", Quantity: 4}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 189}}, Output: Chemical{Name: "KTJDG", Quantity: 9}}, Reaction{Input: []Chemical{Chemical{Name: "MZWV", Quantity: 1}, Chemical{Name: "XDBXC", Quantity: 17}, Chemical{Name: "XCVML", Quantity: 3}}, Output: Chemical{Name: "XMNCP", Quantity: 2}}, Reaction{Input: []Chemical{Chemical{Name: "VRPVC", Quantity: 12}, Chemical{Name: "CNZTR", Quantity: 27}}, Output: Chemical{Name: "XDBXC", Quantity: 2}}, Reaction{Input: []Chemical{Chemical{Name: "KTJDG", Quantity: 15}, Chemical{Name: "BHXH", Quantity: 12}}, Output: Chemical{Name: "XCVML", Quantity: 5}}, Reaction{Input: []Chemical{Chemical{Name: "BHXH", Quantity: 3}, Chemical{Name: "VRPVC", Quantity: 2}}, Output: Chemical{Name: "MZWV", Quantity: 7}}, Reaction{Input: []Chemical{Chemical{Name: "ORE", Quantity: 121}}, Output: Chemical{Name: "VRPVC", Quantity: 7}}, Reaction{Input: []Chemical{Chemical{Name: "XCVML", Quantity: 7}}, Output: Chemical{Name: "RJRHP", Quantity: 6}}, Reaction{Input: []Chemical{Chemical{Name: "BHXH", Quantity: 5}, Chemical{Name: "VRPVC", Quantity: 4}}, Output: Chemical{Name: "LTCX", Quantity: 5}}},
			expected: 460664,
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
