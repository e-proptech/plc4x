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
type FirmataCommand struct {

	// Arguments.
	Response bool
	Child    IFirmataCommandChild
}

// The corresponding interface
type IFirmataCommand interface {
	// CommandCode returns CommandCode
	CommandCode() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IFirmataCommandParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IFirmataCommand, serializeChildFunction func() error) error
	GetTypeName() string
}

type IFirmataCommandChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *FirmataCommand)
	GetTypeName() string
	IFirmataCommand
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewFirmataCommand factory function for FirmataCommand
func NewFirmataCommand(response bool) *FirmataCommand {
	return &FirmataCommand{Response: response}
}

func CastFirmataCommand(structType interface{}) *FirmataCommand {
	if casted, ok := structType.(FirmataCommand); ok {
		return &casted
	}
	if casted, ok := structType.(*FirmataCommand); ok {
		return casted
	}
	return nil
}

func (m *FirmataCommand) GetTypeName() string {
	return "FirmataCommand"
}

func (m *FirmataCommand) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *FirmataCommand) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *FirmataCommand) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (commandCode)
	lengthInBits += 4

	return lengthInBits
}

func (m *FirmataCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func FirmataCommandParse(readBuffer utils.ReadBuffer, response bool) (*FirmataCommand, error) {
	if pullErr := readBuffer.PullContext("FirmataCommand"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Discriminator Field (commandCode) (Used as input to a switch field)
	commandCode, _commandCodeErr := readBuffer.ReadUint8("commandCode", 4)
	if _commandCodeErr != nil {
		return nil, errors.Wrap(_commandCodeErr, "Error parsing 'commandCode' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *FirmataCommand
	var typeSwitchError error
	switch {
	case commandCode == 0x0: // FirmataCommandSysex
		_parent, typeSwitchError = FirmataCommandSysexParse(readBuffer, response)
	case commandCode == 0x4: // FirmataCommandSetPinMode
		_parent, typeSwitchError = FirmataCommandSetPinModeParse(readBuffer, response)
	case commandCode == 0x5: // FirmataCommandSetDigitalPinValue
		_parent, typeSwitchError = FirmataCommandSetDigitalPinValueParse(readBuffer, response)
	case commandCode == 0x9: // FirmataCommandProtocolVersion
		_parent, typeSwitchError = FirmataCommandProtocolVersionParse(readBuffer, response)
	case commandCode == 0xF: // FirmataCommandSystemReset
		_parent, typeSwitchError = FirmataCommandSystemResetParse(readBuffer, response)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("FirmataCommand"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *FirmataCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *FirmataCommand) SerializeParent(writeBuffer utils.WriteBuffer, child IFirmataCommand, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("FirmataCommand"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (commandCode) (Used as input to a switch field)
	commandCode := uint8(child.CommandCode())
	_commandCodeErr := writeBuffer.WriteUint8("commandCode", 4, (commandCode))

	if _commandCodeErr != nil {
		return errors.Wrap(_commandCodeErr, "Error serializing 'commandCode' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("FirmataCommand"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *FirmataCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
