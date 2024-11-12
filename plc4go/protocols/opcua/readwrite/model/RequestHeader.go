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

// RequestHeader is the corresponding interface of RequestHeader
type RequestHeader interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetAuthenticationToken returns AuthenticationToken (property field)
	GetAuthenticationToken() NodeId
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() int64
	// GetRequestHandle returns RequestHandle (property field)
	GetRequestHandle() uint32
	// GetReturnDiagnostics returns ReturnDiagnostics (property field)
	GetReturnDiagnostics() uint32
	// GetAuditEntryId returns AuditEntryId (property field)
	GetAuditEntryId() PascalString
	// GetTimeoutHint returns TimeoutHint (property field)
	GetTimeoutHint() uint32
	// GetAdditionalHeader returns AdditionalHeader (property field)
	GetAdditionalHeader() ExtensionObject
	// IsRequestHeader is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsRequestHeader()
	// CreateBuilder creates a RequestHeaderBuilder
	CreateRequestHeaderBuilder() RequestHeaderBuilder
}

// _RequestHeader is the data-structure of this message
type _RequestHeader struct {
	ExtensionObjectDefinitionContract
	AuthenticationToken NodeId
	Timestamp           int64
	RequestHandle       uint32
	ReturnDiagnostics   uint32
	AuditEntryId        PascalString
	TimeoutHint         uint32
	AdditionalHeader    ExtensionObject
}

var _ RequestHeader = (*_RequestHeader)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_RequestHeader)(nil)

// NewRequestHeader factory function for _RequestHeader
func NewRequestHeader(authenticationToken NodeId, timestamp int64, requestHandle uint32, returnDiagnostics uint32, auditEntryId PascalString, timeoutHint uint32, additionalHeader ExtensionObject) *_RequestHeader {
	if authenticationToken == nil {
		panic("authenticationToken of type NodeId for RequestHeader must not be nil")
	}
	if auditEntryId == nil {
		panic("auditEntryId of type PascalString for RequestHeader must not be nil")
	}
	if additionalHeader == nil {
		panic("additionalHeader of type ExtensionObject for RequestHeader must not be nil")
	}
	_result := &_RequestHeader{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		AuthenticationToken:               authenticationToken,
		Timestamp:                         timestamp,
		RequestHandle:                     requestHandle,
		ReturnDiagnostics:                 returnDiagnostics,
		AuditEntryId:                      auditEntryId,
		TimeoutHint:                       timeoutHint,
		AdditionalHeader:                  additionalHeader,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// RequestHeaderBuilder is a builder for RequestHeader
type RequestHeaderBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(authenticationToken NodeId, timestamp int64, requestHandle uint32, returnDiagnostics uint32, auditEntryId PascalString, timeoutHint uint32, additionalHeader ExtensionObject) RequestHeaderBuilder
	// WithAuthenticationToken adds AuthenticationToken (property field)
	WithAuthenticationToken(NodeId) RequestHeaderBuilder
	// WithAuthenticationTokenBuilder adds AuthenticationToken (property field) which is build by the builder
	WithAuthenticationTokenBuilder(func(NodeIdBuilder) NodeIdBuilder) RequestHeaderBuilder
	// WithTimestamp adds Timestamp (property field)
	WithTimestamp(int64) RequestHeaderBuilder
	// WithRequestHandle adds RequestHandle (property field)
	WithRequestHandle(uint32) RequestHeaderBuilder
	// WithReturnDiagnostics adds ReturnDiagnostics (property field)
	WithReturnDiagnostics(uint32) RequestHeaderBuilder
	// WithAuditEntryId adds AuditEntryId (property field)
	WithAuditEntryId(PascalString) RequestHeaderBuilder
	// WithAuditEntryIdBuilder adds AuditEntryId (property field) which is build by the builder
	WithAuditEntryIdBuilder(func(PascalStringBuilder) PascalStringBuilder) RequestHeaderBuilder
	// WithTimeoutHint adds TimeoutHint (property field)
	WithTimeoutHint(uint32) RequestHeaderBuilder
	// WithAdditionalHeader adds AdditionalHeader (property field)
	WithAdditionalHeader(ExtensionObject) RequestHeaderBuilder
	// WithAdditionalHeaderBuilder adds AdditionalHeader (property field) which is build by the builder
	WithAdditionalHeaderBuilder(func(ExtensionObjectBuilder) ExtensionObjectBuilder) RequestHeaderBuilder
	// Build builds the RequestHeader or returns an error if something is wrong
	Build() (RequestHeader, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() RequestHeader
}

// NewRequestHeaderBuilder() creates a RequestHeaderBuilder
func NewRequestHeaderBuilder() RequestHeaderBuilder {
	return &_RequestHeaderBuilder{_RequestHeader: new(_RequestHeader)}
}

type _RequestHeaderBuilder struct {
	*_RequestHeader

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (RequestHeaderBuilder) = (*_RequestHeaderBuilder)(nil)

func (b *_RequestHeaderBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_RequestHeaderBuilder) WithMandatoryFields(authenticationToken NodeId, timestamp int64, requestHandle uint32, returnDiagnostics uint32, auditEntryId PascalString, timeoutHint uint32, additionalHeader ExtensionObject) RequestHeaderBuilder {
	return b.WithAuthenticationToken(authenticationToken).WithTimestamp(timestamp).WithRequestHandle(requestHandle).WithReturnDiagnostics(returnDiagnostics).WithAuditEntryId(auditEntryId).WithTimeoutHint(timeoutHint).WithAdditionalHeader(additionalHeader)
}

func (b *_RequestHeaderBuilder) WithAuthenticationToken(authenticationToken NodeId) RequestHeaderBuilder {
	b.AuthenticationToken = authenticationToken
	return b
}

func (b *_RequestHeaderBuilder) WithAuthenticationTokenBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) RequestHeaderBuilder {
	builder := builderSupplier(b.AuthenticationToken.CreateNodeIdBuilder())
	var err error
	b.AuthenticationToken, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_RequestHeaderBuilder) WithTimestamp(timestamp int64) RequestHeaderBuilder {
	b.Timestamp = timestamp
	return b
}

func (b *_RequestHeaderBuilder) WithRequestHandle(requestHandle uint32) RequestHeaderBuilder {
	b.RequestHandle = requestHandle
	return b
}

func (b *_RequestHeaderBuilder) WithReturnDiagnostics(returnDiagnostics uint32) RequestHeaderBuilder {
	b.ReturnDiagnostics = returnDiagnostics
	return b
}

func (b *_RequestHeaderBuilder) WithAuditEntryId(auditEntryId PascalString) RequestHeaderBuilder {
	b.AuditEntryId = auditEntryId
	return b
}

func (b *_RequestHeaderBuilder) WithAuditEntryIdBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) RequestHeaderBuilder {
	builder := builderSupplier(b.AuditEntryId.CreatePascalStringBuilder())
	var err error
	b.AuditEntryId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_RequestHeaderBuilder) WithTimeoutHint(timeoutHint uint32) RequestHeaderBuilder {
	b.TimeoutHint = timeoutHint
	return b
}

func (b *_RequestHeaderBuilder) WithAdditionalHeader(additionalHeader ExtensionObject) RequestHeaderBuilder {
	b.AdditionalHeader = additionalHeader
	return b
}

func (b *_RequestHeaderBuilder) WithAdditionalHeaderBuilder(builderSupplier func(ExtensionObjectBuilder) ExtensionObjectBuilder) RequestHeaderBuilder {
	builder := builderSupplier(b.AdditionalHeader.CreateExtensionObjectBuilder())
	var err error
	b.AdditionalHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectBuilder failed"))
	}
	return b
}

