package keeper

import (
	"context"

	"coslms/x/lms/types"
)

// type queryServer struct {
// 	Keeper
// 	types.UnimplementedQueryServer
// }

// func NewQueryServerImpl(k Keeper) types.QueryServer {
// 	return &queryServer{
// 		Keeper: k,
// 	}
// }

var _ types.QueryServer = Keeper{}

//var _ types.QueryServer = queryServer{}

func (k Keeper) GetStudents(goCtx context.Context, req *types.GetStudentsRequest) (*types.GetStudentsResponse, error) {
	//ctx := sdk.UnwrapSDKContext(goCtx)
	// k.GetStudent(ctx, req)
	return &types.GetStudentsResponse{}, nil
}

func (k Keeper) GetLeaveRequests(context.Context, *types.GetLeaveRequestsRequest) (*types.GetLeaveRequestsResponse, error) {
	return &types.GetLeaveRequestsResponse{}, nil
}

func (k Keeper) GetLeaveApprovedRequests(context.Context, *types.GetLeaveApprovedRequestsRequest) (*types.GetLeaveApprovedRequestsResponse, error) {
	return &types.GetLeaveApprovedRequestsResponse{}, nil
}
