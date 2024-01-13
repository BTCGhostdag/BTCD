package protowire

import (
	"github.com/BTCGhostdag/BTCD/app/appmessage"
	"github.com/pkg/errors"
)

func (x *BTCDdMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "BTCDdMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *BTCDdMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
