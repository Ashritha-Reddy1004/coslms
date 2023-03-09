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
	res := k.GetStudentsQuery(ctx, req)
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
func (k Keeper) GetLeaveApprovedRequests(goCtx context.Context, req *types.GetLeaveApprovedRequestsRequest) (*types.GetLeaveApprovedRequestsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	res := k.GetApprovedLeaves(ctx, req)
	result := types.GetLeaveApprovedRequestsResponse{
		Leaverequests: res,
	}
	return &result, nil
}

// Function to get student
func (k Keeper) GetStudent(goCtx context.Context, req *types.GetStudentRequest) (*types.GetStudentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	res, _ := k.GetStudentQuery(ctx, req.Address)
	result := types.GetStudentResponse{
		Student: res,
	}
	return &result, nil
}

// Function to get admin
func (k Keeper) GetAdmin(goctx context.Context, req *types.GetAdminRequest) (*types.GetAdminResponse, error) {

	ctx := sdk.UnwrapSDKContext(goctx)
	res, _ := k.GetAdminQuery(ctx, req.Address)
	result := types.GetAdminResponse{
		Admin: res,
	}
	return &result, nil
}

// Function to get leave status
func (k Keeper) GetLeaveStatus(goctx context.Context, req *types.GetLeaveStatusRequest) (*types.GetLeaveStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goctx)
	res, _ := k.GetLeaveStatusQuery(ctx, req.Admin, req.LeaveId)

	return res, nil
}

// Function to get admins
func (k Keeper) GetAdmins(goCtx context.Context, req *types.GetAdminsRequest) (*types.GetAdminsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	res := k.GetAdminsQuery(ctx, req)
	result := types.GetAdminsResponse{
		Admin: res,
	}
	return &result, nil
}
