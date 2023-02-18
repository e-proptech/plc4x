/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"context"
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// NullCommandResponse is the corresponding interface of NullCommandResponse
type NullCommandResponse interface {
	utils.LengthAware
	utils.Serializable
	EipPacket
}

// NullCommandResponseExactly can be used when we want exactly this type and not a type which fulfills NullCommandResponse.
// This is useful for switch cases.
type NullCommandResponseExactly interface {
	NullCommandResponse
	isNullCommandResponse() bool
}

// _NullCommandResponse is the data-structure of this message
type _NullCommandResponse struct {
	*_EipPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NullCommandResponse) GetCommand() uint16 {
	return 0x0001
}

func (m *_NullCommandResponse) GetResponse() bool {
	return bool(true)
}

func (m *_NullCommandResponse) GetPacketLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NullCommandResponse) InitializeParent(parent EipPacket, sessionHandle uint32, status uint32, senderContext []byte, options uint32) {
	m.SessionHandle = sessionHandle
	m.Status = status
	m.SenderContext = senderContext
	m.Options = options
}

func (m *_NullCommandResponse) GetParent() EipPacket {
	return m._EipPacket
}

// NewNullCommandResponse factory function for _NullCommandResponse
func NewNullCommandResponse(sessionHandle uint32, status uint32, senderContext []byte, options uint32, order IntegerEncoding) *_NullCommandResponse {
	_result := &_NullCommandResponse{
		_EipPacket: NewEipPacket(sessionHandle, status, senderContext, options, order),
	}
	_result._EipPacket._EipPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNullCommandResponse(structType interface{}) NullCommandResponse {
	if casted, ok := structType.(NullCommandResponse); ok {
		return casted
	}
	if casted, ok := structType.(*NullCommandResponse); ok {
		return *casted
	}
	return nil
}

func (m *_NullCommandResponse) GetTypeName() string {
	return "NullCommandResponse"
}

func (m *_NullCommandResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_NullCommandResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NullCommandResponseParse(theBytes []byte, order IntegerEncoding, response bool) (NullCommandResponse, error) {
	return NullCommandResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased((utils.InlineIf(bool((order) == (IntegerEncoding_BIG_ENDIAN)), func() interface{} { return binary.ByteOrder(binary.BigEndian) }, func() interface{} { return binary.ByteOrder(binary.LittleEndian) })).(binary.ByteOrder))), order, response)
}

func NullCommandResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, order IntegerEncoding, response bool) (NullCommandResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NullCommandResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NullCommandResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("NullCommandResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NullCommandResponse")
	}

	// Create a partially initialized instance
	_child := &_NullCommandResponse{
		_EipPacket: &_EipPacket{
			Order: order,
		},
	}
	_child._EipPacket._EipPacketChildRequirements = _child
	return _child, nil
}

func (m *_NullCommandResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer((utils.InlineIf(bool((order) == (IntegerEncoding_BIG_ENDIAN)), func() interface{} { return binary.ByteOrder(binary.BigEndian) }, func() interface{} { return binary.ByteOrder(binary.LittleEndian) })).(binary.ByteOrder)))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NullCommandResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NullCommandResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NullCommandResponse")
		}

		if popErr := writeBuffer.PopContext("NullCommandResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NullCommandResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NullCommandResponse) isNullCommandResponse() bool {
	return true
}

func (m *_NullCommandResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
