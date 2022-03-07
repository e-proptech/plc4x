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
type BACnetContextTagEventState struct {
	*BACnetContextTag
	EventState       BACnetEventState
	ProprietaryValue uint32

	// Arguments.
	TagNumberArgument        uint8
	IsNotOpeningOrClosingTag bool
	ActualLength             uint32
}

// The corresponding interface
type IBACnetContextTagEventState interface {
	IBACnetContextTag
	// GetEventState returns EventState
	GetEventState() BACnetEventState
	// GetProprietaryValue returns ProprietaryValue
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary
	GetIsProprietary() bool
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
func (m *BACnetContextTagEventState) DataType() BACnetDataType {
	return BACnetDataType_EVENT_STATE
}

func (m *BACnetContextTagEventState) GetDataType() BACnetDataType {
	return BACnetDataType_EVENT_STATE
}

func (m *BACnetContextTagEventState) InitializeParent(parent *BACnetContextTag, header *BACnetTagHeader) {
	m.BACnetContextTag.Header = header
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetContextTagEventState) GetEventState() BACnetEventState {
	return m.EventState
}

func (m *BACnetContextTagEventState) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////
func (m *BACnetContextTagEventState) GetIsProprietary() bool {
	return bool((m.GetEventState()) == (BACnetEventState_VENDOR_PROPRIETARY_VALUE))
}

// NewBACnetContextTagEventState factory function for BACnetContextTagEventState
func NewBACnetContextTagEventState(eventState BACnetEventState, proprietaryValue uint32, header *BACnetTagHeader, tagNumberArgument uint8, isNotOpeningOrClosingTag bool, actualLength uint32) *BACnetContextTag {
	child := &BACnetContextTagEventState{
		EventState:       eventState,
		ProprietaryValue: proprietaryValue,
		BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	child.Child = child
	return child.BACnetContextTag
}

func CastBACnetContextTagEventState(structType interface{}) *BACnetContextTagEventState {
	if casted, ok := structType.(BACnetContextTagEventState); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetContextTagEventState); ok {
		return casted
	}
	if casted, ok := structType.(BACnetContextTag); ok {
		return CastBACnetContextTagEventState(casted.Child)
	}
	if casted, ok := structType.(*BACnetContextTag); ok {
		return CastBACnetContextTagEventState(casted.Child)
	}
	return nil
}

func (m *BACnetContextTagEventState) GetTypeName() string {
	return "BACnetContextTagEventState"
}

func (m *BACnetContextTagEventState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetContextTagEventState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Manual Field (eventState)
	lengthInBits += uint16(int32(m.GetActualLength()) * int32(int32(8)))

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(int32(0))

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetContextTagEventState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagEventStateParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType, isNotOpeningOrClosingTag bool, actualLength uint32) (*BACnetContextTag, error) {
	if pullErr := readBuffer.PullContext("BACnetContextTagEventState"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Validation
	if !(isNotOpeningOrClosingTag) {
		return nil, utils.ParseAssertError{"length 6 and 7 reserved for opening and closing tag"}
	}

	// Manual Field (eventState)
	eventState, _eventStateErr := ReadEventState(readBuffer, actualLength)
	if _eventStateErr != nil {
		return nil, errors.Wrap(_eventStateErr, "Error parsing 'eventState' field")
	}

	// Manual Field (proprietaryValue)
	proprietaryValue, _proprietaryValueErr := ReadProprietaryEventState(readBuffer, eventState, actualLength)
	if _proprietaryValueErr != nil {
		return nil, errors.Wrap(_proprietaryValueErr, "Error parsing 'proprietaryValue' field")
	}

	// Virtual field
	_isProprietary := bool((eventState) == (BACnetEventState_VENDOR_PROPRIETARY_VALUE))
	isProprietary := bool(_isProprietary)
	_ = isProprietary

	if closeErr := readBuffer.CloseContext("BACnetContextTagEventState"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetContextTagEventState{
		EventState:       eventState,
		ProprietaryValue: proprietaryValue,
		BACnetContextTag: &BACnetContextTag{},
	}
	_child.BACnetContextTag.Child = _child
	return _child.BACnetContextTag, nil
}

func (m *BACnetContextTagEventState) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagEventState"); pushErr != nil {
			return pushErr
		}

		// Manual Field (eventState)
		_eventStateErr := WriteEventState(writeBuffer, m.GetEventState())
		if _eventStateErr != nil {
			return errors.Wrap(_eventStateErr, "Error serializing 'eventState' field")
		}

		// Manual Field (proprietaryValue)
		_proprietaryValueErr := WriteProprietaryEventState(writeBuffer, m.GetEventState(), m.GetProprietaryValue())
		if _proprietaryValueErr != nil {
			return errors.Wrap(_proprietaryValueErr, "Error serializing 'proprietaryValue' field")
		}
		// Virtual field
		if _isProprietaryErr := writeBuffer.WriteVirtual("isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
			return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagEventState"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetContextTagEventState) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
