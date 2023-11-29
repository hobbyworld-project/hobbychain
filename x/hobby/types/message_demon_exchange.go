package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDemonExchange = "demon_exchange"

var _ sdk.Msg = &MsgDemonExchange{}

func NewMsgDemonExchange(creator string, amount string) *MsgDemonExchange {
	return &MsgDemonExchange{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgDemonExchange) Route() string {
	return RouterKey
}

func (msg *MsgDemonExchange) Type() string {
	return TypeMsgDemonExchange
}

func (msg *MsgDemonExchange) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDemonExchange) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDemonExchange) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
