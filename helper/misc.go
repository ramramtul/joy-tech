package helper

import (
	"errors"
	"log"
	"strconv"
)

func ParseToBool(input string) (res bool, err error) {
	if input != "true" && input != "false" {
		return false, errors.New("Invalid input")
	}

	res, err = strconv.ParseBool(input)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return res, nil
}

func ParseToInt(input string) (res int, err error) {
	i, err := strconv.Atoi(input)
	if err != nil {
		return -1, err
	}

	return i, nil
}
