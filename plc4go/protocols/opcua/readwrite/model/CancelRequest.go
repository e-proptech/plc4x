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

// CancelRequest is the corresponding interface of CancelRequest
type CancelRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() RequestHeader
	// GetRequestHandle returns RequestHandle (property field)
	GetRequestHandle() uint32
	// IsCancelRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCancelRequest()
	// CreateBuilder creates a CancelRequestBuilder
	CreateCancelRequestBuilder() CancelRequestBuilder
}

// _CancelRequest is the data-structure of this message
type _CancelRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader RequestHeader
	RequestHandle uint32
}

var _ CancelRequest = (*_CancelRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CancelRequest)(nil)

// NewCancelRequest factory function for _CancelRequest
func NewCancelRequest(requestHeader RequestHeader, requestHandle uint32) *_CancelRequest {
	if requestHeader == nil {
		panic("requestHeader of type RequestHeader for CancelRequest must not be nil")
	}
	_result := &_CancelRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		RequestHandle:                     requestHandle,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CancelRequestBuilder is a builder for CancelRequest
type CancelRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader RequestHeader, requestHandle uint32) CancelRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(RequestHeader) CancelRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(RequestHeaderBuilder) RequestHeaderBuilder) CancelRequestBuilder
	// WithRequestHandle adds RequestHandle (property field)
	WithRequestHandle(uint32) CancelRequestBuilder
	// Build builds the CancelRequest or returns an error if something is wrong
	Build() (CancelRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CancelRequest
}

// NewCancelRequestBuilder() creates a CancelRequestBuilder
func NewCancelRequestBuilder() CancelRequestBuilder {
	return &_CancelRequestBuilder{_CancelRequest: new(_CancelRequest)}
}

type _CancelRequestBuilder struct {
	*_CancelRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (CancelRequestBuilder) = (*_CancelRequestBuilder)(nil)

func (b *_CancelRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_CancelRequestBuilder) WithMandatoryFields(requestHeader RequestHeader, requestHandle uint32) CancelRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithRequestHandle(requestHandle)
}

func (b *_CancelRequestBuilder) WithRequestHeader(requestHeader RequestHeader) CancelRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_CancelRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(RequestHeaderBuilder) RequestHeaderBuilder) CancelRequestBuilder {
	builder := builderSupplier(b.RequestHeader.CreateRequestHeaderBuilder())
	var err error
	b.RequestHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "RequestHeaderBuilder failed"))
	}
	return b
}

func (b *_CancelRequestBuilder) WithRequestHandle(requestHandle uint32) CancelRequestBuilder {
	b.RequestHandle = requestHandle
	return b
}

func (b *_CancelRequestBuilder) Build() (CancelRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CancelRequest.deepCopy(), nil
}

func (b *_CancelRequestBuilder) MustBuild() CancelRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_CancelRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_CancelRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_CancelRequestBuilder) DeepCopy() any {
	_copy := b.CreateCancelRequestBuilder().(*_CancelRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCancelRequestBuilder creates a CancelRequestBuilder
func (b *_CancelRequest) CreateCancelRequestBuilder() CancelRequestBuilder {
	if b == nil {
		return NewCancelRequestBuilder()
	}
	return &_CancelRequestBuilder{_CancelRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CancelRequest) GetExtensionId() int32 {
	return int32(479)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CancelRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CancelRequest) GetRequestHeader() RequestHeader {
	return m.RequestHeader
}

func (m *_CancelRequest) GetRequestHandle() uint32 {
	return m.RequestHandle
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCancelRequest(structType any) CancelRequest {
	if casted, ok := structType.(CancelRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CancelRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CancelRequest) GetTypeName() string {
	return "CancelRequest"
}

func (m *_CancelRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (requestHandle)
	lengthInBits += 32

	return lengthInBits
}

func (m *_CancelRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CancelRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__cancelRequest CancelRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CancelRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CancelRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[RequestHeader](ctx, "requestHeader", ReadComplex[RequestHeader](ExtensionObjectDefinitionParseWithBufferProducer[RequestHeader]((int32)(int32(391))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	requestHandle, err := ReadSimpleField(ctx, "requestHandle", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHandle' field"))
	}
	m.RequestHandle = requestHandle

	if closeErr := readBuffer.CloseContext("CancelRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CancelRequest")
	}

	return m, nil
}

func (m *_CancelRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CancelRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CancelRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CancelRequest")
		}

		if err := WriteSimpleField[RequestHeader](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[RequestHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[uint32](ctx, "requestHandle", m.GetRequestHandle(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHandle' field")
		}

		if popErr := writeBuffer.PopContext("CancelRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CancelRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CancelRequest) IsCancelRequest() {}

func (m *_CancelRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CancelRequest) deepCopy() *_CancelRequest {
	if m == nil {
		return nil
	}
	_CancelRequestCopy := &_CancelRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[RequestHeader](m.RequestHeader),
		m.RequestHandle,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _CancelRequestCopy
}

func (m *_CancelRequest) String() string {
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
