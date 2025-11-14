package example

import (
	"testing"
)

func TestPersonBomlok(t *testing.T) {
	person := &Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Email:     "john@example.com",
	}

	// Test Bomlok_Fields
	fields := person.Bomlok_Fields()
	expectedFields := []string{"FirstName", "LastName", "Age", "Email"}

	if len(fields) != len(expectedFields) {
		t.Errorf("Expected %d fields, got %d", len(expectedFields), len(fields))
	}

	for i, field := range expectedFields {
		if fields[i] != field {
			t.Errorf("Expected field %s at index %d, got %s", field, i, fields[i])
		}
	}

	// Test Bomlok_GetValue
	tests := []struct {
		field    string
		expected any
	}{
		{"FirstName", "John"},
		{"LastName", "Doe"},
		{"Age", 30},
		{"Email", "john@example.com"},
		{"NonExistent", nil},
	}

	for _, tt := range tests {
		t.Run(tt.field, func(t *testing.T) {
			value := person.Bomlok_GetValue(tt.field)
			if value != tt.expected {
				t.Errorf("Expected %v for field %s, got %v", tt.expected, tt.field, value)
			}
		})
	}
}

func TestAddressBomlok(t *testing.T) {
	address := &Address{
		Street:  "123 Main St",
		City:    "New York",
		State:   "NY",
		ZipCode: "10001",
	}

	fields := address.Bomlok_Fields()
	if len(fields) != 4 {
		t.Errorf("Expected 4 fields, got %d", len(fields))
	}

	street := address.Bomlok_GetValue("Street")
	if street != "123 Main St" {
		t.Errorf("Expected '123 Main St', got %v", street)
	}

	city := address.Bomlok_GetValue("City")
	if city != "New York" {
		t.Errorf("Expected 'New York', got %v", city)
	}
}

func TestCompanyBomlok(t *testing.T) {
	company := &Company{
		Name:      "Tech Corp",
		Employees: 500,
		Revenue:   1000000.50,
		Public:    true,
	}

	// Test getting all values dynamically
	fields := company.Bomlok_Fields()
	for _, field := range fields {
		value := company.Bomlok_GetValue(field)
		if value == nil {
			t.Errorf("Field %s returned nil", field)
		}
	}

	// Test specific values
	if company.Bomlok_GetValue("Name") != "Tech Corp" {
		t.Error("Name mismatch")
	}

	if company.Bomlok_GetValue("Employees") != 500 {
		t.Error("Employees mismatch")
	}

	if company.Bomlok_GetValue("Revenue") != 1000000.50 {
		t.Error("Revenue mismatch")
	}

	if company.Bomlok_GetValue("Public") != true {
		t.Error("Public mismatch")
	}
}
