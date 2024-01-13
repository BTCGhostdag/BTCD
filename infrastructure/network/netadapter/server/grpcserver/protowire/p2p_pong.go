package protowire

import (
	"github.com/BTCGhostdag/BTCD/app/appmessage"
	"github.com/pkg/errors"
)

func (x *BTCDdMessage_Pong) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "BTCDdMessage_Pong is nil")
	}
	return x.Pong.toAppMessage()
}

func (x *PongMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PongMessage is nil")
	}
	return &appmessage.MsgPong{
		Nonce: x.Nonce,
	}, nil
}

func (x *BTCDdMessage_Pong) fromAppMessage(msgPong *appmessage.MsgPong) error {
	x.Pong = &PongMessage{
		Nonce: msgPong.Nonce,
	}
	return nil
}
