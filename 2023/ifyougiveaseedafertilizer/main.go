package ifyougiveaseedafertilizer

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Mapping struct {
	DestStart, SrcStart, Length int
}

const (
	SeedToSoil            = iota
	SoilToFertilizer      = iota
	FertilizerToWater     = iota
	WaterToLight          = iota
	LightToTemperature    = iota
	TemperatureToHumidity = iota
	HumidityToLocation    = iota
)

type Data struct {
	Seeds    []int
	Mappings [][]Mapping
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func parseData(data []string) (Data, error) {
	var result Data = Data{[]int{}, make([][]Mapping, 7)}
	seedsData := strings.Fields(strings.Split(data[0], ":")[1])
	for _, cur := range seedsData {
		tmp, err := strconv.Atoi(cur)
		if err != nil {
			return Data{}, err
		}
		result.Seeds = append(result.Seeds, tmp)
	}
	lineIndex, mappingIndex := 3, 0
	for ; lineIndex < len(data); lineIndex++ {
		if data[lineIndex] == "" {
			lineIndex += 2
			mappingIndex++
		}
		fields := strings.Fields(data[lineIndex])
		dst, err := strconv.Atoi(fields[0])
		if err != nil {
			return Data{}, err
		}
		src, err := strconv.Atoi(fields[1])
		if err != nil {
			return Data{}, err
		}
		rng, err := strconv.Atoi(fields[2])
		if err != nil {
			return Data{}, err
		}
		result.Mappings[mappingIndex] = append(result.Mappings[mappingIndex], Mapping{DestStart: dst, SrcStart: src, Length: rng})
	}
	return result, nil
}

func Challenge1(data []string) (int, error) {
	parsedData, err := parseData(data)
	if err != nil {
		return 0, err
	}
	result := math.MaxInt
	for _, seed := range parsedData.Seeds {
		carry := seed
	OUTER:
		for step := 0; step < len(parsedData.Mappings); step++ {
			for _, mapping := range parsedData.Mappings[step] {
				if carry >= mapping.SrcStart && carry < mapping.SrcStart+mapping.Length {
					carry = mapping.DestStart + (carry - mapping.SrcStart)
					continue OUTER
				}
			}
		}
		if carry < result {
			result = carry
		}
	}
	return result, nil
}

func Challenge2(data []string) (int, error) {
	parsedData, err := parseData(data)
	if err != nil {
		return 0, err
	}
	result := math.MaxInt
	for seedIndex := 0; seedIndex < len(parsedData.Seeds); seedIndex += 2 {
		for seed := parsedData.Seeds[seedIndex]; seed < parsedData.Seeds[seedIndex]+parsedData.Seeds[seedIndex+1]; seed++ {
			carry := seed
		OUTER:
			for step := 0; step < len(parsedData.Mappings); step++ {
				for _, mapping := range parsedData.Mappings[step] {
					if carry >= mapping.SrcStart && carry < mapping.SrcStart+mapping.Length {
						carry = mapping.DestStart + (carry - mapping.SrcStart)
						continue OUTER
					}
				}
			}
			if carry < result {
				result = carry
			}
		}
	}
	return result, nil
}

func Run() {
	fmt.Println("Day 5 - If You Give A Seed A Fertilizer")
	path := "ifyougiveaseedafertilizer/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int
	result, err = Challenge1(data)
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
