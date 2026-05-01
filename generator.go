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

// Hauptfunktion: Schema -> JSON bytes
func (g *Generator) GenerateJSON(schema []byte) ([]byte, error) {
	var s map[string]any

	if err := json.Unmarshal(schema, &s); err != nil {
		return nil, err
	}

	data := g.generate("", s)

	return json.MarshalIndent(data, "", "  ")
}

// Alternativ: direkt Go-Objekt
func (g *Generator) Generate(schema map[string]any) any {
	return g.generate("", schema)
}
