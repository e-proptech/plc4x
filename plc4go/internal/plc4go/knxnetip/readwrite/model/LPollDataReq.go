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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type LPollDataReq struct {
	*CEMI

	// Arguments.
	Size uint16
}

// The corresponding interface
type ILPollDataReq interface {
	ICEMI
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
func (m *LPollDataReq) MessageCode() uint8 {
	return 0x13
}

func (m *LPollDataReq) GetMessageCode() uint8 {
	return 0x13
}

func (m *LPollDataReq) InitializeParent(parent *CEMI) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewLPollDataReq factory function for LPollDataReq
func NewLPollDataReq(size uint16) *CEMI {
	child := &LPollDataReq{
		CEMI: NewCEMI(size),
	}
	child.Child = child
	return child.CEMI
}

func CastLPollDataReq(structType interface{}) *LPollDataReq {
	if casted, ok := structType.(LPollDataReq); ok {
		return &casted
	}
	if casted, ok := structType.(*LPollDataReq); ok {
		return casted
	}
	if casted, ok := structType.(CEMI); ok {
		return CastLPollDataReq(casted.Child)
	}
	if casted, ok := structType.(*CEMI); ok {
		return CastLPollDataReq(casted.Child)
	}
	return nil
}

func (m *LPollDataReq) GetTypeName() string {
	return "LPollDataReq"
}

func (m *LPollDataReq) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *LPollDataReq) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *LPollDataReq) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func LPollDataReqParse(readBuffer utils.ReadBuffer, size uint16) (*CEMI, error) {
	if pullErr := readBuffer.PullContext("LPollDataReq"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("LPollDataReq"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &LPollDataReq{
		CEMI: &CEMI{},
	}
	_child.CEMI.Child = _child
	return _child.CEMI, nil
}

func (m *LPollDataReq) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LPollDataReq"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("LPollDataReq"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *LPollDataReq) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
