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

// OpcuaOpenResponse is the corresponding interface of OpcuaOpenResponse
type OpcuaOpenResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MessagePDU
	// GetOpenResponse returns OpenResponse (property field)
	GetOpenResponse() OpenChannelMessage
	// GetMessage returns Message (property field)
	GetMessage() Payload
	// IsOpcuaOpenResponse is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsOpcuaOpenResponse()
	// CreateBuilder creates a OpcuaOpenResponseBuilder
	CreateOpcuaOpenResponseBuilder() OpcuaOpenResponseBuilder
}

// _OpcuaOpenResponse is the data-structure of this message
type _OpcuaOpenResponse struct {
	MessagePDUContract
	OpenResponse OpenChannelMessage
	Message      Payload

	// Arguments.
	TotalLength uint32
}

var _ OpcuaOpenResponse = (*_OpcuaOpenResponse)(nil)
var _ MessagePDURequirements = (*_OpcuaOpenResponse)(nil)

// NewOpcuaOpenResponse factory function for _OpcuaOpenResponse
func NewOpcuaOpenResponse(chunk ChunkType, openResponse OpenChannelMessage, message Payload, totalLength uint32, binary bool) *_OpcuaOpenResponse {
	if openResponse == nil {
		panic("openResponse of type OpenChannelMessage for OpcuaOpenResponse must not be nil")
	}
	if message == nil {
		panic("message of type Payload for OpcuaOpenResponse must not be nil")
	}
	_result := &_OpcuaOpenResponse{
		MessagePDUContract: NewMessagePDU(chunk, binary),
		OpenResponse:       openResponse,
		Message:            message,
	}
	_result.MessagePDUContract.(*_MessagePDU)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// OpcuaOpenResponseBuilder is a builder for OpcuaOpenResponse
type OpcuaOpenResponseBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openResponse OpenChannelMessage, message Payload) OpcuaOpenResponseBuilder
	// WithOpenResponse adds OpenResponse (property field)
	WithOpenResponse(OpenChannelMessage) OpcuaOpenResponseBuilder
	// WithOpenResponseBuilder adds OpenResponse (property field) which is build by the builder
	WithOpenResponseBuilder(func(OpenChannelMessageBuilder) OpenChannelMessageBuilder) OpcuaOpenResponseBuilder
	// WithMessage adds Message (property field)
	WithMessage(Payload) OpcuaOpenResponseBuilder
	// WithMessageBuilder adds Message (property field) which is build by the builder
	WithMessageBuilder(func(PayloadBuilder) PayloadBuilder) OpcuaOpenResponseBuilder
	// Build builds the OpcuaOpenResponse or returns an error if something is wrong
	Build() (OpcuaOpenResponse, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() OpcuaOpenResponse
}

// NewOpcuaOpenResponseBuilder() creates a OpcuaOpenResponseBuilder
func NewOpcuaOpenResponseBuilder() OpcuaOpenResponseBuilder {
	return &_OpcuaOpenResponseBuilder{_OpcuaOpenResponse: new(_OpcuaOpenResponse)}
}

type _OpcuaOpenResponseBuilder struct {
	*_OpcuaOpenResponse

	parentBuilder *_MessagePDUBuilder

	err *utils.MultiError
}

var _ (OpcuaOpenResponseBuilder) = (*_OpcuaOpenResponseBuilder)(nil)

func (b *_OpcuaOpenResponseBuilder) setParent(contract MessagePDUContract) {
	b.MessagePDUContract = contract
}

func (b *_OpcuaOpenResponseBuilder) WithMandatoryFields(openResponse OpenChannelMessage, message Payload) OpcuaOpenResponseBuilder {
	return b.WithOpenResponse(openResponse).WithMessage(message)
}

func (b *_OpcuaOpenResponseBuilder) WithOpenResponse(openResponse OpenChannelMessage) OpcuaOpenResponseBuilder {
	b.OpenResponse = openResponse
	return b
}

func (b *_OpcuaOpenResponseBuilder) WithOpenResponseBuilder(builderSupplier func(OpenChannelMessageBuilder) OpenChannelMessageBuilder) OpcuaOpenResponseBuilder {
	builder := builderSupplier(b.OpenResponse.CreateOpenChannelMessageBuilder())
	var err error
	b.OpenResponse, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "OpenChannelMessageBuilder failed"))
	}
	return b
}

func (b *_OpcuaOpenResponseBuilder) WithMessage(message Payload) OpcuaOpenResponseBuilder {
	b.Message = message
	return b
}

func (b *_OpcuaOpenResponseBuilder) WithMessageBuilder(builderSupplier func(PayloadBuilder) PayloadBuilder) OpcuaOpenResponseBuilder {
	builder := builderSupplier(b.Message.CreatePayloadBuilder())
	var err error
	b.Message, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PayloadBuilder failed"))
	}
	return b
}

