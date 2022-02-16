package parser

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEmptyKey       = errors.New("empty key")
	ErrUnexpectedRune = errors.New("unexpected rune")
)

type EmitFunc func(string, string) error

func ParseKeyValues(raw string, emitFn EmitFunc) error {
	s := state{parseValues: true, emitFn: emitFn}

	parseFn := parseFunc(parseKeyBegin)
	err := error(nil)

	for index, c := range raw {
		parseFn, err = parseFn(&s, c)
		if err != nil {
			return fmt.Errorf("at index %d: %w", index, err)
		}
	}

	s.emit()

	return nil
}

func ParseKeys(raw string, emitFn EmitFunc) error {
	s := state{parseValues: false, emitFn: emitFn}

	parseFn := parseFunc(parseKeyBegin)
	err := error(nil)

	for index, c := range raw {
		parseFn, err = parseFn(&s, c)
		if err != nil {
			return fmt.Errorf("at index %d: %w", index, err)
		}
	}

	s.emit()

	return nil
}

type parseFunc func(*state, rune) (parseFunc, error)

func parseKeyBegin(s *state, c rune) (parseFunc, error) {
	switch c {
	case ' ': // skip
	case ',':
		if err := s.emit(); err != nil {
			return nil, err
		}
		return parseKeyBegin, nil
	case '\'':
		return parseSingleQuotedKey, nil
	case '"':
		return parseDoubleQuotedKey, nil
	default:
		s.key.WriteRune(c)
		return parseKeyRaw, nil
	}
	return parseKeyBegin, nil
}

func parseKeyEnd(s *state, c rune) (parseFunc, error) {
	switch c {
	case ' ': // skip
	case ':':
		if s.parseValues {
			return parseValueBegin, nil
		}
		return nil, ErrUnexpectedRune
	case ',':
		if err := s.emit(); err != nil {
			return nil, err
		}
		return parseKeyBegin, nil
	default:
		return nil, ErrUnexpectedRune
	}
	return parseValueEnd, nil
}

func parseKeyRaw(s *state, c rune) (parseFunc, error) {
	switch c {
	case ' ': // skip
	case ':':
		if s.parseValues {
			return parseValueBegin, nil
		}
		s.key.WriteRune(c)
	case ',':
		if err := s.emit(); err != nil {
			return nil, err
		}
		return parseKeyBegin, nil
	default:
		s.key.WriteRune(c)
	}
	return parseKeyRaw, nil
}

func parseSingleQuotedKey(s *state, c rune) (parseFunc, error) {
	switch c {
	case '\\':
		return func(s *state, c rune) (parseFunc, error) {
			s.key.WriteRune(c)
			return parseSingleQuotedKey, nil
		}, nil
	case '\'':
		return parseKeyEnd, nil
	default:
		s.key.WriteRune(c)
	}
	return parseSingleQuotedKey, nil
}

func parseDoubleQuotedKey(s *state, c rune) (parseFunc, error) {
	switch c {
	case '\\':
		return func(s *state, c rune) (parseFunc, error) {
			s.key.WriteRune(c)
			return parseDoubleQuotedKey, nil
		}, nil
	case '"':
		return parseKeyEnd, nil
	default:
		s.key.WriteRune(c)
	}
	return parseDoubleQuotedKey, nil
}

func parseValueBegin(s *state, c rune) (parseFunc, error) {
	switch c {
	case ' ': // skip
	case ',':
		if err := s.emit(); err != nil {
			return nil, err
		}
		return parseKeyBegin, nil
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
		return parseKeyBegin, nil
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
		return parseKeyBegin, nil
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
	parseValues bool
	key         strings.Builder
	value       strings.Builder
	emitFn      EmitFunc
}

func (s *state) emit() error {
	if s.key.Len() == 0 {
		return ErrEmptyKey
	}

	if err := s.emitFn(s.key.String(), s.value.String()); err != nil {
		return err
	}

	s.key.Reset()
	s.value.Reset()

	return nil
}
