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
type AlarmMessageAckPushType struct {
	TimeStamp       *DateAndTime
	FunctionId      uint8
	NumberOfObjects uint8
	MessageObjects  []*AlarmMessageAckObjectPushType
}

// The corresponding interface
type IAlarmMessageAckPushType interface {
	// GetTimeStamp returns TimeStamp
	GetTimeStamp() *DateAndTime
	// GetFunctionId returns FunctionId
	GetFunctionId() uint8
	// GetNumberOfObjects returns NumberOfObjects
	GetNumberOfObjects() uint8
	// GetMessageObjects returns MessageObjects
	GetMessageObjects() []*AlarmMessageAckObjectPushType
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
func (m *AlarmMessageAckPushType) GetTimeStamp() *DateAndTime {
	return m.TimeStamp
}

func (m *AlarmMessageAckPushType) GetFunctionId() uint8 {
	return m.FunctionId
}

func (m *AlarmMessageAckPushType) GetNumberOfObjects() uint8 {
	return m.NumberOfObjects
}

func (m *AlarmMessageAckPushType) GetMessageObjects() []*AlarmMessageAckObjectPushType {
	return m.MessageObjects
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewAlarmMessageAckPushType factory function for AlarmMessageAckPushType
func NewAlarmMessageAckPushType(TimeStamp *DateAndTime, functionId uint8, numberOfObjects uint8, messageObjects []*AlarmMessageAckObjectPushType) *AlarmMessageAckPushType {
	return &AlarmMessageAckPushType{TimeStamp: TimeStamp, FunctionId: functionId, NumberOfObjects: numberOfObjects, MessageObjects: messageObjects}
}

func CastAlarmMessageAckPushType(structType interface{}) *AlarmMessageAckPushType {
	if casted, ok := structType.(AlarmMessageAckPushType); ok {
		return &casted
	}
	if casted, ok := structType.(*AlarmMessageAckPushType); ok {
		return casted
	}
	return nil
}

func (m *AlarmMessageAckPushType) GetTypeName() string {
	return "AlarmMessageAckPushType"
}

func (m *AlarmMessageAckPushType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AlarmMessageAckPushType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (TimeStamp)
	lengthInBits += m.TimeStamp.GetLengthInBits()

	// Simple field (functionId)
	lengthInBits += 8

	// Simple field (numberOfObjects)
	lengthInBits += 8

	// Array field
	if len(m.MessageObjects) > 0 {
		for i, element := range m.MessageObjects {
			last := i == len(m.MessageObjects)-1
			lengthInBits += element.GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *AlarmMessageAckPushType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AlarmMessageAckPushTypeParse(readBuffer utils.ReadBuffer) (*AlarmMessageAckPushType, error) {
	if pullErr := readBuffer.PullContext("AlarmMessageAckPushType"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (TimeStamp)
	if pullErr := readBuffer.PullContext("TimeStamp"); pullErr != nil {
		return nil, pullErr
	}
	_TimeStamp, _TimeStampErr := DateAndTimeParse(readBuffer)
	if _TimeStampErr != nil {
		return nil, errors.Wrap(_TimeStampErr, "Error parsing 'TimeStamp' field")
	}
	TimeStamp := CastDateAndTime(_TimeStamp)
	if closeErr := readBuffer.CloseContext("TimeStamp"); closeErr != nil {
		return nil, closeErr
	}

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
	messageObjects := make([]*AlarmMessageAckObjectPushType, numberOfObjects)
	{
		for curItem := uint16(0); curItem < uint16(numberOfObjects); curItem++ {
			_item, _err := AlarmMessageAckObjectPushTypeParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'messageObjects' field")
			}
			messageObjects[curItem] = CastAlarmMessageAckObjectPushType(_item)
		}
	}
	if closeErr := readBuffer.CloseContext("messageObjects", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("AlarmMessageAckPushType"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewAlarmMessageAckPushType(TimeStamp, functionId, numberOfObjects, messageObjects), nil
}

func (m *AlarmMessageAckPushType) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("AlarmMessageAckPushType"); pushErr != nil {
		return pushErr
	}

	// Simple Field (TimeStamp)
	if pushErr := writeBuffer.PushContext("TimeStamp"); pushErr != nil {
		return pushErr
	}
	_TimeStampErr := m.TimeStamp.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("TimeStamp"); popErr != nil {
		return popErr
	}
	if _TimeStampErr != nil {
		return errors.Wrap(_TimeStampErr, "Error serializing 'TimeStamp' field")
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
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'messageObjects' field")
			}
		}
		if popErr := writeBuffer.PopContext("messageObjects", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	if popErr := writeBuffer.PopContext("AlarmMessageAckPushType"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *AlarmMessageAckPushType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
