package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

type Keeper struct {
	storekey storetypes.StoreKey
	cdc      codec.BinaryCodec
}

func NewKeeper(storekey storetypes.StoreKey, cdc codec.BinaryCodec,
) Keeper {
	return Keeper{
		cdc:      cdc,
		storekey: storekey,
	}
}
