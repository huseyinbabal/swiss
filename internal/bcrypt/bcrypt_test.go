package bcrypt

import "testing"

func TestVerify(t *testing.T) {
	hash, err := Hash("s3cret")
	if err != nil {
		t.Fatalf("Hash returned error: %v", err)
	}

	tests := []struct {
		name     string
		hash     string
		password string
		want     string
		wantErr  bool
	}{
		{name: "match", hash: hash, password: "s3cret", want: "match"},
		{name: "no match", hash: hash, password: "wrong", want: "no match"},
		{name: "malformed hash", hash: "not-a-hash", password: "s3cret", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Verify(tt.hash, tt.password)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}
