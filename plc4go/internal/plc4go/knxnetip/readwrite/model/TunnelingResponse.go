/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type TunnelingResponse struct {
	*KnxNetIpMessage
	TunnelingResponseDataBlock *TunnelingResponseDataBlock
}

// The corresponding interface
type ITunnelingResponse interface {
	IKnxNetIpMessage
	// GetTunnelingResponseDataBlock returns TunnelingResponseDataBlock
	GetTunnelingResponseDataBlock() *TunnelingResponseDataBlock
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *TunnelingResponse) MsgType() uint16 {
	return 0x0421
}

func (m *TunnelingResponse) GetMsgType() uint16 {
	return 0x0421
}

func (m *TunnelingResponse) InitializeParent(parent *KnxNetIpMessage) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *TunnelingResponse) GetTunnelingResponseDataBlock() *TunnelingResponseDataBlock {
	return m.TunnelingResponseDataBlock
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewTunnelingResponse factory function for TunnelingResponse
func NewTunnelingResponse(tunnelingResponseDataBlock *TunnelingResponseDataBlock) *KnxNetIpMessage {
	child := &TunnelingResponse{
		TunnelingResponseDataBlock: tunnelingResponseDataBlock,
		KnxNetIpMessage:            NewKnxNetIpMessage(),
	}
	child.Child = child
	return child.KnxNetIpMessage
}

func CastTunnelingResponse(structType interface{}) *TunnelingResponse {
	if casted, ok := structType.(TunnelingResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*TunnelingResponse); ok {
		return casted
	}
	if casted, ok := structType.(KnxNetIpMessage); ok {
		return CastTunnelingResponse(casted.Child)
	}
	if casted, ok := structType.(*KnxNetIpMessage); ok {
		return CastTunnelingResponse(casted.Child)
	}
	return nil
}

func (m *TunnelingResponse) GetTypeName() string {
	return "TunnelingResponse"
}

func (m *TunnelingResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *TunnelingResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (tunnelingResponseDataBlock)
	lengthInBits += m.TunnelingResponseDataBlock.GetLengthInBits()

	return lengthInBits
}

func (m *TunnelingResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TunnelingResponseParse(readBuffer utils.ReadBuffer) (*KnxNetIpMessage, error) {
	if pullErr := readBuffer.PullContext("TunnelingResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (tunnelingResponseDataBlock)
	if pullErr := readBuffer.PullContext("tunnelingResponseDataBlock"); pullErr != nil {
		return nil, pullErr
	}
	_tunnelingResponseDataBlock, _tunnelingResponseDataBlockErr := TunnelingResponseDataBlockParse(readBuffer)
	if _tunnelingResponseDataBlockErr != nil {
		return nil, errors.Wrap(_tunnelingResponseDataBlockErr, "Error parsing 'tunnelingResponseDataBlock' field")
	}
	tunnelingResponseDataBlock := CastTunnelingResponseDataBlock(_tunnelingResponseDataBlock)
	if closeErr := readBuffer.CloseContext("tunnelingResponseDataBlock"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("TunnelingResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &TunnelingResponse{
		TunnelingResponseDataBlock: CastTunnelingResponseDataBlock(tunnelingResponseDataBlock),
		KnxNetIpMessage:            &KnxNetIpMessage{},
	}
	_child.KnxNetIpMessage.Child = _child
	return _child.KnxNetIpMessage, nil
}

func (m *TunnelingResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TunnelingResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (tunnelingResponseDataBlock)
		if pushErr := writeBuffer.PushContext("tunnelingResponseDataBlock"); pushErr != nil {
			return pushErr
		}
		_tunnelingResponseDataBlockErr := m.TunnelingResponseDataBlock.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("tunnelingResponseDataBlock"); popErr != nil {
			return popErr
		}
		if _tunnelingResponseDataBlockErr != nil {
			return errors.Wrap(_tunnelingResponseDataBlockErr, "Error serializing 'tunnelingResponseDataBlock' field")
		}

		if popErr := writeBuffer.PopContext("TunnelingResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *TunnelingResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
