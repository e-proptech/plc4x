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
type DIBDeviceInfo struct {
	DescriptionType                uint8
	KnxMedium                      KnxMedium
	DeviceStatus                   *DeviceStatus
	KnxAddress                     *KnxAddress
	ProjectInstallationIdentifier  *ProjectInstallationIdentifier
	KnxNetIpDeviceSerialNumber     []byte
	KnxNetIpDeviceMulticastAddress *IPAddress
	KnxNetIpDeviceMacAddress       *MACAddress
	DeviceFriendlyName             []byte
}

// The corresponding interface
type IDIBDeviceInfo interface {
	// GetDescriptionType returns DescriptionType
	GetDescriptionType() uint8
	// GetKnxMedium returns KnxMedium
	GetKnxMedium() KnxMedium
	// GetDeviceStatus returns DeviceStatus
	GetDeviceStatus() *DeviceStatus
	// GetKnxAddress returns KnxAddress
	GetKnxAddress() *KnxAddress
	// GetProjectInstallationIdentifier returns ProjectInstallationIdentifier
	GetProjectInstallationIdentifier() *ProjectInstallationIdentifier
	// GetKnxNetIpDeviceSerialNumber returns KnxNetIpDeviceSerialNumber
	GetKnxNetIpDeviceSerialNumber() []byte
	// GetKnxNetIpDeviceMulticastAddress returns KnxNetIpDeviceMulticastAddress
	GetKnxNetIpDeviceMulticastAddress() *IPAddress
	// GetKnxNetIpDeviceMacAddress returns KnxNetIpDeviceMacAddress
	GetKnxNetIpDeviceMacAddress() *MACAddress
	// GetDeviceFriendlyName returns DeviceFriendlyName
	GetDeviceFriendlyName() []byte
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
func (m *DIBDeviceInfo) GetDescriptionType() uint8 {
	return m.DescriptionType
}

func (m *DIBDeviceInfo) GetKnxMedium() KnxMedium {
	return m.KnxMedium
}

func (m *DIBDeviceInfo) GetDeviceStatus() *DeviceStatus {
	return m.DeviceStatus
}

func (m *DIBDeviceInfo) GetKnxAddress() *KnxAddress {
	return m.KnxAddress
}

func (m *DIBDeviceInfo) GetProjectInstallationIdentifier() *ProjectInstallationIdentifier {
	return m.ProjectInstallationIdentifier
}

func (m *DIBDeviceInfo) GetKnxNetIpDeviceSerialNumber() []byte {
	return m.KnxNetIpDeviceSerialNumber
}

func (m *DIBDeviceInfo) GetKnxNetIpDeviceMulticastAddress() *IPAddress {
	return m.KnxNetIpDeviceMulticastAddress
}

func (m *DIBDeviceInfo) GetKnxNetIpDeviceMacAddress() *MACAddress {
	return m.KnxNetIpDeviceMacAddress
}

func (m *DIBDeviceInfo) GetDeviceFriendlyName() []byte {
	return m.DeviceFriendlyName
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewDIBDeviceInfo factory function for DIBDeviceInfo
func NewDIBDeviceInfo(descriptionType uint8, knxMedium KnxMedium, deviceStatus *DeviceStatus, knxAddress *KnxAddress, projectInstallationIdentifier *ProjectInstallationIdentifier, knxNetIpDeviceSerialNumber []byte, knxNetIpDeviceMulticastAddress *IPAddress, knxNetIpDeviceMacAddress *MACAddress, deviceFriendlyName []byte) *DIBDeviceInfo {
	return &DIBDeviceInfo{DescriptionType: descriptionType, KnxMedium: knxMedium, DeviceStatus: deviceStatus, KnxAddress: knxAddress, ProjectInstallationIdentifier: projectInstallationIdentifier, KnxNetIpDeviceSerialNumber: knxNetIpDeviceSerialNumber, KnxNetIpDeviceMulticastAddress: knxNetIpDeviceMulticastAddress, KnxNetIpDeviceMacAddress: knxNetIpDeviceMacAddress, DeviceFriendlyName: deviceFriendlyName}
}

func CastDIBDeviceInfo(structType interface{}) *DIBDeviceInfo {
	if casted, ok := structType.(DIBDeviceInfo); ok {
		return &casted
	}
	if casted, ok := structType.(*DIBDeviceInfo); ok {
		return casted
	}
	return nil
}

func (m *DIBDeviceInfo) GetTypeName() string {
	return "DIBDeviceInfo"
}

func (m *DIBDeviceInfo) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *DIBDeviceInfo) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (descriptionType)
	lengthInBits += 8

	// Simple field (knxMedium)
	lengthInBits += 8

	// Simple field (deviceStatus)
	lengthInBits += m.DeviceStatus.GetLengthInBits()

	// Simple field (knxAddress)
	lengthInBits += m.KnxAddress.GetLengthInBits()

	// Simple field (projectInstallationIdentifier)
	lengthInBits += m.ProjectInstallationIdentifier.GetLengthInBits()

	// Array field
	if len(m.KnxNetIpDeviceSerialNumber) > 0 {
		lengthInBits += 8 * uint16(len(m.KnxNetIpDeviceSerialNumber))
	}

	// Simple field (knxNetIpDeviceMulticastAddress)
	lengthInBits += m.KnxNetIpDeviceMulticastAddress.GetLengthInBits()

	// Simple field (knxNetIpDeviceMacAddress)
	lengthInBits += m.KnxNetIpDeviceMacAddress.GetLengthInBits()

	// Array field
	if len(m.DeviceFriendlyName) > 0 {
		lengthInBits += 8 * uint16(len(m.DeviceFriendlyName))
	}

	return lengthInBits
}

