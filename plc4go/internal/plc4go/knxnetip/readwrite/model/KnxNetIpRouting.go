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
type KnxNetIpRouting struct {
	*ServiceId
	Version uint8
}

// The corresponding interface
type IKnxNetIpRouting interface {
	IServiceId
	// GetVersion returns Version
	GetVersion() uint8
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
func (m *KnxNetIpRouting) ServiceType() uint8 {
	return 0x05
}

func (m *KnxNetIpRouting) GetServiceType() uint8 {
	return 0x05
}

func (m *KnxNetIpRouting) InitializeParent(parent *ServiceId) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *KnxNetIpRouting) GetVersion() uint8 {
	return m.Version
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewKnxNetIpRouting factory function for KnxNetIpRouting
func NewKnxNetIpRouting(version uint8) *ServiceId {
	child := &KnxNetIpRouting{
		Version:   version,
		ServiceId: NewServiceId(),
	}
	child.Child = child
	return child.ServiceId
}

func CastKnxNetIpRouting(structType interface{}) *KnxNetIpRouting {
	if casted, ok := structType.(KnxNetIpRouting); ok {
		return &casted
	}
	if casted, ok := structType.(*KnxNetIpRouting); ok {
		return casted
	}
	if casted, ok := structType.(ServiceId); ok {
		return CastKnxNetIpRouting(casted.Child)
	}
	if casted, ok := structType.(*ServiceId); ok {
		return CastKnxNetIpRouting(casted.Child)
	}
	return nil
}

func (m *KnxNetIpRouting) GetTypeName() string {
	return "KnxNetIpRouting"
}

func (m *KnxNetIpRouting) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *KnxNetIpRouting) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *KnxNetIpRouting) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxNetIpRoutingParse(readBuffer utils.ReadBuffer) (*ServiceId, error) {
	if pullErr := readBuffer.PullContext("KnxNetIpRouting"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (version)
	_version, _versionErr := readBuffer.ReadUint8("version", 8)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field")
	}
	version := _version

	if closeErr := readBuffer.CloseContext("KnxNetIpRouting"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &KnxNetIpRouting{
		Version:   version,
		ServiceId: &ServiceId{},
	}
	_child.ServiceId.Child = _child
	return _child.ServiceId, nil
}

func (m *KnxNetIpRouting) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetIpRouting"); pushErr != nil {
			return pushErr
		}

		// Simple Field (version)
		version := uint8(m.Version)
		_versionErr := writeBuffer.WriteUint8("version", 8, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetIpRouting"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *KnxNetIpRouting) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
