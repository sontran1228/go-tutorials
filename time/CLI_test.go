package poker_test

import (
	"bytes"
	"fmt"
	poker "go-tutorials/time"
	"strings"
	"testing"
	"time"
)

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

var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		blindAlerter := &SpyBlindAlerter{}

		game := poker.NewGame(blindAlerter, dummyPlayerStore)

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		got := stdout.String()

		if got != poker.PlayerPrompt {
			t.Errorf("got '%s', want '%s'", got, poker.PlayerPrompt)
		}

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
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

	// t.Run("record chris win from user input", func(t *testing.T) {
	// 	in := strings.NewReader("Chris wins\n")
	// 	playerStore := &poker.StubPlayerStore{}

	// 	cli := poker.NewCLI(playerStore, in, dummyBlindAlerter)
	// 	cli.PlayPoker()

	// 	poker.AssertPlayerWin(t, playerStore, "Chris")
	// })

	// t.Run("record cleo win from user input", func(t *testing.T) {
	// 	in := strings.NewReader("Cleo wins\n")
	// 	playerStore := &poker.StubPlayerStore{}

	// 	cli := poker.NewCLI(playerStore, in, dummyBlindAlerter)
	// 	cli.PlayPoker()

	// 	poker.AssertPlayerWin(t, playerStore, "Cleo")
	// })
}

func assertScheduledAlert(t *testing.T, got, want scheduledAlert) {
	amountGot := got.amount
	if amountGot != want.amount {
		t.Errorf("got amount %d, want %d", amountGot, want.amount)
	}

	gotScheduledTime := got.at
	if gotScheduledTime != want.at {
		t.Errorf("got scheduled time of %v, want %v", gotScheduledTime, want.at)
	}
}
