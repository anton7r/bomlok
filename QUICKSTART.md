# Bomlok Quick Reference

## Installation

```bash
go install github.com/anton7r/bomlok/cmd/bomlok@latest
```

## Basic Commands

```bash
# Generate in current directory
bomlok

# Include specific directories
bomlok -include=./models -include=./entities

# Exclude directories
bomlok -exclude=vendor -exclude=testdata

# Verbose output
bomlok -verbose

# Combine options
bomlok -include=./src -exclude=test -exclude=vendor -verbose
```

## Generated Methods

For each struct, two methods are generated:

```go
// Get field value by name
func (s *StructName) Bomlok_GetValue(field string) any

// Get all field names
func (s *StructName) Bomlok_Fields() []string
```

## Usage Example

```go
type Person struct {
    Name string
    Age  int
}

// After running bomlok, you can:
person := &Person{Name: "John", Age: 30}

// Get all fields
fields := person.Bomlok_Fields()
// Returns: []string{"Name", "Age"}

// Get field value
name := person.Bomlok_GetValue("Name")
// Returns: "John"

age := person.Bomlok_GetValue("Age")
// Returns: 30

invalid := person.Bomlok_GetValue("Invalid")
// Returns: nil
```

## Integration with go generate

```go
//go:generate bomlok -include=.
package mypackage
```

Then run:
```bash
go generate ./...
```

## File Naming Convention

| Input File | Output File |
|------------|-------------|
| `models.go` | `models_bomlok.go` |
| `user.go` | `user_bomlok.go` |
| `entities.go` | `entities_bomlok.go` |

## Skipped Files

The generator automatically skips:
- Test files (`*_test.go`)
- Generated files (`*_bomlok.go`)
- Excluded directories

## Common Use Cases

### 1. Dynamic Field Access

```go
func PrintAllFields(b Bomlok) {
    for _, field := range b.Bomlok_Fields() {
        fmt.Printf("%s: %v\n", field, b.Bomlok_GetValue(field))
    }
}
```

### 2. Generic Field Mapper

```go
func ToMap(b Bomlok) map[string]any {
    result := make(map[string]any)
    for _, field := range b.Bomlok_Fields() {
        result[field] = b.Bomlok_GetValue(field)
    }
    return result
}
```

### 3. Configuration Loader

```go
func LoadConfig(b Bomlok, config map[string]any) {
    for _, field := range b.Bomlok_Fields() {
        if val, ok := config[field]; ok {
            // Set field value based on config
        }
    }
}
```

## Performance Tips

- Generated code uses switch statements (fast!)
- No reflection overhead at runtime
- Type-safe and compiler-optimized
- Zero external dependencies

## Troubleshooting

### Generator not finding structs?
- Make sure files have `.go` extension
- Check file isn't excluded
- Run with `-verbose` to see what's being processed

### Generated code has compile errors?
- Re-run the generator
- Check if struct definitions changed
- Verify package names are correct

### Want to regenerate all files?
```bash
# Delete old generated files
find . -name "*_bomlok.go" -delete

# Regenerate
bomlok -include=. -verbose
```

## Help

```bash
bomlok -help
```

## Links

- [Full Documentation](./README.md)
- [Architecture Guide](./ARCHITECTURE.md)
- [Example Code](./example/)
- [Demo Application](./demo/)
