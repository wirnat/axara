package ztring

import (
	plural "github.com/gertd/go-pluralize"
	"regexp"
	"strings"
)

// ConvertToHyphenated mengonversi string menjadi format "user-company".
func ConvertToHyphenated(input string) string {
	// Menghapus spasi dan mengganti huruf besar dengan tanda "-"
	re := regexp.MustCompile(`[\s_]+`)
	hyphenated := re.ReplaceAllString(input, "-")

	// Menghapus karakter non-alphanumeric kecuali tanda "-"
	hyphenated = regexp.MustCompile(`[^a-zA-Z0-9\-]+`).ReplaceAllString(hyphenated, "")

	// Menghapus tanda "-" yang berurutan
	hyphenated = strings.Join(removeConsecutiveHyphens(strings.Split(hyphenated, "-")), "-")

	// Mengonversi ke huruf kecil
	hyphenated = strings.ToLower(hyphenated)

	return hyphenated
}

// removeConsecutiveHyphens menghapus tanda "-" yang berurutan dari slice string.
func removeConsecutiveHyphens(parts []string) []string {
	result := []string{}
	prev := ""

	for _, part := range parts {
		if part != "-" || part != prev {
			result = append(result, part)
			prev = part
		}
	}

	return result
}

func Pluralize(s string) string {
	return plural.NewClient().Plural(ConvertToHyphenated(s))
}
