package keeper

import (
	"context"

	"github.com/Ashritha-Reddy1004/coslms/x/lms/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	keeper
	types.UnimplementedMsgServer
}

func (k msgServer) AddStudent(goCtx context.Context, req *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.AddStudents(ctx, req)
	return &types.AddStudentResponse{}, nil
}
func (k msgServer) ApplyLeave(goCtx context.Context, req *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.ApplyLeaves(ctx, req)
	return &types.ApplyLeaveResponse{}, nil

}
func (k msgServer) AcceptLeave(goCtx context.Context, req *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.AcceptLeaves(ctx, req)
	return &types.AcceptLeaveResponse{}, nil
}
func (k msgServer) RegisterAdmin(goCtx context.Context, req *types.RegisterAdminRequest) (*types.RegisterAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.RegisterAdmins(ctx, req)
	return &types.RegisterAdminResponse{}, nil
}
