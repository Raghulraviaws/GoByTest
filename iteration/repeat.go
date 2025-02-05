package iteration

import "strings"

const repeatCount = 5

func Repeat(name string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(name)
	}
	return repeated.String()
}
