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

// BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged is the corresponding interface of BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged
type BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter
}

// BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTaggedExactly interface {
	BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged
	isBACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged() bool
}

// _BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged is the data-structure of this message
type _BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged struct {
	Header BACnetTagHeader
	Value  BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged) GetValue() BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged factory function for _BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged
func NewBACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged(header BACnetTagHeader, value BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter, tagNumber uint8, tagClass TagClass) *_BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged {
	return &_BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged(structType interface{}) BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged {
	if casted, ok := structType.(BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged) GetTypeName() string {
	return "BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged"
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTaggedParse(theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged, error) {
	return BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTaggedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Manual Field (value)
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter_OFFNORMAL)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged")
	}
	var value BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter
	if _value != nil {
		value = _value.(BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilter)
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged")
	}

	// Create the instance
	return &_BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged{
		TagNumber: tagNumber,
		TagClass:  tagClass,
		Header:    header,
		Value:     value,
	}, nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(ctx, m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged) isBACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestGetEnrollmentSummaryEventStateFilterTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
