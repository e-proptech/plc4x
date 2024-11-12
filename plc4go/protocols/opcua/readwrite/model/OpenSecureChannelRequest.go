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

// OpenSecureChannelRequest is the corresponding interface of OpenSecureChannelRequest
type OpenSecureChannelRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() RequestHeader
	// GetClientProtocolVersion returns ClientProtocolVersion (property field)
	GetClientProtocolVersion() uint32
	// GetRequestType returns RequestType (property field)
	GetRequestType() SecurityTokenRequestType
	// GetSecurityMode returns SecurityMode (property field)
	GetSecurityMode() MessageSecurityMode
	// GetClientNonce returns ClientNonce (property field)
	GetClientNonce() PascalByteString
	// GetRequestedLifetime returns RequestedLifetime (property field)
	GetRequestedLifetime() uint32
	// IsOpenSecureChannelRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsOpenSecureChannelRequest()
	// CreateBuilder creates a OpenSecureChannelRequestBuilder
	CreateOpenSecureChannelRequestBuilder() OpenSecureChannelRequestBuilder
}

// _OpenSecureChannelRequest is the data-structure of this message
type _OpenSecureChannelRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader         RequestHeader
	ClientProtocolVersion uint32
	RequestType           SecurityTokenRequestType
	SecurityMode          MessageSecurityMode
	ClientNonce           PascalByteString
	RequestedLifetime     uint32
}

var _ OpenSecureChannelRequest = (*_OpenSecureChannelRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_OpenSecureChannelRequest)(nil)

// NewOpenSecureChannelRequest factory function for _OpenSecureChannelRequest
func NewOpenSecureChannelRequest(requestHeader RequestHeader, clientProtocolVersion uint32, requestType SecurityTokenRequestType, securityMode MessageSecurityMode, clientNonce PascalByteString, requestedLifetime uint32) *_OpenSecureChannelRequest {
	if requestHeader == nil {
		panic("requestHeader of type RequestHeader for OpenSecureChannelRequest must not be nil")
	}
	if clientNonce == nil {
		panic("clientNonce of type PascalByteString for OpenSecureChannelRequest must not be nil")
	}
	_result := &_OpenSecureChannelRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		ClientProtocolVersion:             clientProtocolVersion,
		RequestType:                       requestType,
		SecurityMode:                      securityMode,
		ClientNonce:                       clientNonce,
		RequestedLifetime:                 requestedLifetime,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// OpenSecureChannelRequestBuilder is a builder for OpenSecureChannelRequest
type OpenSecureChannelRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader RequestHeader, clientProtocolVersion uint32, requestType SecurityTokenRequestType, securityMode MessageSecurityMode, clientNonce PascalByteString, requestedLifetime uint32) OpenSecureChannelRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(RequestHeader) OpenSecureChannelRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(RequestHeaderBuilder) RequestHeaderBuilder) OpenSecureChannelRequestBuilder
	// WithClientProtocolVersion adds ClientProtocolVersion (property field)
	WithClientProtocolVersion(uint32) OpenSecureChannelRequestBuilder
	// WithRequestType adds RequestType (property field)
	WithRequestType(SecurityTokenRequestType) OpenSecureChannelRequestBuilder
	// WithSecurityMode adds SecurityMode (property field)
	WithSecurityMode(MessageSecurityMode) OpenSecureChannelRequestBuilder
	// WithClientNonce adds ClientNonce (property field)
	WithClientNonce(PascalByteString) OpenSecureChannelRequestBuilder
	// WithClientNonceBuilder adds ClientNonce (property field) which is build by the builder
	WithClientNonceBuilder(func(PascalByteStringBuilder) PascalByteStringBuilder) OpenSecureChannelRequestBuilder
	// WithRequestedLifetime adds RequestedLifetime (property field)
	WithRequestedLifetime(uint32) OpenSecureChannelRequestBuilder
	// Build builds the OpenSecureChannelRequest or returns an error if something is wrong
	Build() (OpenSecureChannelRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() OpenSecureChannelRequest
}

// NewOpenSecureChannelRequestBuilder() creates a OpenSecureChannelRequestBuilder
func NewOpenSecureChannelRequestBuilder() OpenSecureChannelRequestBuilder {
	return &_OpenSecureChannelRequestBuilder{_OpenSecureChannelRequest: new(_OpenSecureChannelRequest)}
}

