package csv

import "testing"

func TestToJSON(t *testing.T) {
	got, err := ToJSON("name,age\nada,36\nbob,40")
	if err != nil {
		t.Fatal(err)
	}
	want := `[{"age":"36","name":"ada"},{"age":"40","name":"bob"}]`
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestToJSONEmpty(t *testing.T) {
	got, err := ToJSON("")
	if err != nil {
		t.Fatal(err)
	}
	if got != "[]" {
		t.Errorf("got %q", got)
	}
}

func TestToJSONQuoted(t *testing.T) {
	got, err := ToJSON("v\n\"a,b\"")
	if err != nil {
		t.Fatal(err)
	}
	if got != `[{"v":"a,b"}]` {
		t.Errorf("got %q", got)
	}
}

func TestEscapeUnescapeRoundTrip(t *testing.T) {
	raw := `a,b "c"`
	escaped, err := Escape(raw)
	if err != nil {
		t.Fatal(err)
	}
	if escaped != `"a,b ""c"""` {
		t.Errorf("escape got %q", escaped)
	}
	back, err := Unescape(escaped)
	if err != nil {
		t.Fatal(err)
	}
	if back != raw {
		t.Errorf("round trip got %q want %q", back, raw)
	}
}

func TestToXML(t *testing.T) {
	// CSV decodes to an array of row objects, so each row is wrapped in <item>.
	got, err := ToXML("a,b\n1,2")
	if err != nil {
		t.Fatal(err)
	}
	want := "<root>\n    <item>\n        <a>1</a>\n        <b>2</b>\n    </item>\n</root>"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
