package keeper

import (
	"fmt"
	"log"

	// "x/lms/types"

	"coslms/x/lms/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Function to ADD STUDENT
func (k Keeper) AddStudents(ctx sdk.Context, addstudentreq *types.AddStudentRequest) string {

	if addstudentreq.Admin == "" {
		return ("Admin field cannot be null")
	} else {
		students := addstudentreq.Student
		store := ctx.KVStore(k.storekey)
		for _, stud := range students {
			marshalAddStudents, err := k.cdc.Marshal(stud)
			if err != nil {
				log.Fatal(err)
			}
			store.Set(types.StudentStoreKey(stud.Address), marshalAddStudents)
			//k.GetStudent(ctx, sdk.AccAddress("lms1").String())
		}
		return ""

	}
}

//Function to REGISTER ADMIN

func (k Keeper) RegisterAdmins(ctx sdk.Context, registeradminreq *types.RegisterAdminRequest) string {
	if registeradminreq.Address == "" {
		return ("Address field cannot be null")
	} else if registeradminreq.Name == "" {
		return ("Name field cannot be null")
	} else {

		store := ctx.KVStore(k.storekey)
		//key := types.StudentKey
		k.cdc.MustMarshal(registeradminreq)

		marshalAdmin, err := k.cdc.Marshal(registeradminreq)

		if err != nil {
			return "invalid"
		}
		store.Set(types.AdminStoreKey(registeradminreq.Address), marshalAdmin)
		return "Admin Registered Successfully"
	}
}

// Function to APPLY LEAVES
func (k Keeper) ApplyLeaves(ctx sdk.Context, applyleavereq *types.ApplyLeaveRequest) string {

	if applyleavereq.Address == "" {
		return ("Address field cannot be null")
	} else if applyleavereq.Reason == "" {
		return ("Reason field cannot be null")
	} else if applyleavereq.LeaveId == "" {
		return ("Id cannot be empty")
	} else if applyleavereq.From == nil {
		return ("From field cannot be null")
	} else if applyleavereq.To == nil {
		return ("To field cannot be null")
	} else {
		store := ctx.KVStore(k.storekey)
		//key := types.StudentKey

		k.cdc.MustMarshal(applyleavereq)

		marshalApplyLeave, err := k.cdc.Marshal(applyleavereq)
		if err != nil {
			panic(err)
		}
		store.Set(types.ApplyLeavesStoreKey(applyleavereq.LeaveId, applyleavereq.Address), marshalApplyLeave)
	}
	return "Leave Applied Successfully"
}

// Function to ACCEPT LEAVES
func (k Keeper) AcceptLeaves(ctx sdk.Context, acceptleavereq *types.AcceptLeaveRequest) string {
	if acceptleavereq.Admin == "" {
		return ("Admin field cannot be null")
	} else if acceptleavereq.LeaveId == "" {
		return ("LeaveId field cannot be null")
	} else if acceptleavereq.Status == 0 {
		return ("Status field cannot be null")
	} else {
		store := ctx.KVStore(k.storekey)
		//key := types.StudentKey
		k.cdc.MustMarshal(acceptleavereq)

		marshalAcceptLeave, err := k.cdc.Marshal(acceptleavereq)
		if err != nil {
			panic(err)
		}
		store.Set(types.AcceptLeavesStoreKey(acceptleavereq.Admin, acceptleavereq.LeaveId), marshalAcceptLeave)
	}
	return "Leave Accepted"
}

// Function to GET STUDENT
func (k Keeper) GetStudent(ctx sdk.Context) {
	store := ctx.KVStore(k.storekey)
	var t types.Student
	itr := store.Iterator(types.StudentKey, nil)
	for ; itr.Valid(); itr.Next() {
		k.cdc.Unmarshal(itr.Value(), &t)
		fmt.Println(t)

	}

}

// Function to GET ADMIN
func (k Keeper) GetAdmin(ctx sdk.Context, Address string) []byte {
	if _, err := sdk.AccAddressFromBech32(Address); err != nil {
		panic(err)
	}
	store := ctx.KVStore(k.storekey)
	return store.Get(types.AdminStoreKey(Address))
}

func (k Keeper) GetLeaveRequests(ctx sdk.Context, applytleave types.GetLeaveRequestsRequest) {
	store := ctx.KVStore(k.storekey)
	var t types.ApplyLeaveRequest
	itr := store.Iterator(types.ApplyLeavesKey, nil)
	for ; itr.Valid(); itr.Next() {
		k.cdc.Unmarshal(itr.Value(), &t)
		fmt.Println(t)
	}
}

func (k Keeper) GetApprovedLeaves(ctx sdk.Context, approveleaves types.GetLeaveApprovedRequestsRequest) {
	store := ctx.KVStore(k.storekey)
	var t types.AcceptLeaveRequest
	itr := store.Iterator(types.AcceptLeavesKey, nil)
	for ; itr.Valid(); itr.Next() {
		k.cdc.Unmarshal(itr.Value(), &t)
		fmt.Println(t)
	}
}
