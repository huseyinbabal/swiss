package xml

import "testing"

func TestEscapeUnescapeRoundTrip(t *testing.T) {
	raw := `<a> & "b" 'c'`
	escaped, _ := Escape(raw)
	if escaped != `&lt;a&gt; &amp; &quot;b&quot; &apos;c&apos;` {
		t.Errorf("escape got %q", escaped)
	}
	back, _ := Unescape(escaped)
	if back != raw {
		t.Errorf("round trip got %q want %q", back, raw)
	}
}

func TestEscapeAmpersandNotDoubled(t *testing.T) {
	got, _ := Escape("a < b")
	if got != "a &lt; b" {
		t.Errorf("got %q", got)
	}
}

func TestBeautify(t *testing.T) {
	got, err := Beautify(`<a><b>1</b><c>2</c></a>`)
	if err != nil {
		t.Fatal(err)
	}
	want := "<a>\n    <b>1</b>\n    <c>2</c>\n</a>"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestUglify(t *testing.T) {
	got, err := Uglify("<a>\n  <b>1</b>\n</a>")
	if err != nil {
		t.Fatal(err)
	}
	if got != "<a><b>1</b></a>" {
		t.Errorf("got %q", got)
	}
}

func TestBeautifyInvalid(t *testing.T) {
	if _, err := Beautify("<a><b></a>"); err == nil {
		t.Error("expected error for malformed xml")
	}
}
