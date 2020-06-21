package second

import (
	"testing"
)

func TestLearnBazel(t *testing.T) {
	got := LearnBazel("Abhi")
	want := "Hello Abhi, so are you ready to learn bazel?"

    if got != want {
       t.Errorf("Test failed, got: %s, want: %s", got, want)
    }
}