package list

import (
	"fmt"
	"strings"
)

type ListFormat string

const (
	All     ListFormat = "all"
	Partial ListFormat = "partial"
)

func (e *ListFormat) String() string {
	return string(*e)
}

func (e *ListFormat) Set(s string) error {
	switch strings.ToLower(s) {
	case "all", "partial":
		*e = ListFormat(s)
		return nil
	default:
		return fmt.Errorf("invalid list format: %s. Must be one of %v", s, []ListFormat{All, Partial})
	}
}

func (e *ListFormat) Type() string {
	return "ListFormat"
}
