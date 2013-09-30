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
	for len(input) > 0 {
		if white_space_regex.MatchString(input) {
			match := white_space_regex.FindString(input)
			c <- Token{match, t_white_space}
			input = input[len(match):]
		} else if literal_number_regex.MatchString(input) {
			match := literal_number_regex.FindString(input)
			c <- Token{match, t_literal_number}
			input = input[len(match):]
		} else if plus_regex.MatchString(input) {
			match := plus_regex.FindString(input)
			c <- Token{match, t_plus}
			input = input[len(match):]
		} else {
			c <- Token{input, t_garbage}
			input = input[len(input):]
		}
	}
	close(c)
}

func Parse(input string) (int64, error) {

	return strconv.ParseInt(input, 10, 64)
}
