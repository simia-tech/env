package parser

import (
	"errors"
	"strings"
)

var (
	ErrEmptyKey       = errors.New("empty key")
	ErrUnexpectedRune = errors.New("unexpected rune")
)

type EmitFunc func(string, string)

func ParseKeyValues(raw string, emitFn EmitFunc) error {
	s := state{emitFn: emitFn}

	parseFn := parseFunc(parseKey)
	err := error(nil)

	for _, c := range raw {
		parseFn, err = parseFn(&s, c)
		if err != nil {
			return err
		}
	}

	s.emit()

	return nil
}

type parseFunc func(*state, rune) (parseFunc, error)

func parseKey(s *state, c rune) (parseFunc, error) {
	switch c {
	case ' ': // skip
	case ':':
		return parseValueBegin, nil
	case ',':
		if err := s.emit(); err != nil {
			return nil, err
		}
	default:
		s.key.WriteRune(c)
	}
	return parseKey, nil
}

func parseValueBegin(s *state, c rune) (parseFunc, error) {
	switch c {
	case ' ': // skip
	case ',':
		if err := s.emit(); err != nil {
			return nil, err
		}
		return parseKey, nil
	case '\'':
		return parseSingleQuotedValue, nil
	case '"':
		return parseDoubleQuotedValue, nil
	default:
		s.value.WriteRune(c)
		return parseValueRaw, nil
	}
	return parseValueBegin, nil
}

func parseValueEnd(s *state, c rune) (parseFunc, error) {
	switch c {
	case ' ': // skip
	case ',':
		if err := s.emit(); err != nil {
			return nil, err
		}
		return parseKey, nil
	default:
		return nil, ErrUnexpectedRune
	}
	return parseValueEnd, nil
}

func parseValueRaw(s *state, c rune) (parseFunc, error) {
	switch c {
	case ',':
		if err := s.emit(); err != nil {
			return nil, err
		}
		return parseKey, nil
	default:
		s.value.WriteRune(c)
	}
	return parseValueRaw, nil
}

func parseSingleQuotedValue(s *state, c rune) (parseFunc, error) {
	switch c {
	case '\\':
		return func(s *state, c rune) (parseFunc, error) {
			s.value.WriteRune(c)
			return parseSingleQuotedValue, nil
		}, nil
	case '\'':
		return parseValueEnd, nil
	default:
		s.value.WriteRune(c)
	}
	return parseSingleQuotedValue, nil
}

func parseDoubleQuotedValue(s *state, c rune) (parseFunc, error) {
	switch c {
	case '\\':
		return func(s *state, c rune) (parseFunc, error) {
			s.value.WriteRune(c)
			return parseDoubleQuotedValue, nil
		}, nil
	case '"':
		return parseValueEnd, nil
	default:
		s.value.WriteRune(c)
	}
	return parseDoubleQuotedValue, nil
}

type state struct {
	key    strings.Builder
	value  strings.Builder
	emitFn EmitFunc
}

func (s *state) emit() error {
	if s.key.Len() == 0 {
		return ErrEmptyKey
	}

	s.emitFn(s.key.String(), s.value.String())

	s.key.Reset()
	s.value.Reset()

	return nil
}
