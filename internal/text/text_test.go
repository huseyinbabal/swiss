package text

import "testing"

func TestCount(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "empty",
			in:   "",
			want: "Characters:   0\nBytes:        0\nWords:        0\nLines:        0\nReading time: 0 min",
		},
		{
			name: "single line no newline",
			in:   "hello world",
			want: "Characters:   11\nBytes:        11\nWords:        2\nLines:        1\nReading time: 1 min",
		},
		{
			name: "two lines trailing newline",
			in:   "a b\nc d\n",
			want: "Characters:   8\nBytes:        8\nWords:        4\nLines:        2\nReading time: 1 min",
		},
		{
			name: "multibyte",
			in:   "héllo",
			want: "Characters:   5\nBytes:        6\nWords:        1\nLines:        1\nReading time: 1 min",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Count(tt.in); got != tt.want {
				t.Errorf("Count(%q) =\n%q\nwant\n%q", tt.in, got, tt.want)
			}
		})
	}
}
