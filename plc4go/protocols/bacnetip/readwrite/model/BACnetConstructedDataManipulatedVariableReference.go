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

// BACnetConstructedDataManipulatedVariableReference is the data-structure of this message
type BACnetConstructedDataManipulatedVariableReference struct {
	*BACnetConstructedData
	ManipulatedVariableReference *BACnetObjectPropertyReference

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataManipulatedVariableReference is the corresponding interface of BACnetConstructedDataManipulatedVariableReference
type IBACnetConstructedDataManipulatedVariableReference interface {
	IBACnetConstructedData
	// GetManipulatedVariableReference returns ManipulatedVariableReference (property field)
	GetManipulatedVariableReference() *BACnetObjectPropertyReference
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

func (m *BACnetConstructedDataManipulatedVariableReference) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataManipulatedVariableReference) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MANIPULATED_VARIABLE_REFERENCE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataManipulatedVariableReference) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataManipulatedVariableReference) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataManipulatedVariableReference) GetManipulatedVariableReference() *BACnetObjectPropertyReference {
	return m.ManipulatedVariableReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataManipulatedVariableReference factory function for BACnetConstructedDataManipulatedVariableReference
func NewBACnetConstructedDataManipulatedVariableReference(manipulatedVariableReference *BACnetObjectPropertyReference, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataManipulatedVariableReference {
	_result := &BACnetConstructedDataManipulatedVariableReference{
		ManipulatedVariableReference: manipulatedVariableReference,
		BACnetConstructedData:        NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataManipulatedVariableReference(structType interface{}) *BACnetConstructedDataManipulatedVariableReference {
	if casted, ok := structType.(BACnetConstructedDataManipulatedVariableReference); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataManipulatedVariableReference); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataManipulatedVariableReference(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataManipulatedVariableReference(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataManipulatedVariableReference) GetTypeName() string {
	return "BACnetConstructedDataManipulatedVariableReference"
}

func (m *BACnetConstructedDataManipulatedVariableReference) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataManipulatedVariableReference) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (manipulatedVariableReference)
	lengthInBits += m.ManipulatedVariableReference.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataManipulatedVariableReference) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataManipulatedVariableReferenceParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataManipulatedVariableReference, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataManipulatedVariableReference"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (manipulatedVariableReference)
	if pullErr := readBuffer.PullContext("manipulatedVariableReference"); pullErr != nil {
		return nil, pullErr
	}
	_manipulatedVariableReference, _manipulatedVariableReferenceErr := BACnetObjectPropertyReferenceParse(readBuffer)
	if _manipulatedVariableReferenceErr != nil {
		return nil, errors.Wrap(_manipulatedVariableReferenceErr, "Error parsing 'manipulatedVariableReference' field")
	}
	manipulatedVariableReference := CastBACnetObjectPropertyReference(_manipulatedVariableReference)
	if closeErr := readBuffer.CloseContext("manipulatedVariableReference"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataManipulatedVariableReference"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataManipulatedVariableReference{
		ManipulatedVariableReference: CastBACnetObjectPropertyReference(manipulatedVariableReference),
		BACnetConstructedData:        &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataManipulatedVariableReference) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataManipulatedVariableReference"); pushErr != nil {
			return pushErr
		}

		// Simple Field (manipulatedVariableReference)
		if pushErr := writeBuffer.PushContext("manipulatedVariableReference"); pushErr != nil {
			return pushErr
		}
		_manipulatedVariableReferenceErr := m.ManipulatedVariableReference.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("manipulatedVariableReference"); popErr != nil {
			return popErr
		}
		if _manipulatedVariableReferenceErr != nil {
			return errors.Wrap(_manipulatedVariableReferenceErr, "Error serializing 'manipulatedVariableReference' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataManipulatedVariableReference"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataManipulatedVariableReference) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
