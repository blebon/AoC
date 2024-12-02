package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// readFile reads lines of text file f into slices of strings.
func readFile(f string) ([]string, error) {
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

// stringsToInts converts slices of strings to a matrix of ints.
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

// ReadSpaceSeparatedFile reads the text file `input` to a matrix
// of integers. It returns any error it encounters.
func ReadSpaceSeparatedFile(input string) ([][]int, error) {
	records, err := readFile(input)
	if err != nil {
		return nil, err
	}
	return stringsToInts(records)
}
