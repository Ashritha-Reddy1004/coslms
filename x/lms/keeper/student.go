package keeper

import (
	"errors"
	"log"

	//"x/lms/types"

	"github.com/Ashritha-Reddy1004/coslms/x/lms/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AddStudents(ctx sdk.Context, addstudentreq *types.AddStudentRequest) error {

	// if addstudentreq.Admin == "" {
	// 	return errors.New("Admin field cannot be null")
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

func (k Keeper) RegisterAdmins(ctx sdk.Context, registeradminreq *types.RegisterAdminRequest) error {
	if registeradminreq.Address == "" {
		return errors.New("Address field cannot be null")
	} else if registeradminreq.Name == "" {
		return errors.New("Name field cannot be null")
	} else {

		store := ctx.KVStore(k.storekey)
		//key := types.StudentKey
		//k.cdc.MustMarshal(registeradminreq)

		marshalAdmin, err := k.cdc.Marshal(registeradminreq)

		if err != nil {
			return err
		}
		store.Set(types.AdminStoreKey(registeradminreq.Address), marshalAdmin)
		return nil
	}
}

func (k Keeper) ApplyLeaves(ctx sdk.Context, applyleavereq *types.ApplyLeaveRequest) error {

	if applyleavereq.Address == "" {
		return errors.New("Address field cannot be null")
	} else if applyleavereq.Reason == "" {
		return errors.New("Reason field cannot be null")
	} else if applyleavereq.LeaveId == "" {
		return errors.New("Id cannot be empty")
	} else if applyleavereq.From == nil {
		return errors.New("From field cannot be null")
	} else if applyleavereq.To == nil {
		return errors.New("To cannot be null")
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
	return nil
}

func (k Keeper) AcceptLeaves(ctx sdk.Context, acceptleavereq *types.AcceptLeaveRequest) error {
	if acceptleavereq.Admin == "" {
		return errors.New("Admin field cannot be null")
	} else if acceptleavereq.LeaveId == "" {
		return errors.New("LeaveId field cannot be null")
	} else if acceptleavereq.Status == 0 {
		return errors.New("Status field cannot be null")
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
	return nil
}
