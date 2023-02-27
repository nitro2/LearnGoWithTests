package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	out_string := Repeat("a", 5)
	expected_string := "aaaaa"

	if out_string != expected_string {
		t.Errorf("expected %s, got %s", expected_string, out_string)
	}
}

func TestRepeatSpecial(t *testing.T) {
	out_string := Repeat("'", 5)
	expected_string := "'''''"

	if out_string != expected_string {
		t.Errorf("expected %s, got %s", expected_string, out_string)
	}
}

func TestRepeatEscape(t *testing.T) {
	out_string := Repeat("\\", 3)
	expected_string := "\\\\\\"

	if out_string != expected_string {
		t.Errorf("expected %s, got %s", expected_string, out_string)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 3))
	// Output: aaa

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
