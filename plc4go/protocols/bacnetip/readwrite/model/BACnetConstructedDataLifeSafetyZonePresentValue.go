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

// BACnetConstructedDataLifeSafetyZonePresentValue is the corresponding interface of BACnetConstructedDataLifeSafetyZonePresentValue
type BACnetConstructedDataLifeSafetyZonePresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetLifeSafetyStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLifeSafetyStateTagged
	// IsBACnetConstructedDataLifeSafetyZonePresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLifeSafetyZonePresentValue()
	// CreateBuilder creates a BACnetConstructedDataLifeSafetyZonePresentValueBuilder
	CreateBACnetConstructedDataLifeSafetyZonePresentValueBuilder() BACnetConstructedDataLifeSafetyZonePresentValueBuilder
}

// _BACnetConstructedDataLifeSafetyZonePresentValue is the data-structure of this message
type _BACnetConstructedDataLifeSafetyZonePresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetLifeSafetyStateTagged
}

var _ BACnetConstructedDataLifeSafetyZonePresentValue = (*_BACnetConstructedDataLifeSafetyZonePresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLifeSafetyZonePresentValue)(nil)

// NewBACnetConstructedDataLifeSafetyZonePresentValue factory function for _BACnetConstructedDataLifeSafetyZonePresentValue
func NewBACnetConstructedDataLifeSafetyZonePresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue BACnetLifeSafetyStateTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLifeSafetyZonePresentValue {
	if presentValue == nil {
		panic("presentValue of type BACnetLifeSafetyStateTagged for BACnetConstructedDataLifeSafetyZonePresentValue must not be nil")
	}
	_result := &_BACnetConstructedDataLifeSafetyZonePresentValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PresentValue:                  presentValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLifeSafetyZonePresentValueBuilder is a builder for BACnetConstructedDataLifeSafetyZonePresentValue
type BACnetConstructedDataLifeSafetyZonePresentValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(presentValue BACnetLifeSafetyStateTagged) BACnetConstructedDataLifeSafetyZonePresentValueBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(BACnetLifeSafetyStateTagged) BACnetConstructedDataLifeSafetyZonePresentValueBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(BACnetLifeSafetyStateTaggedBuilder) BACnetLifeSafetyStateTaggedBuilder) BACnetConstructedDataLifeSafetyZonePresentValueBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataLifeSafetyZonePresentValue or returns an error if something is wrong
	Build() (BACnetConstructedDataLifeSafetyZonePresentValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLifeSafetyZonePresentValue
}

// NewBACnetConstructedDataLifeSafetyZonePresentValueBuilder() creates a BACnetConstructedDataLifeSafetyZonePresentValueBuilder
func NewBACnetConstructedDataLifeSafetyZonePresentValueBuilder() BACnetConstructedDataLifeSafetyZonePresentValueBuilder {
	return &_BACnetConstructedDataLifeSafetyZonePresentValueBuilder{_BACnetConstructedDataLifeSafetyZonePresentValue: new(_BACnetConstructedDataLifeSafetyZonePresentValue)}
}

type _BACnetConstructedDataLifeSafetyZonePresentValueBuilder struct {
	*_BACnetConstructedDataLifeSafetyZonePresentValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLifeSafetyZonePresentValueBuilder) = (*_BACnetConstructedDataLifeSafetyZonePresentValueBuilder)(nil)

func (b *_BACnetConstructedDataLifeSafetyZonePresentValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataLifeSafetyZonePresentValue
}

func (b *_BACnetConstructedDataLifeSafetyZonePresentValueBuilder) WithMandatoryFields(presentValue BACnetLifeSafetyStateTagged) BACnetConstructedDataLifeSafetyZonePresentValueBuilder {
	return b.WithPresentValue(presentValue)
}

func (b *_BACnetConstructedDataLifeSafetyZonePresentValueBuilder) WithPresentValue(presentValue BACnetLifeSafetyStateTagged) BACnetConstructedDataLifeSafetyZonePresentValueBuilder {
	b.PresentValue = presentValue
	return b
}

