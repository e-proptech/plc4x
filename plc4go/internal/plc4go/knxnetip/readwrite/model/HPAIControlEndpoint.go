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
type HPAIControlEndpoint struct {
	HostProtocolCode HostProtocolCode
	IpAddress        *IPAddress
	IpPort           uint16
}

// The corresponding interface
type IHPAIControlEndpoint interface {
	// GetHostProtocolCode returns HostProtocolCode
	GetHostProtocolCode() HostProtocolCode
	// GetIpAddress returns IpAddress
	GetIpAddress() *IPAddress
	// GetIpPort returns IpPort
	GetIpPort() uint16
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
func (m *HPAIControlEndpoint) GetHostProtocolCode() HostProtocolCode {
	return m.HostProtocolCode
}

func (m *HPAIControlEndpoint) GetIpAddress() *IPAddress {
	return m.IpAddress
}

func (m *HPAIControlEndpoint) GetIpPort() uint16 {
	return m.IpPort
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewHPAIControlEndpoint factory function for HPAIControlEndpoint
func NewHPAIControlEndpoint(hostProtocolCode HostProtocolCode, ipAddress *IPAddress, ipPort uint16) *HPAIControlEndpoint {
	return &HPAIControlEndpoint{HostProtocolCode: hostProtocolCode, IpAddress: ipAddress, IpPort: ipPort}
}

func CastHPAIControlEndpoint(structType interface{}) *HPAIControlEndpoint {
	if casted, ok := structType.(HPAIControlEndpoint); ok {
		return &casted
	}
	if casted, ok := structType.(*HPAIControlEndpoint); ok {
		return casted
	}
	return nil
}

func (m *HPAIControlEndpoint) GetTypeName() string {
	return "HPAIControlEndpoint"
}

func (m *HPAIControlEndpoint) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *HPAIControlEndpoint) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (hostProtocolCode)
	lengthInBits += 8

	// Simple field (ipAddress)
	lengthInBits += m.IpAddress.GetLengthInBits()

	// Simple field (ipPort)
	lengthInBits += 16

	return lengthInBits
}

func (m *HPAIControlEndpoint) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func HPAIControlEndpointParse(readBuffer utils.ReadBuffer) (*HPAIControlEndpoint, error) {
	if pullErr := readBuffer.PullContext("HPAIControlEndpoint"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Implicit Field (structureLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	structureLength, _structureLengthErr := readBuffer.ReadUint8("structureLength", 8)
	_ = structureLength
	if _structureLengthErr != nil {
		return nil, errors.Wrap(_structureLengthErr, "Error parsing 'structureLength' field")
	}

	// Simple Field (hostProtocolCode)
	if pullErr := readBuffer.PullContext("hostProtocolCode"); pullErr != nil {
		return nil, pullErr
	}
	_hostProtocolCode, _hostProtocolCodeErr := HostProtocolCodeParse(readBuffer)
	if _hostProtocolCodeErr != nil {
		return nil, errors.Wrap(_hostProtocolCodeErr, "Error parsing 'hostProtocolCode' field")
	}
	hostProtocolCode := _hostProtocolCode
	if closeErr := readBuffer.CloseContext("hostProtocolCode"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (ipAddress)
	if pullErr := readBuffer.PullContext("ipAddress"); pullErr != nil {
		return nil, pullErr
	}
	_ipAddress, _ipAddressErr := IPAddressParse(readBuffer)
	if _ipAddressErr != nil {
		return nil, errors.Wrap(_ipAddressErr, "Error parsing 'ipAddress' field")
	}
	ipAddress := CastIPAddress(_ipAddress)
	if closeErr := readBuffer.CloseContext("ipAddress"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (ipPort)
	_ipPort, _ipPortErr := readBuffer.ReadUint16("ipPort", 16)
	if _ipPortErr != nil {
		return nil, errors.Wrap(_ipPortErr, "Error parsing 'ipPort' field")
	}
	ipPort := _ipPort

	if closeErr := readBuffer.CloseContext("HPAIControlEndpoint"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewHPAIControlEndpoint(hostProtocolCode, ipAddress, ipPort), nil
}

func (m *HPAIControlEndpoint) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("HPAIControlEndpoint"); pushErr != nil {
		return pushErr
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.GetLengthInBytes()))
	_structureLengthErr := writeBuffer.WriteUint8("structureLength", 8, (structureLength))
	if _structureLengthErr != nil {
		return errors.Wrap(_structureLengthErr, "Error serializing 'structureLength' field")
	}

	// Simple Field (hostProtocolCode)
	if pushErr := writeBuffer.PushContext("hostProtocolCode"); pushErr != nil {
		return pushErr
	}
	_hostProtocolCodeErr := m.HostProtocolCode.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("hostProtocolCode"); popErr != nil {
		return popErr
	}
	if _hostProtocolCodeErr != nil {
		return errors.Wrap(_hostProtocolCodeErr, "Error serializing 'hostProtocolCode' field")
	}

	// Simple Field (ipAddress)
	if pushErr := writeBuffer.PushContext("ipAddress"); pushErr != nil {
		return pushErr
	}
	_ipAddressErr := m.IpAddress.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("ipAddress"); popErr != nil {
		return popErr
	}
	if _ipAddressErr != nil {
		return errors.Wrap(_ipAddressErr, "Error serializing 'ipAddress' field")
	}

	// Simple Field (ipPort)
	ipPort := uint16(m.IpPort)
	_ipPortErr := writeBuffer.WriteUint16("ipPort", 16, (ipPort))
	if _ipPortErr != nil {
		return errors.Wrap(_ipPortErr, "Error serializing 'ipPort' field")
	}

	if popErr := writeBuffer.PopContext("HPAIControlEndpoint"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *HPAIControlEndpoint) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
