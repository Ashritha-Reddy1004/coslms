package keeper_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Ashritha-Reddy1004/coslms/x/lms/types"
	//"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"

	//"github.com/cosmos/cosmos-sdk/x/gov/keeper"
	"github.com/Ashritha-Reddy1004/coslms/x/lms/keeper"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"
)

type TestSuite struct {
	suite.Suite
	ctx           sdk.Context
	studentKeeper keeper.Keeper
	*assert.Assertions
}

func (s *TestSuite) SetupTest() {

	fmt.Println("Setup in progress")

	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	encCfg := simapp.MakeTestEncodingConfig()
	lmsKey := sdk.NewKVStoreKey(types.StoreKey)
	ctx := testutil.DefaultContext(lmsKey, sdk.NewTransientStoreKey("transient_test"))
	keeper := keeper.NewKeeper(lmsKey, encCfg.Codec)

	cms.MountStoreWithDB(lmsKey, storetypes.StoreTypeIAVL, db)
	s.Require().NoError(cms.LoadLatestVersion())

	s.studentKeeper = keeper
	s.ctx = ctx

}

func (s *TestSuite) TestRegisterAdmin() {
	type RegisterAdminTest struct {
		arg1     types.RegisterAdminRequest
		expected string
	}
	var RegisterAdminTests = []RegisterAdminTest{
		{
			arg1: types.RegisterAdminRequest{
				Name:    "Ashritha",
				Address: sdk.AccAddress("a1b2c3").String(),
			},
			expected: "Admin Registered Successfully",
		},
		{
			arg1: types.RegisterAdminRequest{
				Name:    "Ankith",
				Address: sdk.AccAddress("abcd").String(),
			},
			expected: "Admin Registered Successfully",
		},
		{
			arg1: types.RegisterAdminRequest{
				Name:    "Ram",
				Address: "",
			},
			expected: "Address field cannot be null",
		},
		{
			arg1: types.RegisterAdminRequest{
				Name:    "",
				Address: sdk.AccAddress("5678").String(),
			},
			expected: "Name field cannot be null",
		},
	}
	// require := s.Require()
	for _, test := range RegisterAdminTests {
		if output := s.studentKeeper.RegisterAdmins(s.ctx, &test.arg1); output != test.expected {
			// require.Equal(test.expected, output)
			fmt.Println("FAILED")
		}
	}

}

func (s *TestSuite) TestApplyLeave() {
	type applyLeaveTest struct {
		arg1     types.ApplyLeaveRequest
		expected string
	}
	dateString := "2006-Jan-02"
	fromDate, _ := time.Parse(dateString, "2023-Feb-22")
	toDate, _ := time.Parse(dateString, "2023-Feb-26")
	var applyLeaveTests = []applyLeaveTest{
		{
			arg1: types.ApplyLeaveRequest{
				Address: sdk.AccAddress("lms1").String(),
				Reason:  "I am feeling sick",
				From:    &fromDate,
				To:      &toDate,
				LeaveId: "1",
			},
			expected: "Leave Applied Successfully",
		},
		{
			arg1: types.ApplyLeaveRequest{
				Address: sdk.AccAddress("lms2").String(),
				Reason:  "",
				From:    &fromDate,
				To:      &toDate,
				LeaveId: "2",
			},
			expected: "Reason field cannot be null",
		},
		{
			arg1: types.ApplyLeaveRequest{
				Address: "",
				Reason:  "I am feeling sick",
				From:    &fromDate,
				To:      &toDate,
				LeaveId: "3",
			},
			expected: "Address field cannot be null",
		},
		{
			arg1: types.ApplyLeaveRequest{
				Address: sdk.AccAddress("lms2").String(),
				Reason:  "I am feeling sick",
				From:    nil,
				To:      &toDate,
				LeaveId: "4",
			},
			expected: "From field cannot be null",
		},
		{
			arg1: types.ApplyLeaveRequest{
				Address: sdk.AccAddress("lms2").String(),
				Reason:  "I am feeling sick",
				From:    &fromDate,
				To:      nil,
				LeaveId: "5",
			},
			expected: "To field cannot be null",
		},
		{
			arg1: types.ApplyLeaveRequest{
				Address: sdk.AccAddress("lms2").String(),
				Reason:  "I am feeling sick",
				From:    &fromDate,
				To:      &toDate,
				LeaveId: "",
			},
			expected: "Id cannot be empty",
		},
	}
	// require := s.Require()
	for _, test := range applyLeaveTests {
		if output := s.studentKeeper.ApplyLeaves(s.ctx, &test.arg1); output != test.expected {
			fmt.Println(test.expected, output)
			fmt.Println("FAILED")

		}
	}
}

func (s *TestSuite) TestAcceptLeaves() {
	type AcceptLeavesTest struct {
		req      types.AcceptLeaveRequest
		expected string
	}
	var acceptLeavesTest = []AcceptLeavesTest{
		{
			req: types.AcceptLeaveRequest{

				Admin:   sdk.AccAddress("abcdef").String(),
				LeaveId: "1",
				Status:  1,
			},
			expected: "Leave Accepted",
		},
		{
			req: types.AcceptLeaveRequest{
				Admin:   sdk.AccAddress("123658").String(),
				LeaveId: "2",
				Status:  0,
			},
			expected: "Status field cannot be null",
		},
		{
			req: types.AcceptLeaveRequest{
				Admin:   "",
				LeaveId: "3",
				Status:  1,
			},
			expected: "Admin field cannot be null",
		},
		{
			req: types.AcceptLeaveRequest{
				Admin:   sdk.AccAddress("aq1bgd546").String(),
				LeaveId: "",
				Status:  2,
			},
			expected: "LeaveId field cannot be null",
		},
	}

	for _, test := range acceptLeavesTest {
		if output := s.studentKeeper.AcceptLeaves(s.ctx, &test.req); output != test.expected {
			fmt.Println("FAILED")
		}
	}

}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestAddStudents() {
	type AddStudentsTest struct {
		req      types.AddStudentRequest
		expected string
	}
	var students = []AddStudentsTest{
		{
			req: types.AddStudentRequest{
				Admin: "",
			},
			expected: "Admin field cannot be null",
		},
	}
	for _, test1 := range students {
		if output := s.studentKeeper.AddStudents(s.ctx, &test1.req); output != test1.expected {
			fmt.Println("FAILED")
		}
	}
}

func (s *TestSuite) TestGetStudent() {
	s.TestAddStudents()
	s.studentKeeper.GetStudent(s.ctx)
}

func (s *TestSuite) TestGetLeavesRequest() {
	s.TestApplyLeave()
	s.studentKeeper.GetLeaveRequests(s.ctx, types.GetLeaveRequestsRequest{})
}

func (s *TestSuite) TestGetApprovedLeaves() {
	s.TestAcceptLeaves()
	s.studentKeeper.GetApprovedLeaves(s.ctx, types.GetLeaveApprovedRequestsRequest{})

}
