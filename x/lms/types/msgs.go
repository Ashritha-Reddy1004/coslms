package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeRegisterAdmin = "register-admin"
	TypeAddStudent    = "add-student"
	TypeApplyLeave    = "apply-leave"
	TypeAcceptLeave   = "accept-leave"
)

var (
	_ sdk.Msg = &AddStudentRequest{}
	_ sdk.Msg = &AcceptLeaveRequest{}
	_ sdk.Msg = &ApplyLeaveRequest{}
	_ sdk.Msg = &RegisterAdminRequest{}
)

// -----------------------------------ADD STUDENT----------------------------------------------------
func NewAddStudentRequest(admin string, students []*Student) *AddStudentRequest {
	return &AddStudentRequest{
		Admin:   admin,
		Student: students,
	}
}

func (msg AddStudentRequest) GetSignBytes() []byte {
	return []byte{}
}
func (msg AddStudentRequest) Route() string {
	return RouterKey

}

func (msg AddStudentRequest) Type() string {
	return TypeAddStudent
}

// GetSigners Implements Msg.
func (msg AddStudentRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{fromAddress}
}

func (msg AddStudentRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("From address invalid: %s", err)
	}

	// if _, err := sdk.AccAddressFromBech32("Welcome"); err != nil {
	// 	return sdkerrors.ErrInvalidAddress.Wrapf("To address invalid: %s", err)
	// }
	return nil
}

// --------------------------------------------ACCEPT LEAVE REQUEST-----------------------------------------------------
func NewAcceptLeaveRequest(admin string, student string, leaveid string, status LeaveStatus) *AcceptLeaveRequest {
	return &AcceptLeaveRequest{
		Admin:   admin,
		Address: student,
		LeaveId: leaveid,
		Status:  status,
	}
}

func (msg AcceptLeaveRequest) GetSignBytes() []byte {
	return []byte{}
}
func (msg ApplyLeaveRequest) Route() string {
	return RouterKey
}

func (msg ApplyLeaveRequest) Type() string {
	return TypeAcceptLeave
}

// GetSigners Implements Msg.
func (msg AcceptLeaveRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{fromAddress}
}

func (msg AcceptLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("From address invalid: %s", err)
	}

	return nil
}

// ----------------------------APPLY LEAVE REQUESTS-----------------------------------------------------------------------------
func NewApplyLeaveRequest(address string, reason string, from *time.Time, to *time.Time, leaveid string) *ApplyLeaveRequest {
	return &ApplyLeaveRequest{
		Address: address,
		Reason:  reason,
		From:    from,
		To:      to,
		LeaveId: leaveid,
	}
}
func (msg AcceptLeaveRequest) Route() string {
	return RouterKey
}

func (msg AcceptLeaveRequest) Type() string {
	return TypeApplyLeave
}

func (msg ApplyLeaveRequest) GetSignBytes() []byte {
	return []byte{}
}

// GetSigners Implements Msg.
func (msg ApplyLeaveRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Address)
	return []sdk.AccAddress{fromAddress}
}

func (msg ApplyLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("From address invalid: %s", err)
	}

	return nil
}

// -------------------------------------------REGISTER ADMIN------------------------------------------
func NewRegisterAdminRequest(address sdk.AccAddress, name string) *RegisterAdminRequest {
	return &RegisterAdminRequest{
		Address: address.String(),
		Name:    name,
	}
}

func (msg RegisterAdminRequest) Route() string {
	return RouterKey
}
func (msg RegisterAdminRequest) Type() string {
	return TypeRegisterAdmin
}
func (msg RegisterAdminRequest) GetSignBytes() []byte {
	return []byte{}
}

// GetSigners Implements Msg.
func (msg RegisterAdminRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32(msg.Address)
	return []sdk.AccAddress{fromAddress}
}

func (msg RegisterAdminRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("From address invalid: %s", err)
	}

	return nil
}
