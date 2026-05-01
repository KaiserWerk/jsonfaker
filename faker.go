package jsonfaker

import (
	"math/rand"
	"strings"
	"time"
)

func defaultStringFaker(field string, schema map[string]any, r *rand.Rand) string {
	field = strings.ToLower(field)

	if format, ok := schema["format"].(string); ok {
		switch format {

		case "email":
			return "user@example.com"

		case "uri":
			return "https://example.com"

		case "date":
			return randomDate(r).Format("2006-01-02")

		case "date-time":
			return randomDate(r).Format(time.RFC3339)
		}
	}

	switch {
	case strings.Contains(field, "title"):
		return "Awesome Product"

	case strings.Contains(field, "name"):
		return "Sample Name"

	case strings.Contains(field, "description"):
		return "This is a great product."

	default:
		return "example"
	}
}
