package stringfmt

import (
	"fmt"
)

type StringFormatter func(string) string

func Prefixer(prefix string) StringFormatter {
	return func(in string) string {
		return prefix + in
	}
}

func FormatPrefixer(prefix string) StringFormatter {
	return func(in string) string {
		return fmt.Sprintf("%s. %s", prefix, in)
	}
}

func Postfixer(postfix string) StringFormatter {
	return func(in string) string {
		return in + postfix
	}
}
