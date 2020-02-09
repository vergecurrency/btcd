// Copyright (c) 2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package wire

import (
	"fmt"
	"io"
)

// MsgSendCmpct is a placeholder for future reference
type MsgSendCmpct struct{}

// BtcDecode decodes r using the bitcoin protocol encoding into the receiver.
// This is part of the Message interface implementation.
func (msg *MsgSendCmpct) BtcDecode(r io.Reader, pver uint32, enc MessageEncoding) error {
	if pver < SendHeadersVersion {
		str := fmt.Sprintf("sendheaders message invalid for protocol "+
			"version %d", pver)
		return messageError("MsgSendCmpct.BtcDecode", str)
	}

	return nil
}

// BtcEncode encodes the receiver to w using the bitcoin protocol encoding.
// This is part of the Message interface implementation.
func (msg *MsgSendCmpct) BtcEncode(w io.Writer, pver uint32, enc MessageEncoding) error {
	if pver < SendHeadersVersion {
		str := fmt.Sprintf("sendheaders message invalid for protocol "+
			"version %d", pver)
		return messageError("MsgSendCmpct.BtcEncode", str)
	}

	return nil
}

// Command returns the protocol command string for the message.  This is part
// of the Message interface implementation.
func (msg *MsgSendCmpct) Command() string {
	return CmdSendCmpct
}

// MaxPayloadLength returns the maximum length the payload can be for the
// receiver.  This is part of the Message interface implementation.
func (msg *MsgSendCmpct) MaxPayloadLength(pver uint32) uint32 {
	return 1024
}

// NewMsgSendCmpct returns a new bitcoin sendheaders message that conforms to
// the Message interface.  See MsgSendCmpct for details.
func NewMsgSendCmpct() *MsgSendCmpct {
	return &MsgSendCmpct{}
}
