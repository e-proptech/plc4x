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

// CreateMonitoredItemsRequest is the corresponding interface of CreateMonitoredItemsRequest
type CreateMonitoredItemsRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() RequestHeader
	// GetSubscriptionId returns SubscriptionId (property field)
	GetSubscriptionId() uint32
	// GetTimestampsToReturn returns TimestampsToReturn (property field)
	GetTimestampsToReturn() TimestampsToReturn
	// GetItemsToCreate returns ItemsToCreate (property field)
	GetItemsToCreate() []MonitoredItemCreateRequest
	// IsCreateMonitoredItemsRequest is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCreateMonitoredItemsRequest()
	// CreateBuilder creates a CreateMonitoredItemsRequestBuilder
	CreateCreateMonitoredItemsRequestBuilder() CreateMonitoredItemsRequestBuilder
}

// _CreateMonitoredItemsRequest is the data-structure of this message
type _CreateMonitoredItemsRequest struct {
	ExtensionObjectDefinitionContract
	RequestHeader      RequestHeader
	SubscriptionId     uint32
	TimestampsToReturn TimestampsToReturn
	ItemsToCreate      []MonitoredItemCreateRequest
}

var _ CreateMonitoredItemsRequest = (*_CreateMonitoredItemsRequest)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_CreateMonitoredItemsRequest)(nil)

