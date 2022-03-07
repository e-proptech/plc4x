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
type APDUSegmentAck struct {
	*APDU
	NegativeAck        bool
	Server             bool
	OriginalInvokeId   uint8
	SequenceNumber     uint8
	ProposedWindowSize uint8

	// Arguments.
	ApduLength uint16
}

// The corresponding interface
type IAPDUSegmentAck interface {
	IAPDU
	// GetNegativeAck returns NegativeAck
	GetNegativeAck() bool
	// GetServer returns Server
	GetServer() bool
	// GetOriginalInvokeId returns OriginalInvokeId
	GetOriginalInvokeId() uint8
	// GetSequenceNumber returns SequenceNumber
	GetSequenceNumber() uint8
	// GetProposedWindowSize returns ProposedWindowSize
	GetProposedWindowSize() uint8
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
func (m *APDUSegmentAck) ApduType() uint8 {
	return 0x4
}

func (m *APDUSegmentAck) GetApduType() uint8 {
	return 0x4
}

func (m *APDUSegmentAck) InitializeParent(parent *APDU) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *APDUSegmentAck) GetNegativeAck() bool {
	return m.NegativeAck
}

func (m *APDUSegmentAck) GetServer() bool {
	return m.Server
}

func (m *APDUSegmentAck) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *APDUSegmentAck) GetSequenceNumber() uint8 {
	return m.SequenceNumber
}

func (m *APDUSegmentAck) GetProposedWindowSize() uint8 {
	return m.ProposedWindowSize
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewAPDUSegmentAck factory function for APDUSegmentAck
func NewAPDUSegmentAck(negativeAck bool, server bool, originalInvokeId uint8, sequenceNumber uint8, proposedWindowSize uint8, apduLength uint16) *APDU {
	child := &APDUSegmentAck{
		NegativeAck:        negativeAck,
		Server:             server,
		OriginalInvokeId:   originalInvokeId,
		SequenceNumber:     sequenceNumber,
		ProposedWindowSize: proposedWindowSize,
		APDU:               NewAPDU(apduLength),
	}
	child.Child = child
	return child.APDU
}

func CastAPDUSegmentAck(structType interface{}) *APDUSegmentAck {
	if casted, ok := structType.(APDUSegmentAck); ok {
		return &casted
	}
	if casted, ok := structType.(*APDUSegmentAck); ok {
		return casted
	}
	if casted, ok := structType.(APDU); ok {
		return CastAPDUSegmentAck(casted.Child)
	}
	if casted, ok := structType.(*APDU); ok {
		return CastAPDUSegmentAck(casted.Child)
	}
	return nil
}

func (m *APDUSegmentAck) GetTypeName() string {
	return "APDUSegmentAck"
}

func (m *APDUSegmentAck) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *APDUSegmentAck) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (negativeAck)
	lengthInBits += 1

	// Simple field (server)
	lengthInBits += 1

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (sequenceNumber)
	lengthInBits += 8

	// Simple field (proposedWindowSize)
	lengthInBits += 8

	return lengthInBits
}

func (m *APDUSegmentAck) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func APDUSegmentAckParse(readBuffer utils.ReadBuffer, apduLength uint16) (*APDU, error) {
	if pullErr := readBuffer.PullContext("APDUSegmentAck"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 2)
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

	// Simple Field (negativeAck)
	_negativeAck, _negativeAckErr := readBuffer.ReadBit("negativeAck")
	if _negativeAckErr != nil {
		return nil, errors.Wrap(_negativeAckErr, "Error parsing 'negativeAck' field")
	}
	negativeAck := _negativeAck

	// Simple Field (server)
	_server, _serverErr := readBuffer.ReadBit("server")
	if _serverErr != nil {
		return nil, errors.Wrap(_serverErr, "Error parsing 'server' field")
	}
	server := _server

	// Simple Field (originalInvokeId)
	_originalInvokeId, _originalInvokeIdErr := readBuffer.ReadUint8("originalInvokeId", 8)
	if _originalInvokeIdErr != nil {
		return nil, errors.Wrap(_originalInvokeIdErr, "Error parsing 'originalInvokeId' field")
	}
	originalInvokeId := _originalInvokeId

	// Simple Field (sequenceNumber)
	_sequenceNumber, _sequenceNumberErr := readBuffer.ReadUint8("sequenceNumber", 8)
	if _sequenceNumberErr != nil {
		return nil, errors.Wrap(_sequenceNumberErr, "Error parsing 'sequenceNumber' field")
	}
	sequenceNumber := _sequenceNumber

	// Simple Field (proposedWindowSize)
	_proposedWindowSize, _proposedWindowSizeErr := readBuffer.ReadUint8("proposedWindowSize", 8)
	if _proposedWindowSizeErr != nil {
		return nil, errors.Wrap(_proposedWindowSizeErr, "Error parsing 'proposedWindowSize' field")
	}
	proposedWindowSize := _proposedWindowSize

	if closeErr := readBuffer.CloseContext("APDUSegmentAck"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &APDUSegmentAck{
		NegativeAck:        negativeAck,
		Server:             server,
		OriginalInvokeId:   originalInvokeId,
		SequenceNumber:     sequenceNumber,
		ProposedWindowSize: proposedWindowSize,
		APDU:               &APDU{},
	}
	_child.APDU.Child = _child
	return _child.APDU, nil
}

func (m *APDUSegmentAck) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUSegmentAck"); pushErr != nil {
			return pushErr
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 2, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (negativeAck)
		negativeAck := bool(m.NegativeAck)
		_negativeAckErr := writeBuffer.WriteBit("negativeAck", (negativeAck))
		if _negativeAckErr != nil {
			return errors.Wrap(_negativeAckErr, "Error serializing 'negativeAck' field")
		}

		// Simple Field (server)
		server := bool(m.Server)
		_serverErr := writeBuffer.WriteBit("server", (server))
		if _serverErr != nil {
			return errors.Wrap(_serverErr, "Error serializing 'server' field")
		}

		// Simple Field (originalInvokeId)
		originalInvokeId := uint8(m.OriginalInvokeId)
		_originalInvokeIdErr := writeBuffer.WriteUint8("originalInvokeId", 8, (originalInvokeId))
		if _originalInvokeIdErr != nil {
			return errors.Wrap(_originalInvokeIdErr, "Error serializing 'originalInvokeId' field")
		}

		// Simple Field (sequenceNumber)
		sequenceNumber := uint8(m.SequenceNumber)
		_sequenceNumberErr := writeBuffer.WriteUint8("sequenceNumber", 8, (sequenceNumber))
		if _sequenceNumberErr != nil {
			return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
		}

		// Simple Field (proposedWindowSize)
		proposedWindowSize := uint8(m.ProposedWindowSize)
		_proposedWindowSizeErr := writeBuffer.WriteUint8("proposedWindowSize", 8, (proposedWindowSize))
		if _proposedWindowSizeErr != nil {
			return errors.Wrap(_proposedWindowSizeErr, "Error serializing 'proposedWindowSize' field")
		}

		if popErr := writeBuffer.PopContext("APDUSegmentAck"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *APDUSegmentAck) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
