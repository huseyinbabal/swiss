package uuid

import "testing"

func TestGenerateAndValidate(t *testing.T) {
	tests := []struct {
		name string
		gen  func() (string, error)
	}{
		{name: "v4", gen: V4},
		{name: "v7", gen: V7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := tt.gen()
			if err != nil {
				t.Fatalf("generation error: %v", err)
			}
			res, err := Validate(id)
			if err != nil {
				t.Fatalf("Validate error: %v", err)
			}
			if res != "valid" {
				t.Fatalf("generated id %q reported %q, want valid", id, res)
			}
		})
	}
}

func TestValidateInvalid(t *testing.T) {
	res, err := Validate("not-a-uuid")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res != "invalid" {
		t.Fatalf("got %q, want invalid", res)
	}
}
