package game_ID

import "github.com/google/uuid"

/*
IDGen user uuid package to generate new UUID keys
These UUID will be used as gameIDs, to identify every single game match.
*/
func IDGen() uuid.UUID {
	gameID := uuid.New()
	return gameID
}
