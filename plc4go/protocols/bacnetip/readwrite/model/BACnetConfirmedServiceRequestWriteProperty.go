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

// BACnetConfirmedServiceRequestWriteProperty is the corresponding interface of BACnetConfirmedServiceRequestWriteProperty
type BACnetConfirmedServiceRequestWriteProperty interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequest
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() BACnetContextTagUnsignedInteger
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() BACnetConstructedData
	// GetPriority returns Priority (property field)
	GetPriority() BACnetContextTagUnsignedInteger
	// IsBACnetConfirmedServiceRequestWriteProperty is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestWriteProperty()
	// CreateBuilder creates a BACnetConfirmedServiceRequestWritePropertyBuilder
	CreateBACnetConfirmedServiceRequestWritePropertyBuilder() BACnetConfirmedServiceRequestWritePropertyBuilder
}

// _BACnetConfirmedServiceRequestWriteProperty is the data-structure of this message
type _BACnetConfirmedServiceRequestWriteProperty struct {
	BACnetConfirmedServiceRequestContract
	ObjectIdentifier   BACnetContextTagObjectIdentifier
	PropertyIdentifier BACnetPropertyIdentifierTagged
	ArrayIndex         BACnetContextTagUnsignedInteger
	PropertyValue      BACnetConstructedData
	Priority           BACnetContextTagUnsignedInteger
}

var _ BACnetConfirmedServiceRequestWriteProperty = (*_BACnetConfirmedServiceRequestWriteProperty)(nil)
var _ BACnetConfirmedServiceRequestRequirements = (*_BACnetConfirmedServiceRequestWriteProperty)(nil)

