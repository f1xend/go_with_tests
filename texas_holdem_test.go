package poker_test

import (
	"fmt"
	"testing"
	"time"

	poker "github.com/f1xend/go_with_tests"
)

func TestGame_Start(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(5)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100, To: nil},
			{At: 10 * time.Minute, Amount: 200, To: nil},
			{At: 20 * time.Minute, Amount: 300, To: nil},
			{At: 30 * time.Minute, Amount: 400, To: nil},
			{At: 40 * time.Minute, Amount: 500, To: nil},
			{At: 50 * time.Minute, Amount: 600, To: nil},
			{At: 60 * time.Minute, Amount: 800, To: nil},
			{At: 70 * time.Minute, Amount: 1000, To: nil},
			{At: 80 * time.Minute, Amount: 2000, To: nil},
			{At: 90 * time.Minute, Amount: 4000, To: nil},
			{At: 100 * time.Minute, Amount: 8000, To: nil},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(7)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Second, Amount: 100, To: nil},
			{At: 12 * time.Minute, Amount: 200, To: nil},
			{At: 24 * time.Minute, Amount: 300, To: nil},
			{At: 36 * time.Minute, Amount: 400, To: nil},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})

}

func TestGame_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	game := poker.NewTexasHoldem(dummyBlindAlerter, store)
	winner := "Ruth"

	game.Finish(winner)
	poker.AssertPlayerWin(t, store, winner)
}

func checkSchedulingCases(cases []poker.ScheduledAlert, t *testing.T, blindAlerter *poker.SpyBlindAlerter) {
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {

			if len(blindAlerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
			}

			got := blindAlerter.Alerts[i]
			assertScheduledAlert(t, got, want)
		})
	}
}
