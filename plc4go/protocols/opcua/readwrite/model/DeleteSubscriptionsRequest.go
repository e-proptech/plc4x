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

// DeleteSubscriptionsRequest is the corresponding interface of DeleteSubscriptionsRequest
type DeleteSubscriptionsRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfSubscriptionIds returns NoOfSubscriptionIds (property field)
	GetNoOfSubscriptionIds() int32
	// GetSubscriptionIds returns SubscriptionIds (property field)
	GetSubscriptionIds() []uint32
	// IsDeleteSubscriptionsRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDeleteSubscriptionsRequest()
	// CreateBuilder creates a DeleteSubscriptionsRequestBuilder
	CreateDeleteSubscriptionsRequestBuilder() DeleteSubscriptionsRequestBuilder
}

// _DeleteSubscriptionsRequest is the data-structure of this message
type _DeleteSubscriptionsRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader       ExtensionObjectDefinition
	NoOfSubscriptionIds int32
	SubscriptionIds     []uint32
}

var _ DeleteSubscriptionsRequest = (*_DeleteSubscriptionsRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DeleteSubscriptionsRequest)(nil)

// NewDeleteSubscriptionsRequest factory function for _DeleteSubscriptionsRequest
func NewDeleteSubscriptionsRequest(requestHeader ExtensionObjectDefinition, noOfSubscriptionIds int32, subscriptionIds []uint32) *_DeleteSubscriptionsRequest {
	if requestHeader == nil {
		panic("requestHeader of type ExtensionObjectDefinition for DeleteSubscriptionsRequest must not be nil")
	}
	_result := &_DeleteSubscriptionsRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		NoOfSubscriptionIds:               noOfSubscriptionIds,
		SubscriptionIds:                   subscriptionIds,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DeleteSubscriptionsRequestBuilder is a builder for DeleteSubscriptionsRequest
type DeleteSubscriptionsRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader ExtensionObjectDefinition, noOfSubscriptionIds int32, subscriptionIds []uint32) DeleteSubscriptionsRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(ExtensionObjectDefinition) DeleteSubscriptionsRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) DeleteSubscriptionsRequestBuilder
	// WithNoOfSubscriptionIds adds NoOfSubscriptionIds (property field)
	WithNoOfSubscriptionIds(int32) DeleteSubscriptionsRequestBuilder
	// WithSubscriptionIds adds SubscriptionIds (property field)
	WithSubscriptionIds(...uint32) DeleteSubscriptionsRequestBuilder
	// Build builds the DeleteSubscriptionsRequest or returns an error if something is wrong
	Build() (DeleteSubscriptionsRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DeleteSubscriptionsRequest
}

// NewDeleteSubscriptionsRequestBuilder() creates a DeleteSubscriptionsRequestBuilder
func NewDeleteSubscriptionsRequestBuilder() DeleteSubscriptionsRequestBuilder {
	return &_DeleteSubscriptionsRequestBuilder{_DeleteSubscriptionsRequest: new(_DeleteSubscriptionsRequest)}
}

type _DeleteSubscriptionsRequestBuilder struct {
	*_DeleteSubscriptionsRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (DeleteSubscriptionsRequestBuilder) = (*_DeleteSubscriptionsRequestBuilder)(nil)

func (b *_DeleteSubscriptionsRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_DeleteSubscriptionsRequestBuilder) WithMandatoryFields(requestHeader ExtensionObjectDefinition, noOfSubscriptionIds int32, subscriptionIds []uint32) DeleteSubscriptionsRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithNoOfSubscriptionIds(noOfSubscriptionIds).WithSubscriptionIds(subscriptionIds...)
}

func (b *_DeleteSubscriptionsRequestBuilder) WithRequestHeader(requestHeader ExtensionObjectDefinition) DeleteSubscriptionsRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_DeleteSubscriptionsRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(ExtensionObjectDefinitionBuilder) ExtensionObjectDefinitionBuilder) DeleteSubscriptionsRequestBuilder {
	builder := builderSupplier(b.RequestHeader.CreateExtensionObjectDefinitionBuilder())
	var err error
	b.RequestHeader, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExtensionObjectDefinitionBuilder failed"))
	}
	return b
}

func (b *_DeleteSubscriptionsRequestBuilder) WithNoOfSubscriptionIds(noOfSubscriptionIds int32) DeleteSubscriptionsRequestBuilder {
	b.NoOfSubscriptionIds = noOfSubscriptionIds
	return b
}

