package parser

import (
	"strconv"
	"strings"
)

func ParseInts(raw string) ([]int, error) {
	values := []int{}
	emitFn := func(value, _ string) error {
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		values = append(values, int(v))
		return nil
	}
	if err := ParseKeyValues(raw, emitFn); err != nil {
		return nil, err
	}
	return values, nil
}

func FormatInts(values []int) string {
	s := strings.Builder{}
	for _, value := range values {
		s.WriteString(strconv.Itoa(value))
		s.WriteString(",")
	}
	return strings.TrimSuffix(s.String(), ",")
}
