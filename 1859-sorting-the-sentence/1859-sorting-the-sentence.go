package main

import (
	"strconv"
	"strings"
)

func sortSentence(s string) string {
	sentence := strings.Fields(s)
	slice := make([]string, len(sentence))

	for _, v := range sentence {
		idx, _ := strconv.Atoi(v[len(v)-1:])
		word := v[:len(v)-1]
		slice[idx-1] = word
	}
	return strings.Join(slice, " ")
}