func (b *_BACnetConstructedDataLifeSafetyZonePresentValueBuilder) WithPresentValueBuilder(builderSupplier func(BACnetLifeSafetyStateTaggedBuilder) BACnetLifeSafetyStateTaggedBuilder) BACnetConstructedDataLifeSafetyZonePresentValueBuilder {
	builder := builderSupplier(b.PresentValue.CreateBACnetLifeSafetyStateTaggedBuilder())
	var err error
	b.PresentValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLifeSafetyStateTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLifeSafetyZonePresentValueBuilder) Build() (BACnetConstructedDataLifeSafetyZonePresentValue, error) {
	if b.PresentValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLifeSafetyZonePresentValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataLifeSafetyZonePresentValueBuilder) MustBuild() BACnetConstructedDataLifeSafetyZonePresentValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataLifeSafetyZonePresentValueBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLifeSafetyZonePresentValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLifeSafetyZonePresentValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLifeSafetyZonePresentValueBuilder().(*_BACnetConstructedDataLifeSafetyZonePresentValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLifeSafetyZonePresentValueBuilder creates a BACnetConstructedDataLifeSafetyZonePresentValueBuilder
func (b *_BACnetConstructedDataLifeSafetyZonePresentValue) CreateBACnetConstructedDataLifeSafetyZonePresentValueBuilder() BACnetConstructedDataLifeSafetyZonePresentValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataLifeSafetyZonePresentValueBuilder()
	}
	return &_BACnetConstructedDataLifeSafetyZonePresentValueBuilder{_BACnetConstructedDataLifeSafetyZonePresentValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyZonePresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIFE_SAFETY_ZONE
}

func (m *_BACnetConstructedDataLifeSafetyZonePresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLifeSafetyZonePresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyZonePresentValue) GetPresentValue() BACnetLifeSafetyStateTagged {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyZonePresentValue) GetActualValue() BACnetLifeSafetyStateTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLifeSafetyStateTagged(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLifeSafetyZonePresentValue(structType any) BACnetConstructedDataLifeSafetyZonePresentValue {
	if casted, ok := structType.(BACnetConstructedDataLifeSafetyZonePresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLifeSafetyZonePresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLifeSafetyZonePresentValue) GetTypeName() string {
	return "BACnetConstructedDataLifeSafetyZonePresentValue"
}

func (m *_BACnetConstructedDataLifeSafetyZonePresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLifeSafetyZonePresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLifeSafetyZonePresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLifeSafetyZonePresentValue BACnetConstructedDataLifeSafetyZonePresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLifeSafetyZonePresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLifeSafetyZonePresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetLifeSafetyStateTagged](ctx, "presentValue", ReadComplex[BACnetLifeSafetyStateTagged](BACnetLifeSafetyStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetLifeSafetyStateTagged](ctx, "actualValue", (*BACnetLifeSafetyStateTagged)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLifeSafetyZonePresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLifeSafetyZonePresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLifeSafetyZonePresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLifeSafetyZonePresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLifeSafetyZonePresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLifeSafetyZonePresentValue")
		}

		if err := WriteSimpleField[BACnetLifeSafetyStateTagged](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetLifeSafetyStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLifeSafetyZonePresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLifeSafetyZonePresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLifeSafetyZonePresentValue) IsBACnetConstructedDataLifeSafetyZonePresentValue() {
}

func (m *_BACnetConstructedDataLifeSafetyZonePresentValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLifeSafetyZonePresentValue) deepCopy() *_BACnetConstructedDataLifeSafetyZonePresentValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLifeSafetyZonePresentValueCopy := &_BACnetConstructedDataLifeSafetyZonePresentValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetLifeSafetyStateTagged](m.PresentValue),
	}
	_BACnetConstructedDataLifeSafetyZonePresentValueCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLifeSafetyZonePresentValueCopy
}

func (m *_BACnetConstructedDataLifeSafetyZonePresentValue) String() string {
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
