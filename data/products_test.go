package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "bartek",
		Price: 1.00,
		SKU: "abc-abc-ade",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