func (b *_OpcuaOpenResponseBuilder) Build() (OpcuaOpenResponse, error) {
	if b.OpenResponse == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openResponse' not set"))
	}
	if b.Message == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'message' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._OpcuaOpenResponse.deepCopy(), nil
}

func (b *_OpcuaOpenResponseBuilder) MustBuild() OpcuaOpenResponse {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_OpcuaOpenResponseBuilder) Done() MessagePDUBuilder {
	return b.parentBuilder
}

func (b *_OpcuaOpenResponseBuilder) buildForMessagePDU() (MessagePDU, error) {
	return b.Build()
}

func (b *_OpcuaOpenResponseBuilder) DeepCopy() any {
	_copy := b.CreateOpcuaOpenResponseBuilder().(*_OpcuaOpenResponseBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateOpcuaOpenResponseBuilder creates a OpcuaOpenResponseBuilder
func (b *_OpcuaOpenResponse) CreateOpcuaOpenResponseBuilder() OpcuaOpenResponseBuilder {
	if b == nil {
		return NewOpcuaOpenResponseBuilder()
	}
	return &_OpcuaOpenResponseBuilder{_OpcuaOpenResponse: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpcuaOpenResponse) GetMessageType() string {
	return "OPN"
}

func (m *_OpcuaOpenResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpcuaOpenResponse) GetParent() MessagePDUContract {
	return m.MessagePDUContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpcuaOpenResponse) GetOpenResponse() OpenChannelMessage {
	return m.OpenResponse
}

func (m *_OpcuaOpenResponse) GetMessage() Payload {
	return m.Message
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastOpcuaOpenResponse(structType any) OpcuaOpenResponse {
	if casted, ok := structType.(OpcuaOpenResponse); ok {
		return casted
	}
	if casted, ok := structType.(*OpcuaOpenResponse); ok {
		return *casted
	}
	return nil
}

func (m *_OpcuaOpenResponse) GetTypeName() string {
	return "OpcuaOpenResponse"
}

func (m *_OpcuaOpenResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MessagePDUContract.(*_MessagePDU).getLengthInBits(ctx))

	// Simple field (openResponse)
	lengthInBits += m.OpenResponse.GetLengthInBits(ctx)

	// Simple field (message)
	lengthInBits += m.Message.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_OpcuaOpenResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_OpcuaOpenResponse) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MessagePDU, totalLength uint32, response bool, binary bool) (__opcuaOpenResponse OpcuaOpenResponse, err error) {
	m.MessagePDUContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OpcuaOpenResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpcuaOpenResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openResponse, err := ReadSimpleField[OpenChannelMessage](ctx, "openResponse", ReadComplex[OpenChannelMessage](OpenChannelMessageParseWithBufferProducer[OpenChannelMessage]((bool)(response)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openResponse' field"))
	}
	m.OpenResponse = openResponse

	message, err := ReadSimpleField[Payload](ctx, "message", ReadComplex[Payload](PayloadParseWithBufferProducer[Payload]((bool)(binary), (uint32)(uint32(uint32(totalLength)-uint32(openResponse.GetLengthInBytes(ctx)))-uint32(uint32(16)))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'message' field"))
	}
	m.Message = message

	if closeErr := readBuffer.CloseContext("OpcuaOpenResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpcuaOpenResponse")
	}

	return m, nil
}

func (m *_OpcuaOpenResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpcuaOpenResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpcuaOpenResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpcuaOpenResponse")
		}

		if err := WriteSimpleField[OpenChannelMessage](ctx, "openResponse", m.GetOpenResponse(), WriteComplex[OpenChannelMessage](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openResponse' field")
		}

		if err := WriteSimpleField[Payload](ctx, "message", m.GetMessage(), WriteComplex[Payload](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'message' field")
		}

		if popErr := writeBuffer.PopContext("OpcuaOpenResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpcuaOpenResponse")
		}
		return nil
	}
	return m.MessagePDUContract.(*_MessagePDU).serializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_OpcuaOpenResponse) GetTotalLength() uint32 {
	return m.TotalLength
}

//
////

func (m *_OpcuaOpenResponse) IsOpcuaOpenResponse() {}

func (m *_OpcuaOpenResponse) DeepCopy() any {
	return m.deepCopy()
}

func (m *_OpcuaOpenResponse) deepCopy() *_OpcuaOpenResponse {
	if m == nil {
		return nil
	}
	_OpcuaOpenResponseCopy := &_OpcuaOpenResponse{
		m.MessagePDUContract.(*_MessagePDU).deepCopy(),
		utils.DeepCopy[OpenChannelMessage](m.OpenResponse),
		utils.DeepCopy[Payload](m.Message),
		m.TotalLength,
	}
	m.MessagePDUContract.(*_MessagePDU)._SubType = m
	return _OpcuaOpenResponseCopy
}

func (m *_OpcuaOpenResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
