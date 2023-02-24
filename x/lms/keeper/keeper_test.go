package keeper_test

import (
	"fmt"
	"sync"
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
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"
)

type TestSuite struct {
	suite.Suite
	ctx           sdk.Context
	studentKeeper keeper.keeper
	*assert.Assertions
	mutex   sync.RWMutex
	require *require.Assertions
	t       *testing.T
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
func (suite *TestSuite) T() *testing.T {
	suite.mutex.RLock()
	defer suite.mutex.RUnlock()
	return suite.t
}

func (suite *TestSuite) Set(t *testing.T) {
	suite.mutex.Lock()
	defer suite.mutex.Unlock()
	suite.t = t
	suite.Assertions = assert.New(t)
	suite.require = require.New(t)
}

func (suite *TestSuite) Require() *require.Assertions {
	suite.mutex.Lock()
	defer suite.mutex.Unlock()
	if suite.require == nil {
		suite.require = require.New(suite.T())
	}
	return suite.require
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
			expected: "Address cannot be empty",
		},
		{
			arg1: types.RegisterAdminRequest{
				Name:    "",
				Address: sdk.AccAddress("5678").String(),
			},
			expected: "Name cannot be empty",
		},
	}
	require := s.Require()
	for _, test := range RegisterAdminTests {
		if output := s.studentKeeper.RegisterAdmins(s.ctx, &test.arg1); fmt.Sprint(output) != test.expected {
			require.Equal(test.expected, output)
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
			},
			expected: "Leave Applied Successfully",
		},
		{
			arg1: types.ApplyLeaveRequest{
				Address: sdk.AccAddress("lms2").String(),
				Reason:  "I have to attend party",
				From:    &fromDate,
				To:      &toDate,
			},
			expected: "Leave Applied Successfully",
		},
	}
	require := s.Require()
	for _, test := range applyLeaveTests {
		if output := s.studentKeeper.ApplyLeaves(s.ctx, &test.arg1); fmt.Sprint(output) != test.expected {
			require.Equal(test.expected, output)
		}
	}

}
func (s *TestSuite) TestAcceptLeaves() {
	req := types.AcceptLeaveRequest{
		Admin:   sdk.AccAddress("abcdef").String(),
		LeaveId: "1",
		Status:  1,
	}
	res := s.studentKeeper.AcceptLeaves(s.ctx, &req)
	fmt.Println(res)
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestAddStudents() {
	students := []*types.Student{
		{
			Address: sdk.AccAddress("1234").String(),
			Name:    "Ashritha",
			Id:      "1",
		},
		{
			Address: sdk.AccAddress("5678").String(),
			Name:    "Ankith",
			Id:      "2",
		},
		{
			Address: sdk.AccAddress("hgtf").String(),
			Name:    "Sita",
			Id:      "3",
		},
	}
	req := types.AddStudentRequest{
		Admin: "Ashritha",
		Id:    "24",
	}
	res := s.studentKeeper.AddStudents(s.ctx, &req)
	fmt.Println(res)
}
