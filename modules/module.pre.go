package modules

import (
	"errors"
	"strconv"
)

func SingleInt64(args []string) (int64, error) {
	if len(args) != 1 {
		return -1, errors.New(" requires exactly 1 argument")
	}
	return strconv.ParseInt(args[0], 10, 64)

}

func DoubleInt64(args []string) (a, b int64, err error) {
	if len(args) != 3 {
		err = errors.New(" requires exactly 3 argument")
		return
	}

	a, err = strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return
	}

	b, err = strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return
	}
	return
}

func TripleInt64(args []string) (a, b, c int64, err error) {
	if len(args) != 3 {
		err = errors.New(" requires exactly 3 argument")
		return
	}

	a, err = strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return
	}

	b, err = strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return
	}

	c, err = strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return
	}
	return
}

func KInt64s(args []string, k int) (parsed []int64, err error) {
	parsed = make([]int64, k)

	for i := 0; i < k; i++ {
		parsed[i], err = strconv.ParseInt(args[i], 10, 64)
	}
	return
}
