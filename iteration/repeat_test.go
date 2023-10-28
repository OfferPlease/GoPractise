package iteration

import "testing"

func Repeat(character string) (repeat string) {
	for i := 0; i < 5; i++ {
		repeat += character
	}
	return
}
func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}
