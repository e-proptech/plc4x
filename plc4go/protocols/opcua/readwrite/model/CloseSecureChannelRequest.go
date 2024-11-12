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

// CloseSecureChannelRequest is the corresponding interface of CloseSecureChannelRequest
type CloseSecureChannelRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() RequestHeader
	// IsCloseSecureChannelRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCloseSecureChannelRequest()
	// CreateBuilder creates a CloseSecureChannelRequestBuilder
	CreateCloseSecureChannelRequestBuilder() CloseSecureChannelRequestBuilder
}

// _CloseSecureChannelRequest is the data-structure of this message
type _CloseSecureChannelRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader RequestHeader
}

var _ CloseSecureChannelRequest = (*_CloseSecureChannelRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CloseSecureChannelRequest)(nil)

// NewCloseSecureChannelRequest factory function for _CloseSecureChannelRequest
func NewCloseSecureChannelRequest(requestHeader RequestHeader) *_CloseSecureChannelRequest {
	if requestHeader == nil {
		panic("requestHeader of type RequestHeader for CloseSecureChannelRequest must not be nil")
	}
	_result := &_CloseSecureChannelRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CloseSecureChannelRequestBuilder is a builder for CloseSecureChannelRequest
type CloseSecureChannelRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader RequestHeader) CloseSecureChannelRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(RequestHeader) CloseSecureChannelRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(RequestHeaderBuilder) RequestHeaderBuilder) CloseSecureChannelRequestBuilder
	// Build builds the CloseSecureChannelRequest or returns an error if something is wrong
	Build() (CloseSecureChannelRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CloseSecureChannelRequest
}

// NewCloseSecureChannelRequestBuilder() creates a CloseSecureChannelRequestBuilder
func NewCloseSecureChannelRequestBuilder() CloseSecureChannelRequestBuilder {
	return &_CloseSecureChannelRequestBuilder{_CloseSecureChannelRequest: new(_CloseSecureChannelRequest)}
}

type _CloseSecureChannelRequestBuilder struct {
	*_CloseSecureChannelRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (CloseSecureChannelRequestBuilder) = (*_CloseSecureChannelRequestBuilder)(nil)

func (b *_CloseSecureChannelRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_CloseSecureChannelRequestBuilder) WithMandatoryFields(requestHeader RequestHeader) CloseSecureChannelRequestBuilder {
	return b.WithRequestHeader(requestHeader)
}

func (b *_CloseSecureChannelRequestBuilder) WithRequestHeader(requestHeader RequestHeader) CloseSecureChannelRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_CloseSecureChannelRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(RequestHeaderBuilder) RequestHeaderBuilder) CloseSecureChannelRequestBuilder {
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

func (b *_CloseSecureChannelRequestBuilder) Build() (CloseSecureChannelRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CloseSecureChannelRequest.deepCopy(), nil
}

func (b *_CloseSecureChannelRequestBuilder) MustBuild() CloseSecureChannelRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_CloseSecureChannelRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_CloseSecureChannelRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_CloseSecureChannelRequestBuilder) DeepCopy() any {
	_copy := b.CreateCloseSecureChannelRequestBuilder().(*_CloseSecureChannelRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCloseSecureChannelRequestBuilder creates a CloseSecureChannelRequestBuilder
func (b *_CloseSecureChannelRequest) CreateCloseSecureChannelRequestBuilder() CloseSecureChannelRequestBuilder {
	if b == nil {
		return NewCloseSecureChannelRequestBuilder()
	}
	return &_CloseSecureChannelRequestBuilder{_CloseSecureChannelRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CloseSecureChannelRequest) GetExtensionId() int32 {
	return int32(452)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CloseSecureChannelRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CloseSecureChannelRequest) GetRequestHeader() RequestHeader {
	return m.RequestHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCloseSecureChannelRequest(structType any) CloseSecureChannelRequest {
	if casted, ok := structType.(CloseSecureChannelRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CloseSecureChannelRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CloseSecureChannelRequest) GetTypeName() string {
	return "CloseSecureChannelRequest"
}

func (m *_CloseSecureChannelRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CloseSecureChannelRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CloseSecureChannelRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__closeSecureChannelRequest CloseSecureChannelRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CloseSecureChannelRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CloseSecureChannelRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[RequestHeader](ctx, "requestHeader", ReadComplex[RequestHeader](ExtensionObjectDefinitionParseWithBufferProducer[RequestHeader]((int32)(int32(391))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	if closeErr := readBuffer.CloseContext("CloseSecureChannelRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CloseSecureChannelRequest")
	}

	return m, nil
}

func (m *_CloseSecureChannelRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CloseSecureChannelRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CloseSecureChannelRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CloseSecureChannelRequest")
		}

		if err := WriteSimpleField[RequestHeader](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[RequestHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if popErr := writeBuffer.PopContext("CloseSecureChannelRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CloseSecureChannelRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CloseSecureChannelRequest) IsCloseSecureChannelRequest() {}

func (m *_CloseSecureChannelRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CloseSecureChannelRequest) deepCopy() *_CloseSecureChannelRequest {
	if m == nil {
		return nil
	}
	_CloseSecureChannelRequestCopy := &_CloseSecureChannelRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[RequestHeader](m.RequestHeader),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _CloseSecureChannelRequestCopy
}

func (m *_CloseSecureChannelRequest) String() string {
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
