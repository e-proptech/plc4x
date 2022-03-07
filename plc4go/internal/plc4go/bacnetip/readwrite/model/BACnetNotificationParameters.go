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
type BACnetNotificationParameters struct {
	OpeningTag      *BACnetOpeningTag
	PeekedTagHeader *BACnetTagHeader
	ClosingTag      *BACnetClosingTag

	// Arguments.
	TagNumber  uint8
	ObjectType BACnetObjectType
	Child      IBACnetNotificationParametersChild
}

// The corresponding interface
type IBACnetNotificationParameters interface {
	// PeekedTagNumber returns PeekedTagNumber
	PeekedTagNumber() uint8
	// GetOpeningTag returns OpeningTag
	GetOpeningTag() *BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader
	GetPeekedTagHeader() *BACnetTagHeader
	// GetClosingTag returns ClosingTag
	GetClosingTag() *BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber
	GetPeekedTagNumber() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetNotificationParametersParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetNotificationParameters, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetNotificationParametersChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetNotificationParameters, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag)
	GetTypeName() string
	IBACnetNotificationParameters
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetNotificationParameters) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetNotificationParameters) GetPeekedTagHeader() *BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *BACnetNotificationParameters) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////
func (m *BACnetNotificationParameters) GetPeekedTagNumber() uint8 {
	return m.GetPeekedTagHeader().GetActualTagNumber()
}

// NewBACnetNotificationParameters factory function for BACnetNotificationParameters
func NewBACnetNotificationParameters(openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, objectType BACnetObjectType) *BACnetNotificationParameters {
	return &BACnetNotificationParameters{OpeningTag: openingTag, PeekedTagHeader: peekedTagHeader, ClosingTag: closingTag, TagNumber: tagNumber, ObjectType: objectType}
}

func CastBACnetNotificationParameters(structType interface{}) *BACnetNotificationParameters {
	if casted, ok := structType.(BACnetNotificationParameters); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetNotificationParameters); ok {
		return casted
	}
	return nil
}

func (m *BACnetNotificationParameters) GetTypeName() string {
	return "BACnetNotificationParameters"
}

func (m *BACnetNotificationParameters) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNotificationParameters) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetNotificationParameters) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParameters) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectType BACnetObjectType) (*BACnetNotificationParameters, error) {
	if pullErr := readBuffer.PullContext("BACnetNotificationParameters"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetContextTagParse(readBuffer, uint8(tagNumber), BACnetDataType(BACnetDataType_OPENING_TAG))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Peek Field (peekedTagHeader)
	currentPos = readBuffer.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, pullErr
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *BACnetNotificationParameters
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetNotificationParametersChangeOfBitString
		_parent, typeSwitchError = BACnetNotificationParametersChangeOfBitStringParse(readBuffer, tagNumber, objectType, peekedTagNumber)
	case peekedTagNumber == uint8(1): // BACnetNotificationParametersChangeOfState
		_parent, typeSwitchError = BACnetNotificationParametersChangeOfStateParse(readBuffer, tagNumber, objectType, peekedTagNumber)
	case peekedTagNumber == uint8(2): // BACnetNotificationParametersChangeOfValue
		_parent, typeSwitchError = BACnetNotificationParametersChangeOfValueParse(readBuffer, tagNumber, objectType, peekedTagNumber)
	case peekedTagNumber == uint8(4): // BACnetNotificationParametersFloatingLimit
		_parent, typeSwitchError = BACnetNotificationParametersFloatingLimitParse(readBuffer, tagNumber, objectType, peekedTagNumber)
	case peekedTagNumber == uint8(5): // BACnetNotificationParametersOutOfRange
		_parent, typeSwitchError = BACnetNotificationParametersOutOfRangeParse(readBuffer, tagNumber, objectType, peekedTagNumber)
	case peekedTagNumber == uint8(6): // BACnetNotificationParametersComplexEventType
		_parent, typeSwitchError = BACnetNotificationParametersComplexEventTypeParse(readBuffer, tagNumber, objectType, peekedTagNumber)
	case peekedTagNumber == uint8(9): // BACnetNotificationParametersExtended
		_parent, typeSwitchError = BACnetNotificationParametersExtendedParse(readBuffer, tagNumber, objectType, peekedTagNumber)
	case peekedTagNumber == uint8(10): // BACnetNotificationParametersBufferReady
		_parent, typeSwitchError = BACnetNotificationParametersBufferReadyParse(readBuffer, tagNumber, objectType, peekedTagNumber)
	case peekedTagNumber == uint8(11): // BACnetNotificationParametersUnsignedRange
		_parent, typeSwitchError = BACnetNotificationParametersUnsignedRangeParse(readBuffer, tagNumber, objectType, peekedTagNumber)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetContextTagParse(readBuffer, uint8(tagNumber), BACnetDataType(BACnetDataType_CLOSING_TAG))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParameters"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, openingTag, peekedTagHeader, closingTag)
	return _parent, nil
}

func (m *BACnetNotificationParameters) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetNotificationParameters) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetNotificationParameters, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetNotificationParameters"); pushErr != nil {
		return pushErr
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return pushErr
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return popErr
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return pushErr
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return popErr
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetNotificationParameters"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetNotificationParameters) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
