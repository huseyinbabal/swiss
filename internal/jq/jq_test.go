package jq

import "testing"

func TestQuery(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		value   string
		want    string
		wantErr bool
	}{
		{"array index", "a.b[1]", `{"a":{"b":[10,20]}}`, "20", false},
		{"leading dot", ".a", `{"a":5}`, "5", false},
		{"nested", "items[2]", `{"items":[1,2,3]}`, "3", false},
		{"missing key", "a.c", `{"a":{"b":1}}`, "", true},
		{"out of range", "a[5]", `{"a":[1,2]}`, "", true},
		{"type mismatch", "a.b", `{"a":[1,2]}`, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Query(tt.path, tt.value)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Query() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("Query() = %q, want %q", got, tt.want)
			}
		})
	}
}
