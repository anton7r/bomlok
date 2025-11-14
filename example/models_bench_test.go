package example

import (
	"reflect"
	"testing"
)

// Benchmark bomlok generated GetValue method vs reflection for Person struct
func BenchmarkPerson_Bomlok_GetValue(b *testing.B) {
	p := &Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Email:     "john.doe@example.com",
	}

	b.Run("FirstField", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = p.Bomlok_GetValue("FirstName")
		}
	})

	b.Run("MiddleField", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = p.Bomlok_GetValue("Age")
		}
	})

	b.Run("LastField", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = p.Bomlok_GetValue("Email")
		}
	})

	b.Run("InvalidField", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = p.Bomlok_GetValue("NonExistent")
		}
	})
}

func BenchmarkPerson_Reflection_GetValue(b *testing.B) {
	p := &Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Email:     "john.doe@example.com",
	}

	getValue := func(obj interface{}, fieldName string) interface{} {
		val := reflect.ValueOf(obj).Elem()
		field := val.FieldByName(fieldName)
		if !field.IsValid() {
			return nil
		}
		return field.Interface()
	}

	b.Run("FirstField", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = getValue(p, "FirstName")
		}
	})

	b.Run("MiddleField", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = getValue(p, "Age")
		}
	})

	b.Run("LastField", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = getValue(p, "Email")
		}
	})

	b.Run("InvalidField", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = getValue(p, "NonExistent")
		}
	})
}

// Benchmark for Address struct
func BenchmarkAddress_Bomlok_GetValue(b *testing.B) {
	a := &Address{
		Street:  "123 Main St",
		City:    "Springfield",
		State:   "IL",
		ZipCode: "62701",
	}

	b.Run("AllFields", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = a.Bomlok_GetValue("Street")
			_ = a.Bomlok_GetValue("City")
			_ = a.Bomlok_GetValue("State")
			_ = a.Bomlok_GetValue("ZipCode")
		}
	})
}

func BenchmarkAddress_Reflection_GetValue(b *testing.B) {
	a := &Address{
		Street:  "123 Main St",
		City:    "Springfield",
		State:   "IL",
		ZipCode: "62701",
	}

	getValue := func(obj interface{}, fieldName string) interface{} {
		val := reflect.ValueOf(obj).Elem()
		field := val.FieldByName(fieldName)
		if !field.IsValid() {
			return nil
		}
		return field.Interface()
	}

	b.Run("AllFields", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = getValue(a, "Street")
			_ = getValue(a, "City")
			_ = getValue(a, "State")
			_ = getValue(a, "ZipCode")
		}
	})
}

// Benchmark for Company struct with mixed types
func BenchmarkCompany_Bomlok_GetValue(b *testing.B) {
	c := &Company{
		Name:      "Acme Corp",
		Employees: 1000,
		Revenue:   5000000.50,
		Public:    true,
	}

	b.Run("StringField", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = c.Bomlok_GetValue("Name")
		}
	})

	b.Run("IntField", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = c.Bomlok_GetValue("Employees")
		}
	})

	b.Run("Float64Field", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = c.Bomlok_GetValue("Revenue")
		}
	})

	b.Run("BoolField", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = c.Bomlok_GetValue("Public")
		}
	})
}

func BenchmarkCompany_Reflection_GetValue(b *testing.B) {
	c := &Company{
		Name:      "Acme Corp",
		Employees: 1000,
		Revenue:   5000000.50,
		Public:    true,
	}

	getValue := func(obj interface{}, fieldName string) interface{} {
		val := reflect.ValueOf(obj).Elem()
		field := val.FieldByName(fieldName)
		if !field.IsValid() {
			return nil
		}
		return field.Interface()
	}

	b.Run("StringField", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = getValue(c, "Name")
		}
	})

	b.Run("IntField", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = getValue(c, "Employees")
		}
	})

	b.Run("Float64Field", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = getValue(c, "Revenue")
		}
	})

	b.Run("BoolField", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = getValue(c, "Public")
		}
	})
}

// Benchmark Bomlok_Fields method
func BenchmarkPerson_Bomlok_Fields(b *testing.B) {
	p := &Person{}

	for i := 0; i < b.N; i++ {
		_ = p.Bomlok_Fields()
	}
}

func BenchmarkAddress_Bomlok_Fields(b *testing.B) {
	a := &Address{}

	for i := 0; i < b.N; i++ {
		_ = a.Bomlok_Fields()
	}
}

func BenchmarkCompany_Bomlok_Fields(b *testing.B) {
	c := &Company{}

	for i := 0; i < b.N; i++ {
		_ = c.Bomlok_Fields()
	}
}

// Benchmark sequential field access pattern
func BenchmarkPerson_SequentialAccess(b *testing.B) {
	p := &Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Email:     "john.doe@example.com",
	}

	b.Run("Bomlok", func(b *testing.B) {
		fields := p.Bomlok_Fields()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, field := range fields {
				_ = p.Bomlok_GetValue(field)
			}
		}
	})

	b.Run("Reflection", func(b *testing.B) {
		val := reflect.ValueOf(p).Elem()
		typ := val.Type()
		numFields := val.NumField()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := 0; j < numFields; j++ {
				fieldName := typ.Field(j).Name
				field := val.FieldByName(fieldName)
				_ = field.Interface()
			}
		}
	})
}
