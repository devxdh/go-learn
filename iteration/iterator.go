package iteration

import "strings"

func Repeat(c string) string {
	var repeated strings.Builder
	for i := 0; i < 5; i++ {
		repeated.WriteString(c)
	}

	return repeated.String()
}
