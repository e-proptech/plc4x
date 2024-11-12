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

// BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple is the corresponding interface of BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
type BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetSubscriberProcessIdentifier returns SubscriberProcessIdentifier (property field)
	GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetIssueConfirmedNotifications returns IssueConfirmedNotifications (property field)
	GetIssueConfirmedNotifications() BACnetContextTagBoolean
	// GetLifetime returns Lifetime (property field)
	GetLifetime() BACnetContextTagUnsignedInteger
	// GetMaxNotificationDelay returns MaxNotificationDelay (property field)
	GetMaxNotificationDelay() BACnetContextTagUnsignedInteger
	// GetListOfCovSubscriptionSpecifications returns ListOfCovSubscriptionSpecifications (property field)
	GetListOfCovSubscriptionSpecifications() BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList
	// IsBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple()
	// CreateBuilder creates a BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder
	CreateBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder() BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder
}

// _BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple is the data-structure of this message
type _BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple struct {
	BACnetConfirmedServiceRequestContract
	SubscriberProcessIdentifier         BACnetContextTagUnsignedInteger
	IssueConfirmedNotifications         BACnetContextTagBoolean
	Lifetime                            BACnetContextTagUnsignedInteger
	MaxNotificationDelay                BACnetContextTagUnsignedInteger
	ListOfCovSubscriptionSpecifications BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList
}

var _ BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple = (*_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple)(nil)

// NewBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple factory function for _BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
func NewBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, issueConfirmedNotifications BACnetContextTagBoolean, lifetime BACnetContextTagUnsignedInteger, maxNotificationDelay BACnetContextTagUnsignedInteger, listOfCovSubscriptionSpecifications BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple {
	if subscriberProcessIdentifier == nil {
		panic("subscriberProcessIdentifier of type BACnetContextTagUnsignedInteger for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple must not be nil")
	}
	if listOfCovSubscriptionSpecifications == nil {
		panic("listOfCovSubscriptionSpecifications of type BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		SubscriberProcessIdentifier:           subscriberProcessIdentifier,
		IssueConfirmedNotifications:           issueConfirmedNotifications,
		Lifetime:                              lifetime,
		MaxNotificationDelay:                  maxNotificationDelay,
		ListOfCovSubscriptionSpecifications:   listOfCovSubscriptionSpecifications,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder is a builder for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
type BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, listOfCovSubscriptionSpecifications BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder
	// WithSubscriberProcessIdentifier adds SubscriberProcessIdentifier (property field)
	WithSubscriberProcessIdentifier(BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder
	// WithSubscriberProcessIdentifierBuilder adds SubscriberProcessIdentifier (property field) which is build by the builder
	WithSubscriberProcessIdentifierBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder
	// WithIssueConfirmedNotifications adds IssueConfirmedNotifications (property field)
	WithOptionalIssueConfirmedNotifications(BACnetContextTagBoolean) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder
	// WithOptionalIssueConfirmedNotificationsBuilder adds IssueConfirmedNotifications (property field) which is build by the builder
	WithOptionalIssueConfirmedNotificationsBuilder(func(BACnetContextTagBooleanBuilder) BACnetContextTagBooleanBuilder) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder
	// WithLifetime adds Lifetime (property field)
	WithOptionalLifetime(BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder
	// WithOptionalLifetimeBuilder adds Lifetime (property field) which is build by the builder
	WithOptionalLifetimeBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder
	// WithMaxNotificationDelay adds MaxNotificationDelay (property field)
	WithOptionalMaxNotificationDelay(BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder
	// WithOptionalMaxNotificationDelayBuilder adds MaxNotificationDelay (property field) which is build by the builder
	WithOptionalMaxNotificationDelayBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder
	// WithListOfCovSubscriptionSpecifications adds ListOfCovSubscriptionSpecifications (property field)
	WithListOfCovSubscriptionSpecifications(BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder
	// WithListOfCovSubscriptionSpecificationsBuilder adds ListOfCovSubscriptionSpecifications (property field) which is build by the builder
	WithListOfCovSubscriptionSpecificationsBuilder(func(BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsListBuilder) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsListBuilder) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder
	// Build builds the BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple
}

// NewBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder() creates a BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder
func NewBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder() BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder {
	return &_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder{_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple: new(_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple)}
}

type _BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder struct {
	*_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple

	parentBuilder *_BACnetConfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder) = (*_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder) setParent(contract BACnetConfirmedServiceRequestContract) {
	b.BACnetConfirmedServiceRequestContract = contract
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder) WithMandatoryFields(subscriberProcessIdentifier BACnetContextTagUnsignedInteger, listOfCovSubscriptionSpecifications BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder {
	return b.WithSubscriberProcessIdentifier(subscriberProcessIdentifier).WithListOfCovSubscriptionSpecifications(listOfCovSubscriptionSpecifications)
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder) WithSubscriberProcessIdentifier(subscriberProcessIdentifier BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder {
	b.SubscriberProcessIdentifier = subscriberProcessIdentifier
	return b
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder) WithSubscriberProcessIdentifierBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder {
	builder := builderSupplier(b.SubscriberProcessIdentifier.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.SubscriberProcessIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder) WithOptionalIssueConfirmedNotifications(issueConfirmedNotifications BACnetContextTagBoolean) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder {
	b.IssueConfirmedNotifications = issueConfirmedNotifications
	return b
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder) WithOptionalIssueConfirmedNotificationsBuilder(builderSupplier func(BACnetContextTagBooleanBuilder) BACnetContextTagBooleanBuilder) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder {
	builder := builderSupplier(b.IssueConfirmedNotifications.CreateBACnetContextTagBooleanBuilder())
	var err error
	b.IssueConfirmedNotifications, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder) WithOptionalLifetime(lifetime BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder {
	b.Lifetime = lifetime
	return b
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder) WithOptionalLifetimeBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder {
	builder := builderSupplier(b.Lifetime.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.Lifetime, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder) WithOptionalMaxNotificationDelay(maxNotificationDelay BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder {
	b.MaxNotificationDelay = maxNotificationDelay
	return b
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder) WithOptionalMaxNotificationDelayBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder {
	builder := builderSupplier(b.MaxNotificationDelay.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.MaxNotificationDelay, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder) WithListOfCovSubscriptionSpecifications(listOfCovSubscriptionSpecifications BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder {
	b.ListOfCovSubscriptionSpecifications = listOfCovSubscriptionSpecifications
	return b
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder) WithListOfCovSubscriptionSpecificationsBuilder(builderSupplier func(BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsListBuilder) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsListBuilder) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder {
	builder := builderSupplier(b.ListOfCovSubscriptionSpecifications.CreateBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsListBuilder())
	var err error
	b.ListOfCovSubscriptionSpecifications, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsListBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder) Build() (BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple, error) {
	if b.SubscriberProcessIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'subscriberProcessIdentifier' not set"))
	}
	if b.ListOfCovSubscriptionSpecifications == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'listOfCovSubscriptionSpecifications' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder) MustBuild() BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder) Done() BACnetConfirmedServiceRequestBuilder {
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder) buildForBACnetConfirmedServiceRequest() (BACnetConfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder().(*_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder creates a BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder
func (b *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) CreateBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder() BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder()
	}
	return &_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleBuilder{_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_SUBSCRIBE_COV_PROPERTY_MULTIPLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetSubscriberProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.SubscriberProcessIdentifier
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetIssueConfirmedNotifications() BACnetContextTagBoolean {
	return m.IssueConfirmedNotifications
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetLifetime() BACnetContextTagUnsignedInteger {
	return m.Lifetime
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetMaxNotificationDelay() BACnetContextTagUnsignedInteger {
	return m.MaxNotificationDelay
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetListOfCovSubscriptionSpecifications() BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList {
	return m.ListOfCovSubscriptionSpecifications
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple(structType any) BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple {
	if casted, ok := structType.(BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetTypeName() string {
	return "BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple"
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (subscriberProcessIdentifier)
	lengthInBits += m.SubscriberProcessIdentifier.GetLengthInBits(ctx)

	// Optional Field (issueConfirmedNotifications)
	if m.IssueConfirmedNotifications != nil {
		lengthInBits += m.IssueConfirmedNotifications.GetLengthInBits(ctx)
	}

	// Optional Field (lifetime)
	if m.Lifetime != nil {
		lengthInBits += m.Lifetime.GetLengthInBits(ctx)
	}

	// Optional Field (maxNotificationDelay)
	if m.MaxNotificationDelay != nil {
		lengthInBits += m.MaxNotificationDelay.GetLengthInBits(ctx)
	}

	// Simple field (listOfCovSubscriptionSpecifications)
	lengthInBits += m.ListOfCovSubscriptionSpecifications.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	subscriberProcessIdentifier, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "subscriberProcessIdentifier", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'subscriberProcessIdentifier' field"))
	}
	m.SubscriberProcessIdentifier = subscriberProcessIdentifier

	var issueConfirmedNotifications BACnetContextTagBoolean
	_issueConfirmedNotifications, err := ReadOptionalField[BACnetContextTagBoolean](ctx, "issueConfirmedNotifications", ReadComplex[BACnetContextTagBoolean](BACnetContextTagParseWithBufferProducer[BACnetContextTagBoolean]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_BOOLEAN)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'issueConfirmedNotifications' field"))
	}
	if _issueConfirmedNotifications != nil {
		issueConfirmedNotifications = *_issueConfirmedNotifications
		m.IssueConfirmedNotifications = issueConfirmedNotifications
	}

	var lifetime BACnetContextTagUnsignedInteger
	_lifetime, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "lifetime", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lifetime' field"))
	}
	if _lifetime != nil {
		lifetime = *_lifetime
		m.Lifetime = lifetime
	}

	var maxNotificationDelay BACnetContextTagUnsignedInteger
	_maxNotificationDelay, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "maxNotificationDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxNotificationDelay' field"))
	}
	if _maxNotificationDelay != nil {
		maxNotificationDelay = *_maxNotificationDelay
		m.MaxNotificationDelay = maxNotificationDelay
	}

	listOfCovSubscriptionSpecifications, err := ReadSimpleField[BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList](ctx, "listOfCovSubscriptionSpecifications", ReadComplex[BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList](BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsListParseWithBufferProducer((uint8)(uint8(4))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'listOfCovSubscriptionSpecifications' field"))
	}
	m.ListOfCovSubscriptionSpecifications = listOfCovSubscriptionSpecifications

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "subscriberProcessIdentifier", m.GetSubscriberProcessIdentifier(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'subscriberProcessIdentifier' field")
		}

		if err := WriteOptionalField[BACnetContextTagBoolean](ctx, "issueConfirmedNotifications", GetRef(m.GetIssueConfirmedNotifications()), WriteComplex[BACnetContextTagBoolean](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'issueConfirmedNotifications' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "lifetime", GetRef(m.GetLifetime()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'lifetime' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "maxNotificationDelay", GetRef(m.GetMaxNotificationDelay()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'maxNotificationDelay' field")
		}

		if err := WriteSimpleField[BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList](ctx, "listOfCovSubscriptionSpecifications", m.GetListOfCovSubscriptionSpecifications(), WriteComplex[BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'listOfCovSubscriptionSpecifications' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) IsBACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple() {
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) deepCopy() *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleCopy := &_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.SubscriberProcessIdentifier),
		utils.DeepCopy[BACnetContextTagBoolean](m.IssueConfirmedNotifications),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.Lifetime),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.MaxNotificationDelay),
		utils.DeepCopy[BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsList](m.ListOfCovSubscriptionSpecifications),
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleCopy
}

func (m *_BACnetConfirmedServiceRequestSubscribeCOVPropertyMultiple) String() string {
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
