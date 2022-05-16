package types

import (
	"errors"
	fmt "fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"gopkg.in/yaml.v2"
)

const (
	TypeMsgCreateRewardProgramRequest = "create_reward_program"
)

// Compile time interface checks.
var _ sdk.Msg = &MsgCreateRewardProgramRequest{}

// MsgCreateRewardProgramRequest creates a new create reward program request
func NewMsgCreateRewardProgramRequest(
	title string,
	description string,
	distributeFromAddress string,
	coin sdk.Coin,
	maxRewardByAddress sdk.Coin,
	programStartTime time.Time,
	epochType string,
	numberEpochs uint64,
	eligibilityCriteria EligibilityCriteria,
) *MsgCreateRewardProgramRequest {
	return &MsgCreateRewardProgramRequest{
		Title:                 title,
		Description:           description,
		DistributeFromAddress: distributeFromAddress,
		Coin:                  coin,
		MaxRewardByAddress:    maxRewardByAddress,
		ProgramStartTime:      programStartTime,
		EpochType:             epochType,
		NumberEpochs:          numberEpochs,
		EligibilityCriteria:   eligibilityCriteria,
	}
}

// Route implements Msg
func (msg MsgCreateRewardProgramRequest) Route() string { return ModuleName }

// Type implements Msg
func (msg MsgCreateRewardProgramRequest) Type() string { return TypeMsgCreateRewardProgramRequest }

// ValidateBasic runs stateless validation checks on the message.
func (msg MsgCreateRewardProgramRequest) ValidateBasic() error {
	title := msg.GetTitle()
	if len(strings.TrimSpace(title)) == 0 {
		return errors.New("reward program title cannot be blank")
	}
	if len(title) > MaxTitleLength {
		return fmt.Errorf("reward program title is longer than max length of %d", MaxTitleLength)
	}
	description := msg.GetDescription()
	if len(description) == 0 {
		return errors.New("reward program description cannot be blank")
	}
	if len(description) > MaxDescriptionLength {
		return fmt.Errorf("reward program description is longer than max length of %d", MaxDescriptionLength)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateRewardProgramRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners indicates that the message must have been signed by the parent.
func (msg MsgCreateRewardProgramRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.DistributeFromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateRewardProgramRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}
