package tcnksm

import (
	"fmt"
	"testing"
	"time"
)

func TestAge(t *testing.T) {
	t.Cleanup(func() {
		now = time.Now
	})
	now = func() time.Time {
		return time.Date(2018, time.July, 11, 0, 0, 0, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
	}
	if got, want := Age(), 30; got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestName(t *testing.T) {
	if got, want := Name(), "Taichi Nakashima"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}
}

func ExampleName() {
	fmt.Println(Name())
	// Output: Taichi Nakashima
}
