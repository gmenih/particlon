package particlon

type GameState int8

const (
	STATE_INIT GameState = iota
	STATE_PLAY
	STATE_PAUSE
	STATE_STOP
)
