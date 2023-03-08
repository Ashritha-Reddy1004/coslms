package keeper

import (
	"fmt"
	"log"
	"strconv"

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
			std := types.AddStudentRequest{}
			k.cdc.Unmarshal(store.Get(types.StudentStoreKey(stud.Address)), &std)

		}
		return "Students Added Successfully"

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
		address := types.LeavesCounterKey(applyleavereq.Address)
		counter := store.Get(address)
		if counter == nil {
			store.Set(address, []byte("1"))
		} else {
			c, err := strconv.Atoi(string(counter))
			if err != nil {
				panic(err)
			}
			c = c + 1
			store.Set(address, []byte(fmt.Sprint(c)))
		}
		counter = store.Get(address)
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

// Function to GET STUDENTS
func (k Keeper) GetStudentsQuery(ctx sdk.Context, req *types.GetStudentsRequest) []*types.AddStudentRequest {
	store := ctx.KVStore(k.storekey)
	var students []*types.AddStudentRequest
	itr := store.Iterator(types.StudentKey, nil)
	for ; itr.Valid(); itr.Next() {
		var t types.AddStudentRequest
		k.cdc.Unmarshal(itr.Value(), &t)
		students = append(students, &t)

	}
	return students
}

// Function to GET ADMIN
func (k Keeper) GetAdminQuery(ctx sdk.Context, Address string) (req *types.RegisterAdminRequest, err error) {
	if _, err := sdk.AccAddressFromBech32(Address); err != nil {
		panic(err)
	}
	store := ctx.KVStore(k.storekey)
	admin := store.Get(types.AdminStoreKey(Address))
	if admin == nil {
		panic("Admin not found")
	}
	k.cdc.MustUnmarshal(admin, req)
	return req, err
}

// Function to GET LEAVES
func (k Keeper) GetLeaveRequestsQuery(ctx sdk.Context, req *types.GetLeaveRequestsRequest) []*types.ApplyLeaveRequest {
	store := ctx.KVStore(k.storekey)
	var leaves []*types.ApplyLeaveRequest
	itr := store.Iterator(types.ApplyLeavesKey, nil)
	for ; itr.Valid(); itr.Next() {
		var leave types.ApplyLeaveRequest
		k.cdc.Unmarshal(itr.Value(), &leave)
		leaves = append(leaves, &leave)
	}
	return leaves
}

// Function to GET APPROVED LEAVES
func (k Keeper) GetApprovedLeaves(ctx sdk.Context, req *types.GetLeaveApprovedRequestsRequest) []*types.AcceptLeaveRequest {
	store := ctx.KVStore(k.storekey)
	var leaves []*types.AcceptLeaveRequest
	itr := store.Iterator(types.AcceptLeavesKey, nil)
	for ; itr.Valid(); itr.Next() {
		var leave types.AcceptLeaveRequest
		k.cdc.Unmarshal(itr.Value(), &leave)
		leaves = append(leaves, &leave)
	}
	return leaves
}

// Function to GET STUDENT
func (k Keeper) GetStudentQuery(ctx sdk.Context, address string) (req *types.AddStudentRequest, err error) {
	store := ctx.KVStore(k.storekey)
	student := store.Get(types.StudentStoreKey(address))
	fmt.Println(student)
	if student == nil {
		panic("Student not found")
	}
	fmt.Println(student)
	k.cdc.MustUnmarshal(student, req)
	return req, err
}

// Function to GET LEAVE STATUS
func (k Keeper) GetLeaveStatusQuery(ctx sdk.Context, admin string, LeaveId string) (req *types.GetLeaveStatusResponse, err error) {
	store := ctx.KVStore(k.storekey)
	leave := store.Get(types.AcceptLeavesStoreKey(admin, LeaveId))
	if leave == nil {
		panic("Leave with mentioned Id is not found")

	}
	k.cdc.MustUnmarshal(leave, req)
	return req, err
}

// Function to GET ADMINS
func (k Keeper) GetAdminsQuery(ctx sdk.Context, req *types.GetAdminsRequest) []*types.RegisterAdminRequest {
	store := ctx.KVStore(k.storekey)
	var admins []*types.RegisterAdminRequest
	itr := store.Iterator(types.AdminKey, nil)
	for ; itr.Valid(); itr.Next() {
		var t types.RegisterAdminRequest
		k.cdc.Unmarshal(itr.Value(), &t)
		admins = append(admins, &t)

	}
	return admins

}
