/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// APDUReject is the corresponding interface of APDUReject
type APDUReject interface {
	utils.LengthAware
	utils.Serializable
	APDU
	// GetOriginalInvokeId returns OriginalInvokeId (property field)
	GetOriginalInvokeId() uint8
	// GetRejectReason returns RejectReason (property field)
	GetRejectReason() BACnetRejectReasonTagged
}

// APDURejectExactly can be used when we want exactly this type and not a type which fulfills APDUReject.
// This is useful for switch cases.
type APDURejectExactly interface {
	APDUReject
	isAPDUReject() bool
}

// _APDUReject is the data-structure of this message
type _APDUReject struct {
	*_APDU
	OriginalInvokeId uint8
	RejectReason     BACnetRejectReasonTagged
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUReject) GetApduType() ApduType {
	return ApduType_REJECT_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUReject) InitializeParent(parent APDU) {}

func (m *_APDUReject) GetParent() APDU {
	return m._APDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUReject) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *_APDUReject) GetRejectReason() BACnetRejectReasonTagged {
	return m.RejectReason
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAPDUReject factory function for _APDUReject
func NewAPDUReject(originalInvokeId uint8, rejectReason BACnetRejectReasonTagged, apduLength uint16) *_APDUReject {
	_result := &_APDUReject{
		OriginalInvokeId: originalInvokeId,
		RejectReason:     rejectReason,
		_APDU:            NewAPDU(apduLength),
	}
	_result._APDU._APDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAPDUReject(structType interface{}) APDUReject {
	if casted, ok := structType.(APDUReject); ok {
		return casted
	}
	if casted, ok := structType.(*APDUReject); ok {
		return *casted
	}
	return nil
}

func (m *_APDUReject) GetTypeName() string {
	return "APDUReject"
}

func (m *_APDUReject) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (originalInvokeId)
	lengthInBits += 8

	// Simple field (rejectReason)
	lengthInBits += m.RejectReason.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_APDUReject) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func APDURejectParse(theBytes []byte, apduLength uint16) (APDUReject, error) {
	return APDURejectParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), apduLength)
}

func APDURejectParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (APDUReject, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUReject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUReject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 4)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of APDUReject")
		}
		if reserved != uint8(0x00) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (originalInvokeId)
	_originalInvokeId, _originalInvokeIdErr := readBuffer.ReadUint8("originalInvokeId", 8)
	if _originalInvokeIdErr != nil {
		return nil, errors.Wrap(_originalInvokeIdErr, "Error parsing 'originalInvokeId' field of APDUReject")
	}
	originalInvokeId := _originalInvokeId

	// Simple Field (rejectReason)
	if pullErr := readBuffer.PullContext("rejectReason"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for rejectReason")
	}
	_rejectReason, _rejectReasonErr := BACnetRejectReasonTaggedParseWithBuffer(ctx, readBuffer, uint32(uint32(1)))
	if _rejectReasonErr != nil {
		return nil, errors.Wrap(_rejectReasonErr, "Error parsing 'rejectReason' field of APDUReject")
	}
	rejectReason := _rejectReason.(BACnetRejectReasonTagged)
	if closeErr := readBuffer.CloseContext("rejectReason"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for rejectReason")
	}

	if closeErr := readBuffer.CloseContext("APDUReject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUReject")
	}

	// Create a partially initialized instance
	_child := &_APDUReject{
		_APDU: &_APDU{
			ApduLength: apduLength,
		},
		OriginalInvokeId: originalInvokeId,
		RejectReason:     rejectReason,
		reservedField0:   reservedField0,
	}
	_child._APDU._APDUChildRequirements = _child
	return _child, nil
}

func (m *_APDUReject) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_APDUReject) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUReject"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUReject")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 4, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (originalInvokeId)
		originalInvokeId := uint8(m.GetOriginalInvokeId())
		_originalInvokeIdErr := writeBuffer.WriteUint8("originalInvokeId", 8, (originalInvokeId))
		if _originalInvokeIdErr != nil {
			return errors.Wrap(_originalInvokeIdErr, "Error serializing 'originalInvokeId' field")
		}

		// Simple Field (rejectReason)
		if pushErr := writeBuffer.PushContext("rejectReason"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for rejectReason")
		}
		_rejectReasonErr := writeBuffer.WriteSerializable(ctx, m.GetRejectReason())
		if popErr := writeBuffer.PopContext("rejectReason"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for rejectReason")
		}
		if _rejectReasonErr != nil {
			return errors.Wrap(_rejectReasonErr, "Error serializing 'rejectReason' field")
		}

		if popErr := writeBuffer.PopContext("APDUReject"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUReject")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_APDUReject) isAPDUReject() bool {
	return true
}

func (m *_APDUReject) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
