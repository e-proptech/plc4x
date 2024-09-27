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
	"fmt"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// COTPPacketDisconnectRequest is the corresponding interface of COTPPacketDisconnectRequest
type COTPPacketDisconnectRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	COTPPacket
	// GetDestinationReference returns DestinationReference (property field)
	GetDestinationReference() uint16
	// GetSourceReference returns SourceReference (property field)
	GetSourceReference() uint16
	// GetProtocolClass returns ProtocolClass (property field)
	GetProtocolClass() COTPProtocolClass
	// IsCOTPPacketDisconnectRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCOTPPacketDisconnectRequest()
	// CreateBuilder creates a COTPPacketDisconnectRequestBuilder
	CreateCOTPPacketDisconnectRequestBuilder() COTPPacketDisconnectRequestBuilder
}

// _COTPPacketDisconnectRequest is the data-structure of this message
type _COTPPacketDisconnectRequest struct {
	COTPPacketContract
	DestinationReference uint16
	SourceReference      uint16
	ProtocolClass        COTPProtocolClass
}

var _ COTPPacketDisconnectRequest = (*_COTPPacketDisconnectRequest)(nil)
var _ COTPPacketRequirements = (*_COTPPacketDisconnectRequest)(nil)

// NewCOTPPacketDisconnectRequest factory function for _COTPPacketDisconnectRequest
func NewCOTPPacketDisconnectRequest(parameters []COTPParameter, payload S7Message, destinationReference uint16, sourceReference uint16, protocolClass COTPProtocolClass, cotpLen uint16) *_COTPPacketDisconnectRequest {
	_result := &_COTPPacketDisconnectRequest{
		COTPPacketContract:   NewCOTPPacket(parameters, payload, cotpLen),
		DestinationReference: destinationReference,
		SourceReference:      sourceReference,
		ProtocolClass:        protocolClass,
	}
	_result.COTPPacketContract.(*_COTPPacket)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// COTPPacketDisconnectRequestBuilder is a builder for COTPPacketDisconnectRequest
type COTPPacketDisconnectRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(destinationReference uint16, sourceReference uint16, protocolClass COTPProtocolClass) COTPPacketDisconnectRequestBuilder
	// WithDestinationReference adds DestinationReference (property field)
	WithDestinationReference(uint16) COTPPacketDisconnectRequestBuilder
	// WithSourceReference adds SourceReference (property field)
	WithSourceReference(uint16) COTPPacketDisconnectRequestBuilder
	// WithProtocolClass adds ProtocolClass (property field)
	WithProtocolClass(COTPProtocolClass) COTPPacketDisconnectRequestBuilder
	// Build builds the COTPPacketDisconnectRequest or returns an error if something is wrong
	Build() (COTPPacketDisconnectRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() COTPPacketDisconnectRequest
}

// NewCOTPPacketDisconnectRequestBuilder() creates a COTPPacketDisconnectRequestBuilder
func NewCOTPPacketDisconnectRequestBuilder() COTPPacketDisconnectRequestBuilder {
	return &_COTPPacketDisconnectRequestBuilder{_COTPPacketDisconnectRequest: new(_COTPPacketDisconnectRequest)}
}

type _COTPPacketDisconnectRequestBuilder struct {
	*_COTPPacketDisconnectRequest

	parentBuilder *_COTPPacketBuilder

	err *utils.MultiError
}

var _ (COTPPacketDisconnectRequestBuilder) = (*_COTPPacketDisconnectRequestBuilder)(nil)

func (b *_COTPPacketDisconnectRequestBuilder) setParent(contract COTPPacketContract) {
	b.COTPPacketContract = contract
}

func (b *_COTPPacketDisconnectRequestBuilder) WithMandatoryFields(destinationReference uint16, sourceReference uint16, protocolClass COTPProtocolClass) COTPPacketDisconnectRequestBuilder {
	return b.WithDestinationReference(destinationReference).WithSourceReference(sourceReference).WithProtocolClass(protocolClass)
}

func (b *_COTPPacketDisconnectRequestBuilder) WithDestinationReference(destinationReference uint16) COTPPacketDisconnectRequestBuilder {
	b.DestinationReference = destinationReference
	return b
}

func (b *_COTPPacketDisconnectRequestBuilder) WithSourceReference(sourceReference uint16) COTPPacketDisconnectRequestBuilder {
	b.SourceReference = sourceReference
	return b
}

func (b *_COTPPacketDisconnectRequestBuilder) WithProtocolClass(protocolClass COTPProtocolClass) COTPPacketDisconnectRequestBuilder {
	b.ProtocolClass = protocolClass
	return b
}

func (b *_COTPPacketDisconnectRequestBuilder) Build() (COTPPacketDisconnectRequest, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._COTPPacketDisconnectRequest.deepCopy(), nil
}

func (b *_COTPPacketDisconnectRequestBuilder) MustBuild() COTPPacketDisconnectRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_COTPPacketDisconnectRequestBuilder) Done() COTPPacketBuilder {
	return b.parentBuilder
}

