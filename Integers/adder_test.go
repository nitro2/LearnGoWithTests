package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected %d, got %d", expected, sum)
	}

}

// The `go test -v` will execute this Example* func
// We MUST have `// Output: ` format to get the test passed
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
