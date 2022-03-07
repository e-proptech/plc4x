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
type ComObjectTableRealisationType6 struct {
	*ComObjectTable
	ComObjectDescriptors *GroupObjectDescriptorRealisationType6
}

// The corresponding interface
type IComObjectTableRealisationType6 interface {
	IComObjectTable
	// GetComObjectDescriptors returns ComObjectDescriptors
	GetComObjectDescriptors() *GroupObjectDescriptorRealisationType6
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
func (m *ComObjectTableRealisationType6) FirmwareType() FirmwareType {
	return FirmwareType_SYSTEM_300
}

func (m *ComObjectTableRealisationType6) GetFirmwareType() FirmwareType {
	return FirmwareType_SYSTEM_300
}

func (m *ComObjectTableRealisationType6) InitializeParent(parent *ComObjectTable) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ComObjectTableRealisationType6) GetComObjectDescriptors() *GroupObjectDescriptorRealisationType6 {
	return m.ComObjectDescriptors
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewComObjectTableRealisationType6 factory function for ComObjectTableRealisationType6
func NewComObjectTableRealisationType6(comObjectDescriptors *GroupObjectDescriptorRealisationType6) *ComObjectTable {
	child := &ComObjectTableRealisationType6{
		ComObjectDescriptors: comObjectDescriptors,
		ComObjectTable:       NewComObjectTable(),
	}
	child.Child = child
	return child.ComObjectTable
}

func CastComObjectTableRealisationType6(structType interface{}) *ComObjectTableRealisationType6 {
	if casted, ok := structType.(ComObjectTableRealisationType6); ok {
		return &casted
	}
	if casted, ok := structType.(*ComObjectTableRealisationType6); ok {
		return casted
	}
	if casted, ok := structType.(ComObjectTable); ok {
		return CastComObjectTableRealisationType6(casted.Child)
	}
	if casted, ok := structType.(*ComObjectTable); ok {
		return CastComObjectTableRealisationType6(casted.Child)
	}
	return nil
}

func (m *ComObjectTableRealisationType6) GetTypeName() string {
	return "ComObjectTableRealisationType6"
}

func (m *ComObjectTableRealisationType6) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ComObjectTableRealisationType6) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (comObjectDescriptors)
	lengthInBits += m.ComObjectDescriptors.GetLengthInBits()

	return lengthInBits
}

func (m *ComObjectTableRealisationType6) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ComObjectTableRealisationType6Parse(readBuffer utils.ReadBuffer, firmwareType FirmwareType) (*ComObjectTable, error) {
	if pullErr := readBuffer.PullContext("ComObjectTableRealisationType6"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (comObjectDescriptors)
	if pullErr := readBuffer.PullContext("comObjectDescriptors"); pullErr != nil {
		return nil, pullErr
	}
	_comObjectDescriptors, _comObjectDescriptorsErr := GroupObjectDescriptorRealisationType6Parse(readBuffer)
	if _comObjectDescriptorsErr != nil {
		return nil, errors.Wrap(_comObjectDescriptorsErr, "Error parsing 'comObjectDescriptors' field")
	}
	comObjectDescriptors := CastGroupObjectDescriptorRealisationType6(_comObjectDescriptors)
	if closeErr := readBuffer.CloseContext("comObjectDescriptors"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ComObjectTableRealisationType6"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ComObjectTableRealisationType6{
		ComObjectDescriptors: CastGroupObjectDescriptorRealisationType6(comObjectDescriptors),
		ComObjectTable:       &ComObjectTable{},
	}
	_child.ComObjectTable.Child = _child
	return _child.ComObjectTable, nil
}

func (m *ComObjectTableRealisationType6) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ComObjectTableRealisationType6"); pushErr != nil {
			return pushErr
		}

		// Simple Field (comObjectDescriptors)
		if pushErr := writeBuffer.PushContext("comObjectDescriptors"); pushErr != nil {
			return pushErr
		}
		_comObjectDescriptorsErr := m.ComObjectDescriptors.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("comObjectDescriptors"); popErr != nil {
			return popErr
		}
		if _comObjectDescriptorsErr != nil {
			return errors.Wrap(_comObjectDescriptorsErr, "Error serializing 'comObjectDescriptors' field")
		}

		if popErr := writeBuffer.PopContext("ComObjectTableRealisationType6"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ComObjectTableRealisationType6) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
