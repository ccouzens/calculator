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
	t_garbage
)

type Token struct {
	Lexeme string
	Type   token_type
}

func Lex(input string, c chan Token) {
	if white_space_regex.MatchString(input) {
		c <- Token{input, t_white_space}
	} else if literal_number_regex.MatchString(input) {
		c <- Token{input, t_literal_number}
	} else if plus_regex.MatchString(input) {
		c <- Token{input, t_plus}
	} else {
		c <- Token{input, t_garbage}
	}
	close(c)
}

func Parse(input string) (int64, error) {

	return strconv.ParseInt(input, 10, 64)
}
