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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataAckedTransitions is the data-structure of this message
type BACnetConstructedDataAckedTransitions struct {
	*BACnetConstructedData
	AckedTransitions *BACnetEventTransitionBitsTagged

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataAckedTransitions is the corresponding interface of BACnetConstructedDataAckedTransitions
type IBACnetConstructedDataAckedTransitions interface {
	IBACnetConstructedData
	// GetAckedTransitions returns AckedTransitions (property field)
	GetAckedTransitions() *BACnetEventTransitionBitsTagged
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataAckedTransitions) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataAckedTransitions) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACKED_TRANSITIONS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAckedTransitions) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAckedTransitions) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAckedTransitions) GetAckedTransitions() *BACnetEventTransitionBitsTagged {
	return m.AckedTransitions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAckedTransitions factory function for BACnetConstructedDataAckedTransitions
func NewBACnetConstructedDataAckedTransitions(ackedTransitions *BACnetEventTransitionBitsTagged, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataAckedTransitions {
	_result := &BACnetConstructedDataAckedTransitions{
		AckedTransitions:      ackedTransitions,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAckedTransitions(structType interface{}) *BACnetConstructedDataAckedTransitions {
	if casted, ok := structType.(BACnetConstructedDataAckedTransitions); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAckedTransitions); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAckedTransitions(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAckedTransitions(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAckedTransitions) GetTypeName() string {
	return "BACnetConstructedDataAckedTransitions"
}

func (m *BACnetConstructedDataAckedTransitions) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAckedTransitions) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ackedTransitions)
	lengthInBits += m.AckedTransitions.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataAckedTransitions) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAckedTransitionsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataAckedTransitions, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAckedTransitions"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ackedTransitions)
	if pullErr := readBuffer.PullContext("ackedTransitions"); pullErr != nil {
		return nil, pullErr
	}
	_ackedTransitions, _ackedTransitionsErr := BACnetEventTransitionBitsTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _ackedTransitionsErr != nil {
		return nil, errors.Wrap(_ackedTransitionsErr, "Error parsing 'ackedTransitions' field")
	}
	ackedTransitions := CastBACnetEventTransitionBitsTagged(_ackedTransitions)
	if closeErr := readBuffer.CloseContext("ackedTransitions"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAckedTransitions"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAckedTransitions{
		AckedTransitions:      CastBACnetEventTransitionBitsTagged(ackedTransitions),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAckedTransitions) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAckedTransitions"); pushErr != nil {
			return pushErr
		}

		// Simple Field (ackedTransitions)
		if pushErr := writeBuffer.PushContext("ackedTransitions"); pushErr != nil {
			return pushErr
		}
		_ackedTransitionsErr := m.AckedTransitions.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("ackedTransitions"); popErr != nil {
			return popErr
		}
		if _ackedTransitionsErr != nil {
			return errors.Wrap(_ackedTransitionsErr, "Error serializing 'ackedTransitions' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAckedTransitions"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAckedTransitions) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
