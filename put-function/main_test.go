package main

import (
	"testing"
)

func TestHandler(t *testing.T) {
	t.Run("Mock test", func(t *testing.T) {
		if true {
			t.Fatal("Error failed to trigger with an invalid request")
		}
	})
}
