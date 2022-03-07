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
type CipWriteRequest struct {
	*CipService
	RequestPathSize int8
	Tag             []byte
	DataType        CIPDataTypeCode
	ElementNb       uint16
	Data            []byte

	// Arguments.
	ServiceLen uint16
}

// The corresponding interface
type ICipWriteRequest interface {
	ICipService
	// GetRequestPathSize returns RequestPathSize
	GetRequestPathSize() int8
	// GetTag returns Tag
	GetTag() []byte
	// GetDataType returns DataType
	GetDataType() CIPDataTypeCode
	// GetElementNb returns ElementNb
	GetElementNb() uint16
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
func (m *CipWriteRequest) Service() uint8 {
	return 0x4D
}

func (m *CipWriteRequest) GetService() uint8 {
	return 0x4D
}

func (m *CipWriteRequest) InitializeParent(parent *CipService) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *CipWriteRequest) GetRequestPathSize() int8 {
	return m.RequestPathSize
}

func (m *CipWriteRequest) GetTag() []byte {
	return m.Tag
}

func (m *CipWriteRequest) GetDataType() CIPDataTypeCode {
	return m.DataType
}

func (m *CipWriteRequest) GetElementNb() uint16 {
	return m.ElementNb
}

func (m *CipWriteRequest) GetData() []byte {
	return m.Data
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewCipWriteRequest factory function for CipWriteRequest
func NewCipWriteRequest(requestPathSize int8, tag []byte, dataType CIPDataTypeCode, elementNb uint16, data []byte, serviceLen uint16) *CipService {
	child := &CipWriteRequest{
		RequestPathSize: requestPathSize,
		Tag:             tag,
		DataType:        dataType,
		ElementNb:       elementNb,
		Data:            data,
		CipService:      NewCipService(serviceLen),
	}
	child.Child = child
	return child.CipService
}

func CastCipWriteRequest(structType interface{}) *CipWriteRequest {
	if casted, ok := structType.(CipWriteRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*CipWriteRequest); ok {
		return casted
	}
	if casted, ok := structType.(CipService); ok {
		return CastCipWriteRequest(casted.Child)
	}
	if casted, ok := structType.(*CipService); ok {
		return CastCipWriteRequest(casted.Child)
	}
	return nil
}

func (m *CipWriteRequest) GetTypeName() string {
	return "CipWriteRequest"
}

func (m *CipWriteRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CipWriteRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (requestPathSize)
	lengthInBits += 8

	// Array field
	if len(m.Tag) > 0 {
		lengthInBits += 8 * uint16(len(m.Tag))
	}

	// Simple field (dataType)
	lengthInBits += 16

	// Simple field (elementNb)
	lengthInBits += 16

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *CipWriteRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CipWriteRequestParse(readBuffer utils.ReadBuffer, serviceLen uint16) (*CipService, error) {
	if pullErr := readBuffer.PullContext("CipWriteRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (requestPathSize)
	_requestPathSize, _requestPathSizeErr := readBuffer.ReadInt8("requestPathSize", 8)
	if _requestPathSizeErr != nil {
		return nil, errors.Wrap(_requestPathSizeErr, "Error parsing 'requestPathSize' field")
	}
	requestPathSize := _requestPathSize
	// Byte Array field (tag)
	numberOfBytestag := int(uint16(requestPathSize) * uint16(uint16(2)))
	tag, _readArrayErr := readBuffer.ReadByteArray("tag", numberOfBytestag)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'tag' field")
	}

	// Simple Field (dataType)
	if pullErr := readBuffer.PullContext("dataType"); pullErr != nil {
		return nil, pullErr
	}
	_dataType, _dataTypeErr := CIPDataTypeCodeParse(readBuffer)
	if _dataTypeErr != nil {
		return nil, errors.Wrap(_dataTypeErr, "Error parsing 'dataType' field")
	}
	dataType := _dataType
	if closeErr := readBuffer.CloseContext("dataType"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (elementNb)
	_elementNb, _elementNbErr := readBuffer.ReadUint16("elementNb", 16)
	if _elementNbErr != nil {
		return nil, errors.Wrap(_elementNbErr, "Error parsing 'elementNb' field")
	}
	elementNb := _elementNb
	// Byte Array field (data)
	numberOfBytesdata := int(uint16(dataType.Size()) * uint16(elementNb))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	if closeErr := readBuffer.CloseContext("CipWriteRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CipWriteRequest{
		RequestPathSize: requestPathSize,
		Tag:             tag,
		DataType:        dataType,
		ElementNb:       elementNb,
		Data:            data,
		CipService:      &CipService{},
	}
	_child.CipService.Child = _child
	return _child.CipService, nil
}

func (m *CipWriteRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipWriteRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (requestPathSize)
		requestPathSize := int8(m.RequestPathSize)
		_requestPathSizeErr := writeBuffer.WriteInt8("requestPathSize", 8, (requestPathSize))
		if _requestPathSizeErr != nil {
			return errors.Wrap(_requestPathSizeErr, "Error serializing 'requestPathSize' field")
		}

		// Array Field (tag)
		if m.Tag != nil {
			// Byte Array field (tag)
			_writeArrayErr := writeBuffer.WriteByteArray("tag", m.Tag)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'tag' field")
			}
		}

		// Simple Field (dataType)
		if pushErr := writeBuffer.PushContext("dataType"); pushErr != nil {
			return pushErr
		}
		_dataTypeErr := m.DataType.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("dataType"); popErr != nil {
			return popErr
		}
		if _dataTypeErr != nil {
			return errors.Wrap(_dataTypeErr, "Error serializing 'dataType' field")
		}

		// Simple Field (elementNb)
		elementNb := uint16(m.ElementNb)
		_elementNbErr := writeBuffer.WriteUint16("elementNb", 16, (elementNb))
		if _elementNbErr != nil {
			return errors.Wrap(_elementNbErr, "Error serializing 'elementNb' field")
		}

		// Array Field (data)
		if m.Data != nil {
			// Byte Array field (data)
			_writeArrayErr := writeBuffer.WriteByteArray("data", m.Data)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
			}
		}

		if popErr := writeBuffer.PopContext("CipWriteRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CipWriteRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
