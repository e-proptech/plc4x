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
type StatusByte struct {
	Gav3 GAVState
	Gav2 GAVState
	Gav1 GAVState
	Gav0 GAVState
}

// The corresponding interface
type IStatusByte interface {
	// GetGav3 returns Gav3
	GetGav3() GAVState
	// GetGav2 returns Gav2
	GetGav2() GAVState
	// GetGav1 returns Gav1
	GetGav1() GAVState
	// GetGav0 returns Gav0
	GetGav0() GAVState
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
func (m *StatusByte) GetGav3() GAVState {
	return m.Gav3
}

func (m *StatusByte) GetGav2() GAVState {
	return m.Gav2
}

func (m *StatusByte) GetGav1() GAVState {
	return m.Gav1
}

func (m *StatusByte) GetGav0() GAVState {
	return m.Gav0
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewStatusByte factory function for StatusByte
func NewStatusByte(gav3 GAVState, gav2 GAVState, gav1 GAVState, gav0 GAVState) *StatusByte {
	return &StatusByte{Gav3: gav3, Gav2: gav2, Gav1: gav1, Gav0: gav0}
}

func CastStatusByte(structType interface{}) *StatusByte {
	if casted, ok := structType.(StatusByte); ok {
		return &casted
	}
	if casted, ok := structType.(*StatusByte); ok {
		return casted
	}
	return nil
}

func (m *StatusByte) GetTypeName() string {
	return "StatusByte"
}

func (m *StatusByte) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *StatusByte) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (gav3)
	lengthInBits += 2

	// Simple field (gav2)
	lengthInBits += 2

	// Simple field (gav1)
	lengthInBits += 2

	// Simple field (gav0)
	lengthInBits += 2

	return lengthInBits
}

func (m *StatusByte) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func StatusByteParse(readBuffer utils.ReadBuffer) (*StatusByte, error) {
	if pullErr := readBuffer.PullContext("StatusByte"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (gav3)
	if pullErr := readBuffer.PullContext("gav3"); pullErr != nil {
		return nil, pullErr
	}
	_gav3, _gav3Err := GAVStateParse(readBuffer)
	if _gav3Err != nil {
		return nil, errors.Wrap(_gav3Err, "Error parsing 'gav3' field")
	}
	gav3 := _gav3
	if closeErr := readBuffer.CloseContext("gav3"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (gav2)
	if pullErr := readBuffer.PullContext("gav2"); pullErr != nil {
		return nil, pullErr
	}
	_gav2, _gav2Err := GAVStateParse(readBuffer)
	if _gav2Err != nil {
		return nil, errors.Wrap(_gav2Err, "Error parsing 'gav2' field")
	}
	gav2 := _gav2
	if closeErr := readBuffer.CloseContext("gav2"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (gav1)
	if pullErr := readBuffer.PullContext("gav1"); pullErr != nil {
		return nil, pullErr
	}
	_gav1, _gav1Err := GAVStateParse(readBuffer)
	if _gav1Err != nil {
		return nil, errors.Wrap(_gav1Err, "Error parsing 'gav1' field")
	}
	gav1 := _gav1
	if closeErr := readBuffer.CloseContext("gav1"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (gav0)
	if pullErr := readBuffer.PullContext("gav0"); pullErr != nil {
		return nil, pullErr
	}
	_gav0, _gav0Err := GAVStateParse(readBuffer)
	if _gav0Err != nil {
		return nil, errors.Wrap(_gav0Err, "Error parsing 'gav0' field")
	}
	gav0 := _gav0
	if closeErr := readBuffer.CloseContext("gav0"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("StatusByte"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewStatusByte(gav3, gav2, gav1, gav0), nil
}

func (m *StatusByte) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("StatusByte"); pushErr != nil {
		return pushErr
	}

	// Simple Field (gav3)
	if pushErr := writeBuffer.PushContext("gav3"); pushErr != nil {
		return pushErr
	}
	_gav3Err := m.Gav3.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("gav3"); popErr != nil {
		return popErr
	}
	if _gav3Err != nil {
		return errors.Wrap(_gav3Err, "Error serializing 'gav3' field")
	}

	// Simple Field (gav2)
	if pushErr := writeBuffer.PushContext("gav2"); pushErr != nil {
		return pushErr
	}
	_gav2Err := m.Gav2.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("gav2"); popErr != nil {
		return popErr
	}
	if _gav2Err != nil {
		return errors.Wrap(_gav2Err, "Error serializing 'gav2' field")
	}

	// Simple Field (gav1)
	if pushErr := writeBuffer.PushContext("gav1"); pushErr != nil {
		return pushErr
	}
	_gav1Err := m.Gav1.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("gav1"); popErr != nil {
		return popErr
	}
	if _gav1Err != nil {
		return errors.Wrap(_gav1Err, "Error serializing 'gav1' field")
	}

	// Simple Field (gav0)
	if pushErr := writeBuffer.PushContext("gav0"); pushErr != nil {
		return pushErr
	}
	_gav0Err := m.Gav0.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("gav0"); popErr != nil {
		return popErr
	}
	if _gav0Err != nil {
		return errors.Wrap(_gav0Err, "Error serializing 'gav0' field")
	}

	if popErr := writeBuffer.PopContext("StatusByte"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *StatusByte) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
