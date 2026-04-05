package main

import (
	"strings"
	"unicode"
)

// stopWords are filtered from entity extraction.
var stopWords = map[string]bool{
	"a": true, "an": true, "the": true, "and": true, "or": true,
	"but": true, "in": true, "on": true, "at": true, "to": true,
	"for": true, "of": true, "with": true, "by": true, "from": true,
	"is": true, "are": true, "was": true, "were": true, "be": true,
	"how": true, "what": true, "when": true, "where": true, "why": true,
	"this": true, "that": true, "it": true, "its": true, "not": true,
	"no": true, "do": true, "does": true, "did": true, "has": true,
	"have": true, "had": true, "will": true, "would": true, "can": true,
	"could": true, "should": true, "may": true, "might": true, "new": true,
	"via": true, "vs": true, "about": true, "into": true, "just": true,
	"i": true, "me": true, "my": true, "we": true, "you": true, "your": true,
	"he": true, "she": true, "they": true, "them": true, "their": true,
	"there": true, "here": true, "some": true, "any": true, "all": true,
	"been": true, "being": true, "get": true, "got": true, "also": true,
	"like": true, "want": true, "need": true, "know": true, "think": true,
	"make": true, "made": true, "use": true, "used": true, "using": true,
	"one": true, "two": true, "more": true, "than": true, "very": true,
}

// extractTitleEntities pulls meaningful terms from a title string.
// Returns up to maxEntities cleaned tokens, filtering stop words,
// punctuation, short tokens, and pure numerics.
func extractTitleEntities(title string, maxEntities int) []string {
	words := strings.Fields(title)
	seen := make(map[string]bool)
	var entities []string

	for _, w := range words {
		cleaned := strings.TrimFunc(w, func(r rune) bool {
			return unicode.IsPunct(r) || unicode.IsSymbol(r)
		})
		lower := strings.ToLower(cleaned)

		if len(cleaned) < 3 {
			continue
		}
		if stopWords[lower] {
			continue
		}
		if isNumericToken(cleaned) {
			continue
		}
		if seen[lower] {
			continue
		}
		seen[lower] = true
		entities = append(entities, cleaned)
		if len(entities) >= maxEntities {
			break
		}
	}
	return entities
}

// isNumericToken returns true if s contains only digits and date separators.
func isNumericToken(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) && r != '-' && r != '/' && r != '.' {
			return false
		}
	}
	return true
}
