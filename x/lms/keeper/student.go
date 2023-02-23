package keeper

import (
	"errors"

	"github.com/Ashritha-Reddy1004/coslms/x/lms/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k keeper) AddStudent(ctx sdk.Context, addstudent *types.AddStudentRequest) error {

	if addstudent.Name == "" {
		return errors.New("name cant be null")
	} else if addstudent.Address == "" {
		return errors.New("address cant be null")
	} else if addstudent.Id == "" {
		return errors.New("Id cant be null")
	} else {
		store := ctx.KVStore(k.storekey)
		// key := types.StudentKey

		k.cdc.MustMarshal(addstudent)

		marshalAddStudent, err := k.cdc.Marshal(addstudent)
		if err != nil {
			panic(err)
		}
		store.Set(types.StudentStoreKey(addstudent.Id), marshalAddStudent)
	}
	return nil
}

func (k keeper) RegisterAdmin(ctx sdk.Context, registeradminreq *types.RegisterAdminRequest) error {

	store := ctx.KVStore(k.storekey)
	// key := types.StudentKey

	//k.cdc.MustMarshal(registeradminreq)

	marshalAdmin, err := k.cdc.Marshal(registeradminreq)
	if err != nil {
		return err
	}
	//store.Set(types.StudentStoreKey((addstudent.Admin), marshalAddStudent))
	store.Set(types.AdminStoreKey(registeradminreq.Address), marshalAdmin)
	return nil
}

func (k keeper) ApplyLeaves(ctx sdk.Context, applyleavereq *types.ApplyLeaveRequest) error {

	if applyleavereq.Address == "" {
		return errors.New("Address cant be null")
	} else if applyleavereq.Reason == "" {
		return errors.New("Reason cant be null")
	} else if applyleavereq.from == nil {
		return errors.New("From date cant be null")
	} else if applyleavereq.to == nil {
		return errors.New("To date cant be null")
	} else {
		store := ctx.KVStore(k.storekey)
		// key := types.StudentKey

		//k.cdc.MustMarshal(applyleavereq)

		marshalApplyLeave, err := k.cdc.Marshal(applyleavereq)
		if err != nil {
			panic(err)
		}
		store.Set(types.LeavesStoreKey(applyleavereq.Address, applyleavereq.LeaveId), marshalApplyLeave)
	}
	return nil
}

func (k keeper) AcceptLeave(ctx sdk.Context, acceptleavereq *types.AcceptLeaveRequest) error {
	if acceptleavereq.Admin == "" {
		return errors.New("Admin cant be null")
	} else if acceptleavereq.LeaveId == "" {
		return errors.New("LeaveId cant be null")
	} else if acceptleavereq.Status == 0 {
		return errors.New("Status cant be null")
	} else {
		store := ctx.KVStore(k.storekey)
		// key := types.StudentKey

		//k.cdc.MustMarshal(applyleavereq)

		marshalAcceptLeave, err := k.cdc.Marshal(acceptleavereq)
		if err != nil {
			panic(err)
		}
		store.Set(types.LeavesStoreKey(acceptleavereq.Admin, acceptleavereq.LeaveId), marshalAcceptLeave)
	}
	return nil
}
