/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// NLMSecurityPayload is the corresponding interface of NLMSecurityPayload
type NLMSecurityPayload interface {
	utils.LengthAware
	utils.Serializable
	NLM
	// GetPayloadLength returns PayloadLength (property field)
	GetPayloadLength() uint16
	// GetPayload returns Payload (property field)
	GetPayload() []byte
}

// NLMSecurityPayloadExactly can be used when we want exactly this type and not a type which fulfills NLMSecurityPayload.
// This is useful for switch cases.
type NLMSecurityPayloadExactly interface {
	NLMSecurityPayload
	isNLMSecurityPayload() bool
}

// _NLMSecurityPayload is the data-structure of this message
type _NLMSecurityPayload struct {
	*_NLM
	PayloadLength uint16
	Payload       []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMSecurityPayload) GetMessageType() uint8 {
	return 0x0B
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMSecurityPayload) InitializeParent(parent NLM) {}

func (m *_NLMSecurityPayload) GetParent() NLM {
	return m._NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMSecurityPayload) GetPayloadLength() uint16 {
	return m.PayloadLength
}

func (m *_NLMSecurityPayload) GetPayload() []byte {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMSecurityPayload factory function for _NLMSecurityPayload
func NewNLMSecurityPayload(payloadLength uint16, payload []byte, apduLength uint16) *_NLMSecurityPayload {
	_result := &_NLMSecurityPayload{
		PayloadLength: payloadLength,
		Payload:       payload,
		_NLM:          NewNLM(apduLength),
	}
	_result._NLM._NLMChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMSecurityPayload(structType interface{}) NLMSecurityPayload {
	if casted, ok := structType.(NLMSecurityPayload); ok {
		return casted
	}
	if casted, ok := structType.(*NLMSecurityPayload); ok {
		return *casted
	}
	return nil
}

func (m *_NLMSecurityPayload) GetTypeName() string {
	return "NLMSecurityPayload"
}

func (m *_NLMSecurityPayload) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (payloadLength)
	lengthInBits += 16

	// Array field
	if len(m.Payload) > 0 {
		lengthInBits += 8 * uint16(len(m.Payload))
	}

	return lengthInBits
}

func (m *_NLMSecurityPayload) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NLMSecurityPayloadParse(theBytes []byte, apduLength uint16) (NLMSecurityPayload, error) {
	return NLMSecurityPayloadParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), apduLength)
}

func NLMSecurityPayloadParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, apduLength uint16) (NLMSecurityPayload, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMSecurityPayload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMSecurityPayload")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (payloadLength)
	_payloadLength, _payloadLengthErr := readBuffer.ReadUint16("payloadLength", 16)
	if _payloadLengthErr != nil {
		return nil, errors.Wrap(_payloadLengthErr, "Error parsing 'payloadLength' field of NLMSecurityPayload")
	}
	payloadLength := _payloadLength
	// Byte Array field (payload)
	numberOfBytespayload := int(payloadLength)
	payload, _readArrayErr := readBuffer.ReadByteArray("payload", numberOfBytespayload)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'payload' field of NLMSecurityPayload")
	}

	if closeErr := readBuffer.CloseContext("NLMSecurityPayload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMSecurityPayload")
	}

	// Create a partially initialized instance
	_child := &_NLMSecurityPayload{
		_NLM: &_NLM{
			ApduLength: apduLength,
		},
		PayloadLength: payloadLength,
		Payload:       payload,
	}
	_child._NLM._NLMChildRequirements = _child
	return _child, nil
}

func (m *_NLMSecurityPayload) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMSecurityPayload) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMSecurityPayload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMSecurityPayload")
		}

		// Simple Field (payloadLength)
		payloadLength := uint16(m.GetPayloadLength())
		_payloadLengthErr := writeBuffer.WriteUint16("payloadLength", 16, (payloadLength))
		if _payloadLengthErr != nil {
			return errors.Wrap(_payloadLengthErr, "Error serializing 'payloadLength' field")
		}

		// Array Field (payload)
		// Byte Array field (payload)
		if err := writeBuffer.WriteByteArray("payload", m.GetPayload()); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("NLMSecurityPayload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMSecurityPayload")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NLMSecurityPayload) isNLMSecurityPayload() bool {
	return true
}

func (m *_NLMSecurityPayload) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
