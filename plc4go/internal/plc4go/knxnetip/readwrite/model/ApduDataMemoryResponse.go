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
type ApduDataMemoryResponse struct {
	*ApduData
	Address uint16
	Data    []byte

	// Arguments.
	DataLength uint8
}

// The corresponding interface
type IApduDataMemoryResponse interface {
	IApduData
	// GetAddress returns Address
	GetAddress() uint16
	// GetData returns Data
	GetData() []byte
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
func (m *ApduDataMemoryResponse) ApciType() uint8 {
	return 0x9
}

func (m *ApduDataMemoryResponse) GetApciType() uint8 {
	return 0x9
}

func (m *ApduDataMemoryResponse) InitializeParent(parent *ApduData) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ApduDataMemoryResponse) GetAddress() uint16 {
	return m.Address
}

func (m *ApduDataMemoryResponse) GetData() []byte {
	return m.Data
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewApduDataMemoryResponse factory function for ApduDataMemoryResponse
func NewApduDataMemoryResponse(address uint16, data []byte, dataLength uint8) *ApduData {
	child := &ApduDataMemoryResponse{
		Address:  address,
		Data:     data,
		ApduData: NewApduData(dataLength),
	}
	child.Child = child
	return child.ApduData
}

func CastApduDataMemoryResponse(structType interface{}) *ApduDataMemoryResponse {
	if casted, ok := structType.(ApduDataMemoryResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataMemoryResponse); ok {
		return casted
	}
	if casted, ok := structType.(ApduData); ok {
		return CastApduDataMemoryResponse(casted.Child)
	}
	if casted, ok := structType.(*ApduData); ok {
		return CastApduDataMemoryResponse(casted.Child)
	}
	return nil
}

func (m *ApduDataMemoryResponse) GetTypeName() string {
	return "ApduDataMemoryResponse"
}

func (m *ApduDataMemoryResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataMemoryResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Implicit Field (numBytes)
	lengthInBits += 6

	// Simple field (address)
	lengthInBits += 16

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *ApduDataMemoryResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataMemoryResponseParse(readBuffer utils.ReadBuffer, dataLength uint8) (*ApduData, error) {
	if pullErr := readBuffer.PullContext("ApduDataMemoryResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Implicit Field (numBytes) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	numBytes, _numBytesErr := readBuffer.ReadUint8("numBytes", 6)
	_ = numBytes
	if _numBytesErr != nil {
		return nil, errors.Wrap(_numBytesErr, "Error parsing 'numBytes' field")
	}

	// Simple Field (address)
	_address, _addressErr := readBuffer.ReadUint16("address", 16)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field")
	}
	address := _address
	// Byte Array field (data)
	numberOfBytesdata := int(numBytes)
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	if closeErr := readBuffer.CloseContext("ApduDataMemoryResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataMemoryResponse{
		Address:  address,
		Data:     data,
		ApduData: &ApduData{},
	}
	_child.ApduData.Child = _child
	return _child.ApduData, nil
}

func (m *ApduDataMemoryResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataMemoryResponse"); pushErr != nil {
			return pushErr
		}

		// Implicit Field (numBytes) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		numBytes := uint8(uint8(len(m.GetData())))
		_numBytesErr := writeBuffer.WriteUint8("numBytes", 6, (numBytes))
		if _numBytesErr != nil {
			return errors.Wrap(_numBytesErr, "Error serializing 'numBytes' field")
		}

		// Simple Field (address)
		address := uint16(m.Address)
		_addressErr := writeBuffer.WriteUint16("address", 16, (address))
		if _addressErr != nil {
			return errors.Wrap(_addressErr, "Error serializing 'address' field")
		}

		// Array Field (data)
		if m.Data != nil {
			// Byte Array field (data)
			_writeArrayErr := writeBuffer.WriteByteArray("data", m.Data)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
			}
		}

		if popErr := writeBuffer.PopContext("ApduDataMemoryResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataMemoryResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
