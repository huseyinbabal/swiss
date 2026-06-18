package textcase

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []string
	}{
		{"camel", "HelloWorld", []string{"hello", "world"}},
		{"snake", "hello_world", []string{"hello", "world"}},
		{"kebab", "hello-world", []string{"hello", "world"}},
		{"space", "hello world", []string{"hello", "world"}},
		{"acronym", "HTTPServer", []string{"http", "server"}},
		{"empty", "", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tokenize(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tokenize(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestConversions(t *testing.T) {
	tests := []struct {
		name string
		fn   func(string) string
		in   string
		want string
	}{
		{"camel", ToCamel, "hello_world foo", "helloWorldFoo"},
		{"pascal", ToPascal, "hello-world", "HelloWorld"},
		{"snake", ToSnake, "HelloWorld", "hello_world"},
		{"kebab", ToKebab, "Hello World", "hello-world"},
		{"title", ToTitle, "hello_world", "Hello World"},
		{"upper", ToUpper, "Hello World", "HELLO WORLD"},
		{"lower", ToLower, "Hello World", "hello world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fn(tt.in); got != tt.want {
				t.Errorf("%s(%q) = %q, want %q", tt.name, tt.in, got, tt.want)
			}
		})
	}
}
