package enter_room

import (
	gameModel "game-go/core/model/game"
	userModel "game-go/core/model/user"
)

type Output struct {
	User  *userModel.Output
	Games []*gameModel.Output
}
