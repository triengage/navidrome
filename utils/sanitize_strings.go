package utils

import (
	"regexp"
	"sort"
	"strings"

	"github.com/kennygrant/sanitize"
)

var quotesRegex = regexp.MustCompile("[“”‘’'\"\\[\\(\\{\\]\\)\\}]")

func SanitizeStrings(text ...string) string {
	sanitizedText := strings.Builder{}
	for _, txt := range text {
		sanitizedText.WriteString(strings.TrimSpace(sanitize.Accents(strings.ToLower(txt))) + " ")
	}
	words := make(map[string]struct{})
	for _, w := range strings.Fields(sanitizedText.String()) {
		words[w] = struct{}{}
	}
	var fullText []string
	for w := range words {
		w = quotesRegex.ReplaceAllString(w, "")
		if w != "" {
			fullText = append(fullText, w)
		}
	}
	sort.Strings(fullText)
	return strings.Join(fullText, " ")
}
