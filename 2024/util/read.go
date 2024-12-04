package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ReadFile reads lines of text file f into slices of strings.
// It returns any error it encounters.
func ReadFile(f string) ([]string, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var records []string
	for scanner.Scan() {
		records = append(records, scanner.Text())
	}

	return records, nil
}

// readLine reads the text file f and flattens it as a
// single string. It returns any error it encounters.
func ReadLine(f string) (string, error) {
	records, err := ReadFile(f)
	if err != nil {
		return "", err
	}
	return strings.Join(records, ""), nil
}

// stringsToStrings converts slices of strings to a matrix of ints.
// It returns any error it encounters.
func stringsToStrings(s []string) ([][]string, error) {
	res := make([][]string, 0, len(s))

	for _, line := range s {
		l := strings.Fields(line)
		res = append(res, l)
	}

	return res, nil
}

// stringsToInts converts slices of strings to a matrix of ints.
// It returns any error it encounters.
func stringsToInts(s []string) ([][]int, error) {
	res := make([][]int, 0, len(s))

	for _, line := range s {
		l := strings.Fields(line)
		ints := make([]int, 0, len(l))
		for _, v := range l {
			d, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			ints = append(ints, d)
		}
		res = append(res, ints)
	}

	return res, nil
}

// ReadSpaceSeparatedFileToStr reads the text file `input` to a matrix
// of strings. It returns any error it encounters.
func ReadSpaceSeparatedFileToStr(input string) ([][]string, error) {
	records, err := ReadFile(input)
	if err != nil {
		return nil, err
	}
	return stringsToStrings(records)
}

// ReadSpaceSeparatedFileToInt reads the text file `input` to a matrix
// of integers. It returns any error it encounters.
func ReadSpaceSeparatedFileToInt(input string) ([][]int, error) {
	records, err := ReadFile(input)
	if err != nil {
		return nil, err
	}
	return stringsToInts(records)
}
