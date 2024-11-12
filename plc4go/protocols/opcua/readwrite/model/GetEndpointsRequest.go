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

// GetEndpointsRequest is the corresponding interface of GetEndpointsRequest
type GetEndpointsRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() RequestHeader
	// GetEndpointUrl returns EndpointUrl (property field)
	GetEndpointUrl() PascalString
	// GetLocaleIds returns LocaleIds (property field)
	GetLocaleIds() []PascalString
	// GetProfileUris returns ProfileUris (property field)
	GetProfileUris() []PascalString
	// IsGetEndpointsRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsGetEndpointsRequest()
	// CreateBuilder creates a GetEndpointsRequestBuilder
	CreateGetEndpointsRequestBuilder() GetEndpointsRequestBuilder
}

// _GetEndpointsRequest is the data-structure of this message
type _GetEndpointsRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader RequestHeader
	EndpointUrl   PascalString
	LocaleIds     []PascalString
	ProfileUris   []PascalString
}

var _ GetEndpointsRequest = (*_GetEndpointsRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_GetEndpointsRequest)(nil)

// NewGetEndpointsRequest factory function for _GetEndpointsRequest
func NewGetEndpointsRequest(requestHeader RequestHeader, endpointUrl PascalString, localeIds []PascalString, profileUris []PascalString) *_GetEndpointsRequest {
	if requestHeader == nil {
		panic("requestHeader of type RequestHeader for GetEndpointsRequest must not be nil")
	}
	if endpointUrl == nil {
		panic("endpointUrl of type PascalString for GetEndpointsRequest must not be nil")
	}
	_result := &_GetEndpointsRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		EndpointUrl:                       endpointUrl,
		LocaleIds:                         localeIds,
		ProfileUris:                       profileUris,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// GetEndpointsRequestBuilder is a builder for GetEndpointsRequest
type GetEndpointsRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader RequestHeader, endpointUrl PascalString, localeIds []PascalString, profileUris []PascalString) GetEndpointsRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(RequestHeader) GetEndpointsRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(RequestHeaderBuilder) RequestHeaderBuilder) GetEndpointsRequestBuilder
	// WithEndpointUrl adds EndpointUrl (property field)
	WithEndpointUrl(PascalString) GetEndpointsRequestBuilder
	// WithEndpointUrlBuilder adds EndpointUrl (property field) which is build by the builder
	WithEndpointUrlBuilder(func(PascalStringBuilder) PascalStringBuilder) GetEndpointsRequestBuilder
	// WithLocaleIds adds LocaleIds (property field)
	WithLocaleIds(...PascalString) GetEndpointsRequestBuilder
	// WithProfileUris adds ProfileUris (property field)
	WithProfileUris(...PascalString) GetEndpointsRequestBuilder
	// Build builds the GetEndpointsRequest or returns an error if something is wrong
	Build() (GetEndpointsRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() GetEndpointsRequest
}

// NewGetEndpointsRequestBuilder() creates a GetEndpointsRequestBuilder
func NewGetEndpointsRequestBuilder() GetEndpointsRequestBuilder {
	return &_GetEndpointsRequestBuilder{_GetEndpointsRequest: new(_GetEndpointsRequest)}
}

type _GetEndpointsRequestBuilder struct {
	*_GetEndpointsRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (GetEndpointsRequestBuilder) = (*_GetEndpointsRequestBuilder)(nil)

func (b *_GetEndpointsRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_GetEndpointsRequestBuilder) WithMandatoryFields(requestHeader RequestHeader, endpointUrl PascalString, localeIds []PascalString, profileUris []PascalString) GetEndpointsRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithEndpointUrl(endpointUrl).WithLocaleIds(localeIds...).WithProfileUris(profileUris...)
}

func (b *_GetEndpointsRequestBuilder) WithRequestHeader(requestHeader RequestHeader) GetEndpointsRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_GetEndpointsRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(RequestHeaderBuilder) RequestHeaderBuilder) GetEndpointsRequestBuilder {
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

func (b *_GetEndpointsRequestBuilder) WithEndpointUrl(endpointUrl PascalString) GetEndpointsRequestBuilder {
	b.EndpointUrl = endpointUrl
	return b
}

func (b *_GetEndpointsRequestBuilder) WithEndpointUrlBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) GetEndpointsRequestBuilder {
	builder := builderSupplier(b.EndpointUrl.CreatePascalStringBuilder())
	var err error
	b.EndpointUrl, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_GetEndpointsRequestBuilder) WithLocaleIds(localeIds ...PascalString) GetEndpointsRequestBuilder {
	b.LocaleIds = localeIds
	return b
}

func (b *_GetEndpointsRequestBuilder) WithProfileUris(profileUris ...PascalString) GetEndpointsRequestBuilder {
	b.ProfileUris = profileUris
	return b
}

func (b *_GetEndpointsRequestBuilder) Build() (GetEndpointsRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.EndpointUrl == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'endpointUrl' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._GetEndpointsRequest.deepCopy(), nil
}

