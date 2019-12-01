package input

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// GetInputFileAsFloat parses the input file and returns a slice of floats
func GetInputFileAsFloat(file string) ([]float64, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	cleaned := strings.TrimSpace(string(data))
	lines := strings.Split(cleaned, "\n")
	result := make([]float64, len(lines))

	for i, l := range lines {
		num, err := strconv.ParseFloat(l, 64)
		if err != nil {
			return result, err
		}

		result[i] = num
	}

	return result, nil
}

// GetInputFileAsInt parses the input file and returns a slice of ints
func GetInputFileAsInt(file string) ([]int, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	cleaned := strings.TrimSpace(string(data))
	lines := strings.Split(cleaned, "\n")
	result := make([]int, len(lines))

	for i, l := range lines {
		num, err := strconv.Atoi(l)
		if err != nil {
			return result, err
		}

		result[i] = num
	}

	return result, nil
}
