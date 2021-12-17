package packetdecoder

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	content := scanner.Text()

	return content, nil
}

func hexToBinary(data string) string {
	result := ""

	var mapping map[rune]string = map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}

	for _, val := range data {
		result = result + mapping[val]
	}
	return result
}

func binaryStringToInt(data string) int64 {
	intVersion, _ := strconv.ParseInt(data, 2, 64)
	return intVersion
}

func binarySliceToInt(data []string) int64 {
	newSlice := make([]string, len(data))
	for i, cur := range data {
		newSlice[i] = cur[1:]
	}
	return binaryStringToInt(strings.Join(newSlice, ""))
}

func sumValues(values []int64) int64 {
	var result int64 = 0
	for _, cur := range values {
		result = result + cur
	}
	return result
}

func prodValues(values []int64) int64 {
	var result int64 = 1
	for _, cur := range values {
		result = result * cur
	}
	return result
}

func minValues(values []int64) int64 {
	var result int64 = values[0]
	for _, cur := range values {
		if cur < result {
			result = cur
		}
	}
	return result
}

func maxValues(values []int64) int64 {
	var result int64 = values[0]
	for _, cur := range values {
		if cur > result {
			result = cur
		}
	}
	return result
}

func solve(data string, start int) (int64, int, int64) {
	curVersion := binaryStringToInt(data[start : start+3])
	curType := data[start+3 : start+6]
	curIndex := start + 6
	curValue := int64(0)

	if curType == "100" {
		var groups []string = make([]string, 0)
		for len(data)-curIndex >= 5 {
			group := data[curIndex : curIndex+5]
			curIndex = curIndex + 5
			groups = append(groups, group)
			if group[0] == '0' {
				curValue = binarySliceToInt(groups)
				return curVersion, curIndex - start, curValue
			}
		}
	} else {
		lengthType := data[curIndex]
		curIndex++
		tmpIndex := curIndex
		subValues := []int64{}
		if lengthType == '0' {
			length := int(binaryStringToInt(data[curIndex : curIndex+15]))
			curIndex = curIndex + 15
			tmpIndex = curIndex

			for tmpIndex < curIndex+length {
				tmpVersion, tmpLength, tmpValue := solve(data, tmpIndex)
				curVersion = curVersion + tmpVersion
				tmpIndex = tmpIndex + tmpLength
				subValues = append(subValues, tmpValue)
			}
		} else {
			numSubPackets := int(binaryStringToInt(data[curIndex : curIndex+11]))
			curIndex = curIndex + 11
			tmpIndex = curIndex

			for i := 0; i < numSubPackets; i++ {
				tmpVersion, tmpLength, tmpValue := solve(data, tmpIndex)
				curVersion = curVersion + tmpVersion
				tmpIndex = tmpIndex + tmpLength
				subValues = append(subValues, tmpValue)
			}
		}
		curIndex = tmpIndex
		switch curType {
		case "000":
			curValue = sumValues(subValues)
		case "001":
			curValue = prodValues(subValues)
		case "010":
			curValue = minValues(subValues)
		case "011":
			curValue = maxValues(subValues)
		case "101":
			if subValues[0] > subValues[1] {
				curValue = 1
			}
		case "110":
			if subValues[0] < subValues[1] {
				curValue = 1
			}
		case "111":
			if subValues[0] == subValues[1] {
				curValue = 1
			}
		}
	}

	return curVersion, curIndex - start, curValue
}

func Challenge1(data string) (int64, error) {
	result, _, _ := solve(hexToBinary(data), 0)
	return result, nil
}

func Challenge2(data string) (int64, error) {
	_, _, result := solve(hexToBinary(data), 0)
	return result, nil
}

func Run() {
	fmt.Println("Day 16 - Packet Decoder")
	path := "packetdecoder/input.txt"
	data, err := readLines(path)
	if err != nil {
		fmt.Printf("Failed with error '%v'\n", err)
	}

	var result int64
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