func (b *_DeleteSubscriptionsRequestBuilder) WithSubscriptionIds(subscriptionIds ...uint32) DeleteSubscriptionsRequestBuilder {
	b.SubscriptionIds = subscriptionIds
	return b
}

func (b *_DeleteSubscriptionsRequestBuilder) Build() (DeleteSubscriptionsRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DeleteSubscriptionsRequest.deepCopy(), nil
}

func (b *_DeleteSubscriptionsRequestBuilder) MustBuild() DeleteSubscriptionsRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_DeleteSubscriptionsRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_DeleteSubscriptionsRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_DeleteSubscriptionsRequestBuilder) DeepCopy() any {
	_copy := b.CreateDeleteSubscriptionsRequestBuilder().(*_DeleteSubscriptionsRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDeleteSubscriptionsRequestBuilder creates a DeleteSubscriptionsRequestBuilder
func (b *_DeleteSubscriptionsRequest) CreateDeleteSubscriptionsRequestBuilder() DeleteSubscriptionsRequestBuilder {
	if b == nil {
		return NewDeleteSubscriptionsRequestBuilder()
	}
	return &_DeleteSubscriptionsRequestBuilder{_DeleteSubscriptionsRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteSubscriptionsRequest) GetIdentifier() string {
	return "847"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteSubscriptionsRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteSubscriptionsRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_DeleteSubscriptionsRequest) GetNoOfSubscriptionIds() int32 {
	return m.NoOfSubscriptionIds
}

func (m *_DeleteSubscriptionsRequest) GetSubscriptionIds() []uint32 {
	return m.SubscriptionIds
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDeleteSubscriptionsRequest(structType any) DeleteSubscriptionsRequest {
	if casted, ok := structType.(DeleteSubscriptionsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteSubscriptionsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteSubscriptionsRequest) GetTypeName() string {
	return "DeleteSubscriptionsRequest"
}

func (m *_DeleteSubscriptionsRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfSubscriptionIds)
	lengthInBits += 32

	// Array field
	if len(m.SubscriptionIds) > 0 {
		lengthInBits += 32 * uint16(len(m.SubscriptionIds))
	}

	return lengthInBits
}

func (m *_DeleteSubscriptionsRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DeleteSubscriptionsRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__deleteSubscriptionsRequest DeleteSubscriptionsRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeleteSubscriptionsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteSubscriptionsRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("391")), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	noOfSubscriptionIds, err := ReadSimpleField(ctx, "noOfSubscriptionIds", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfSubscriptionIds' field"))
	}
	m.NoOfSubscriptionIds = noOfSubscriptionIds

	subscriptionIds, err := ReadCountArrayField[uint32](ctx, "subscriptionIds", ReadUnsignedInt(readBuffer, uint8(32)), uint64(noOfSubscriptionIds))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriptionIds' field"))
	}
	m.SubscriptionIds = subscriptionIds

	if closeErr := readBuffer.CloseContext("DeleteSubscriptionsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteSubscriptionsRequest")
	}

	return m, nil
}

func (m *_DeleteSubscriptionsRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteSubscriptionsRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteSubscriptionsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteSubscriptionsRequest")
		}

		if err := WriteSimpleField[ExtensionObjectDefinition](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[ExtensionObjectDefinition](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfSubscriptionIds", m.GetNoOfSubscriptionIds(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfSubscriptionIds' field")
		}

		if err := WriteSimpleTypeArrayField(ctx, "subscriptionIds", m.GetSubscriptionIds(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriptionIds' field")
		}

		if popErr := writeBuffer.PopContext("DeleteSubscriptionsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteSubscriptionsRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteSubscriptionsRequest) IsDeleteSubscriptionsRequest() {}

func (m *_DeleteSubscriptionsRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DeleteSubscriptionsRequest) deepCopy() *_DeleteSubscriptionsRequest {
	if m == nil {
		return nil
	}
	_DeleteSubscriptionsRequestCopy := &_DeleteSubscriptionsRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.RequestHeader.DeepCopy().(ExtensionObjectDefinition),
		m.NoOfSubscriptionIds,
		utils.DeepCopySlice[uint32, uint32](m.SubscriptionIds),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DeleteSubscriptionsRequestCopy
}

func (m *_DeleteSubscriptionsRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
