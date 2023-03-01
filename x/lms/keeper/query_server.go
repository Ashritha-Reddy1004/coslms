package keeper

import (
	"context"

	"coslms/x/lms/types"
)

var _ types.QueryServer = Keeper{}

// Method to get students
func (k Keeper) GetStudents(goCtx context.Context, req *types.GetStudentsRequest) (*types.GetStudentsResponse, error) {
	return &types.GetStudentsResponse{}, nil
}

// Method to get applied leaves
func (k Keeper) GetLeaveRequests(context.Context, *types.GetLeaveRequestsRequest) (*types.GetLeaveRequestsResponse, error) {
	return &types.GetLeaveRequestsResponse{}, nil
}

// Method to get approved leaves
func (k Keeper) GetLeaveApprovedRequests(context.Context, *types.GetLeaveApprovedRequestsRequest) (*types.GetLeaveApprovedRequestsResponse, error) {
	return &types.GetLeaveApprovedRequestsResponse{}, nil
}