func (b *_GetEndpointsRequestBuilder) MustBuild() GetEndpointsRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_GetEndpointsRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_GetEndpointsRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_GetEndpointsRequestBuilder) DeepCopy() any {
	_copy := b.CreateGetEndpointsRequestBuilder().(*_GetEndpointsRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateGetEndpointsRequestBuilder creates a GetEndpointsRequestBuilder
func (b *_GetEndpointsRequest) CreateGetEndpointsRequestBuilder() GetEndpointsRequestBuilder {
	if b == nil {
		return NewGetEndpointsRequestBuilder()
	}
	return &_GetEndpointsRequestBuilder{_GetEndpointsRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_GetEndpointsRequest) GetExtensionId() int32 {
	return int32(428)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_GetEndpointsRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GetEndpointsRequest) GetRequestHeader() RequestHeader {
	return m.RequestHeader
}

func (m *_GetEndpointsRequest) GetEndpointUrl() PascalString {
	return m.EndpointUrl
}

func (m *_GetEndpointsRequest) GetLocaleIds() []PascalString {
	return m.LocaleIds
}

func (m *_GetEndpointsRequest) GetProfileUris() []PascalString {
	return m.ProfileUris
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastGetEndpointsRequest(structType any) GetEndpointsRequest {
	if casted, ok := structType.(GetEndpointsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*GetEndpointsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_GetEndpointsRequest) GetTypeName() string {
	return "GetEndpointsRequest"
}

func (m *_GetEndpointsRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (endpointUrl)
	lengthInBits += m.EndpointUrl.GetLengthInBits(ctx)

	// Implicit Field (noOfLocaleIds)
	lengthInBits += 32

	// Array field
	if len(m.LocaleIds) > 0 {
		for _curItem, element := range m.LocaleIds {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.LocaleIds), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Implicit Field (noOfProfileUris)
	lengthInBits += 32

	// Array field
	if len(m.ProfileUris) > 0 {
		for _curItem, element := range m.ProfileUris {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ProfileUris), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_GetEndpointsRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_GetEndpointsRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__getEndpointsRequest GetEndpointsRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GetEndpointsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GetEndpointsRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[RequestHeader](ctx, "requestHeader", ReadComplex[RequestHeader](ExtensionObjectDefinitionParseWithBufferProducer[RequestHeader]((int32)(int32(391))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	endpointUrl, err := ReadSimpleField[PascalString](ctx, "endpointUrl", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'endpointUrl' field"))
	}
	m.EndpointUrl = endpointUrl

	noOfLocaleIds, err := ReadImplicitField[int32](ctx, "noOfLocaleIds", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfLocaleIds' field"))
	}
	_ = noOfLocaleIds

	localeIds, err := ReadCountArrayField[PascalString](ctx, "localeIds", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfLocaleIds))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localeIds' field"))
	}
	m.LocaleIds = localeIds

	noOfProfileUris, err := ReadImplicitField[int32](ctx, "noOfProfileUris", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfProfileUris' field"))
	}
	_ = noOfProfileUris

	profileUris, err := ReadCountArrayField[PascalString](ctx, "profileUris", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer), uint64(noOfProfileUris))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'profileUris' field"))
	}
	m.ProfileUris = profileUris

	if closeErr := readBuffer.CloseContext("GetEndpointsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GetEndpointsRequest")
	}

	return m, nil
}

func (m *_GetEndpointsRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GetEndpointsRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("GetEndpointsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for GetEndpointsRequest")
		}

		if err := WriteSimpleField[RequestHeader](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[RequestHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "endpointUrl", m.GetEndpointUrl(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'endpointUrl' field")
		}
		noOfLocaleIds := int32(utils.InlineIf(bool((m.GetLocaleIds()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetLocaleIds()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfLocaleIds", noOfLocaleIds, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfLocaleIds' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "localeIds", m.GetLocaleIds(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'localeIds' field")
		}
		noOfProfileUris := int32(utils.InlineIf(bool((m.GetProfileUris()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetProfileUris()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfProfileUris", noOfProfileUris, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfProfileUris' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "profileUris", m.GetProfileUris(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'profileUris' field")
		}

		if popErr := writeBuffer.PopContext("GetEndpointsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for GetEndpointsRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_GetEndpointsRequest) IsGetEndpointsRequest() {}

func (m *_GetEndpointsRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_GetEndpointsRequest) deepCopy() *_GetEndpointsRequest {
	if m == nil {
		return nil
	}
	_GetEndpointsRequestCopy := &_GetEndpointsRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[RequestHeader](m.RequestHeader),
		utils.DeepCopy[PascalString](m.EndpointUrl),
		utils.DeepCopySlice[PascalString, PascalString](m.LocaleIds),
		utils.DeepCopySlice[PascalString, PascalString](m.ProfileUris),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _GetEndpointsRequestCopy
}

func (m *_GetEndpointsRequest) String() string {
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