type _OpenSecureChannelRequestBuilder struct {
	*_OpenSecureChannelRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (OpenSecureChannelRequestBuilder) = (*_OpenSecureChannelRequestBuilder)(nil)

func (b *_OpenSecureChannelRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_OpenSecureChannelRequestBuilder) WithMandatoryFields(requestHeader RequestHeader, clientProtocolVersion uint32, requestType SecurityTokenRequestType, securityMode MessageSecurityMode, clientNonce PascalByteString, requestedLifetime uint32) OpenSecureChannelRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithClientProtocolVersion(clientProtocolVersion).WithRequestType(requestType).WithSecurityMode(securityMode).WithClientNonce(clientNonce).WithRequestedLifetime(requestedLifetime)
}

func (b *_OpenSecureChannelRequestBuilder) WithRequestHeader(requestHeader RequestHeader) OpenSecureChannelRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_OpenSecureChannelRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(RequestHeaderBuilder) RequestHeaderBuilder) OpenSecureChannelRequestBuilder {
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

func (b *_OpenSecureChannelRequestBuilder) WithClientProtocolVersion(clientProtocolVersion uint32) OpenSecureChannelRequestBuilder {
	b.ClientProtocolVersion = clientProtocolVersion
	return b
}

func (b *_OpenSecureChannelRequestBuilder) WithRequestType(requestType SecurityTokenRequestType) OpenSecureChannelRequestBuilder {
	b.RequestType = requestType
	return b
}

func (b *_OpenSecureChannelRequestBuilder) WithSecurityMode(securityMode MessageSecurityMode) OpenSecureChannelRequestBuilder {
	b.SecurityMode = securityMode
	return b
}

func (b *_OpenSecureChannelRequestBuilder) WithClientNonce(clientNonce PascalByteString) OpenSecureChannelRequestBuilder {
	b.ClientNonce = clientNonce
	return b
}

func (b *_OpenSecureChannelRequestBuilder) WithClientNonceBuilder(builderSupplier func(PascalByteStringBuilder) PascalByteStringBuilder) OpenSecureChannelRequestBuilder {
	builder := builderSupplier(b.ClientNonce.CreatePascalByteStringBuilder())
	var err error
	b.ClientNonce, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalByteStringBuilder failed"))
	}
	return b
}

func (b *_OpenSecureChannelRequestBuilder) WithRequestedLifetime(requestedLifetime uint32) OpenSecureChannelRequestBuilder {
	b.RequestedLifetime = requestedLifetime
	return b
}

func (b *_OpenSecureChannelRequestBuilder) Build() (OpenSecureChannelRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.ClientNonce == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'clientNonce' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._OpenSecureChannelRequest.deepCopy(), nil
}

