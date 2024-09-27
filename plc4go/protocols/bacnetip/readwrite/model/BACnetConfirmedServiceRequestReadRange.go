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

// BACnetConfirmedServiceRequestReadRange is the corresponding interface of BACnetConfirmedServiceRequestReadRange
type BACnetConfirmedServiceRequestReadRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetPropertyArrayIndex returns PropertyArrayIndex (property field)
	GetPropertyArrayIndex() BACnetContextTagUnsignedInteger
	// GetReadRange returns ReadRange (property field)
	GetReadRange() BACnetConfirmedServiceRequestReadRangeRange
	// IsBACnetConfirmedServiceRequestReadRange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestReadRange()
	// CreateBuilder creates a BACnetConfirmedServiceRequestReadRangeBuilder
	CreateBACnetConfirmedServiceRequestReadRangeBuilder() BACnetConfirmedServiceRequestReadRangeBuilder
}

// _BACnetConfirmedServiceRequestReadRange is the data-structure of this message
type _BACnetConfirmedServiceRequestReadRange struct {
	BACnetConfirmedServiceRequestContract
	ObjectIdentifier   BACnetContextTagObjectIdentifier
	PropertyIdentifier BACnetPropertyIdentifierTagged
	PropertyArrayIndex BACnetContextTagUnsignedInteger
	ReadRange          BACnetConfirmedServiceRequestReadRangeRange
}

var _ BACnetConfirmedServiceRequestReadRange = (*_BACnetConfirmedServiceRequestReadRange)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestReadRange)(nil)

