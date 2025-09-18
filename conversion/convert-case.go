package conversion

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ฟังก์ชันกลาง: แปลง string → slice ของ tokens
func normalizeTokens(input string) []string {
	// แทน space และ dash เป็น underscore
	input = strings.ReplaceAll(input, "-", "_")
	input = strings.ReplaceAll(input, " ", "_")

	// แทรก _ ระหว่าง lower/number + Upper
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	input = re.ReplaceAllString(input, "${1}_${2}")

	// แปลงเป็น lowercase
	input = strings.ToLower(input)

	// split ด้วย underscore
	tokens := strings.Split(input, "_")

	// ลบช่องว่าง
	clean := tokens[:0]
	for _, t := range tokens {
		if t != "" {
			clean = append(clean, t)
		}
	}
	return clean
}

var titleCaser = cases.Title(language.English) // ใช้สำหรับ uppercase ตัวแรก

// -------------------- Case Functions --------------------

// to snake_case
func ToSnake(input string) string {
	return strings.Join(normalizeTokens(input), "_")
}

// to camelCase
func ToCamel(input string) string {
	tokens := normalizeTokens(input)
	for i := range tokens {
		if i == 0 {
			continue
		}
		tokens[i] = titleCaser.String(tokens[i])
	}
	return strings.Join(tokens, "")
}

// to PascalCase
func ToPascal(input string) string {
	tokens := normalizeTokens(input)
	for i := range tokens {
		tokens[i] = titleCaser.String(tokens[i])
	}
	return strings.Join(tokens, "")
}

// to kebab-case
func ToKebab(input string) string {
	return strings.Join(normalizeTokens(input), "-")
}

// to CONSTANT_CASE
func ToConstant(input string) string {
	return strings.ToUpper(strings.Join(normalizeTokens(input), "_"))
}
