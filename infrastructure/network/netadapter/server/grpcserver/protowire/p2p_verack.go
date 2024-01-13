package protowire

import (
	"github.com/BTCGhostdag/BTCD/app/appmessage"
	"github.com/pkg/errors"
)

func (x *BTCDdMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "BTCDdMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *BTCDdMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
