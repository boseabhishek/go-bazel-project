package second

import (
	"testing"
)

func TestLearnBazel(t *testing.T) {
	want := "Hello Abhi, so are you ready to learn bazel?"
	got := LearnBazel("Abhi")

	if want != got {
		t.Errorf("Test failed, got: %s, want: %s", got, want)
	}
}
