package toml

import (
	"strings"
	"testing"
)

func TestToJSON(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{"simple", "a = 1", `"a": 1`, false},
		{"invalid", "a = =", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToJSON(tt.in)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ToJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !strings.Contains(got, tt.want) {
				t.Errorf("ToJSON() = %q, want contains %q", got, tt.want)
			}
		})
	}
}

func TestRoundTrip(t *testing.T) {
	in := `{"a":1,"b":"two"}`
	tomlOut, err := FromJSON(in)
	if err != nil {
		t.Fatalf("FromJSON() error = %v", err)
	}
	jsonOut, err := ToJSON(tomlOut)
	if err != nil {
		t.Fatalf("ToJSON() error = %v", err)
	}
	if !strings.Contains(jsonOut, `"a": 1`) {
		t.Errorf("round trip = %q, want contains %q", jsonOut, `"a": 1`)
	}
	if !strings.Contains(jsonOut, `"b": "two"`) {
		t.Errorf("round trip = %q, want contains %q", jsonOut, `"b": "two"`)
	}
}

func TestToYAML(t *testing.T) {
	got, err := ToYAML("a = 1")
	if err != nil {
		t.Fatalf("ToYAML() error = %v", err)
	}
	if !strings.Contains(got, "a: 1") {
		t.Errorf("ToYAML() = %q, want contains %q", got, "a: 1")
	}
}