func (m *DIBDeviceInfo) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DIBDeviceInfoParse(readBuffer utils.ReadBuffer) (*DIBDeviceInfo, error) {
	if pullErr := readBuffer.PullContext("DIBDeviceInfo"); pullErr != nil {
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

	// Simple Field (descriptionType)
	_descriptionType, _descriptionTypeErr := readBuffer.ReadUint8("descriptionType", 8)
	if _descriptionTypeErr != nil {
		return nil, errors.Wrap(_descriptionTypeErr, "Error parsing 'descriptionType' field")
	}
	descriptionType := _descriptionType

	// Simple Field (knxMedium)
	if pullErr := readBuffer.PullContext("knxMedium"); pullErr != nil {
		return nil, pullErr
	}
	_knxMedium, _knxMediumErr := KnxMediumParse(readBuffer)
	if _knxMediumErr != nil {
		return nil, errors.Wrap(_knxMediumErr, "Error parsing 'knxMedium' field")
	}
	knxMedium := _knxMedium
	if closeErr := readBuffer.CloseContext("knxMedium"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (deviceStatus)
	if pullErr := readBuffer.PullContext("deviceStatus"); pullErr != nil {
		return nil, pullErr
	}
	_deviceStatus, _deviceStatusErr := DeviceStatusParse(readBuffer)
	if _deviceStatusErr != nil {
		return nil, errors.Wrap(_deviceStatusErr, "Error parsing 'deviceStatus' field")
	}
	deviceStatus := CastDeviceStatus(_deviceStatus)
	if closeErr := readBuffer.CloseContext("deviceStatus"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (knxAddress)
	if pullErr := readBuffer.PullContext("knxAddress"); pullErr != nil {
		return nil, pullErr
	}
	_knxAddress, _knxAddressErr := KnxAddressParse(readBuffer)
	if _knxAddressErr != nil {
		return nil, errors.Wrap(_knxAddressErr, "Error parsing 'knxAddress' field")
	}
	knxAddress := CastKnxAddress(_knxAddress)
	if closeErr := readBuffer.CloseContext("knxAddress"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (projectInstallationIdentifier)
	if pullErr := readBuffer.PullContext("projectInstallationIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_projectInstallationIdentifier, _projectInstallationIdentifierErr := ProjectInstallationIdentifierParse(readBuffer)
	if _projectInstallationIdentifierErr != nil {
		return nil, errors.Wrap(_projectInstallationIdentifierErr, "Error parsing 'projectInstallationIdentifier' field")
	}
	projectInstallationIdentifier := CastProjectInstallationIdentifier(_projectInstallationIdentifier)
	if closeErr := readBuffer.CloseContext("projectInstallationIdentifier"); closeErr != nil {
		return nil, closeErr
	}
	// Byte Array field (knxNetIpDeviceSerialNumber)
	numberOfBytesknxNetIpDeviceSerialNumber := int(uint16(6))
	knxNetIpDeviceSerialNumber, _readArrayErr := readBuffer.ReadByteArray("knxNetIpDeviceSerialNumber", numberOfBytesknxNetIpDeviceSerialNumber)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'knxNetIpDeviceSerialNumber' field")
	}

	// Simple Field (knxNetIpDeviceMulticastAddress)
	if pullErr := readBuffer.PullContext("knxNetIpDeviceMulticastAddress"); pullErr != nil {
		return nil, pullErr
	}
	_knxNetIpDeviceMulticastAddress, _knxNetIpDeviceMulticastAddressErr := IPAddressParse(readBuffer)
	if _knxNetIpDeviceMulticastAddressErr != nil {
		return nil, errors.Wrap(_knxNetIpDeviceMulticastAddressErr, "Error parsing 'knxNetIpDeviceMulticastAddress' field")
	}
	knxNetIpDeviceMulticastAddress := CastIPAddress(_knxNetIpDeviceMulticastAddress)
	if closeErr := readBuffer.CloseContext("knxNetIpDeviceMulticastAddress"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (knxNetIpDeviceMacAddress)
	if pullErr := readBuffer.PullContext("knxNetIpDeviceMacAddress"); pullErr != nil {
		return nil, pullErr
	}
	_knxNetIpDeviceMacAddress, _knxNetIpDeviceMacAddressErr := MACAddressParse(readBuffer)
	if _knxNetIpDeviceMacAddressErr != nil {
		return nil, errors.Wrap(_knxNetIpDeviceMacAddressErr, "Error parsing 'knxNetIpDeviceMacAddress' field")
	}
	knxNetIpDeviceMacAddress := CastMACAddress(_knxNetIpDeviceMacAddress)
	if closeErr := readBuffer.CloseContext("knxNetIpDeviceMacAddress"); closeErr != nil {
		return nil, closeErr
	}
	// Byte Array field (deviceFriendlyName)
	numberOfBytesdeviceFriendlyName := int(uint16(30))
	deviceFriendlyName, _readArrayErr := readBuffer.ReadByteArray("deviceFriendlyName", numberOfBytesdeviceFriendlyName)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'deviceFriendlyName' field")
	}

	if closeErr := readBuffer.CloseContext("DIBDeviceInfo"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewDIBDeviceInfo(descriptionType, knxMedium, deviceStatus, knxAddress, projectInstallationIdentifier, knxNetIpDeviceSerialNumber, knxNetIpDeviceMulticastAddress, knxNetIpDeviceMacAddress, deviceFriendlyName), nil
}

func (m *DIBDeviceInfo) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("DIBDeviceInfo"); pushErr != nil {
		return pushErr
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.GetLengthInBytes()))
	_structureLengthErr := writeBuffer.WriteUint8("structureLength", 8, (structureLength))
	if _structureLengthErr != nil {
		return errors.Wrap(_structureLengthErr, "Error serializing 'structureLength' field")
	}

	// Simple Field (descriptionType)
	descriptionType := uint8(m.DescriptionType)
	_descriptionTypeErr := writeBuffer.WriteUint8("descriptionType", 8, (descriptionType))
	if _descriptionTypeErr != nil {
		return errors.Wrap(_descriptionTypeErr, "Error serializing 'descriptionType' field")
	}

	// Simple Field (knxMedium)
	if pushErr := writeBuffer.PushContext("knxMedium"); pushErr != nil {
		return pushErr
	}
	_knxMediumErr := m.KnxMedium.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("knxMedium"); popErr != nil {
		return popErr
	}
	if _knxMediumErr != nil {
		return errors.Wrap(_knxMediumErr, "Error serializing 'knxMedium' field")
	}

	// Simple Field (deviceStatus)
	if pushErr := writeBuffer.PushContext("deviceStatus"); pushErr != nil {
		return pushErr
	}
	_deviceStatusErr := m.DeviceStatus.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("deviceStatus"); popErr != nil {
		return popErr
	}
	if _deviceStatusErr != nil {
		return errors.Wrap(_deviceStatusErr, "Error serializing 'deviceStatus' field")
	}

	// Simple Field (knxAddress)
	if pushErr := writeBuffer.PushContext("knxAddress"); pushErr != nil {
		return pushErr
	}
	_knxAddressErr := m.KnxAddress.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("knxAddress"); popErr != nil {
		return popErr
	}
	if _knxAddressErr != nil {
		return errors.Wrap(_knxAddressErr, "Error serializing 'knxAddress' field")
	}

	// Simple Field (projectInstallationIdentifier)
	if pushErr := writeBuffer.PushContext("projectInstallationIdentifier"); pushErr != nil {
		return pushErr
	}
	_projectInstallationIdentifierErr := m.ProjectInstallationIdentifier.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("projectInstallationIdentifier"); popErr != nil {
		return popErr
	}
	if _projectInstallationIdentifierErr != nil {
		return errors.Wrap(_projectInstallationIdentifierErr, "Error serializing 'projectInstallationIdentifier' field")
	}

	// Array Field (knxNetIpDeviceSerialNumber)
	if m.KnxNetIpDeviceSerialNumber != nil {
		// Byte Array field (knxNetIpDeviceSerialNumber)
		_writeArrayErr := writeBuffer.WriteByteArray("knxNetIpDeviceSerialNumber", m.KnxNetIpDeviceSerialNumber)
		if _writeArrayErr != nil {
			return errors.Wrap(_writeArrayErr, "Error serializing 'knxNetIpDeviceSerialNumber' field")
		}
	}

	// Simple Field (knxNetIpDeviceMulticastAddress)
	if pushErr := writeBuffer.PushContext("knxNetIpDeviceMulticastAddress"); pushErr != nil {
		return pushErr
	}
	_knxNetIpDeviceMulticastAddressErr := m.KnxNetIpDeviceMulticastAddress.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("knxNetIpDeviceMulticastAddress"); popErr != nil {
		return popErr
	}
	if _knxNetIpDeviceMulticastAddressErr != nil {
		return errors.Wrap(_knxNetIpDeviceMulticastAddressErr, "Error serializing 'knxNetIpDeviceMulticastAddress' field")
	}

	// Simple Field (knxNetIpDeviceMacAddress)
	if pushErr := writeBuffer.PushContext("knxNetIpDeviceMacAddress"); pushErr != nil {
		return pushErr
	}
	_knxNetIpDeviceMacAddressErr := m.KnxNetIpDeviceMacAddress.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("knxNetIpDeviceMacAddress"); popErr != nil {
		return popErr
	}
	if _knxNetIpDeviceMacAddressErr != nil {
		return errors.Wrap(_knxNetIpDeviceMacAddressErr, "Error serializing 'knxNetIpDeviceMacAddress' field")
	}

	// Array Field (deviceFriendlyName)
	if m.DeviceFriendlyName != nil {
		// Byte Array field (deviceFriendlyName)
		_writeArrayErr := writeBuffer.WriteByteArray("deviceFriendlyName", m.DeviceFriendlyName)
		if _writeArrayErr != nil {
			return errors.Wrap(_writeArrayErr, "Error serializing 'deviceFriendlyName' field")
		}
	}

	if popErr := writeBuffer.PopContext("DIBDeviceInfo"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *DIBDeviceInfo) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