func (b *_OpenSecureChannelRequestBuilder) MustBuild() OpenSecureChannelRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_OpenSecureChannelRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_OpenSecureChannelRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_OpenSecureChannelRequestBuilder) DeepCopy() any {
	_copy := b.CreateOpenSecureChannelRequestBuilder().(*_OpenSecureChannelRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateOpenSecureChannelRequestBuilder creates a OpenSecureChannelRequestBuilder
func (b *_OpenSecureChannelRequest) CreateOpenSecureChannelRequestBuilder() OpenSecureChannelRequestBuilder {
	if b == nil {
		return NewOpenSecureChannelRequestBuilder()
	}
	return &_OpenSecureChannelRequestBuilder{_OpenSecureChannelRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_OpenSecureChannelRequest) GetExtensionId() int32 {
	return int32(446)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_OpenSecureChannelRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_OpenSecureChannelRequest) GetRequestHeader() RequestHeader {
	return m.RequestHeader
}

func (m *_OpenSecureChannelRequest) GetClientProtocolVersion() uint32 {
	return m.ClientProtocolVersion
}

func (m *_OpenSecureChannelRequest) GetRequestType() SecurityTokenRequestType {
	return m.RequestType
}

func (m *_OpenSecureChannelRequest) GetSecurityMode() MessageSecurityMode {
	return m.SecurityMode
}

func (m *_OpenSecureChannelRequest) GetClientNonce() PascalByteString {
	return m.ClientNonce
}

func (m *_OpenSecureChannelRequest) GetRequestedLifetime() uint32 {
	return m.RequestedLifetime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastOpenSecureChannelRequest(structType any) OpenSecureChannelRequest {
	if casted, ok := structType.(OpenSecureChannelRequest); ok {
		return casted
	}
	if casted, ok := structType.(*OpenSecureChannelRequest); ok {
		return *casted
	}
	return nil
}

func (m *_OpenSecureChannelRequest) GetTypeName() string {
	return "OpenSecureChannelRequest"
}

func (m *_OpenSecureChannelRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (clientProtocolVersion)
	lengthInBits += 32

	// Simple field (requestType)
	lengthInBits += 32

	// Simple field (securityMode)
	lengthInBits += 32

	// Simple field (clientNonce)
	lengthInBits += m.ClientNonce.GetLengthInBits(ctx)

	// Simple field (requestedLifetime)
	lengthInBits += 32

	return lengthInBits
}

func (m *_OpenSecureChannelRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_OpenSecureChannelRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__openSecureChannelRequest OpenSecureChannelRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("OpenSecureChannelRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for OpenSecureChannelRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[RequestHeader](ctx, "requestHeader", ReadComplex[RequestHeader](ExtensionObjectDefinitionParseWithBufferProducer[RequestHeader]((int32)(int32(391))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	clientProtocolVersion, err := ReadSimpleField(ctx, "clientProtocolVersion", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'clientProtocolVersion' field"))
	}
	m.ClientProtocolVersion = clientProtocolVersion

	requestType, err := ReadEnumField[SecurityTokenRequestType](ctx, "requestType", "SecurityTokenRequestType", ReadEnum(SecurityTokenRequestTypeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestType' field"))
	}
	m.RequestType = requestType

	securityMode, err := ReadEnumField[MessageSecurityMode](ctx, "securityMode", "MessageSecurityMode", ReadEnum(MessageSecurityModeByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'securityMode' field"))
	}
	m.SecurityMode = securityMode

	clientNonce, err := ReadSimpleField[PascalByteString](ctx, "clientNonce", ReadComplex[PascalByteString](PascalByteStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'clientNonce' field"))
	}
	m.ClientNonce = clientNonce

	requestedLifetime, err := ReadSimpleField(ctx, "requestedLifetime", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestedLifetime' field"))
	}
	m.RequestedLifetime = requestedLifetime

	if closeErr := readBuffer.CloseContext("OpenSecureChannelRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for OpenSecureChannelRequest")
	}

	return m, nil
}

func (m *_OpenSecureChannelRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_OpenSecureChannelRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("OpenSecureChannelRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for OpenSecureChannelRequest")
		}

		if err := WriteSimpleField[RequestHeader](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[RequestHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[uint32](ctx, "clientProtocolVersion", m.GetClientProtocolVersion(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'clientProtocolVersion' field")
		}

		if err := WriteSimpleEnumField[SecurityTokenRequestType](ctx, "requestType", "SecurityTokenRequestType", m.GetRequestType(), WriteEnum[SecurityTokenRequestType, uint32](SecurityTokenRequestType.GetValue, SecurityTokenRequestType.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'requestType' field")
		}

		if err := WriteSimpleEnumField[MessageSecurityMode](ctx, "securityMode", "MessageSecurityMode", m.GetSecurityMode(), WriteEnum[MessageSecurityMode, uint32](MessageSecurityMode.GetValue, MessageSecurityMode.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'securityMode' field")
		}

		if err := WriteSimpleField[PascalByteString](ctx, "clientNonce", m.GetClientNonce(), WriteComplex[PascalByteString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'clientNonce' field")
		}

		if err := WriteSimpleField[uint32](ctx, "requestedLifetime", m.GetRequestedLifetime(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestedLifetime' field")
		}

		if popErr := writeBuffer.PopContext("OpenSecureChannelRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for OpenSecureChannelRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_OpenSecureChannelRequest) IsOpenSecureChannelRequest() {}

func (m *_OpenSecureChannelRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_OpenSecureChannelRequest) deepCopy() *_OpenSecureChannelRequest {
	if m == nil {
		return nil
	}
	_OpenSecureChannelRequestCopy := &_OpenSecureChannelRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[RequestHeader](m.RequestHeader),
		m.ClientProtocolVersion,
		m.RequestType,
		m.SecurityMode,
		utils.DeepCopy[PascalByteString](m.ClientNonce),
		m.RequestedLifetime,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _OpenSecureChannelRequestCopy
}

func (m *_OpenSecureChannelRequest) String() string {
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