func (b *_COTPPacketDisconnectRequestBuilder) buildForCOTPPacket() (COTPPacket, error) {
	return b.Build()
}

func (b *_COTPPacketDisconnectRequestBuilder) DeepCopy() any {
	_copy := b.CreateCOTPPacketDisconnectRequestBuilder().(*_COTPPacketDisconnectRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCOTPPacketDisconnectRequestBuilder creates a COTPPacketDisconnectRequestBuilder
func (b *_COTPPacketDisconnectRequest) CreateCOTPPacketDisconnectRequestBuilder() COTPPacketDisconnectRequestBuilder {
	if b == nil {
		return NewCOTPPacketDisconnectRequestBuilder()
	}
	return &_COTPPacketDisconnectRequestBuilder{_COTPPacketDisconnectRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPPacketDisconnectRequest) GetTpduCode() uint8 {
	return 0x80
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPPacketDisconnectRequest) GetParent() COTPPacketContract {
	return m.COTPPacketContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPPacketDisconnectRequest) GetDestinationReference() uint16 {
	return m.DestinationReference
}

func (m *_COTPPacketDisconnectRequest) GetSourceReference() uint16 {
	return m.SourceReference
}

func (m *_COTPPacketDisconnectRequest) GetProtocolClass() COTPProtocolClass {
	return m.ProtocolClass
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCOTPPacketDisconnectRequest(structType any) COTPPacketDisconnectRequest {
	if casted, ok := structType.(COTPPacketDisconnectRequest); ok {
		return casted
	}
	if casted, ok := structType.(*COTPPacketDisconnectRequest); ok {
		return *casted
	}
	return nil
}

func (m *_COTPPacketDisconnectRequest) GetTypeName() string {
	return "COTPPacketDisconnectRequest"
}

func (m *_COTPPacketDisconnectRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.COTPPacketContract.(*_COTPPacket).getLengthInBits(ctx))

	// Simple field (destinationReference)
	lengthInBits += 16

	// Simple field (sourceReference)
	lengthInBits += 16

	// Simple field (protocolClass)
	lengthInBits += 8

	return lengthInBits
}

func (m *_COTPPacketDisconnectRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_COTPPacketDisconnectRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_COTPPacket, cotpLen uint16) (__cOTPPacketDisconnectRequest COTPPacketDisconnectRequest, err error) {
	m.COTPPacketContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPPacketDisconnectRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPPacketDisconnectRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	destinationReference, err := ReadSimpleField(ctx, "destinationReference", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'destinationReference' field"))
	}
	m.DestinationReference = destinationReference

	sourceReference, err := ReadSimpleField(ctx, "sourceReference", ReadUnsignedShort(readBuffer, uint8(16)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceReference' field"))
	}
	m.SourceReference = sourceReference

	protocolClass, err := ReadEnumField[COTPProtocolClass](ctx, "protocolClass", "COTPProtocolClass", ReadEnum(COTPProtocolClassByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'protocolClass' field"))
	}
	m.ProtocolClass = protocolClass

	if closeErr := readBuffer.CloseContext("COTPPacketDisconnectRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPPacketDisconnectRequest")
	}

	return m, nil
}

func (m *_COTPPacketDisconnectRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_COTPPacketDisconnectRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPPacketDisconnectRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPPacketDisconnectRequest")
		}

		if err := WriteSimpleField[uint16](ctx, "destinationReference", m.GetDestinationReference(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'destinationReference' field")
		}

		if err := WriteSimpleField[uint16](ctx, "sourceReference", m.GetSourceReference(), WriteUnsignedShort(writeBuffer, 16)); err != nil {
			return errors.Wrap(err, "Error serializing 'sourceReference' field")
		}

		if err := WriteSimpleEnumField[COTPProtocolClass](ctx, "protocolClass", "COTPProtocolClass", m.GetProtocolClass(), WriteEnum[COTPProtocolClass, uint8](COTPProtocolClass.GetValue, COTPProtocolClass.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'protocolClass' field")
		}

		if popErr := writeBuffer.PopContext("COTPPacketDisconnectRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPPacketDisconnectRequest")
		}
		return nil
	}
	return m.COTPPacketContract.(*_COTPPacket).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_COTPPacketDisconnectRequest) IsCOTPPacketDisconnectRequest() {}

func (m *_COTPPacketDisconnectRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_COTPPacketDisconnectRequest) deepCopy() *_COTPPacketDisconnectRequest {
	if m == nil {
		return nil
	}
	_COTPPacketDisconnectRequestCopy := &_COTPPacketDisconnectRequest{
		m.COTPPacketContract.(*_COTPPacket).deepCopy(),
		m.DestinationReference,
		m.SourceReference,
		m.ProtocolClass,
	}
	m.COTPPacketContract.(*_COTPPacket)._SubType = m
	return _COTPPacketDisconnectRequestCopy
}

func (m *_COTPPacketDisconnectRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
