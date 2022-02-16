package parser

import (
	"strconv"
	"strings"
)

func ParseStrings(raw string) ([]string, error) {
	values := []string{}
	emitFn := func(key, _ string) error {
		values = append(values, key)
		return nil
	}
	if err := ParseKeys(raw, emitFn); err != nil {
		return nil, err
	}
	return values, nil
}

func FormatStrings(values []string) string {
	s := strings.Builder{}
	for _, value := range values {
		s.WriteString(strconv.Quote(value))
		s.WriteString(",")
	}
	return strings.TrimSuffix(s.String(), ",")
}
