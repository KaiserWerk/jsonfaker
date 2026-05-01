package jsonfaker

import (
	"encoding/json"
)

type Generator struct {
	opts Options
}

func New(opts ...Option) *Generator {
	g := &Generator{
		opts: defaultOptions(),
	}

	for _, opt := range opts {
		opt(&g.opts)
	}

	return g
}

// GenerateJSON accepts a JSON schema as byte slice and returns generated partially pseudo-random JSON data as byte slice.
func (g *Generator) GenerateJSON(schema []byte) ([]byte, error) {
	var s map[string]any

	if err := json.Unmarshal(schema, &s); err != nil {
		return nil, err
	}

	data := g.generate("", s)

	return json.MarshalIndent(data, "", "  ")
}

// Generate accepts a JSON schema as a map and returns generated partially pseudo-random Go object.
func (g *Generator) Generate(schema map[string]any) any {
	return g.generate("", schema)
}