// NewCreateMonitoredItemsRequest factory function for _CreateMonitoredItemsRequest
func NewCreateMonitoredItemsRequest(requestHeader RequestHeader, subscriptionId uint32, timestampsToReturn TimestampsToReturn, itemsToCreate []MonitoredItemCreateRequest) *_CreateMonitoredItemsRequest {
	if requestHeader == nil {
		panic("requestHeader of type RequestHeader for CreateMonitoredItemsRequest must not be nil")
	}
	_result := &_CreateMonitoredItemsRequest{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		RequestHeader:                     requestHeader,
		SubscriptionId:                    subscriptionId,
		TimestampsToReturn:                timestampsToReturn,
		ItemsToCreate:                     itemsToCreate,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// CreateMonitoredItemsRequestBuilder is a builder for CreateMonitoredItemsRequest
type CreateMonitoredItemsRequestBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(requestHeader RequestHeader, subscriptionId uint32, timestampsToReturn TimestampsToReturn, itemsToCreate []MonitoredItemCreateRequest) CreateMonitoredItemsRequestBuilder
	// WithRequestHeader adds RequestHeader (property field)
	WithRequestHeader(RequestHeader) CreateMonitoredItemsRequestBuilder
	// WithRequestHeaderBuilder adds RequestHeader (property field) which is build by the builder
	WithRequestHeaderBuilder(func(RequestHeaderBuilder) RequestHeaderBuilder) CreateMonitoredItemsRequestBuilder
	// WithSubscriptionId adds SubscriptionId (property field)
	WithSubscriptionId(uint32) CreateMonitoredItemsRequestBuilder
	// WithTimestampsToReturn adds TimestampsToReturn (property field)
	WithTimestampsToReturn(TimestampsToReturn) CreateMonitoredItemsRequestBuilder
	// WithItemsToCreate adds ItemsToCreate (property field)
	WithItemsToCreate(...MonitoredItemCreateRequest) CreateMonitoredItemsRequestBuilder
	// Build builds the CreateMonitoredItemsRequest or returns an error if something is wrong
	Build() (CreateMonitoredItemsRequest, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() CreateMonitoredItemsRequest
}

// NewCreateMonitoredItemsRequestBuilder() creates a CreateMonitoredItemsRequestBuilder
func NewCreateMonitoredItemsRequestBuilder() CreateMonitoredItemsRequestBuilder {
	return &_CreateMonitoredItemsRequestBuilder{_CreateMonitoredItemsRequest: new(_CreateMonitoredItemsRequest)}
}

type _CreateMonitoredItemsRequestBuilder struct {
	*_CreateMonitoredItemsRequest

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (CreateMonitoredItemsRequestBuilder) = (*_CreateMonitoredItemsRequestBuilder)(nil)

func (b *_CreateMonitoredItemsRequestBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_CreateMonitoredItemsRequestBuilder) WithMandatoryFields(requestHeader RequestHeader, subscriptionId uint32, timestampsToReturn TimestampsToReturn, itemsToCreate []MonitoredItemCreateRequest) CreateMonitoredItemsRequestBuilder {
	return b.WithRequestHeader(requestHeader).WithSubscriptionId(subscriptionId).WithTimestampsToReturn(timestampsToReturn).WithItemsToCreate(itemsToCreate...)
}

func (b *_CreateMonitoredItemsRequestBuilder) WithRequestHeader(requestHeader RequestHeader) CreateMonitoredItemsRequestBuilder {
	b.RequestHeader = requestHeader
	return b
}

func (b *_CreateMonitoredItemsRequestBuilder) WithRequestHeaderBuilder(builderSupplier func(RequestHeaderBuilder) RequestHeaderBuilder) CreateMonitoredItemsRequestBuilder {
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

func (b *_CreateMonitoredItemsRequestBuilder) WithSubscriptionId(subscriptionId uint32) CreateMonitoredItemsRequestBuilder {
	b.SubscriptionId = subscriptionId
	return b
}

func (b *_CreateMonitoredItemsRequestBuilder) WithTimestampsToReturn(timestampsToReturn TimestampsToReturn) CreateMonitoredItemsRequestBuilder {
	b.TimestampsToReturn = timestampsToReturn
	return b
}

func (b *_CreateMonitoredItemsRequestBuilder) WithItemsToCreate(itemsToCreate ...MonitoredItemCreateRequest) CreateMonitoredItemsRequestBuilder {
	b.ItemsToCreate = itemsToCreate
	return b
}

func (b *_CreateMonitoredItemsRequestBuilder) Build() (CreateMonitoredItemsRequest, error) {
	if b.RequestHeader == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'requestHeader' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._CreateMonitoredItemsRequest.deepCopy(), nil
}

func (b *_CreateMonitoredItemsRequestBuilder) MustBuild() CreateMonitoredItemsRequest {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_CreateMonitoredItemsRequestBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_CreateMonitoredItemsRequestBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_CreateMonitoredItemsRequestBuilder) DeepCopy() any {
	_copy := b.CreateCreateMonitoredItemsRequestBuilder().(*_CreateMonitoredItemsRequestBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateCreateMonitoredItemsRequestBuilder creates a CreateMonitoredItemsRequestBuilder
func (b *_CreateMonitoredItemsRequest) CreateCreateMonitoredItemsRequestBuilder() CreateMonitoredItemsRequestBuilder {
	if b == nil {
		return NewCreateMonitoredItemsRequestBuilder()
	}
	return &_CreateMonitoredItemsRequestBuilder{_CreateMonitoredItemsRequest: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CreateMonitoredItemsRequest) GetExtensionId() int32 {
	return int32(751)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CreateMonitoredItemsRequest) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CreateMonitoredItemsRequest) GetRequestHeader() RequestHeader {
	return m.RequestHeader
}

func (m *_CreateMonitoredItemsRequest) GetSubscriptionId() uint32 {
	return m.SubscriptionId
}

func (m *_CreateMonitoredItemsRequest) GetTimestampsToReturn() TimestampsToReturn {
	return m.TimestampsToReturn
}

func (m *_CreateMonitoredItemsRequest) GetItemsToCreate() []MonitoredItemCreateRequest {
	return m.ItemsToCreate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastCreateMonitoredItemsRequest(structType any) CreateMonitoredItemsRequest {
	if casted, ok := structType.(CreateMonitoredItemsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CreateMonitoredItemsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CreateMonitoredItemsRequest) GetTypeName() string {
	return "CreateMonitoredItemsRequest"
}

func (m *_CreateMonitoredItemsRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (subscriptionId)
	lengthInBits += 32

	// Simple field (timestampsToReturn)
	lengthInBits += 32

	// Implicit Field (noOfItemsToCreate)
	lengthInBits += 32

	// Array field
	if len(m.ItemsToCreate) > 0 {
		for _curItem, element := range m.ItemsToCreate {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ItemsToCreate), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_CreateMonitoredItemsRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CreateMonitoredItemsRequest) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, extensionId int32) (__createMonitoredItemsRequest CreateMonitoredItemsRequest, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CreateMonitoredItemsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CreateMonitoredItemsRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	requestHeader, err := ReadSimpleField[RequestHeader](ctx, "requestHeader", ReadComplex[RequestHeader](ExtensionObjectDefinitionParseWithBufferProducer[RequestHeader]((int32)(int32(391))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'requestHeader' field"))
	}
	m.RequestHeader = requestHeader

	subscriptionId, err := ReadSimpleField(ctx, "subscriptionId", ReadUnsignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriptionId' field"))
	}
	m.SubscriptionId = subscriptionId

	timestampsToReturn, err := ReadEnumField[TimestampsToReturn](ctx, "timestampsToReturn", "TimestampsToReturn", ReadEnum(TimestampsToReturnByValue, ReadUnsignedInt(readBuffer, uint8(32))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timestampsToReturn' field"))
	}
	m.TimestampsToReturn = timestampsToReturn

	noOfItemsToCreate, err := ReadImplicitField[int32](ctx, "noOfItemsToCreate", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfItemsToCreate' field"))
	}
	_ = noOfItemsToCreate

	itemsToCreate, err := ReadCountArrayField[MonitoredItemCreateRequest](ctx, "itemsToCreate", ReadComplex[MonitoredItemCreateRequest](ExtensionObjectDefinitionParseWithBufferProducer[MonitoredItemCreateRequest]((int32)(int32(745))), readBuffer), uint64(noOfItemsToCreate))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'itemsToCreate' field"))
	}
	m.ItemsToCreate = itemsToCreate

	if closeErr := readBuffer.CloseContext("CreateMonitoredItemsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CreateMonitoredItemsRequest")
	}

	return m, nil
}

func (m *_CreateMonitoredItemsRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CreateMonitoredItemsRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CreateMonitoredItemsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CreateMonitoredItemsRequest")
		}

		if err := WriteSimpleField[RequestHeader](ctx, "requestHeader", m.GetRequestHeader(), WriteComplex[RequestHeader](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'requestHeader' field")
		}

		if err := WriteSimpleField[uint32](ctx, "subscriptionId", m.GetSubscriptionId(), WriteUnsignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriptionId' field")
		}

		if err := WriteSimpleEnumField[TimestampsToReturn](ctx, "timestampsToReturn", "TimestampsToReturn", m.GetTimestampsToReturn(), WriteEnum[TimestampsToReturn, uint32](TimestampsToReturn.GetValue, TimestampsToReturn.PLC4XEnumName, WriteUnsignedInt(writeBuffer, 32))); err != nil {
			return errors.Wrap(err, "Error serializing 'timestampsToReturn' field")
		}
		noOfItemsToCreate := int32(utils.InlineIf(bool((m.GetItemsToCreate()) == (nil)), func() any { return int32(-(int32(1))) }, func() any { return int32(int32(len(m.GetItemsToCreate()))) }).(int32))
		if err := WriteImplicitField(ctx, "noOfItemsToCreate", noOfItemsToCreate, WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfItemsToCreate' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "itemsToCreate", m.GetItemsToCreate(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'itemsToCreate' field")
		}

		if popErr := writeBuffer.PopContext("CreateMonitoredItemsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CreateMonitoredItemsRequest")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CreateMonitoredItemsRequest) IsCreateMonitoredItemsRequest() {}

func (m *_CreateMonitoredItemsRequest) DeepCopy() any {
	return m.deepCopy()
}

func (m *_CreateMonitoredItemsRequest) deepCopy() *_CreateMonitoredItemsRequest {
	if m == nil {
		return nil
	}
	_CreateMonitoredItemsRequestCopy := &_CreateMonitoredItemsRequest{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		utils.DeepCopy[RequestHeader](m.RequestHeader),
		m.SubscriptionId,
		m.TimestampsToReturn,
		utils.DeepCopySlice[MonitoredItemCreateRequest, MonitoredItemCreateRequest](m.ItemsToCreate),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _CreateMonitoredItemsRequestCopy
}

func (m *_CreateMonitoredItemsRequest) String() string {
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
