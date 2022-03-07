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
type NLMICouldBeRouterToNetwork struct {
	*NLM
	DestinationNetworkAddress uint16
	PerformanceIndex          uint8

	// Arguments.
	ApduLength uint16
}

// The corresponding interface
type INLMICouldBeRouterToNetwork interface {
	INLM
	// GetDestinationNetworkAddress returns DestinationNetworkAddress
	GetDestinationNetworkAddress() uint16
	// GetPerformanceIndex returns PerformanceIndex
	GetPerformanceIndex() uint8
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
func (m *NLMICouldBeRouterToNetwork) MessageType() uint8 {
	return 0x02
}

func (m *NLMICouldBeRouterToNetwork) GetMessageType() uint8 {
	return 0x02
}

func (m *NLMICouldBeRouterToNetwork) InitializeParent(parent *NLM, vendorId *uint16) {
	m.NLM.VendorId = vendorId
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *NLMICouldBeRouterToNetwork) GetDestinationNetworkAddress() uint16 {
	return m.DestinationNetworkAddress
}

func (m *NLMICouldBeRouterToNetwork) GetPerformanceIndex() uint8 {
	return m.PerformanceIndex
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewNLMICouldBeRouterToNetwork factory function for NLMICouldBeRouterToNetwork
func NewNLMICouldBeRouterToNetwork(destinationNetworkAddress uint16, performanceIndex uint8, vendorId *uint16, apduLength uint16) *NLM {
	child := &NLMICouldBeRouterToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		PerformanceIndex:          performanceIndex,
		NLM:                       NewNLM(vendorId, apduLength),
	}
	child.Child = child
	return child.NLM
}

func CastNLMICouldBeRouterToNetwork(structType interface{}) *NLMICouldBeRouterToNetwork {
	if casted, ok := structType.(NLMICouldBeRouterToNetwork); ok {
		return &casted
	}
	if casted, ok := structType.(*NLMICouldBeRouterToNetwork); ok {
		return casted
	}
	if casted, ok := structType.(NLM); ok {
		return CastNLMICouldBeRouterToNetwork(casted.Child)
	}
	if casted, ok := structType.(*NLM); ok {
		return CastNLMICouldBeRouterToNetwork(casted.Child)
	}
	return nil
}

func (m *NLMICouldBeRouterToNetwork) GetTypeName() string {
	return "NLMICouldBeRouterToNetwork"
}

func (m *NLMICouldBeRouterToNetwork) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *NLMICouldBeRouterToNetwork) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (destinationNetworkAddress)
	lengthInBits += 16

	// Simple field (performanceIndex)
	lengthInBits += 8

	return lengthInBits
}

func (m *NLMICouldBeRouterToNetwork) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NLMICouldBeRouterToNetworkParse(readBuffer utils.ReadBuffer, apduLength uint16, messageType uint8) (*NLM, error) {
	if pullErr := readBuffer.PullContext("NLMICouldBeRouterToNetwork"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (destinationNetworkAddress)
	_destinationNetworkAddress, _destinationNetworkAddressErr := readBuffer.ReadUint16("destinationNetworkAddress", 16)
	if _destinationNetworkAddressErr != nil {
		return nil, errors.Wrap(_destinationNetworkAddressErr, "Error parsing 'destinationNetworkAddress' field")
	}
	destinationNetworkAddress := _destinationNetworkAddress

	// Simple Field (performanceIndex)
	_performanceIndex, _performanceIndexErr := readBuffer.ReadUint8("performanceIndex", 8)
	if _performanceIndexErr != nil {
		return nil, errors.Wrap(_performanceIndexErr, "Error parsing 'performanceIndex' field")
	}
	performanceIndex := _performanceIndex

	if closeErr := readBuffer.CloseContext("NLMICouldBeRouterToNetwork"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &NLMICouldBeRouterToNetwork{
		DestinationNetworkAddress: destinationNetworkAddress,
		PerformanceIndex:          performanceIndex,
		NLM:                       &NLM{},
	}
	_child.NLM.Child = _child
	return _child.NLM, nil
}

func (m *NLMICouldBeRouterToNetwork) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMICouldBeRouterToNetwork"); pushErr != nil {
			return pushErr
		}

		// Simple Field (destinationNetworkAddress)
		destinationNetworkAddress := uint16(m.DestinationNetworkAddress)
		_destinationNetworkAddressErr := writeBuffer.WriteUint16("destinationNetworkAddress", 16, (destinationNetworkAddress))
		if _destinationNetworkAddressErr != nil {
			return errors.Wrap(_destinationNetworkAddressErr, "Error serializing 'destinationNetworkAddress' field")
		}

		// Simple Field (performanceIndex)
		performanceIndex := uint8(m.PerformanceIndex)
		_performanceIndexErr := writeBuffer.WriteUint8("performanceIndex", 8, (performanceIndex))
		if _performanceIndexErr != nil {
			return errors.Wrap(_performanceIndexErr, "Error serializing 'performanceIndex' field")
		}

		if popErr := writeBuffer.PopContext("NLMICouldBeRouterToNetwork"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *NLMICouldBeRouterToNetwork) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