// NewBACnetConfirmedServiceRequestReadRange factory function for _BACnetConfirmedServiceRequestReadRange
func NewBACnetConfirmedServiceRequestReadRange(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, propertyArrayIndex BACnetContextTagUnsignedInteger, readRange BACnetConfirmedServiceRequestReadRangeRange, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestReadRange {
	if objectIdentifier == nil {
		panic("objectIdentifier of type BACnetContextTagObjectIdentifier for BACnetConfirmedServiceRequestReadRange must not be nil")
	}
	if propertyIdentifier == nil {
		panic("propertyIdentifier of type BACnetPropertyIdentifierTagged for BACnetConfirmedServiceRequestReadRange must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestReadRange{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		ObjectIdentifier:                      objectIdentifier,
		PropertyIdentifier:                    propertyIdentifier,
		PropertyArrayIndex:                    propertyArrayIndex,
		ReadRange:                             readRange,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestReadRangeBuilder is a builder for BACnetConfirmedServiceRequestReadRange
type BACnetConfirmedServiceRequestReadRangeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged) BACnetConfirmedServiceRequestReadRangeBuilder
	// WithObjectIdentifier adds ObjectIdentifier (property field)
	WithObjectIdentifier(BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestReadRangeBuilder
	// WithObjectIdentifierBuilder adds ObjectIdentifier (property field) which is build by the builder
	WithObjectIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestReadRangeBuilder
	// WithPropertyIdentifier adds PropertyIdentifier (property field)
	WithPropertyIdentifier(BACnetPropertyIdentifierTagged) BACnetConfirmedServiceRequestReadRangeBuilder
	// WithPropertyIdentifierBuilder adds PropertyIdentifier (property field) which is build by the builder
	WithPropertyIdentifierBuilder(func(BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyIdentifierTaggedBuilder) BACnetConfirmedServiceRequestReadRangeBuilder
	// WithPropertyArrayIndex adds PropertyArrayIndex (property field)
	WithOptionalPropertyArrayIndex(BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestReadRangeBuilder
	// WithOptionalPropertyArrayIndexBuilder adds PropertyArrayIndex (property field) which is build by the builder
	WithOptionalPropertyArrayIndexBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestReadRangeBuilder
	// WithReadRange adds ReadRange (property field)
	WithOptionalReadRange(BACnetConfirmedServiceRequestReadRangeRange) BACnetConfirmedServiceRequestReadRangeBuilder
	// WithOptionalReadRangeBuilder adds ReadRange (property field) which is build by the builder
	WithOptionalReadRangeBuilder(func(BACnetConfirmedServiceRequestReadRangeRangeBuilder) BACnetConfirmedServiceRequestReadRangeRangeBuilder) BACnetConfirmedServiceRequestReadRangeBuilder
	// Build builds the BACnetConfirmedServiceRequestReadRange or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestReadRange, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestReadRange
}

// NewBACnetConfirmedServiceRequestReadRangeBuilder() creates a BACnetConfirmedServiceRequestReadRangeBuilder
func NewBACnetConfirmedServiceRequestReadRangeBuilder() BACnetConfirmedServiceRequestReadRangeBuilder {
	return &_BACnetConfirmedServiceRequestReadRangeBuilder{_BACnetConfirmedServiceRequestReadRange: new(_BACnetConfirmedServiceRequestReadRange)}
}

type _BACnetConfirmedServiceRequestReadRangeBuilder struct {
	*_BACnetConfirmedServiceRequestReadRange

	parentBuilder *_BACnetConfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestReadRangeBuilder) = (*_BACnetConfirmedServiceRequestReadRangeBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestReadRangeBuilder) setParent(contract BACnetConfirmedServiceRequestContract) {
	b.BACnetConfirmedServiceRequestContract = contract
}

func (b *_BACnetConfirmedServiceRequestReadRangeBuilder) WithMandatoryFields(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged) BACnetConfirmedServiceRequestReadRangeBuilder {
	return b.WithObjectIdentifier(objectIdentifier).WithPropertyIdentifier(propertyIdentifier)
}

func (b *_BACnetConfirmedServiceRequestReadRangeBuilder) WithObjectIdentifier(objectIdentifier BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestReadRangeBuilder {
	b.ObjectIdentifier = objectIdentifier
	return b
}

func (b *_BACnetConfirmedServiceRequestReadRangeBuilder) WithObjectIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestReadRangeBuilder {
	builder := builderSupplier(b.ObjectIdentifier.CreateBACnetContextTagObjectIdentifierBuilder())
	var err error
	b.ObjectIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagObjectIdentifierBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestReadRangeBuilder) WithPropertyIdentifier(propertyIdentifier BACnetPropertyIdentifierTagged) BACnetConfirmedServiceRequestReadRangeBuilder {
	b.PropertyIdentifier = propertyIdentifier
	return b
}

func (b *_BACnetConfirmedServiceRequestReadRangeBuilder) WithPropertyIdentifierBuilder(builderSupplier func(BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyIdentifierTaggedBuilder) BACnetConfirmedServiceRequestReadRangeBuilder {
	builder := builderSupplier(b.PropertyIdentifier.CreateBACnetPropertyIdentifierTaggedBuilder())
	var err error
	b.PropertyIdentifier, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetPropertyIdentifierTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestReadRangeBuilder) WithOptionalPropertyArrayIndex(propertyArrayIndex BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestReadRangeBuilder {
	b.PropertyArrayIndex = propertyArrayIndex
	return b
}

func (b *_BACnetConfirmedServiceRequestReadRangeBuilder) WithOptionalPropertyArrayIndexBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestReadRangeBuilder {
	builder := builderSupplier(b.PropertyArrayIndex.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.PropertyArrayIndex, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestReadRangeBuilder) WithOptionalReadRange(readRange BACnetConfirmedServiceRequestReadRangeRange) BACnetConfirmedServiceRequestReadRangeBuilder {
	b.ReadRange = readRange
	return b
}

func (b *_BACnetConfirmedServiceRequestReadRangeBuilder) WithOptionalReadRangeBuilder(builderSupplier func(BACnetConfirmedServiceRequestReadRangeRangeBuilder) BACnetConfirmedServiceRequestReadRangeRangeBuilder) BACnetConfirmedServiceRequestReadRangeBuilder {
	builder := builderSupplier(b.ReadRange.CreateBACnetConfirmedServiceRequestReadRangeRangeBuilder())
	var err error
	b.ReadRange, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetConfirmedServiceRequestReadRangeRangeBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestReadRangeBuilder) Build() (BACnetConfirmedServiceRequestReadRange, error) {
	if b.ObjectIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'objectIdentifier' not set"))
	}
	if b.PropertyIdentifier == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'propertyIdentifier' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestReadRange.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestReadRangeBuilder) MustBuild() BACnetConfirmedServiceRequestReadRange {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConfirmedServiceRequestReadRangeBuilder) Done() BACnetConfirmedServiceRequestBuilder {
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestReadRangeBuilder) buildForBACnetConfirmedServiceRequest() (BACnetConfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestReadRangeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestReadRangeBuilder().(*_BACnetConfirmedServiceRequestReadRangeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestReadRangeBuilder creates a BACnetConfirmedServiceRequestReadRangeBuilder
func (b *_BACnetConfirmedServiceRequestReadRange) CreateBACnetConfirmedServiceRequestReadRangeBuilder() BACnetConfirmedServiceRequestReadRangeBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestReadRangeBuilder()
	}
	return &_BACnetConfirmedServiceRequestReadRangeBuilder{_BACnetConfirmedServiceRequestReadRange: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadRange) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_READ_RANGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestReadRange) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadRange) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestReadRange) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetConfirmedServiceRequestReadRange) GetPropertyArrayIndex() BACnetContextTagUnsignedInteger {
	return m.PropertyArrayIndex
}

func (m *_BACnetConfirmedServiceRequestReadRange) GetReadRange() BACnetConfirmedServiceRequestReadRangeRange {
	return m.ReadRange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReadRange(structType any) BACnetConfirmedServiceRequestReadRange {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReadRange) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadRange"
}

func (m *_BACnetConfirmedServiceRequestReadRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (propertyArrayIndex)
	if m.PropertyArrayIndex != nil {
		lengthInBits += m.PropertyArrayIndex.GetLengthInBits(ctx)
	}

	// Optional Field (readRange)
	if m.ReadRange != nil {
		lengthInBits += m.ReadRange.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReadRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestReadRange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestReadRange BACnetConfirmedServiceRequestReadRange, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReadRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	objectIdentifier, err := ReadSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", ReadComplex[BACnetContextTagObjectIdentifier](BACnetContextTagParseWithBufferProducer[BACnetContextTagObjectIdentifier]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BACNET_OBJECT_IDENTIFIER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'objectIdentifier' field"))
	}
	m.ObjectIdentifier = objectIdentifier

	propertyIdentifier, err := ReadSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", ReadComplex[BACnetPropertyIdentifierTagged](BACnetPropertyIdentifierTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyIdentifier' field"))
	}
	m.PropertyIdentifier = propertyIdentifier

	var propertyArrayIndex BACnetContextTagUnsignedInteger
	_propertyArrayIndex, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "propertyArrayIndex", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyArrayIndex' field"))
	}
	if _propertyArrayIndex != nil {
		propertyArrayIndex = *_propertyArrayIndex
		m.PropertyArrayIndex = propertyArrayIndex
	}

	var readRange BACnetConfirmedServiceRequestReadRangeRange
	_readRange, err := ReadOptionalField[BACnetConfirmedServiceRequestReadRangeRange](ctx, "readRange", ReadComplex[BACnetConfirmedServiceRequestReadRangeRange](BACnetConfirmedServiceRequestReadRangeRangeParseWithBuffer, readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'readRange' field"))
	}
	if _readRange != nil {
		readRange = *_readRange
		m.ReadRange = readRange
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReadRange")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestReadRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestReadRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReadRange")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
		}

		if err := WriteSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", m.GetPropertyIdentifier(), WriteComplex[BACnetPropertyIdentifierTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyIdentifier' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "propertyArrayIndex", GetRef(m.GetPropertyArrayIndex()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyArrayIndex' field")
		}

		if err := WriteOptionalField[BACnetConfirmedServiceRequestReadRangeRange](ctx, "readRange", GetRef(m.GetReadRange()), WriteComplex[BACnetConfirmedServiceRequestReadRangeRange](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'readRange' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReadRange")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestReadRange) IsBACnetConfirmedServiceRequestReadRange() {}

func (m *_BACnetConfirmedServiceRequestReadRange) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestReadRange) deepCopy() *_BACnetConfirmedServiceRequestReadRange {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestReadRangeCopy := &_BACnetConfirmedServiceRequestReadRange{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		m.ObjectIdentifier.DeepCopy().(BACnetContextTagObjectIdentifier),
		m.PropertyIdentifier.DeepCopy().(BACnetPropertyIdentifierTagged),
		m.PropertyArrayIndex.DeepCopy().(BACnetContextTagUnsignedInteger),
		m.ReadRange.DeepCopy().(BACnetConfirmedServiceRequestReadRangeRange),
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestReadRangeCopy
}

func (m *_BACnetConfirmedServiceRequestReadRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
