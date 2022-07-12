package keeper

import (
	"github.com/alice/tictactoe/x/tictactoe/types"
)

var _ types.QueryServer = Keeper{}
