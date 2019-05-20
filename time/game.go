package poker

import "time"

// Game manages the state of a game
type Game struct {
	alerter BlindAlerter
	store   PlayerStore
}

// NewGame creates a game for playing poker
func NewGame(alerter BlindAlerter, store PlayerStore) *Game {
	return &Game{
		alerter: alerter,
		store:   store,
	}
}

// Start starts the game
func (p *Game) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

// Finish finishs the game
func (p *Game) Finish(winner string) {
	p.store.RecordWin(winner)
}
