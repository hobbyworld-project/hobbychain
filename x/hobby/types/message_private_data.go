package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPrivateData = "private_data"

var _ sdk.Msg = &MsgPrivateData{}

func NewMsgPrivateData(creator string, key string, value string) *MsgPrivateData {
  return &MsgPrivateData{
		Creator: creator,
    Key: key,
    Value: value,
	}
}

func (msg *MsgPrivateData) Route() string {
  return RouterKey
}

func (msg *MsgPrivateData) Type() string {
  return TypeMsgPrivateData
}

func (msg *MsgPrivateData) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgPrivateData) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgPrivateData) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