func (b *_RequestHeaderBuilder) Build() (RequestHeader, error) {
	if b.AuthenticationToken == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'authenticationToken' not set"))
	}
	if b.AuditEntryId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'auditEntryId' not set"))
	}
	if b.AdditionalHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'additionalHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._RequestHeader.deepCopy(), nil
}

func (b *_RequestHeaderBuilder) MustBuild() RequestHeader {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_RequestHeaderBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_RequestHeaderBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_RequestHeaderBuilder) DeepCopy() any {
	_copy := b.CreateRequestHeaderBuilder().(*_RequestHeaderBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateRequestHeaderBuilder creates a RequestHeaderBuilder
func (b *_RequestHeader) CreateRequestHeaderBuilder() RequestHeaderBuilder {
	if b == nil {
		return NewRequestHeaderBuilder()
	}
	return &_RequestHeaderBuilder{_RequestHeader: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RequestHeader) GetExtensionId() int32 {
	return int32(391)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RequestHeader) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RequestHeader) GetAuthenticationToken() NodeId {
	return m.AuthenticationToken
}

func (m *_RequestHeader) GetTimestamp() int64 {
	return m.Timestamp
}

func (m *_RequestHeader) GetRequestHandle() uint32 {
	return m.RequestHandle
}

func (m *_RequestHeader) GetReturnDiagnostics() uint32 {
	return m.ReturnDiagnostics
}

func (m *_RequestHeader) GetAuditEntryId() PascalString {
	return m.AuditEntryId
}

func (m *_RequestHeader) GetTimeoutHint() uint32 {
	return m.TimeoutHint
}

func (m *_RequestHeader) GetAdditionalHeader() ExtensionObject {
	return m.AdditionalHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastRequestHeader(structType any) RequestHeader {
	if casted, ok := structType.(RequestHeader); ok {
		return casted
	}
	if casted, ok := structType.(*RequestHeader); ok {
		return *casted
	}
	return nil
}

func (m *_RequestHeader) GetTypeName() string {
	return "RequestHeader"
}

func (m *_RequestHeader) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (authenticationToken)
	lengthInBits += m.AuthenticationToken.GetLengthInBits(ctx)

	// Simple field (timestamp)
	lengthInBits += 64

	// Simple field (requestHandle)
	lengthInBits += 32

	// Simple field (returnDiagnostics)
	lengthInBits += 32

	// Simple field (auditEntryId)
	lengthInBits += m.AuditEntryId.GetLengthInBits(ctx)

	// Simple field (timeoutHint)
	lengthInBits += 32

	// Simple field (additionalHeader)
	lengthInBits += m.AdditionalHeader.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_RequestHeader) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_RequestHeader) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__requestHeader RequestHeader, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RequestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestHeader")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	authenticationToken, err := ReadSimpleField[NodeId](ctx, "authenticationToken", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'authenticationToken' field"))
	}
	m.AuthenticationToken = authenticationToken

	timestamp, err := ReadSimpleField(ctx, "timestamp", ReadSignedLong(readBuffer, uint8(64)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestamp' field"))
	}
	m.Timestamp = timestamp

	requestHandle, err := ReadSimpleField(ctx, "requestHandle", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHandle' field"))
	}
	m.RequestHandle = requestHandle

	returnDiagnostics, err := ReadSimpleField(ctx, "returnDiagnostics", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'returnDiagnostics' field"))
	}
	m.ReturnDiagnostics = returnDiagnostics

	auditEntryId, err := ReadSimpleField[PascalString](ctx, "auditEntryId", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'auditEntryId' field"))
	}
	m.AuditEntryId = auditEntryId

	timeoutHint, err := ReadSimpleField(ctx, "timeoutHint", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeoutHint' field"))
	}
	m.TimeoutHint = timeoutHint

	additionalHeader, err := ReadSimpleField[ExtensionObject](ctx, "additionalHeader", ReadComplex[ExtensionObject](ExtensionObjectParseWithBufferProducer[ExtensionObject]((bool)(bool(true))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'additionalHeader' field"))
	}
	m.AdditionalHeader = additionalHeader

	if closeErr := readBuffer.CloseContext("RequestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestHeader")
	}

	return m, nil
}

func (m *_RequestHeader) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RequestHeader) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RequestHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RequestHeader")
		}

		if err := WriteSimpleField[NodeId](ctx, "authenticationToken", m.GetAuthenticationToken(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'authenticationToken' field")
		}

		if err := WriteSimpleField[int64](ctx, "timestamp", m.GetTimestamp(), WriteSignedLong(writeBuffer, 64)); err != nil {
			return errors.Wrap(err, "Error serializing 'timestamp' field")
		}

		if err := WriteSimpleField[uint32](ctx, "requestHandle", m.GetRequestHandle(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHandle' field")
		}

		if err := WriteSimpleField[uint32](ctx, "returnDiagnostics", m.GetReturnDiagnostics(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'returnDiagnostics' field")
		}

		if err := WriteSimpleField[PascalString](ctx, "auditEntryId", m.GetAuditEntryId(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'auditEntryId' field")
		}

		if err := WriteSimpleField[uint32](ctx, "timeoutHint", m.GetTimeoutHint(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeoutHint' field")
		}

		if err := WriteSimpleField[ExtensionObject](ctx, "additionalHeader", m.GetAdditionalHeader(), WriteComplex[ExtensionObject](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'additionalHeader' field")
		}

		if popErr := writeBuffer.PopContext("RequestHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RequestHeader")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RequestHeader) IsRequestHeader() {}

func (m *_RequestHeader) DeepCopy() any {
	return m.deepCopy()
}

func (m *_RequestHeader) deepCopy() *_RequestHeader {
	if m == nil {
		return nil
	}
	_RequestHeaderCopy := &_RequestHeader{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[NodeId](m.AuthenticationToken),
		m.Timestamp,
		m.RequestHandle,
		m.ReturnDiagnostics,
		utils.DeepCopy[PascalString](m.AuditEntryId),
		m.TimeoutHint,
		utils.DeepCopy[ExtensionObject](m.AdditionalHeader),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _RequestHeaderCopy
}

func (m *_RequestHeader) String() string {
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
