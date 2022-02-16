package parser

import (
	"strconv"
	"strings"
)

func ParseStringMap(raw string) (map[string]string, error) {
	m := map[string]string{}
	emitFn := func(key, value string) error {
		m[key] = value
		return nil
	}
	if err := ParseKeyValues(raw, emitFn); err != nil {
		return nil, err
	}
	return m, nil
}

func FormatStringMap(m map[string]string) string {
	s := strings.Builder{}
	for key, value := range m {
		s.WriteString(key)
		if value != "" {
			s.WriteString(":")
			s.WriteString(strconv.Quote(value))
		}
		s.WriteString(",")
	}
	return strings.TrimSuffix(s.String(), ",")
}
