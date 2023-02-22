package keeper

import (
	"context"
<<<<<<< HEAD
=======
	//"x/lms/types"
>>>>>>> f206accc8a09d55bec9142889f6ebed7a801eccc

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
	if _, err := k.AddStudent(ctx, req); err != nil {
		return nil, err
	}
	return &types.AddStudentResponse{}, nil
}
func (k msgServer) Applyleave(c context.Context, a *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	//stdnt := types.ApplyLeaveRequest{}
	//k.cdc.MustMarshal(&stdnt)
	return &types.AcceptLeaveResponse{}, nil

}
func (k msgServer) Acceptleave(c context.Context, a *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	return &types.AcceptLeaveResponse{}, nil
}
func (k msgServer) RegisterAdmin(c context.Context, a *types.RegisterAdminRequest) (*types.RegisterAdminResponse, error) {
	return &types.RegisterAdminResponse{}, nil
}
