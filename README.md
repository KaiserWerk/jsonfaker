# jsonfaker

`jsonfaker` is a lightweight Go package for generating sample data from a JSON Schema.

It is designed for simple to moderately complex schemas, such as product data, forms, or API payloads, and produces valid, realistic-looking JSON structures without external dependencies.

## Features

* Generates JSON from JSON Schema (`type`, `properties`, `items`)
* Supports basic data types
* Respects simple constraints (`minimum`, `maximum`, `enum`)
* Detects field names (for example `price`, `title`) to generate more meaningful fake data
* Supports string `format` values (for example `date-time`)
* Produces random values for numbers and booleans
* Can be deterministic via seed, which is useful for tests
* Easily extensible through configurable faker functions

---

## Installation

```bash
go get github.com/KaiserWerk/jsonfaker
```

## Usage

```go
gen := jsonfaker.New()

result, err := gen.GenerateJSON(schemaBytes)
if err != nil {
    panic(err)
}

fmt.Println(string(result))
```

## Supported Data Types

| JSON Schema Type | Support               |
| ---------------- | --------------------- |
| `object`         | recursive          |
| `array`          | (1 element)        |
| `string`         | yes                |
| `number`         | (random float64)   |
| `integer`        | (random int)       |
| `boolean`        | (random)           |

## Supported Schema Features

| Feature      | Support             |
| ------------ | ------------------- |
| `properties` | yes                 |
| `items`      | yes                 |
| `enum`       | yes (first element) |
| `minimum`    | yes                 |
| `maximum`    | yes                 |
| `format`     | yes (partial)       |

## Supported String Formats

| Format      | Example                                    |
| ----------- | ------------------------------------------- |
| `email`     | [user@example.com](mailto:user@example.com) |
| `uri`       | https://example.com                         |
| `date`      | 2026-01-02                                  |
| `date-time` | 2026-01-02T15:04:05Z                        |


## Random Data And Determinism

By default, random values are generated.

For reproducible results, use a fixed seed:

```go
gen := jsonfaker.New(
    jsonfaker.WithSeed(42),
)
```

## Customization

All data generators can be overridden:

```go
gen := jsonfaker.New(func(o *jsonfaker.Options) {
    o.StringFaker = func(field string, schema map[string]interface{}, r *rand.Rand) string {
        return "custom"
    }
})
```

## Limitations

This package is intentionally simple and does **not** support:

* `$ref` resolution
* `oneOf`, `anyOf`, `allOf`
* complex validation rules
* the full JSON Schema specification
