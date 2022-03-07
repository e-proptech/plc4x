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
type AlarmMessageAckResponseType struct {
	FunctionId      uint8
	NumberOfObjects uint8
	MessageObjects  []uint8
}

// The corresponding interface
type IAlarmMessageAckResponseType interface {
	// GetFunctionId returns FunctionId
	GetFunctionId() uint8
	// GetNumberOfObjects returns NumberOfObjects
	GetNumberOfObjects() uint8
	// GetMessageObjects returns MessageObjects
	GetMessageObjects() []uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *AlarmMessageAckResponseType) GetFunctionId() uint8 {
	return m.FunctionId
}

func (m *AlarmMessageAckResponseType) GetNumberOfObjects() uint8 {
	return m.NumberOfObjects
}

func (m *AlarmMessageAckResponseType) GetMessageObjects() []uint8 {
	return m.MessageObjects
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewAlarmMessageAckResponseType factory function for AlarmMessageAckResponseType
func NewAlarmMessageAckResponseType(functionId uint8, numberOfObjects uint8, messageObjects []uint8) *AlarmMessageAckResponseType {
	return &AlarmMessageAckResponseType{FunctionId: functionId, NumberOfObjects: numberOfObjects, MessageObjects: messageObjects}
}

func CastAlarmMessageAckResponseType(structType interface{}) *AlarmMessageAckResponseType {
	if casted, ok := structType.(AlarmMessageAckResponseType); ok {
		return &casted
	}
	if casted, ok := structType.(*AlarmMessageAckResponseType); ok {
		return casted
	}
	return nil
}

func (m *AlarmMessageAckResponseType) GetTypeName() string {
	return "AlarmMessageAckResponseType"
}

func (m *AlarmMessageAckResponseType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AlarmMessageAckResponseType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (functionId)
	lengthInBits += 8

	// Simple field (numberOfObjects)
	lengthInBits += 8

	// Array field
	if len(m.MessageObjects) > 0 {
		lengthInBits += 8 * uint16(len(m.MessageObjects))
	}

	return lengthInBits
}

func (m *AlarmMessageAckResponseType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AlarmMessageAckResponseTypeParse(readBuffer utils.ReadBuffer) (*AlarmMessageAckResponseType, error) {
	if pullErr := readBuffer.PullContext("AlarmMessageAckResponseType"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (functionId)
	_functionId, _functionIdErr := readBuffer.ReadUint8("functionId", 8)
	if _functionIdErr != nil {
		return nil, errors.Wrap(_functionIdErr, "Error parsing 'functionId' field")
	}
	functionId := _functionId

	// Simple Field (numberOfObjects)
	_numberOfObjects, _numberOfObjectsErr := readBuffer.ReadUint8("numberOfObjects", 8)
	if _numberOfObjectsErr != nil {
		return nil, errors.Wrap(_numberOfObjectsErr, "Error parsing 'numberOfObjects' field")
	}
	numberOfObjects := _numberOfObjects

	// Array field (messageObjects)
	if pullErr := readBuffer.PullContext("messageObjects", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	messageObjects := make([]uint8, numberOfObjects)
	{
		for curItem := uint16(0); curItem < uint16(numberOfObjects); curItem++ {
			_item, _err := readBuffer.ReadUint8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'messageObjects' field")
			}
			messageObjects[curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("messageObjects", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("AlarmMessageAckResponseType"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewAlarmMessageAckResponseType(functionId, numberOfObjects, messageObjects), nil
}

func (m *AlarmMessageAckResponseType) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("AlarmMessageAckResponseType"); pushErr != nil {
		return pushErr
	}

	// Simple Field (functionId)
	functionId := uint8(m.FunctionId)
	_functionIdErr := writeBuffer.WriteUint8("functionId", 8, (functionId))
	if _functionIdErr != nil {
		return errors.Wrap(_functionIdErr, "Error serializing 'functionId' field")
	}

	// Simple Field (numberOfObjects)
	numberOfObjects := uint8(m.NumberOfObjects)
	_numberOfObjectsErr := writeBuffer.WriteUint8("numberOfObjects", 8, (numberOfObjects))
	if _numberOfObjectsErr != nil {
		return errors.Wrap(_numberOfObjectsErr, "Error serializing 'numberOfObjects' field")
	}

	// Array Field (messageObjects)
	if m.MessageObjects != nil {
		if pushErr := writeBuffer.PushContext("messageObjects", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.MessageObjects {
			_elementErr := writeBuffer.WriteUint8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'messageObjects' field")
			}
		}
		if popErr := writeBuffer.PopContext("messageObjects", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	if popErr := writeBuffer.PopContext("AlarmMessageAckResponseType"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *AlarmMessageAckResponseType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
