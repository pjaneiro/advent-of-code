package conwaycubes

import (
	"fmt"
)

func Challenge1(data [][][]rune) (int, error) {
	result := 0
	var actualmap map[[3]int]rune = make(map[[3]int]rune, 0)
	for x := 0; x < len(data[0]); x++ {
		for y := 0; y < len(data[0][0]); y++ {
			for z := 0; z < len(data); z++ {
				var index [3]int = [3]int{x, y, z}
				actualmap[index] = data[z][y][x]
			}
		}
	}

	for round := 0; round < 6; round++ {
		var newmap map[[3]int]rune = make(map[[3]int]rune, 0)
		for x := -6; x < 16; x++ {
			for y := -6; y < 16; y++ {
				for z := -6; z < 7; z++ {
					var index [3]int = [3]int{x, y, z}
					activeneighbours := 0
					for dx := -1; dx < 2; dx++ {
						for dy := -1; dy < 2; dy++ {
							for dz := -1; dz < 2; dz++ {
								if dx == 0 && dy == 0 && dz == 0 {
									continue
								}
								var neighbourindex [3]int = [3]int{x + dx, y + dy, z + dz}
								if val, ok := actualmap[neighbourindex]; ok && val == '#' {
									activeneighbours++
								}
							}
						}
					}
					val, ok := actualmap[index]
					if ok && val == '#' {
						if activeneighbours == 2 || activeneighbours == 3 {
							newmap[index] = '#'
						} else {
							newmap[index] = '.'
						}
					} else if activeneighbours == 3 {
						newmap[index] = '#'
					} else {
						newmap[index] = '.'
					}
				}
			}
		}
		actualmap = make(map[[3]int]rune, 0)
		for k, v := range newmap {
			actualmap[k] = v
		}
	}

	for _, v := range actualmap {
		if v == '#' {
			result++
		}
	}

	return result, nil
}

func Challenge2(data [][][][]rune) (int, error) {
	result := 0
	var actualmap map[[4]int]rune = make(map[[4]int]rune, 0)
	for x := 0; x < len(data[0][0]); x++ {
		for y := 0; y < len(data[0][0][0]); y++ {
			for z := 0; z < len(data[0]); z++ {
				for w := 0; w < len(data); w++ {
					var index [4]int = [4]int{x, y, z, w}
					actualmap[index] = data[w][z][y][x]
				}
			}
		}
	}

	for round := 0; round < 6; round++ {
		var newmap map[[4]int]rune = make(map[[4]int]rune, 0)
		for x := -6; x < 16; x++ {
			for y := -6; y < 16; y++ {
				for z := -6; z < 7; z++ {
					for w := -6; w < 7; w++ {
						var index [4]int = [4]int{x, y, z, w}
						activeneighbours := 0
						for dx := -1; dx < 2; dx++ {
							for dy := -1; dy < 2; dy++ {
								for dz := -1; dz < 2; dz++ {
									for dw := -1; dw < 2; dw++ {
										if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
											continue
										}
										var neighbourindex [4]int = [4]int{x + dx, y + dy, z + dz, w + dw}
										if val, ok := actualmap[neighbourindex]; ok && val == '#' {
											activeneighbours++
										}
									}
								}
							}
						}
						val, ok := actualmap[index]
						if ok && val == '#' {
							if activeneighbours == 2 || activeneighbours == 3 {
								newmap[index] = '#'
							} else {
								newmap[index] = '.'
							}
						} else if activeneighbours == 3 {
							newmap[index] = '#'
						} else {
							newmap[index] = '.'
						}
					}
				}
			}
		}
		actualmap = make(map[[4]int]rune, 0)
		for k, v := range newmap {
			actualmap[k] = v
		}
	}

	for _, v := range actualmap {
		if v == '#' {
			result++
		}
	}

	return result, nil
}

func Run() {
	fmt.Println("Day 17 - Conway Cubes")

	data := [][][][]rune{
		{
			{
				{'#', '#', '#', '.', '.', '#', '.', '.'},
				{'.', '#', '#', '#', '#', '#', '#', '#'},
				{'#', '#', '#', '#', '#', '.', '.', '.'},
				{'#', '.', '.', '#', '#', '.', '#', '.'},
				{'#', '#', '#', '.', '.', '#', '#', '.'},
				{'#', '#', '.', '.', '.', '#', '.', '.'},
				{'.', '.', '#', '.', '.', '.', '#', '.'},
				{'.', '#', '.', '.', '.', '.', '#', '#'},
			},
		},
	}

	var result int
	result, err := Challenge1(data[0])
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result, err = Challenge2(data)
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result)
	}
}
