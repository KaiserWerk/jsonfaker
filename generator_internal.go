package jsonfaker

func (g *Generator) generate(field string, schema map[string]any) any {

	if enum, ok := schema["enum"].([]any); ok && len(enum) > 0 {
		return enum[0]
	}

	t, _ := schema["type"].(string)

	switch t {

	case "object":
		result := map[string]any{}
		props, _ := schema["properties"].(map[string]any)

		for key, val := range props {
			result[key] = g.generate(key, val.(map[string]any))
		}
		return result

	case "array":
		items, _ := schema["items"].(map[string]any)
		return []any{
			g.generate(field, items),
		}

	case "string":
		return g.opts.StringFaker(field, schema, g.opts.Rand)

	case "number":
		return g.opts.NumberFaker(field, schema, g.opts.Rand)

	case "integer":
		return g.opts.IntegerFaker(field, schema, g.opts.Rand)

	case "boolean":
		return g.opts.BoolFaker(field, schema, g.opts.Rand)

	default:
		return nil
	}
}
