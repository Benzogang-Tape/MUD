package utils

import (
	"fmt"
)

type StrFuncTypeStr func(string) string

func Prefixer(prefix string) StrFuncTypeStr {
	return func(in string) string {
		return prefix + in
	}
}

func FormatPrefixer(prefix string) StrFuncTypeStr {
	return func(in string) string {
		return fmt.Sprintf("%s. %s", prefix, in)
	}
}

func Postfixer(postfix string) StrFuncTypeStr {
	return func(in string) string {
		return in + postfix
	}
}

func ContainerExists(gear [][]string, container string) bool {
	var flag bool
	for _, containerContent := range gear {
		if len(containerContent) != 0 && containerContent[0] == container {
			flag = true
			break
		}
	}
	return flag
}
