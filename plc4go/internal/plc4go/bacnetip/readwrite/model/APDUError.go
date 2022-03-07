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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type APDUError struct {
	*APDU
	OriginalInvokeId uint8
	Error            *BACnetError

	// Arguments.
	ApduLength uint16
}

// The corresponding interface
type IAPDUError interface {
	IAPDU
	// GetOriginalInvokeId returns OriginalInvokeId
	GetOriginalInvokeId() uint8
	// GetError returns Error
	GetError() *BACnetError
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
func (m *APDUError) ApduType() uint8 {
	return 0x5
}

func (m *APDUError) GetApduType() uint8 {
	return 0x5
}

func (m *APDUError) InitializeParent(parent *APDU) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *APDUError) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *APDUError) GetError() *BACnetError {
	return m.Error
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewAPDUError factory function for APDUError
func NewAPDUError(originalInvokeId uint8, error *BACnetError, apduLength uint16) *APDU {
	child := &APDUError{
		OriginalInvokeId: originalInvokeId,
		Error:            error,
		APDU:             NewAPDU(apduLength),
	}
	child.Child = child
	return child.APDU
}

func CastAPDUError(structType interface{}) *APDUError {
	if casted, ok := structType.(APDUError); ok {
		return &casted
	}
	if casted, ok := structType.(*APDUError); ok {
		return casted
	}
	if casted, ok := structType.(APDU); ok {
		return CastAPDUError(casted.Child)
	}
	if casted, ok := structType.(*APDU); ok {
		return CastAPDUError(casted.Child)
	}
	return nil
}

func (m *APDUError) GetTypeName() string {
	return "APDUError"
}

func (m *APDUError) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *APDUError) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (error)
	lengthInBits += m.Error.GetLengthInBits()

	return lengthInBits
}

func (m *APDUError) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func APDUErrorParse(readBuffer utils.ReadBuffer, apduLength uint16) (*APDU, error) {
	if pullErr := readBuffer.PullContext("APDUError"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 4)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (originalInvokeId)
	_originalInvokeId, _originalInvokeIdErr := readBuffer.ReadUint8("originalInvokeId", 8)
	if _originalInvokeIdErr != nil {
		return nil, errors.Wrap(_originalInvokeIdErr, "Error parsing 'originalInvokeId' field")
	}
	originalInvokeId := _originalInvokeId

	// Simple Field (error)
	if pullErr := readBuffer.PullContext("error"); pullErr != nil {
		return nil, pullErr
	}
	_error, _errorErr := BACnetErrorParse(readBuffer)
	if _errorErr != nil {
		return nil, errors.Wrap(_errorErr, "Error parsing 'error' field")
	}
	error := CastBACnetError(_error)
	if closeErr := readBuffer.CloseContext("error"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("APDUError"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &APDUError{
		OriginalInvokeId: originalInvokeId,
		Error:            CastBACnetError(error),
		APDU:             &APDU{},
	}
	_child.APDU.Child = _child
	return _child.APDU, nil
}

func (m *APDUError) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUError"); pushErr != nil {
			return pushErr
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 4, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (originalInvokeId)
		originalInvokeId := uint8(m.OriginalInvokeId)
		_originalInvokeIdErr := writeBuffer.WriteUint8("originalInvokeId", 8, (originalInvokeId))
		if _originalInvokeIdErr != nil {
			return errors.Wrap(_originalInvokeIdErr, "Error serializing 'originalInvokeId' field")
		}

		// Simple Field (error)
		if pushErr := writeBuffer.PushContext("error"); pushErr != nil {
			return pushErr
		}
		_errorErr := m.Error.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("error"); popErr != nil {
			return popErr
		}
		if _errorErr != nil {
			return errors.Wrap(_errorErr, "Error serializing 'error' field")
		}

		if popErr := writeBuffer.PopContext("APDUError"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *APDUError) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
