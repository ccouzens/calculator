package calculator

import (
	"errors"
	"regexp"
	"strconv"
)

type token_type int

var white_space_regex = regexp.MustCompile("^\\s+")
var literal_number_regex = regexp.MustCompile("^\\d+")
var plus_regex = regexp.MustCompile("^\\+")

const (
	t_white_space token_type = iota
	t_literal_number
	t_plus
)

type Token struct {
	Lexeme string
	Type   token_type
}

func Lex(input string) ([]Token, error) {
	if white_space_regex.MatchString(input) {
		return []Token{{input, t_white_space}}, nil
	}
	if literal_number_regex.MatchString(input) {
		return []Token{{input, t_literal_number}}, nil
	}
	if plus_regex.MatchString(input) {
		return []Token{{input, t_plus}}, nil
	}

	return nil, errors.New("Input not recognized")
}

func Parse(input string) (int64, error) {

	return strconv.ParseInt(input, 10, 64)
}
