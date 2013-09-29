package calculator

import (
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
	return []Token{}, nil
}

func Parse(input string) (int64, error) {

	return strconv.ParseInt(input, 10, 64)
}
