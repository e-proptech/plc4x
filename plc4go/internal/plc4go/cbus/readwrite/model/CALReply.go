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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const CALReply_CR byte = 0x0D
const CALReply_LF byte = 0x0A

// The data-structure of this message
type CALReply struct {
	CalType byte
	CalData *CALData
	Child   ICALReplyChild
}

// The corresponding interface
type ICALReply interface {
	// GetCalType returns CalType
	GetCalType() byte
	// GetCalData returns CalData
	GetCalData() *CALData
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type ICALReplyParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child ICALReply, serializeChildFunction func() error) error
	GetTypeName() string
}

type ICALReplyChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *CALReply, calType byte, calData *CALData)
	GetTypeName() string
	ICALReply
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *CALReply) GetCalType() byte {
	return m.CalType
}

func (m *CALReply) GetCalData() *CALData {
	return m.CalData
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewCALReply factory function for CALReply
func NewCALReply(calType byte, calData *CALData) *CALReply {
	return &CALReply{CalType: calType, CalData: calData}
}

func CastCALReply(structType interface{}) *CALReply {
	if casted, ok := structType.(CALReply); ok {
		return &casted
	}
	if casted, ok := structType.(*CALReply); ok {
		return casted
	}
	return nil
}

func (m *CALReply) GetTypeName() string {
	return "CALReply"
}

func (m *CALReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CALReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *CALReply) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (calData)
	lengthInBits += m.CalData.GetLengthInBits()

	// Const Field (cr)
	lengthInBits += 8

	// Const Field (lf)
	lengthInBits += 8

	return lengthInBits
}

func (m *CALReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALReplyParse(readBuffer utils.ReadBuffer) (*CALReply, error) {
	if pullErr := readBuffer.PullContext("CALReply"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Peek Field (calType)
	currentPos = readBuffer.GetPos()
	calType, _err := readBuffer.ReadByte("calType")
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'calType' field")
	}

	readBuffer.Reset(currentPos)

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *CALReply
	var typeSwitchError error
	switch {
	case calType == 0x86: // CALReplyLong
		_parent, typeSwitchError = CALReplyLongParse(readBuffer)
	case true: // CALReplyShort
		_parent, typeSwitchError = CALReplyShortParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Simple Field (calData)
	if pullErr := readBuffer.PullContext("calData"); pullErr != nil {
		return nil, pullErr
	}
	_calData, _calDataErr := CALDataParse(readBuffer)
	if _calDataErr != nil {
		return nil, errors.Wrap(_calDataErr, "Error parsing 'calData' field")
	}
	calData := CastCALData(_calData)
	if closeErr := readBuffer.CloseContext("calData"); closeErr != nil {
		return nil, closeErr
	}

	// Const Field (cr)
	cr, _crErr := readBuffer.ReadByte("cr")
	if _crErr != nil {
		return nil, errors.Wrap(_crErr, "Error parsing 'cr' field")
	}
	if cr != CALReply_CR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CALReply_CR) + " but got " + fmt.Sprintf("%d", cr))
	}

	// Const Field (lf)
	lf, _lfErr := readBuffer.ReadByte("lf")
	if _lfErr != nil {
		return nil, errors.Wrap(_lfErr, "Error parsing 'lf' field")
	}
	if lf != CALReply_LF {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CALReply_LF) + " but got " + fmt.Sprintf("%d", lf))
	}

	if closeErr := readBuffer.CloseContext("CALReply"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, calType, calData)
	return _parent, nil
}

func (m *CALReply) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *CALReply) SerializeParent(writeBuffer utils.WriteBuffer, child ICALReply, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("CALReply"); pushErr != nil {
		return pushErr
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (calData)
	if pushErr := writeBuffer.PushContext("calData"); pushErr != nil {
		return pushErr
	}
	_calDataErr := m.CalData.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("calData"); popErr != nil {
		return popErr
	}
	if _calDataErr != nil {
		return errors.Wrap(_calDataErr, "Error serializing 'calData' field")
	}

	// Const Field (cr)
	_crErr := writeBuffer.WriteByte("cr", 0x0D)
	if _crErr != nil {
		return errors.Wrap(_crErr, "Error serializing 'cr' field")
	}

	// Const Field (lf)
	_lfErr := writeBuffer.WriteByte("lf", 0x0A)
	if _lfErr != nil {
		return errors.Wrap(_lfErr, "Error serializing 'lf' field")
	}

	if popErr := writeBuffer.PopContext("CALReply"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *CALReply) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
