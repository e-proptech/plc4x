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
type CALDataReplyReply struct {
	*CALData
	ParamNumber uint8
	Data        []byte
}

// The corresponding interface
type ICALDataReplyReply interface {
	ICALData
	// GetParamNumber returns ParamNumber
	GetParamNumber() uint8
	// GetData returns Data
	GetData() []byte
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
func (m *CALDataReplyReply) CommandType() CALCommandType {
	return CALCommandType_REPLY
}

func (m *CALDataReplyReply) GetCommandType() CALCommandType {
	return CALCommandType_REPLY
}

func (m *CALDataReplyReply) InitializeParent(parent *CALData, commandTypeContainer CALCommandTypeContainer) {
	m.CALData.CommandTypeContainer = commandTypeContainer
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *CALDataReplyReply) GetParamNumber() uint8 {
	return m.ParamNumber
}

func (m *CALDataReplyReply) GetData() []byte {
	return m.Data
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewCALDataReplyReply factory function for CALDataReplyReply
func NewCALDataReplyReply(paramNumber uint8, data []byte, commandTypeContainer CALCommandTypeContainer) *CALData {
	child := &CALDataReplyReply{
		ParamNumber: paramNumber,
		Data:        data,
		CALData:     NewCALData(commandTypeContainer),
	}
	child.Child = child
	return child.CALData
}

func CastCALDataReplyReply(structType interface{}) *CALDataReplyReply {
	if casted, ok := structType.(CALDataReplyReply); ok {
		return &casted
	}
	if casted, ok := structType.(*CALDataReplyReply); ok {
		return casted
	}
	if casted, ok := structType.(CALData); ok {
		return CastCALDataReplyReply(casted.Child)
	}
	if casted, ok := structType.(*CALData); ok {
		return CastCALDataReplyReply(casted.Child)
	}
	return nil
}

func (m *CALDataReplyReply) GetTypeName() string {
	return "CALDataReplyReply"
}

func (m *CALDataReplyReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CALDataReplyReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (paramNumber)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *CALDataReplyReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataReplyReplyParse(readBuffer utils.ReadBuffer, commandTypeContainer CALCommandTypeContainer) (*CALData, error) {
	if pullErr := readBuffer.PullContext("CALDataReplyReply"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (paramNumber)
	_paramNumber, _paramNumberErr := readBuffer.ReadUint8("paramNumber", 8)
	if _paramNumberErr != nil {
		return nil, errors.Wrap(_paramNumberErr, "Error parsing 'paramNumber' field")
	}
	paramNumber := _paramNumber
	// Byte Array field (data)
	numberOfBytesdata := int(commandTypeContainer.NumBytes())
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	if closeErr := readBuffer.CloseContext("CALDataReplyReply"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CALDataReplyReply{
		ParamNumber: paramNumber,
		Data:        data,
		CALData:     &CALData{},
	}
	_child.CALData.Child = _child
	return _child.CALData, nil
}

func (m *CALDataReplyReply) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataReplyReply"); pushErr != nil {
			return pushErr
		}

		// Simple Field (paramNumber)
		paramNumber := uint8(m.ParamNumber)
		_paramNumberErr := writeBuffer.WriteUint8("paramNumber", 8, (paramNumber))
		if _paramNumberErr != nil {
			return errors.Wrap(_paramNumberErr, "Error serializing 'paramNumber' field")
		}

		// Array Field (data)
		if m.Data != nil {
			// Byte Array field (data)
			_writeArrayErr := writeBuffer.WriteByteArray("data", m.Data)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
			}
		}

		if popErr := writeBuffer.PopContext("CALDataReplyReply"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CALDataReplyReply) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
