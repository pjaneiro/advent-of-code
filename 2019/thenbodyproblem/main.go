package thenbodyproblem

import (
	"errors"
	"fmt"
)

type Moon struct {
	X  int
	Y  int
	Z  int
	Vx int
	Vy int
	Vz int
}

type PosSet struct {
	m1 int
	m2 int
	m3 int
	m4 int
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func gcd(a int64, b int64) int64 {
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}
	return a
}

func lcm(a int64, b int64) int64 {
	return (a * b) / gcd(a, b)
}

func lcmslice(data []int64) int64 {
	var result int64 = 1
	for _, cur := range data {
		result = lcm(result, cur)
	}
	return result
}

func (m Moon) energy() int {
	return (abs(m.X) + abs(m.Y) + abs(m.Z)) * (abs(m.Vx) + abs(m.Vy) + abs(m.Vz))
}

func Challenge1(data []Moon, steps int) (int, error) {
	var state []Moon = make([]Moon, len(data))
	copy(state, data)
	for step := 0; step < steps; step++ {
		var newState []Moon = make([]Moon, len(data))

		for i, cur := range state {
			var newMoon Moon

			newMoon.Vx = cur.Vx
			newMoon.Vy = cur.Vy
			newMoon.Vz = cur.Vz

			for j, other := range state {
				if i == j {
					continue
				}

				if other.X < cur.X {
					newMoon.Vx = newMoon.Vx - 1
				} else if other.X > cur.X {
					newMoon.Vx = newMoon.Vx + 1
				}
				if other.Y < cur.Y {
					newMoon.Vy = newMoon.Vy - 1
				} else if other.Y > cur.Y {
					newMoon.Vy = newMoon.Vy + 1
				}
				if other.Z < cur.Z {
					newMoon.Vz = newMoon.Vz - 1
				} else if other.Z > cur.Z {
					newMoon.Vz = newMoon.Vz + 1
				}

				newMoon.X = cur.X + newMoon.Vx
				newMoon.Y = cur.Y + newMoon.Vy
				newMoon.Z = cur.Z + newMoon.Vz
			}

			newState[i] = newMoon
		}

		state = newState
	}
	var result int = 0
	for _, cur := range state {
		result = result + cur.energy()
	}
	return result, nil
}

func Challenge2(data []Moon) (int64, error) {
	var periods []int64 = make([]int64, 3)
	var state []Moon = make([]Moon, len(data))
	copy(state, data)
	var step int64
	for step = 0; true; step++ {
		var newState []Moon = make([]Moon, len(data))

		if periods[0] == 0 {
			var foundX bool = true
			for i := 0; i < len(data); i++ {
				if state[i].Vx != 0 || state[i].X != data[i].X {
					foundX = false
				}
			}
			if foundX {
				periods[0] = step
			}
		}
		if periods[1] == 0 {
			var foundY bool = true
			for i := 0; i < len(data); i++ {
				if state[i].Vy != 0 || state[i].Y != data[i].Y {
					foundY = false
				}
			}
			if foundY {
				periods[1] = step
			}
		}
		if periods[2] == 0 {
			var foundZ bool = true
			for i := 0; i < len(data); i++ {
				if state[i].Vz != 0 || state[i].Z != data[i].Z {
					foundZ = false
				}
			}
			if foundZ {
				periods[2] = step
			}
		}

		if periods[0] != 0 && periods[1] != 0 && periods[2] != 0 {
			return lcmslice(periods), nil
		}

		for i, cur := range state {
			var newMoon Moon

			newMoon.Vx = cur.Vx
			newMoon.Vy = cur.Vy
			newMoon.Vz = cur.Vz

			for j, other := range state {
				if i == j {
					continue
				}

				if other.X < cur.X {
					newMoon.Vx = newMoon.Vx - 1
				} else if other.X > cur.X {
					newMoon.Vx = newMoon.Vx + 1
				}
				if other.Y < cur.Y {
					newMoon.Vy = newMoon.Vy - 1
				} else if other.Y > cur.Y {
					newMoon.Vy = newMoon.Vy + 1
				}
				if other.Z < cur.Z {
					newMoon.Vz = newMoon.Vz - 1
				} else if other.Z > cur.Z {
					newMoon.Vz = newMoon.Vz + 1
				}

				newMoon.X = cur.X + newMoon.Vx
				newMoon.Y = cur.Y + newMoon.Vy
				newMoon.Z = cur.Z + newMoon.Vz
			}

			newState[i] = newMoon
		}

		state = newState
	}

	return 0, errors.New("something went wrong")
}

func Run() {
	fmt.Println("Day 12 - The N-Body Problem")

	result, err := Challenge1([]Moon{Moon{X: -6, Y: 2, Z: -9, Vx: 0, Vy: 0, Vz: 0}, Moon{X: 12, Y: -14, Z: -4, Vx: 0, Vy: 0, Vz: 0}, Moon{X: 9, Y: 5, Z: -6, Vx: 0, Vy: 0, Vz: 0}, Moon{X: -1, Y: -4, Z: 9, Vx: 0, Vy: 0, Vz: 0}}, 1000)
	if err != nil {
		fmt.Printf("Error running challenge 1: %v\n", err)
	} else {
		fmt.Printf("Challenge 1: %d\n", result)
	}

	result2, err := Challenge2([]Moon{Moon{X: -6, Y: 2, Z: -9, Vx: 0, Vy: 0, Vz: 0}, Moon{X: 12, Y: -14, Z: -4, Vx: 0, Vy: 0, Vz: 0}, Moon{X: 9, Y: 5, Z: -6, Vx: 0, Vy: 0, Vz: 0}, Moon{X: -1, Y: -4, Z: 9, Vx: 0, Vy: 0, Vz: 0}})
	if err != nil {
		fmt.Printf("Error running challenge 2: %v\n", err)
	} else {
		fmt.Printf("Challenge 2: %d\n", result2)
	}
}
