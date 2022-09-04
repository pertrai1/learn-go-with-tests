package poker_test

import (
	"strings"
	"testing"

	poker "example.com/hello-world/http-server"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		want := "Chris"

		poker.AssertPlayerWin(t, playerStore, want)
	})
	t.Run("record rob win from user input", func(t *testing.T) {
		in := strings.NewReader("Rob wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		want := "Rob"

		poker.AssertPlayerWin(t, playerStore, want)
	})
}
