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
type BVLCForwardedNPDU struct {
	*BVLC
	Ip   []uint8
	Port uint16
	Npdu *NPDU

	// Arguments.
	BvlcPayloadLength uint16
}

// The corresponding interface
type IBVLCForwardedNPDU interface {
	IBVLC
	// GetIp returns Ip
	GetIp() []uint8
	// GetPort returns Port
	GetPort() uint16
	// GetNpdu returns Npdu
	GetNpdu() *NPDU
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
func (m *BVLCForwardedNPDU) BvlcFunction() uint8 {
	return 0x04
}

func (m *BVLCForwardedNPDU) GetBvlcFunction() uint8 {
	return 0x04
}

func (m *BVLCForwardedNPDU) InitializeParent(parent *BVLC) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BVLCForwardedNPDU) GetIp() []uint8 {
	return m.Ip
}

func (m *BVLCForwardedNPDU) GetPort() uint16 {
	return m.Port
}

func (m *BVLCForwardedNPDU) GetNpdu() *NPDU {
	return m.Npdu
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBVLCForwardedNPDU factory function for BVLCForwardedNPDU
func NewBVLCForwardedNPDU(ip []uint8, port uint16, npdu *NPDU, bvlcPayloadLength uint16) *BVLC {
	child := &BVLCForwardedNPDU{
		Ip:   ip,
		Port: port,
		Npdu: npdu,
		BVLC: NewBVLC(),
	}
	child.Child = child
	return child.BVLC
}

func CastBVLCForwardedNPDU(structType interface{}) *BVLCForwardedNPDU {
	if casted, ok := structType.(BVLCForwardedNPDU); ok {
		return &casted
	}
	if casted, ok := structType.(*BVLCForwardedNPDU); ok {
		return casted
	}
	if casted, ok := structType.(BVLC); ok {
		return CastBVLCForwardedNPDU(casted.Child)
	}
	if casted, ok := structType.(*BVLC); ok {
		return CastBVLCForwardedNPDU(casted.Child)
	}
	return nil
}

func (m *BVLCForwardedNPDU) GetTypeName() string {
	return "BVLCForwardedNPDU"
}

func (m *BVLCForwardedNPDU) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BVLCForwardedNPDU) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Ip) > 0 {
		lengthInBits += 8 * uint16(len(m.Ip))
	}

	// Simple field (port)
	lengthInBits += 16

	// Simple field (npdu)
	lengthInBits += m.Npdu.GetLengthInBits()

	return lengthInBits
}

func (m *BVLCForwardedNPDU) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BVLCForwardedNPDUParse(readBuffer utils.ReadBuffer, bvlcPayloadLength uint16) (*BVLC, error) {
	if pullErr := readBuffer.PullContext("BVLCForwardedNPDU"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Array field (ip)
	if pullErr := readBuffer.PullContext("ip", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	ip := make([]uint8, uint16(4))
	{
		for curItem := uint16(0); curItem < uint16(uint16(4)); curItem++ {
			_item, _err := readBuffer.ReadUint8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'ip' field")
			}
			ip[curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("ip", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (port)
	_port, _portErr := readBuffer.ReadUint16("port", 16)
	if _portErr != nil {
		return nil, errors.Wrap(_portErr, "Error parsing 'port' field")
	}
	port := _port

	// Simple Field (npdu)
	if pullErr := readBuffer.PullContext("npdu"); pullErr != nil {
		return nil, pullErr
	}
	_npdu, _npduErr := NPDUParse(readBuffer, uint16(uint16(bvlcPayloadLength)-uint16(uint16(6))))
	if _npduErr != nil {
		return nil, errors.Wrap(_npduErr, "Error parsing 'npdu' field")
	}
	npdu := CastNPDU(_npdu)
	if closeErr := readBuffer.CloseContext("npdu"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BVLCForwardedNPDU"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BVLCForwardedNPDU{
		Ip:   ip,
		Port: port,
		Npdu: CastNPDU(npdu),
		BVLC: &BVLC{},
	}
	_child.BVLC.Child = _child
	return _child.BVLC, nil
}

func (m *BVLCForwardedNPDU) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCForwardedNPDU"); pushErr != nil {
			return pushErr
		}

		// Array Field (ip)
		if m.Ip != nil {
			if pushErr := writeBuffer.PushContext("ip", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Ip {
				_elementErr := writeBuffer.WriteUint8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'ip' field")
				}
			}
			if popErr := writeBuffer.PopContext("ip", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		// Simple Field (port)
		port := uint16(m.Port)
		_portErr := writeBuffer.WriteUint16("port", 16, (port))
		if _portErr != nil {
			return errors.Wrap(_portErr, "Error serializing 'port' field")
		}

		// Simple Field (npdu)
		if pushErr := writeBuffer.PushContext("npdu"); pushErr != nil {
			return pushErr
		}
		_npduErr := m.Npdu.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("npdu"); popErr != nil {
			return popErr
		}
		if _npduErr != nil {
			return errors.Wrap(_npduErr, "Error serializing 'npdu' field")
		}

		if popErr := writeBuffer.PopContext("BVLCForwardedNPDU"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BVLCForwardedNPDU) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
