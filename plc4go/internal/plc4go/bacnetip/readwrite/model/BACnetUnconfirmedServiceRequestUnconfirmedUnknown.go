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
type BACnetUnconfirmedServiceRequestUnconfirmedUnknown struct {
	*BACnetUnconfirmedServiceRequest
	UnknownBytes []byte

	// Arguments.
	Len uint16
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequestUnconfirmedUnknown interface {
	IBACnetUnconfirmedServiceRequest
	// GetUnknownBytes returns UnknownBytes
	GetUnknownBytes() []byte
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
func (m *BACnetUnconfirmedServiceRequestUnconfirmedUnknown) ServiceChoice() uint8 {
	return 0
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedUnknown) GetServiceChoice() uint8 {
	return 0
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedUnknown) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetUnconfirmedServiceRequestUnconfirmedUnknown) GetUnknownBytes() []byte {
	return m.UnknownBytes
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestUnconfirmedUnknown factory function for BACnetUnconfirmedServiceRequestUnconfirmedUnknown
func NewBACnetUnconfirmedServiceRequestUnconfirmedUnknown(unknownBytes []byte, len uint16) *BACnetUnconfirmedServiceRequest {
	child := &BACnetUnconfirmedServiceRequestUnconfirmedUnknown{
		UnknownBytes:                    unknownBytes,
		BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(len),
	}
	child.Child = child
	return child.BACnetUnconfirmedServiceRequest
}

func CastBACnetUnconfirmedServiceRequestUnconfirmedUnknown(structType interface{}) *BACnetUnconfirmedServiceRequestUnconfirmedUnknown {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestUnconfirmedUnknown); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestUnconfirmedUnknown); ok {
		return casted
	}
	if casted, ok := structType.(BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestUnconfirmedUnknown(casted.Child)
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestUnconfirmedUnknown(casted.Child)
	}
	return nil
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedUnknown) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnconfirmedUnknown"
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedUnknown) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedUnknown) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.UnknownBytes) > 0 {
		lengthInBits += 8 * uint16(len(m.UnknownBytes))
	}

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedUnknown) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestUnconfirmedUnknownParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetUnconfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUnconfirmedUnknown"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos
	// Byte Array field (unknownBytes)
	numberOfBytesunknownBytes := int(utils.InlineIf(bool(bool((len) > (0))), func() interface{} { return uint16(uint16(uint16(len) - uint16(uint16(1)))) }, func() interface{} { return uint16(uint16(0)) }).(uint16))
	unknownBytes, _readArrayErr := readBuffer.ReadByteArray("unknownBytes", numberOfBytesunknownBytes)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'unknownBytes' field")
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUnconfirmedUnknown"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestUnconfirmedUnknown{
		UnknownBytes:                    unknownBytes,
		BACnetUnconfirmedServiceRequest: &BACnetUnconfirmedServiceRequest{},
	}
	_child.BACnetUnconfirmedServiceRequest.Child = _child
	return _child.BACnetUnconfirmedServiceRequest, nil
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedUnknown) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUnconfirmedUnknown"); pushErr != nil {
			return pushErr
		}

		// Array Field (unknownBytes)
		if m.UnknownBytes != nil {
			// Byte Array field (unknownBytes)
			_writeArrayErr := writeBuffer.WriteByteArray("unknownBytes", m.UnknownBytes)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'unknownBytes' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUnconfirmedUnknown"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedUnknown) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
