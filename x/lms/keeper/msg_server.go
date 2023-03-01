package keeper

import (
	"context"

	"coslms/x/lms/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
	types.UnimplementedMsgServer
}

func (k msgServer) AddStudent(goCtx context.Context, addstudentreq *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.AddStudents(ctx, addstudentreq)
	return &types.AddStudentResponse{}, nil
}
func (k msgServer) ApplyLeave(goCtx context.Context, applyleavereq *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.ApplyLeaves(ctx, applyleavereq)
	return &types.ApplyLeaveResponse{}, nil

}
func (k msgServer) AcceptLeave(goCtx context.Context, acceptleavereq *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.AcceptLeaves(ctx, acceptleavereq)
	return &types.AcceptLeaveResponse{}, nil
}
func (k msgServer) RegisterAdmin(goCtx context.Context, registeradminreq *types.RegisterAdminRequest) (*types.RegisterAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.RegisterAdmins(ctx, registeradminreq)
	return &types.RegisterAdminResponse{}, nil
}
