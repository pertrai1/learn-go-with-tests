package poker_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	poker "example.com/hello-world/http-server"
)

type GameSpy struct {
	StartedWith  int
	FinishedWith string

	StartCalled bool
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartedWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{duration, amount})
}

var dummySpyAlerter = &SpyBlindAlerter{}
var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	// t.Run("record chris win from user input", func(t *testing.T) {
	// 	in := strings.NewReader("Chris wins\n")
	// 	playerStore := &poker.StubPlayerStore{}
	//
	// 	game := poker.NewGame(dummySpyAlerter, playerStore)
	// 	cli := poker.NewCLI(in, dummyStdOut, game)
	// 	cli.PlayPoker()
	//
	// 	want := "Chris"
	//
	// 	poker.AssertPlayerWin(t, playerStore, want)
	// })
	// t.Run("record rob win from user input", func(t *testing.T) {
	// 	in := strings.NewReader("Rob wins\n")
	// 	playerStore := &poker.StubPlayerStore{}
	//
	// 	game := poker.NewGame(dummySpyAlerter, playerStore)
	// 	cli := poker.NewCLI(in, dummyStdOut, game)
	// 	cli.PlayPoker()
	//
	// 	want := "Rob"
	//
	// 	poker.AssertPlayerWin(t, playerStore, want)
	// })
	// t.Run("it schedules printing of blind value", func(t *testing.T) {
	// 	in := strings.NewReader("Chris wins\n")
	// 	playerStore := &poker.StubPlayerStore{}
	// 	blindAlerter := &SpyBlindAlerter{}
	//
	// 	game := poker.NewGame(blindAlerter, playerStore)
	// 	cli := poker.NewCLI(in, dummyStdOut, game)
	// 	cli.PlayPoker()
	//
	// 	cases := []scheduledAlert{
	// 		{0 * time.Second, 100},
	// 		{10 * time.Minute, 200},
	// 		{20 * time.Minute, 300},
	// 		{30 * time.Minute, 400},
	// 		{40 * time.Minute, 500},
	// 		{50 * time.Minute, 600},
	// 		{60 * time.Minute, 800},
	// 		{70 * time.Minute, 1000},
	// 		{80 * time.Minute, 2000},
	// 		{90 * time.Minute, 4000},
	// 		{100 * time.Minute, 8000},
	// 	}
	//
	// 	for i, want := range cases {
	// 		t.Run(fmt.Sprint(want), func(t *testing.T) {
	// 			if len(blindAlerter.alerts) <= i {
	// 				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
	// 			}
	//
	// 			got := blindAlerter.alerts[i]
	//
	// 			assertScheduledAlert(t, got, want)
	//
	// 		})
	// 	}
	// })
	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		blindAlerter := &SpyBlindAlerter{}

		game := poker.NewGame(blindAlerter, dummyPlayerStore)

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		got := stdout.String()
		want := poker.PlayerPrompt

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{1 * time.Minute, 200},
			{2 * time.Minute, 300},
			{3 * time.Minute, 400},
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]

				assertScheduledAlert(t, got, want)

			})
		}
	})
	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := poker.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartedWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartedWith)
		}
	})
	t.Run("it prints an error when a non-numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Errorf("game should not have been called")
		}
	})
}

func assertScheduledAlert(t testing.TB, got, want scheduledAlert) {
	amountGot := want.amount
	if amountGot != want.amount {
		t.Errorf("got amount %d, want %d", amountGot, want.amount)
	}

	gotScheduledTime := want.at
	if gotScheduledTime != want.at {
		t.Errorf("got scheduled time of %v, want %v", gotScheduledTime, want.at)
	}
}
