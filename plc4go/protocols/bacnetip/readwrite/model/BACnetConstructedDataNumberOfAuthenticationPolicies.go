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

// BACnetConstructedDataNumberOfAuthenticationPolicies is the corresponding interface of BACnetConstructedDataNumberOfAuthenticationPolicies
type BACnetConstructedDataNumberOfAuthenticationPolicies interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfAuthenticationPolicies returns NumberOfAuthenticationPolicies (property field)
	GetNumberOfAuthenticationPolicies() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataNumberOfAuthenticationPoliciesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataNumberOfAuthenticationPolicies.
// This is useful for switch cases.
type BACnetConstructedDataNumberOfAuthenticationPoliciesExactly interface {
	BACnetConstructedDataNumberOfAuthenticationPolicies
	isBACnetConstructedDataNumberOfAuthenticationPolicies() bool
}

// _BACnetConstructedDataNumberOfAuthenticationPolicies is the data-structure of this message
type _BACnetConstructedDataNumberOfAuthenticationPolicies struct {
	*_BACnetConstructedData
	NumberOfAuthenticationPolicies BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNumberOfAuthenticationPolicies) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNumberOfAuthenticationPolicies) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NUMBER_OF_AUTHENTICATION_POLICIES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNumberOfAuthenticationPolicies) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataNumberOfAuthenticationPolicies) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNumberOfAuthenticationPolicies) GetNumberOfAuthenticationPolicies() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfAuthenticationPolicies
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNumberOfAuthenticationPolicies) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetNumberOfAuthenticationPolicies())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataNumberOfAuthenticationPolicies factory function for _BACnetConstructedDataNumberOfAuthenticationPolicies
func NewBACnetConstructedDataNumberOfAuthenticationPolicies(numberOfAuthenticationPolicies BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNumberOfAuthenticationPolicies {
	_result := &_BACnetConstructedDataNumberOfAuthenticationPolicies{
		NumberOfAuthenticationPolicies: numberOfAuthenticationPolicies,
		_BACnetConstructedData:         NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNumberOfAuthenticationPolicies(structType interface{}) BACnetConstructedDataNumberOfAuthenticationPolicies {
	if casted, ok := structType.(BACnetConstructedDataNumberOfAuthenticationPolicies); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNumberOfAuthenticationPolicies); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNumberOfAuthenticationPolicies) GetTypeName() string {
	return "BACnetConstructedDataNumberOfAuthenticationPolicies"
}

func (m *_BACnetConstructedDataNumberOfAuthenticationPolicies) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (numberOfAuthenticationPolicies)
	lengthInBits += m.NumberOfAuthenticationPolicies.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNumberOfAuthenticationPolicies) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataNumberOfAuthenticationPoliciesParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataNumberOfAuthenticationPolicies, error) {
	return BACnetConstructedDataNumberOfAuthenticationPoliciesParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataNumberOfAuthenticationPoliciesParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataNumberOfAuthenticationPolicies, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNumberOfAuthenticationPolicies"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNumberOfAuthenticationPolicies")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (numberOfAuthenticationPolicies)
	if pullErr := readBuffer.PullContext("numberOfAuthenticationPolicies"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for numberOfAuthenticationPolicies")
	}
	_numberOfAuthenticationPolicies, _numberOfAuthenticationPoliciesErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _numberOfAuthenticationPoliciesErr != nil {
		return nil, errors.Wrap(_numberOfAuthenticationPoliciesErr, "Error parsing 'numberOfAuthenticationPolicies' field of BACnetConstructedDataNumberOfAuthenticationPolicies")
	}
	numberOfAuthenticationPolicies := _numberOfAuthenticationPolicies.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("numberOfAuthenticationPolicies"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for numberOfAuthenticationPolicies")
	}

	// Virtual field
	_actualValue := numberOfAuthenticationPolicies
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNumberOfAuthenticationPolicies"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNumberOfAuthenticationPolicies")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataNumberOfAuthenticationPolicies{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NumberOfAuthenticationPolicies: numberOfAuthenticationPolicies,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataNumberOfAuthenticationPolicies) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNumberOfAuthenticationPolicies) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNumberOfAuthenticationPolicies"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNumberOfAuthenticationPolicies")
		}

		// Simple Field (numberOfAuthenticationPolicies)
		if pushErr := writeBuffer.PushContext("numberOfAuthenticationPolicies"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for numberOfAuthenticationPolicies")
		}
		_numberOfAuthenticationPoliciesErr := writeBuffer.WriteSerializable(ctx, m.GetNumberOfAuthenticationPolicies())
		if popErr := writeBuffer.PopContext("numberOfAuthenticationPolicies"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for numberOfAuthenticationPolicies")
		}
		if _numberOfAuthenticationPoliciesErr != nil {
			return errors.Wrap(_numberOfAuthenticationPoliciesErr, "Error serializing 'numberOfAuthenticationPolicies' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNumberOfAuthenticationPolicies"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNumberOfAuthenticationPolicies")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNumberOfAuthenticationPolicies) isBACnetConstructedDataNumberOfAuthenticationPolicies() bool {
	return true
}

func (m *_BACnetConstructedDataNumberOfAuthenticationPolicies) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