// NewBACnetConfirmedServiceRequestWriteProperty factory function for _BACnetConfirmedServiceRequestWriteProperty
func NewBACnetConfirmedServiceRequestWriteProperty(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, arrayIndex BACnetContextTagUnsignedInteger, propertyValue BACnetConstructedData, priority BACnetContextTagUnsignedInteger, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestWriteProperty {
	if objectIdentifier == nil {
		panic("objectIdentifier of type BACnetContextTagObjectIdentifier for BACnetConfirmedServiceRequestWriteProperty must not be nil")
	}
	if propertyIdentifier == nil {
		panic("propertyIdentifier of type BACnetPropertyIdentifierTagged for BACnetConfirmedServiceRequestWriteProperty must not be nil")
	}
	if propertyValue == nil {
		panic("propertyValue of type BACnetConstructedData for BACnetConfirmedServiceRequestWriteProperty must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestWriteProperty{
		BACnetConfirmedServiceRequestContract: NewBACnetConfirmedServiceRequest(serviceRequestLength),
		ObjectIdentifier:                      objectIdentifier,
		PropertyIdentifier:                    propertyIdentifier,
		ArrayIndex:                            arrayIndex,
		PropertyValue:                         propertyValue,
		Priority:                              priority,
	}
	_result.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestWritePropertyBuilder is a builder for BACnetConfirmedServiceRequestWriteProperty
type BACnetConfirmedServiceRequestWritePropertyBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, propertyValue BACnetConstructedData) BACnetConfirmedServiceRequestWritePropertyBuilder
	// WithObjectIdentifier adds ObjectIdentifier (property field)
	WithObjectIdentifier(BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestWritePropertyBuilder
	// WithObjectIdentifierBuilder adds ObjectIdentifier (property field) which is build by the builder
	WithObjectIdentifierBuilder(func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestWritePropertyBuilder
	// WithPropertyIdentifier adds PropertyIdentifier (property field)
	WithPropertyIdentifier(BACnetPropertyIdentifierTagged) BACnetConfirmedServiceRequestWritePropertyBuilder
	// WithPropertyIdentifierBuilder adds PropertyIdentifier (property field) which is build by the builder
	WithPropertyIdentifierBuilder(func(BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyIdentifierTaggedBuilder) BACnetConfirmedServiceRequestWritePropertyBuilder
	// WithArrayIndex adds ArrayIndex (property field)
	WithOptionalArrayIndex(BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestWritePropertyBuilder
	// WithOptionalArrayIndexBuilder adds ArrayIndex (property field) which is build by the builder
	WithOptionalArrayIndexBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestWritePropertyBuilder
	// WithPropertyValue adds PropertyValue (property field)
	WithPropertyValue(BACnetConstructedData) BACnetConfirmedServiceRequestWritePropertyBuilder
	// WithPropertyValueBuilder adds PropertyValue (property field) which is build by the builder
	WithPropertyValueBuilder(func(BACnetConstructedDataBuilder) BACnetConstructedDataBuilder) BACnetConfirmedServiceRequestWritePropertyBuilder
	// WithPriority adds Priority (property field)
	WithOptionalPriority(BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestWritePropertyBuilder
	// WithOptionalPriorityBuilder adds Priority (property field) which is build by the builder
	WithOptionalPriorityBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestWritePropertyBuilder
	// Build builds the BACnetConfirmedServiceRequestWriteProperty or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestWriteProperty, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestWriteProperty
}

// NewBACnetConfirmedServiceRequestWritePropertyBuilder() creates a BACnetConfirmedServiceRequestWritePropertyBuilder
func NewBACnetConfirmedServiceRequestWritePropertyBuilder() BACnetConfirmedServiceRequestWritePropertyBuilder {
	return &_BACnetConfirmedServiceRequestWritePropertyBuilder{_BACnetConfirmedServiceRequestWriteProperty: new(_BACnetConfirmedServiceRequestWriteProperty)}
}

type _BACnetConfirmedServiceRequestWritePropertyBuilder struct {
	*_BACnetConfirmedServiceRequestWriteProperty

	parentBuilder *_BACnetConfirmedServiceRequestBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestWritePropertyBuilder) = (*_BACnetConfirmedServiceRequestWritePropertyBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestWritePropertyBuilder) setParent(contract BACnetConfirmedServiceRequestContract) {
	b.BACnetConfirmedServiceRequestContract = contract
}

func (b *_BACnetConfirmedServiceRequestWritePropertyBuilder) WithMandatoryFields(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, propertyValue BACnetConstructedData) BACnetConfirmedServiceRequestWritePropertyBuilder {
	return b.WithObjectIdentifier(objectIdentifier).WithPropertyIdentifier(propertyIdentifier).WithPropertyValue(propertyValue)
}

func (b *_BACnetConfirmedServiceRequestWritePropertyBuilder) WithObjectIdentifier(objectIdentifier BACnetContextTagObjectIdentifier) BACnetConfirmedServiceRequestWritePropertyBuilder {
	b.ObjectIdentifier = objectIdentifier
	return b
}

func (b *_BACnetConfirmedServiceRequestWritePropertyBuilder) WithObjectIdentifierBuilder(builderSupplier func(BACnetContextTagObjectIdentifierBuilder) BACnetContextTagObjectIdentifierBuilder) BACnetConfirmedServiceRequestWritePropertyBuilder {
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

func (b *_BACnetConfirmedServiceRequestWritePropertyBuilder) WithPropertyIdentifier(propertyIdentifier BACnetPropertyIdentifierTagged) BACnetConfirmedServiceRequestWritePropertyBuilder {
	b.PropertyIdentifier = propertyIdentifier
	return b
}

func (b *_BACnetConfirmedServiceRequestWritePropertyBuilder) WithPropertyIdentifierBuilder(builderSupplier func(BACnetPropertyIdentifierTaggedBuilder) BACnetPropertyIdentifierTaggedBuilder) BACnetConfirmedServiceRequestWritePropertyBuilder {
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

func (b *_BACnetConfirmedServiceRequestWritePropertyBuilder) WithOptionalArrayIndex(arrayIndex BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestWritePropertyBuilder {
	b.ArrayIndex = arrayIndex
	return b
}

func (b *_BACnetConfirmedServiceRequestWritePropertyBuilder) WithOptionalArrayIndexBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestWritePropertyBuilder {
	builder := builderSupplier(b.ArrayIndex.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.ArrayIndex, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestWritePropertyBuilder) WithPropertyValue(propertyValue BACnetConstructedData) BACnetConfirmedServiceRequestWritePropertyBuilder {
	b.PropertyValue = propertyValue
	return b
}

func (b *_BACnetConfirmedServiceRequestWritePropertyBuilder) WithPropertyValueBuilder(builderSupplier func(BACnetConstructedDataBuilder) BACnetConstructedDataBuilder) BACnetConfirmedServiceRequestWritePropertyBuilder {
	builder := builderSupplier(b.PropertyValue.CreateBACnetConstructedDataBuilder())
	var err error
	b.PropertyValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetConstructedDataBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestWritePropertyBuilder) WithOptionalPriority(priority BACnetContextTagUnsignedInteger) BACnetConfirmedServiceRequestWritePropertyBuilder {
	b.Priority = priority
	return b
}

func (b *_BACnetConfirmedServiceRequestWritePropertyBuilder) WithOptionalPriorityBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestWritePropertyBuilder {
	builder := builderSupplier(b.Priority.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.Priority, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestWritePropertyBuilder) Build() (BACnetConfirmedServiceRequestWriteProperty, error) {
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
	if b.PropertyValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'propertyValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestWriteProperty.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestWritePropertyBuilder) MustBuild() BACnetConfirmedServiceRequestWriteProperty {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConfirmedServiceRequestWritePropertyBuilder) Done() BACnetConfirmedServiceRequestBuilder {
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestWritePropertyBuilder) buildForBACnetConfirmedServiceRequest() (BACnetConfirmedServiceRequest, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestWritePropertyBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestWritePropertyBuilder().(*_BACnetConfirmedServiceRequestWritePropertyBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestWritePropertyBuilder creates a BACnetConfirmedServiceRequestWritePropertyBuilder
func (b *_BACnetConfirmedServiceRequestWriteProperty) CreateBACnetConfirmedServiceRequestWritePropertyBuilder() BACnetConfirmedServiceRequestWritePropertyBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestWritePropertyBuilder()
	}
	return &_BACnetConfirmedServiceRequestWritePropertyBuilder{_BACnetConfirmedServiceRequestWriteProperty: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_WRITE_PROPERTY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetParent() BACnetConfirmedServiceRequestContract {
	return m.BACnetConfirmedServiceRequestContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetArrayIndex() BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetPropertyValue() BACnetConstructedData {
	return m.PropertyValue
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetPriority() BACnetContextTagUnsignedInteger {
	return m.Priority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestWriteProperty(structType any) BACnetConfirmedServiceRequestWriteProperty {
	if casted, ok := structType.(BACnetConfirmedServiceRequestWriteProperty); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestWriteProperty); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetTypeName() string {
	return "BACnetConfirmedServiceRequestWriteProperty"
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).getLengthInBits(ctx))

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += m.ArrayIndex.GetLengthInBits(ctx)
	}

	// Simple field (propertyValue)
	lengthInBits += m.PropertyValue.GetLengthInBits(ctx)

	// Optional Field (priority)
	if m.Priority != nil {
		lengthInBits += m.Priority.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequest, serviceRequestLength uint32) (__bACnetConfirmedServiceRequestWriteProperty BACnetConfirmedServiceRequestWriteProperty, err error) {
	m.BACnetConfirmedServiceRequestContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestWriteProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestWriteProperty")
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

	var arrayIndex BACnetContextTagUnsignedInteger
	_arrayIndex, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "arrayIndex", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'arrayIndex' field"))
	}
	if _arrayIndex != nil {
		arrayIndex = *_arrayIndex
		m.ArrayIndex = arrayIndex
	}

	propertyValue, err := ReadSimpleField[BACnetConstructedData](ctx, "propertyValue", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(uint8(3)), (BACnetObjectType)(objectIdentifier.GetObjectType()), (BACnetPropertyIdentifier)(propertyIdentifier.GetValue()), (BACnetTagPayloadUnsignedInteger)((CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((arrayIndex) != (nil)), func() any { return CastBACnetTagPayloadUnsignedInteger((arrayIndex).GetPayload()) }, func() any { return CastBACnetTagPayloadUnsignedInteger(nil) }))))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyValue' field"))
	}
	m.PropertyValue = propertyValue

	var priority BACnetContextTagUnsignedInteger
	_priority, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "priority", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(4)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'priority' field"))
	}
	if _priority != nil {
		priority = *_priority
		m.Priority = priority
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestWriteProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestWriteProperty")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestWriteProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestWriteProperty")
		}

		if err := WriteSimpleField[BACnetContextTagObjectIdentifier](ctx, "objectIdentifier", m.GetObjectIdentifier(), WriteComplex[BACnetContextTagObjectIdentifier](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'objectIdentifier' field")
		}

		if err := WriteSimpleField[BACnetPropertyIdentifierTagged](ctx, "propertyIdentifier", m.GetPropertyIdentifier(), WriteComplex[BACnetPropertyIdentifierTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyIdentifier' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "arrayIndex", GetRef(m.GetArrayIndex()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'arrayIndex' field")
		}

		if err := WriteSimpleField[BACnetConstructedData](ctx, "propertyValue", m.GetPropertyValue(), WriteComplex[BACnetConstructedData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyValue' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "priority", GetRef(m.GetPriority()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'priority' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestWriteProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestWriteProperty")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) IsBACnetConfirmedServiceRequestWriteProperty() {
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) deepCopy() *_BACnetConfirmedServiceRequestWriteProperty {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestWritePropertyCopy := &_BACnetConfirmedServiceRequestWriteProperty{
		m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest).deepCopy(),
		utils.DeepCopy[BACnetContextTagObjectIdentifier](m.ObjectIdentifier),
		utils.DeepCopy[BACnetPropertyIdentifierTagged](m.PropertyIdentifier),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.ArrayIndex),
		utils.DeepCopy[BACnetConstructedData](m.PropertyValue),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.Priority),
	}
	m.BACnetConfirmedServiceRequestContract.(*_BACnetConfirmedServiceRequest)._SubType = m
	return _BACnetConfirmedServiceRequestWritePropertyCopy
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) String() string {
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
