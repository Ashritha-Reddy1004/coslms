package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&AddStudentRequest{}, "coslms/AddStudent", nil)
	cdc.RegisterConcrete(&RegisterAdminRequest{}, "coslms/RegisterAdmin", nil)
	cdc.RegisterConcrete(&ApplyLeaveRequest{}, "coslms/ApplyLeave", nil)
	cdc.RegisterConcrete(&AcceptLeaveRequest{}, "AcceptLeave", nil)
}

// Registering Interfaces
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&AddStudentRequest{},
		&RegisterAdminRequest{},
		&ApplyLeaveRequest{},
		&AcceptLeaveRequest{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

}
