package amphipod

import (
	"fmt"
	"math"
)

type Column struct {
	Data []byte
	Size int
	Type byte
}

type State struct {
	Cols []Column
}

func isAmphipod(val byte) bool {
	return val == 'A' || val == 'B' || val == 'C' || val == 'D'
}

func cost(val byte) int {
	switch val {
	case 'A':
		return 1
	case 'B':
		return 10
	case 'C':
		return 100
	case 'D':
		return 1000
	}
	return 10000
}

func (s State) solved() bool {
	return s.Cols[2].solved() && s.Cols[4].solved() && s.Cols[6].solved() && s.Cols[8].solved()
}

func (c Column) solved() bool {
	if c.Type == '.' || len(c.Data) != c.Size {
		return false
	}
	for _, cur := range c.Data {
		if cur != c.Type {
			return false
		}
	}
	return true
}

func (c Column) dirty() bool {
	if c.Type == '.' {
		return true
	}
	for _, cur := range c.Data {
		if cur != c.Type {
			return true
		}
	}
	return false
}

func (c Column) full() bool {
	return len(c.Data) == c.Size
}

func (s State) hash() string {
	var result []byte = make([]byte, 0)
	for _, cur := range s.Cols {
		result = append(result, cur.Data...)
		for i := 0; i < cur.Size-len(cur.Data); i++ {
			result = append(result, '.')
		}
		result = append(result, ' ')
	}
	return string(result)
}

func (s State) copy() State {
	var result State
	result.Cols = make([]Column, 11)
	for i := 0; i < 11; i++ {
		result.Cols[i].Data = make([]byte, 0)
		for j := 0; j < len(s.Cols[i].Data); j++ {
			result.Cols[i].Data = append(result.Cols[i].Data, s.Cols[i].Data[j])
		}
		result.Cols[i].Size = s.Cols[i].Size
		result.Cols[i].Type = s.Cols[i].Type
	}
	return result
}

func (s State) move(toMove byte, col1 int, col2 int) State {
	var result State = s.copy()
	result.Cols[col1].Data = result.Cols[col1].Data[0 : len(result.Cols[col1].Data)-1]
	result.Cols[col2].Data = append(result.Cols[col2].Data, toMove)
	return result
}

func minSlice(data []int) int {
	result := math.MaxInt32
	for _, cur := range data {
		if cur < result {
			result = cur
		}
	}
	return result
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func solve(state State, path map[string]struct{}, cache map[string]int) int {
	if val, ok := cache[state.hash()]; ok {
		return val
	}
	if state.solved() {
		return 0
	}
	var counts []int = make([]int, 0)
	for i1 := 0; i1 < len(state.Cols); i1++ {
		col1 := state.Cols[i1]
		if len(col1.Data) == 0 || col1.solved() || !col1.dirty() {
			continue
		}
		toMove := col1.Data[len(col1.Data)-1]
		for i2 := 0; i2 < len(state.Cols); i2++ {
			if i2 == i1 {
				continue
			}
			blocked := false
			for j := min(i1, i2) + 1; j < max(i1, i2); j++ {
				if state.Cols[j].Type == '.' && state.Cols[j].full() {
					blocked = true
				}
			}
			if blocked {
				continue
			}
			col2 := state.Cols[i2]
			if col2.full() || (col1.Type == '.' && col2.Type == '.') || (col1.Type != '.' && col2.Type != '.') || (col2.Type != '.' && (col2.Type != toMove || col2.dirty())) {
				continue
			}
			new_state := state.move(toMove, i1, i2)
			if _, ok := path[new_state.hash()]; ok {
				continue
			}
			distance := abs(i2 - i1)
			if col1.Type != '.' {
				distance = distance + 1 + col1.Size - len(col1.Data)
			}
			if col2.Type != '.' {
				distance = distance + col2.Size - len(col2.Data)
			}
			path[new_state.hash()] = struct{}{}
			subCost := solve(new_state, path, cache)
			if subCost >= 0 {
				counts = append(counts, (distance*cost(toMove))+subCost)
			}
			delete(path, new_state.hash())
		}
	}
	result := -1
	if len(counts) > 0 {
		result = minSlice(counts)
	}
	cache[state.hash()] = result
	return result
}

func Challenge1(state State) (int, error) {
	var path map[string]struct{} = make(map[string]struct{})
	path[state.hash()] = struct{}{}
	return solve(state, path, make(map[string]int)), nil
}

func Challenge2(state State) (int, error) {
	var path map[string]struct{} = make(map[string]struct{})
	path[state.hash()] = struct{}{}
	return solve(state, path, make(map[string]int)), nil
}

func Run() {
	fmt.Println("Day 23 - Amphipod")

	var result int
	var err error

	result, err = Challenge1(
		State{
			Cols: []Column{
				Column{Data: []byte{}, Size: 1, Type: '.'},
				Column{Data: []byte{}, Size: 1, Type: '.'},
				Column{Data: []byte{'D', 'B'}, Size: 2, Type: 'A'},
				Column{Data: []byte{}, Size: 1, Type: '.'},
				Column{Data: []byte{'D', 'C'}, Size: 2, Type: 'B'},
				Column{Data: []byte{}, Size: 1, Type: '.'},
				Column{Data: []byte{'A', 'C'}, Size: 2, Type: 'C'},
				Column{Data: []byte{}, Size: 1, Type: '.'},
				Column{Data: []byte{'A', 'B'}, Size: 2, Type: 'D'},
				Column{Data: []byte{}, Size: 1, Type: '.'},
				Column{Data: []byte{}, Size: 1, Type: '.'},
			},
		},
	)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result, err = Challenge2(
		State{
			Cols: []Column{
				Column{Data: []byte{}, Size: 1, Type: '.'},
				Column{Data: []byte{}, Size: 1, Type: '.'},
				Column{Data: []byte{'D', 'D', 'D', 'B'}, Size: 4, Type: 'A'},
				Column{Data: []byte{}, Size: 1, Type: '.'},
				Column{Data: []byte{'D', 'B', 'C', 'C'}, Size: 4, Type: 'B'},
				Column{Data: []byte{}, Size: 1, Type: '.'},
				Column{Data: []byte{'A', 'A', 'B', 'C'}, Size: 4, Type: 'C'},
				Column{Data: []byte{}, Size: 1, Type: '.'},
				Column{Data: []byte{'A', 'C', 'A', 'B'}, Size: 4, Type: 'D'},
				Column{Data: []byte{}, Size: 1, Type: '.'},
				Column{Data: []byte{}, Size: 1, Type: '.'},
			},
		},
	)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
