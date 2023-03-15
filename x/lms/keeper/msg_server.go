package keeper

import (
	"context"

	"coslms/x/lms/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = Keeper{}

// type msgServer struct {
// 	Keeper
// 	types.UnimplementedMsgServer
// }

// func NewMsgServerImpl(k Keeper) types.MsgServer {
// 	return &msgServer{
// 		Keeper: k,
// 	}
// }

// Add student method
func (k Keeper) StudentAdd(goCtx context.Context, addstudentreq *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.AddStudents(ctx, addstudentreq)
	return &types.AddStudentResponse{}, nil
}

// Method to apply leave
func (k Keeper) LeaveApply(goCtx context.Context, applyleavereq *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.ApplyLeaves(ctx, applyleavereq)
	return &types.ApplyLeaveResponse{}, nil

}

// Method to accept leave
func (k Keeper) LeaveAccept(goCtx context.Context, acceptleavereq *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.AcceptLeaves(ctx, acceptleavereq)
	return &types.AcceptLeaveResponse{}, nil
}

// Method to register admin
func (k Keeper) AdminRegister(goCtx context.Context, registeradminreq *types.RegisterAdminRequest) (*types.RegisterAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.RegisterAdmins(ctx, registeradminreq)
	return &types.RegisterAdminResponse{}, nil
}
