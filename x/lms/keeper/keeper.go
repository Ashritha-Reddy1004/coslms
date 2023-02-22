package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

type keeper struct {
	storekey storetypes.StoreKey
	cdc      codec.BinaryCodec
}

func NewKeeper(
	storekey storetypes.StoreKey,
	cdc codec.BinaryCodec,
) keeper {
	return keeper{
		cdc:      cdc,
		storekey: storekey,
	}
}
