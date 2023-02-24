package keeper

import (
	"context"

	"github.com/Ashritha-Reddy1004/coslms/x/lms/types"
)

type queryServer struct {
	Keeper
	types.UnimplementedQueryServer
}

//var _ types.QueryServer = queryServer{}

func (k queryServer) GetStudents(context.Context, *types.GetStudentsRequest) (*types.GetStudentsResponse, error) {
	return &types.GetStudentsResponse{}, nil
}

// func (k queryServer) GetStudent(context.Context, *types.GetStudentRequest) (*types.GetStudentResponse, error) {
// 	return &types.GetStudentResponse{}, nil
// }

func (k queryServer) GetLeaveRequests(context.Context, *types.GetLeaveRequestsRequest) (*types.GetLeaveRequestsResponse, error) {
	return &types.GetLeaveRequestsResponse{}, nil
}

func (k queryServer) GetLeaveApproves(context.Context, *types.GetLeaveApprovedRequestsRequest) (*types.GetLeaveApprovedRequestsResponse, error) {
	return &types.GetLeaveApprovedRequestsResponse{}, nil
}
