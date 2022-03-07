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
type DF1CommandResponseMessageProtectedTypedLogicalRead struct {
	*DF1ResponseMessage
	Data []uint8

	// Arguments.
	PayloadLength uint16
}

// The corresponding interface
type IDF1CommandResponseMessageProtectedTypedLogicalRead interface {
	IDF1ResponseMessage
	// GetData returns Data
	GetData() []uint8
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
func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) CommandCode() uint8 {
	return 0x4F
}

func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) GetCommandCode() uint8 {
	return 0x4F
}

func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) InitializeParent(parent *DF1ResponseMessage, destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16) {
	m.DF1ResponseMessage.DestinationAddress = destinationAddress
	m.DF1ResponseMessage.SourceAddress = sourceAddress
	m.DF1ResponseMessage.Status = status
	m.DF1ResponseMessage.TransactionCounter = transactionCounter
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) GetData() []uint8 {
	return m.Data
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewDF1CommandResponseMessageProtectedTypedLogicalRead factory function for DF1CommandResponseMessageProtectedTypedLogicalRead
func NewDF1CommandResponseMessageProtectedTypedLogicalRead(data []uint8, destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16, payloadLength uint16) *DF1ResponseMessage {
	child := &DF1CommandResponseMessageProtectedTypedLogicalRead{
		Data:               data,
		DF1ResponseMessage: NewDF1ResponseMessage(destinationAddress, sourceAddress, status, transactionCounter, payloadLength),
	}
	child.Child = child
	return child.DF1ResponseMessage
}

func CastDF1CommandResponseMessageProtectedTypedLogicalRead(structType interface{}) *DF1CommandResponseMessageProtectedTypedLogicalRead {
	if casted, ok := structType.(DF1CommandResponseMessageProtectedTypedLogicalRead); ok {
		return &casted
	}
	if casted, ok := structType.(*DF1CommandResponseMessageProtectedTypedLogicalRead); ok {
		return casted
	}
	if casted, ok := structType.(DF1ResponseMessage); ok {
		return CastDF1CommandResponseMessageProtectedTypedLogicalRead(casted.Child)
	}
	if casted, ok := structType.(*DF1ResponseMessage); ok {
		return CastDF1CommandResponseMessageProtectedTypedLogicalRead(casted.Child)
	}
	return nil
}

func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) GetTypeName() string {
	return "DF1CommandResponseMessageProtectedTypedLogicalRead"
}

func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DF1CommandResponseMessageProtectedTypedLogicalReadParse(readBuffer utils.ReadBuffer, payloadLength uint16) (*DF1ResponseMessage, error) {
	if pullErr := readBuffer.PullContext("DF1CommandResponseMessageProtectedTypedLogicalRead"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	data := make([]uint8, 0)
	{
		_dataLength := uint16(payloadLength) - uint16(uint16(8))
		_dataEndPos := readBuffer.GetPos() + uint16(_dataLength)
		for readBuffer.GetPos() < _dataEndPos {
			_item, _err := readBuffer.ReadUint8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'data' field")
			}
			data = append(data, _item)
		}
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("DF1CommandResponseMessageProtectedTypedLogicalRead"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &DF1CommandResponseMessageProtectedTypedLogicalRead{
		Data:               data,
		DF1ResponseMessage: &DF1ResponseMessage{},
	}
	_child.DF1ResponseMessage.Child = _child
	return _child.DF1ResponseMessage, nil
}

func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1CommandResponseMessageProtectedTypedLogicalRead"); pushErr != nil {
			return pushErr
		}

		// Array Field (data)
		if m.Data != nil {
			if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Data {
				_elementErr := writeBuffer.WriteUint8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'data' field")
				}
			}
			if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("DF1CommandResponseMessageProtectedTypedLogicalRead"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *DF1CommandResponseMessageProtectedTypedLogicalRead) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
