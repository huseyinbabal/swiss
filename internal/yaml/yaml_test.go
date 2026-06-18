package yaml

import "testing"

func TestToXML(t *testing.T) {
	got, err := ToXML("name: ada\nage: 36")
	if err != nil {
		t.Fatal(err)
	}
	want := "<root>\n    <age>36</age>\n    <name>ada</name>\n</root>"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestToCSV(t *testing.T) {
	got, err := ToCSV("- name: ada\n  age: 36\n- name: bob")
	if err != nil {
		t.Fatal(err)
	}
	want := "age,name\n36,ada\n,bob"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
