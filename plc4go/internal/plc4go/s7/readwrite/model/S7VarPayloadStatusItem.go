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
type S7VarPayloadStatusItem struct {
	ReturnCode DataTransportErrorCode
}

// The corresponding interface
type IS7VarPayloadStatusItem interface {
	// GetReturnCode returns ReturnCode
	GetReturnCode() DataTransportErrorCode
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
func (m *S7VarPayloadStatusItem) GetReturnCode() DataTransportErrorCode {
	return m.ReturnCode
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewS7VarPayloadStatusItem factory function for S7VarPayloadStatusItem
func NewS7VarPayloadStatusItem(returnCode DataTransportErrorCode) *S7VarPayloadStatusItem {
	return &S7VarPayloadStatusItem{ReturnCode: returnCode}
}

func CastS7VarPayloadStatusItem(structType interface{}) *S7VarPayloadStatusItem {
	if casted, ok := structType.(S7VarPayloadStatusItem); ok {
		return &casted
	}
	if casted, ok := structType.(*S7VarPayloadStatusItem); ok {
		return casted
	}
	return nil
}

func (m *S7VarPayloadStatusItem) GetTypeName() string {
	return "S7VarPayloadStatusItem"
}

func (m *S7VarPayloadStatusItem) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7VarPayloadStatusItem) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (returnCode)
	lengthInBits += 8

	return lengthInBits
}

func (m *S7VarPayloadStatusItem) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7VarPayloadStatusItemParse(readBuffer utils.ReadBuffer) (*S7VarPayloadStatusItem, error) {
	if pullErr := readBuffer.PullContext("S7VarPayloadStatusItem"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (returnCode)
	if pullErr := readBuffer.PullContext("returnCode"); pullErr != nil {
		return nil, pullErr
	}
	_returnCode, _returnCodeErr := DataTransportErrorCodeParse(readBuffer)
	if _returnCodeErr != nil {
		return nil, errors.Wrap(_returnCodeErr, "Error parsing 'returnCode' field")
	}
	returnCode := _returnCode
	if closeErr := readBuffer.CloseContext("returnCode"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("S7VarPayloadStatusItem"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewS7VarPayloadStatusItem(returnCode), nil
}

func (m *S7VarPayloadStatusItem) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("S7VarPayloadStatusItem"); pushErr != nil {
		return pushErr
	}

	// Simple Field (returnCode)
	if pushErr := writeBuffer.PushContext("returnCode"); pushErr != nil {
		return pushErr
	}
	_returnCodeErr := m.ReturnCode.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("returnCode"); popErr != nil {
		return popErr
	}
	if _returnCodeErr != nil {
		return errors.Wrap(_returnCodeErr, "Error serializing 'returnCode' field")
	}

	if popErr := writeBuffer.PopContext("S7VarPayloadStatusItem"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *S7VarPayloadStatusItem) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
