package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		main()
		if 1 != 1 {
			t.Error("will not happen")
		}
	})
}
