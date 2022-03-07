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
type BACnetTagPayloadTime struct {
	Hour       uint8
	Minute     uint8
	Second     uint8
	Fractional uint8
}

// The corresponding interface
type IBACnetTagPayloadTime interface {
	// GetHour returns Hour
	GetHour() uint8
	// GetMinute returns Minute
	GetMinute() uint8
	// GetSecond returns Second
	GetSecond() uint8
	// GetFractional returns Fractional
	GetFractional() uint8
	// GetWildcard returns Wildcard
	GetWildcard() uint8
	// GetHourIsWildcard returns HourIsWildcard
	GetHourIsWildcard() bool
	// GetMinuteIsWildcard returns MinuteIsWildcard
	GetMinuteIsWildcard() bool
	// GetSecondIsWildcard returns SecondIsWildcard
	GetSecondIsWildcard() bool
	// GetFractionalIsWildcard returns FractionalIsWildcard
	GetFractionalIsWildcard() bool
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetTagPayloadTime) GetHour() uint8 {
	return m.Hour
}

func (m *BACnetTagPayloadTime) GetMinute() uint8 {
	return m.Minute
}

func (m *BACnetTagPayloadTime) GetSecond() uint8 {
	return m.Second
}

func (m *BACnetTagPayloadTime) GetFractional() uint8 {
	return m.Fractional
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////
func (m *BACnetTagPayloadTime) GetWildcard() uint8 {
	return 0xFF
}

func (m *BACnetTagPayloadTime) GetHourIsWildcard() bool {
	return bool((m.GetHour()) == (m.GetWildcard()))
}

func (m *BACnetTagPayloadTime) GetMinuteIsWildcard() bool {
	return bool((m.GetMinute()) == (m.GetWildcard()))
}

func (m *BACnetTagPayloadTime) GetSecondIsWildcard() bool {
	return bool((m.GetSecond()) == (m.GetWildcard()))
}

func (m *BACnetTagPayloadTime) GetFractionalIsWildcard() bool {
	return bool((m.GetFractional()) == (m.GetWildcard()))
}

// NewBACnetTagPayloadTime factory function for BACnetTagPayloadTime
func NewBACnetTagPayloadTime(hour uint8, minute uint8, second uint8, fractional uint8) *BACnetTagPayloadTime {
	return &BACnetTagPayloadTime{Hour: hour, Minute: minute, Second: second, Fractional: fractional}
}

func CastBACnetTagPayloadTime(structType interface{}) *BACnetTagPayloadTime {
	if casted, ok := structType.(BACnetTagPayloadTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetTagPayloadTime); ok {
		return casted
	}
	return nil
}

func (m *BACnetTagPayloadTime) GetTypeName() string {
	return "BACnetTagPayloadTime"
}

func (m *BACnetTagPayloadTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetTagPayloadTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// Simple field (hour)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (minute)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (second)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Simple field (fractional)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetTagPayloadTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTagPayloadTimeParse(readBuffer utils.ReadBuffer) (*BACnetTagPayloadTime, error) {
	if pullErr := readBuffer.PullContext("BACnetTagPayloadTime"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Virtual field
	_wildcard := 0xFF
	wildcard := uint8(_wildcard)
	_ = wildcard

	// Simple Field (hour)
	_hour, _hourErr := readBuffer.ReadUint8("hour", 8)
	if _hourErr != nil {
		return nil, errors.Wrap(_hourErr, "Error parsing 'hour' field")
	}
	hour := _hour

	// Virtual field
	_hourIsWildcard := bool((hour) == (wildcard))
	hourIsWildcard := bool(_hourIsWildcard)
	_ = hourIsWildcard

	// Simple Field (minute)
	_minute, _minuteErr := readBuffer.ReadUint8("minute", 8)
	if _minuteErr != nil {
		return nil, errors.Wrap(_minuteErr, "Error parsing 'minute' field")
	}
	minute := _minute

	// Virtual field
	_minuteIsWildcard := bool((minute) == (wildcard))
	minuteIsWildcard := bool(_minuteIsWildcard)
	_ = minuteIsWildcard

	// Simple Field (second)
	_second, _secondErr := readBuffer.ReadUint8("second", 8)
	if _secondErr != nil {
		return nil, errors.Wrap(_secondErr, "Error parsing 'second' field")
	}
	second := _second

	// Virtual field
	_secondIsWildcard := bool((second) == (wildcard))
	secondIsWildcard := bool(_secondIsWildcard)
	_ = secondIsWildcard

	// Simple Field (fractional)
	_fractional, _fractionalErr := readBuffer.ReadUint8("fractional", 8)
	if _fractionalErr != nil {
		return nil, errors.Wrap(_fractionalErr, "Error parsing 'fractional' field")
	}
	fractional := _fractional

	// Virtual field
	_fractionalIsWildcard := bool((fractional) == (wildcard))
	fractionalIsWildcard := bool(_fractionalIsWildcard)
	_ = fractionalIsWildcard

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadTime"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetTagPayloadTime(hour, minute, second, fractional), nil
}

func (m *BACnetTagPayloadTime) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadTime"); pushErr != nil {
		return pushErr
	}
	// Virtual field
	if _wildcardErr := writeBuffer.WriteVirtual("wildcard", m.GetWildcard()); _wildcardErr != nil {
		return errors.Wrap(_wildcardErr, "Error serializing 'wildcard' field")
	}

	// Simple Field (hour)
	hour := uint8(m.Hour)
	_hourErr := writeBuffer.WriteUint8("hour", 8, (hour))
	if _hourErr != nil {
		return errors.Wrap(_hourErr, "Error serializing 'hour' field")
	}
	// Virtual field
	if _hourIsWildcardErr := writeBuffer.WriteVirtual("hourIsWildcard", m.GetHourIsWildcard()); _hourIsWildcardErr != nil {
		return errors.Wrap(_hourIsWildcardErr, "Error serializing 'hourIsWildcard' field")
	}

	// Simple Field (minute)
	minute := uint8(m.Minute)
	_minuteErr := writeBuffer.WriteUint8("minute", 8, (minute))
	if _minuteErr != nil {
		return errors.Wrap(_minuteErr, "Error serializing 'minute' field")
	}
	// Virtual field
	if _minuteIsWildcardErr := writeBuffer.WriteVirtual("minuteIsWildcard", m.GetMinuteIsWildcard()); _minuteIsWildcardErr != nil {
		return errors.Wrap(_minuteIsWildcardErr, "Error serializing 'minuteIsWildcard' field")
	}

	// Simple Field (second)
	second := uint8(m.Second)
	_secondErr := writeBuffer.WriteUint8("second", 8, (second))
	if _secondErr != nil {
		return errors.Wrap(_secondErr, "Error serializing 'second' field")
	}
	// Virtual field
	if _secondIsWildcardErr := writeBuffer.WriteVirtual("secondIsWildcard", m.GetSecondIsWildcard()); _secondIsWildcardErr != nil {
		return errors.Wrap(_secondIsWildcardErr, "Error serializing 'secondIsWildcard' field")
	}

	// Simple Field (fractional)
	fractional := uint8(m.Fractional)
	_fractionalErr := writeBuffer.WriteUint8("fractional", 8, (fractional))
	if _fractionalErr != nil {
		return errors.Wrap(_fractionalErr, "Error serializing 'fractional' field")
	}
	// Virtual field
	if _fractionalIsWildcardErr := writeBuffer.WriteVirtual("fractionalIsWildcard", m.GetFractionalIsWildcard()); _fractionalIsWildcardErr != nil {
		return errors.Wrap(_fractionalIsWildcardErr, "Error serializing 'fractionalIsWildcard' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadTime"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetTagPayloadTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
