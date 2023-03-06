package keeper

import (
	"context"

	"coslms/x/lms/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

// Method to get students
func (k Keeper) GetStudents(goCtx context.Context, req *types.GetStudentsRequest) (*types.GetStudentsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	res := k.GetStudent(ctx, req)
	result := types.GetStudentsResponse{
		Students: res,
	}
	return &result, nil
}

// Method to get applied leaves
func (k Keeper) GetLeaveRequests(goCtx context.Context, req *types.GetLeaveRequestsRequest) (*types.GetLeaveRequestsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	res := k.GetLeaveRequestsQuery(ctx, req)
	result := types.GetLeaveRequestsResponse{
		Leaverequests: res,
	}
	return &result, nil
}

// Method to get approved leaves
func (k Keeper) GetLeaveApprovedRequests(context.Context, *types.GetLeaveApprovedRequestsRequest) (*types.GetLeaveApprovedRequestsResponse, error) {
	return &types.GetLeaveApprovedRequestsResponse{}, nil
}
