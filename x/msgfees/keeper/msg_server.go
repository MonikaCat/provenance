package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/provenance-io/provenance/x/msgfees/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the msgfees MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (m msgServer) AssessCustomMsgFee(goCtx context.Context, req *types.MsgAssessCustomMsgFeeRequest) (*types.MsgAssessCustomMsgFeeResponse, error) {
	// method only emits that the event has been submitted, all logic is handled in the provenance custom msg handlers
	ctx := sdk.UnwrapSDKContext(goCtx)

	// if there is a recipient and bips are not set, we will want to emit the default bips with event
	recipientBips := req.RecipientBasisPoints
	if len(req.Recipient) > 0 && len(req.RecipientBasisPoints) == 0 {
		recipientBips = fmt.Sprintf("%v", types.AssessCustomMsgFeeBips)
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAssessCustomMsgFee,
			sdk.NewAttribute(types.KeyAttributeName, req.Name),
			sdk.NewAttribute(types.KeyAttributeAmount, req.Amount.String()),
			sdk.NewAttribute(types.KeyAttributeRecipient, req.Recipient),
			sdk.NewAttribute(types.KeyAttributeBips, recipientBips),
		),
	)
	return &types.MsgAssessCustomMsgFeeResponse{}, nil
}
