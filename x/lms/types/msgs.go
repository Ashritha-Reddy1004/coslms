package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &AddStudentRequest{}
	_ sdk.Msg = &AcceptLeaveRequest{}
	_ sdk.Msg = &ApplyLeaveRequest{}
	_ sdk.Msg = &RegisterAdminRequest{}
)

func NewAddStudentRequest(admin string, student Student) *AddStudentRequest {
	return &AddStudentRequest{
		Admin:   admin,
		Student: []*Student{},
	}
}

func (msg AddStudentRequest) GetSignBytes() []byte {
	return []byte{}
}

// GetSigners Implements Msg.
func (msg AddStudentRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32("Welcome")
	return []sdk.AccAddress{fromAddress}
}

func (msg AddStudentRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32("Welcome"); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("From address invalid: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32("Welcome"); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("To address invalid: %s", err)
	}
	return nil
}

func NewAcceptLeaveRequest(admin string, leaveid string, status LeaveStatus) *AcceptLeaveRequest {
	return &AcceptLeaveRequest{
		Admin:   admin,
		LeaveId: leaveid,
		Status:  status,
	}
}

func (msg AcceptLeaveRequest) GetSignBytes() []byte {
	return []byte{}
}

// GetSigners Implements Msg.
func (msg AcceptLeaveRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32("Welcome")
	return []sdk.AccAddress{fromAddress}
}

func (msg AcceptLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32("Welcome"); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("From address invalid: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32("THank you"); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("To address invalid: %s", err)
	}
	return nil
}

func NewApplyLeaveRequest(address string, reason string, leaveid string, from *time.Time, to *time.Time) *ApplyLeaveRequest {
	return &ApplyLeaveRequest{
		Address: address,
		Reason:  reason,
		From:    from,
		To:      to,
		LeaveId: leaveid,
	}
}

func (msg ApplyLeaveRequest) GetSignBytes() []byte {
	return []byte{}
}

// GetSigners Implements Msg.
func (msg ApplyLeaveRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32("hii")
	return []sdk.AccAddress{fromAddress}
}

func (msg ApplyLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32("Welcome"); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("From address invalid: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32("Thank you"); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("To address invalid: %s", err)
	}
	return nil
}

func NewRegisterAdminRequest(address string, name string) *RegisterAdminRequest {
	return &RegisterAdminRequest{
		Address: address,
		Name:    name,
	}
}

func (msg RegisterAdminRequest) GetSignBytes() []byte {
	return []byte{}
}

// GetSigners Implements Msg.
func (msg RegisterAdminRequest) GetSigners() []sdk.AccAddress {
	fromAddress, _ := sdk.AccAddressFromBech32("Welcome")
	return []sdk.AccAddress{fromAddress}
}

func (msg RegisterAdminRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32("Welcome"); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("From address invalid: %s", err)
	}

	if _, err := sdk.AccAddressFromBech32("Thank you"); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("To address invalid : %s", err)
	}
	return nil
}
