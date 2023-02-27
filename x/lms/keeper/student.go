package keeper

import (
	"log"

	//"x/lms/types"

	"github.com/Ashritha-Reddy1004/coslms/x/lms/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AddStudents(ctx sdk.Context, addstudentreq *types.AddStudentRequest) error {

	// if addstudentreq.Admin == "" {
	// 	return errors.New("Admin
	// field cannot be null")
	// } else if addstudentreq.Name == "" {
	// 	return errors.New("Address field cannot be null")
	// } else if addstudentreq.Id == "" {
	// 	return errors.New("Id field cannot be null")
	// } else {}
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
	return nil

}

func (k Keeper) RegisterAdmins(ctx sdk.Context, registeradminreq *types.RegisterAdminRequest) string {
	if registeradminreq.Address == "" {
		return ("Address field cannot be null")
	} else if registeradminreq.Name == "" {
		return ("Name field cannot be null")
	} else {

		store := ctx.KVStore(k.storekey)
		//key := types.StudentKey
		//k.cdc.MustMarshal(registeradminreq)

		marshalAdmin, err := k.cdc.Marshal(registeradminreq)

		if err != nil {
			return "invalid"
		}
		store.Set(types.AdminStoreKey(registeradminreq.Address), marshalAdmin)
		return "Admin Registered Successfully"
	}
}

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
