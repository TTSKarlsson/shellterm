package shellterm

import "strings"

func RepeatRune(r rune, count int) string {
	sb := strings.Builder{}
	for range count {
		sb.WriteRune(r)
	}
	return sb.String()
}
