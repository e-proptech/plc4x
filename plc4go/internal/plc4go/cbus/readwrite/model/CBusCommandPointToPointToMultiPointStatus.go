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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const CBusCommandPointToPointToMultiPointStatus_CR byte = 0xD

// The data-structure of this message
type CBusCommandPointToPointToMultiPointStatus struct {
	*CBusPointToPointToMultipointCommand
	StatusRequest *StatusRequest
	Crc           *Checksum
	PeekAlpha     byte
	Alpha         *Alpha

	// Arguments.
	Srchk bool
}

// The corresponding interface
type ICBusCommandPointToPointToMultiPointStatus interface {
	ICBusPointToPointToMultipointCommand
	// GetStatusRequest returns StatusRequest
	GetStatusRequest() *StatusRequest
	// GetCrc returns Crc
	GetCrc() *Checksum
	// GetPeekAlpha returns PeekAlpha
	GetPeekAlpha() byte
	// GetAlpha returns Alpha
	GetAlpha() *Alpha
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

func (m *CBusCommandPointToPointToMultiPointStatus) InitializeParent(parent *CBusPointToPointToMultipointCommand, bridgeAddress *BridgeAddress, networkRoute *NetworkRoute, peekedApplication byte) {
	m.CBusPointToPointToMultipointCommand.BridgeAddress = bridgeAddress
	m.CBusPointToPointToMultipointCommand.NetworkRoute = networkRoute
	m.CBusPointToPointToMultipointCommand.PeekedApplication = peekedApplication
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *CBusCommandPointToPointToMultiPointStatus) GetStatusRequest() *StatusRequest {
	return m.StatusRequest
}

func (m *CBusCommandPointToPointToMultiPointStatus) GetCrc() *Checksum {
	return m.Crc
}

func (m *CBusCommandPointToPointToMultiPointStatus) GetPeekAlpha() byte {
	return m.PeekAlpha
}

func (m *CBusCommandPointToPointToMultiPointStatus) GetAlpha() *Alpha {
	return m.Alpha
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewCBusCommandPointToPointToMultiPointStatus factory function for CBusCommandPointToPointToMultiPointStatus
func NewCBusCommandPointToPointToMultiPointStatus(statusRequest *StatusRequest, crc *Checksum, peekAlpha byte, alpha *Alpha, bridgeAddress *BridgeAddress, networkRoute *NetworkRoute, peekedApplication byte, srchk bool) *CBusPointToPointToMultipointCommand {
	child := &CBusCommandPointToPointToMultiPointStatus{
		StatusRequest:                       statusRequest,
		Crc:                                 crc,
		PeekAlpha:                           peekAlpha,
		Alpha:                               alpha,
		CBusPointToPointToMultipointCommand: NewCBusPointToPointToMultipointCommand(bridgeAddress, networkRoute, peekedApplication, srchk),
	}
	child.Child = child
	return child.CBusPointToPointToMultipointCommand
}

func CastCBusCommandPointToPointToMultiPointStatus(structType interface{}) *CBusCommandPointToPointToMultiPointStatus {
	if casted, ok := structType.(CBusCommandPointToPointToMultiPointStatus); ok {
		return &casted
	}
	if casted, ok := structType.(*CBusCommandPointToPointToMultiPointStatus); ok {
		return casted
	}
	if casted, ok := structType.(CBusPointToPointToMultipointCommand); ok {
		return CastCBusCommandPointToPointToMultiPointStatus(casted.Child)
	}
	if casted, ok := structType.(*CBusPointToPointToMultipointCommand); ok {
		return CastCBusCommandPointToPointToMultiPointStatus(casted.Child)
	}
	return nil
}

func (m *CBusCommandPointToPointToMultiPointStatus) GetTypeName() string {
	return "CBusCommandPointToPointToMultiPointStatus"
}

func (m *CBusCommandPointToPointToMultiPointStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CBusCommandPointToPointToMultiPointStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (statusRequest)
	lengthInBits += m.StatusRequest.GetLengthInBits()

	// Optional Field (crc)
	if m.Crc != nil {
		lengthInBits += (*m.Crc).GetLengthInBits()
	}

	// Optional Field (alpha)
	if m.Alpha != nil {
		lengthInBits += (*m.Alpha).GetLengthInBits()
	}

	// Const Field (cr)
	lengthInBits += 8

	return lengthInBits
}

func (m *CBusCommandPointToPointToMultiPointStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusCommandPointToPointToMultiPointStatusParse(readBuffer utils.ReadBuffer, srchk bool) (*CBusPointToPointToMultipointCommand, error) {
	if pullErr := readBuffer.PullContext("CBusCommandPointToPointToMultiPointStatus"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != byte(0xFF) {
			log.Info().Fields(map[string]interface{}{
				"expected value": byte(0xFF),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (statusRequest)
	if pullErr := readBuffer.PullContext("statusRequest"); pullErr != nil {
		return nil, pullErr
	}
	_statusRequest, _statusRequestErr := StatusRequestParse(readBuffer)
	if _statusRequestErr != nil {
		return nil, errors.Wrap(_statusRequestErr, "Error parsing 'statusRequest' field")
	}
	statusRequest := CastStatusRequest(_statusRequest)
	if closeErr := readBuffer.CloseContext("statusRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (crc) (Can be skipped, if a given expression evaluates to false)
	var crc *Checksum = nil
	if srchk {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("crc"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := ChecksumParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'crc' field")
		default:
			crc = CastChecksum(_val)
			if closeErr := readBuffer.CloseContext("crc"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Peek Field (peekAlpha)
	currentPos = readBuffer.GetPos()
	peekAlpha, _err := readBuffer.ReadByte("peekAlpha")
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'peekAlpha' field")
	}

	readBuffer.Reset(currentPos)

	// Optional Field (alpha) (Can be skipped, if a given expression evaluates to false)
	var alpha *Alpha = nil
	if bool(bool(bool((peekAlpha) >= (0x67)))) && bool(bool(bool((peekAlpha) <= (0x7A)))) {
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("alpha"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := AlphaParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'alpha' field")
		default:
			alpha = CastAlpha(_val)
			if closeErr := readBuffer.CloseContext("alpha"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Const Field (cr)
	cr, _crErr := readBuffer.ReadByte("cr")
	if _crErr != nil {
		return nil, errors.Wrap(_crErr, "Error parsing 'cr' field")
	}
	if cr != CBusCommandPointToPointToMultiPointStatus_CR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CBusCommandPointToPointToMultiPointStatus_CR) + " but got " + fmt.Sprintf("%d", cr))
	}

	if closeErr := readBuffer.CloseContext("CBusCommandPointToPointToMultiPointStatus"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CBusCommandPointToPointToMultiPointStatus{
		StatusRequest:                       CastStatusRequest(statusRequest),
		Crc:                                 CastChecksum(crc),
		PeekAlpha:                           peekAlpha,
		Alpha:                               CastAlpha(alpha),
		CBusPointToPointToMultipointCommand: &CBusPointToPointToMultipointCommand{},
	}
	_child.CBusPointToPointToMultipointCommand.Child = _child
	return _child.CBusPointToPointToMultipointCommand, nil
}

func (m *CBusCommandPointToPointToMultiPointStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusCommandPointToPointToMultiPointStatus"); pushErr != nil {
			return pushErr
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteByte("reserved", byte(0xFF))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (statusRequest)
		if pushErr := writeBuffer.PushContext("statusRequest"); pushErr != nil {
			return pushErr
		}
		_statusRequestErr := m.StatusRequest.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("statusRequest"); popErr != nil {
			return popErr
		}
		if _statusRequestErr != nil {
			return errors.Wrap(_statusRequestErr, "Error serializing 'statusRequest' field")
		}

		// Optional Field (crc) (Can be skipped, if the value is null)
		var crc *Checksum = nil
		if m.Crc != nil {
			if pushErr := writeBuffer.PushContext("crc"); pushErr != nil {
				return pushErr
			}
			crc = m.Crc
			_crcErr := crc.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("crc"); popErr != nil {
				return popErr
			}
			if _crcErr != nil {
				return errors.Wrap(_crcErr, "Error serializing 'crc' field")
			}
		}

		// Optional Field (alpha) (Can be skipped, if the value is null)
		var alpha *Alpha = nil
		if m.Alpha != nil {
			if pushErr := writeBuffer.PushContext("alpha"); pushErr != nil {
				return pushErr
			}
			alpha = m.Alpha
			_alphaErr := alpha.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("alpha"); popErr != nil {
				return popErr
			}
			if _alphaErr != nil {
				return errors.Wrap(_alphaErr, "Error serializing 'alpha' field")
			}
		}

		// Const Field (cr)
		_crErr := writeBuffer.WriteByte("cr", 0xD)
		if _crErr != nil {
			return errors.Wrap(_crErr, "Error serializing 'cr' field")
		}

		if popErr := writeBuffer.PopContext("CBusCommandPointToPointToMultiPointStatus"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CBusCommandPointToPointToMultiPointStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
