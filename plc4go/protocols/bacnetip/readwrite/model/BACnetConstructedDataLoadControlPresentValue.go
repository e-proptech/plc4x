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

// BACnetConstructedDataLoadControlPresentValue is the corresponding interface of BACnetConstructedDataLoadControlPresentValue
type BACnetConstructedDataLoadControlPresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetShedStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetShedStateTagged
	// IsBACnetConstructedDataLoadControlPresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLoadControlPresentValue()
	// CreateBuilder creates a BACnetConstructedDataLoadControlPresentValueBuilder
	CreateBACnetConstructedDataLoadControlPresentValueBuilder() BACnetConstructedDataLoadControlPresentValueBuilder
}

// _BACnetConstructedDataLoadControlPresentValue is the data-structure of this message
type _BACnetConstructedDataLoadControlPresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetShedStateTagged
}

var _ BACnetConstructedDataLoadControlPresentValue = (*_BACnetConstructedDataLoadControlPresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLoadControlPresentValue)(nil)

// NewBACnetConstructedDataLoadControlPresentValue factory function for _BACnetConstructedDataLoadControlPresentValue
func NewBACnetConstructedDataLoadControlPresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue BACnetShedStateTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLoadControlPresentValue {
	if presentValue == nil {
		panic("presentValue of type BACnetShedStateTagged for BACnetConstructedDataLoadControlPresentValue must not be nil")
	}
	_result := &_BACnetConstructedDataLoadControlPresentValue{
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

// BACnetConstructedDataLoadControlPresentValueBuilder is a builder for BACnetConstructedDataLoadControlPresentValue
type BACnetConstructedDataLoadControlPresentValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(presentValue BACnetShedStateTagged) BACnetConstructedDataLoadControlPresentValueBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(BACnetShedStateTagged) BACnetConstructedDataLoadControlPresentValueBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(BACnetShedStateTaggedBuilder) BACnetShedStateTaggedBuilder) BACnetConstructedDataLoadControlPresentValueBuilder
	// Build builds the BACnetConstructedDataLoadControlPresentValue or returns an error if something is wrong
	Build() (BACnetConstructedDataLoadControlPresentValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLoadControlPresentValue
}

// NewBACnetConstructedDataLoadControlPresentValueBuilder() creates a BACnetConstructedDataLoadControlPresentValueBuilder
func NewBACnetConstructedDataLoadControlPresentValueBuilder() BACnetConstructedDataLoadControlPresentValueBuilder {
	return &_BACnetConstructedDataLoadControlPresentValueBuilder{_BACnetConstructedDataLoadControlPresentValue: new(_BACnetConstructedDataLoadControlPresentValue)}
}

type _BACnetConstructedDataLoadControlPresentValueBuilder struct {
	*_BACnetConstructedDataLoadControlPresentValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLoadControlPresentValueBuilder) = (*_BACnetConstructedDataLoadControlPresentValueBuilder)(nil)

func (b *_BACnetConstructedDataLoadControlPresentValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataLoadControlPresentValueBuilder) WithMandatoryFields(presentValue BACnetShedStateTagged) BACnetConstructedDataLoadControlPresentValueBuilder {
	return b.WithPresentValue(presentValue)
}

func (b *_BACnetConstructedDataLoadControlPresentValueBuilder) WithPresentValue(presentValue BACnetShedStateTagged) BACnetConstructedDataLoadControlPresentValueBuilder {
	b.PresentValue = presentValue
	return b
}

func (b *_BACnetConstructedDataLoadControlPresentValueBuilder) WithPresentValueBuilder(builderSupplier func(BACnetShedStateTaggedBuilder) BACnetShedStateTaggedBuilder) BACnetConstructedDataLoadControlPresentValueBuilder {
	builder := builderSupplier(b.PresentValue.CreateBACnetShedStateTaggedBuilder())
	var err error
	b.PresentValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetShedStateTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLoadControlPresentValueBuilder) Build() (BACnetConstructedDataLoadControlPresentValue, error) {
	if b.PresentValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLoadControlPresentValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataLoadControlPresentValueBuilder) MustBuild() BACnetConstructedDataLoadControlPresentValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataLoadControlPresentValueBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLoadControlPresentValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLoadControlPresentValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLoadControlPresentValueBuilder().(*_BACnetConstructedDataLoadControlPresentValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLoadControlPresentValueBuilder creates a BACnetConstructedDataLoadControlPresentValueBuilder
func (b *_BACnetConstructedDataLoadControlPresentValue) CreateBACnetConstructedDataLoadControlPresentValueBuilder() BACnetConstructedDataLoadControlPresentValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataLoadControlPresentValueBuilder()
	}
	return &_BACnetConstructedDataLoadControlPresentValueBuilder{_BACnetConstructedDataLoadControlPresentValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLoadControlPresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LOAD_CONTROL
}

func (m *_BACnetConstructedDataLoadControlPresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLoadControlPresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLoadControlPresentValue) GetPresentValue() BACnetShedStateTagged {
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

func (m *_BACnetConstructedDataLoadControlPresentValue) GetActualValue() BACnetShedStateTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetShedStateTagged(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLoadControlPresentValue(structType any) BACnetConstructedDataLoadControlPresentValue {
	if casted, ok := structType.(BACnetConstructedDataLoadControlPresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLoadControlPresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLoadControlPresentValue) GetTypeName() string {
	return "BACnetConstructedDataLoadControlPresentValue"
}

func (m *_BACnetConstructedDataLoadControlPresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLoadControlPresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLoadControlPresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLoadControlPresentValue BACnetConstructedDataLoadControlPresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLoadControlPresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLoadControlPresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetShedStateTagged](ctx, "presentValue", ReadComplex[BACnetShedStateTagged](BACnetShedStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetShedStateTagged](ctx, "actualValue", (*BACnetShedStateTagged)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLoadControlPresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLoadControlPresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLoadControlPresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLoadControlPresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLoadControlPresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLoadControlPresentValue")
		}

		if err := WriteSimpleField[BACnetShedStateTagged](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetShedStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLoadControlPresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLoadControlPresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLoadControlPresentValue) IsBACnetConstructedDataLoadControlPresentValue() {
}

func (m *_BACnetConstructedDataLoadControlPresentValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLoadControlPresentValue) deepCopy() *_BACnetConstructedDataLoadControlPresentValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLoadControlPresentValueCopy := &_BACnetConstructedDataLoadControlPresentValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetShedStateTagged](m.PresentValue),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLoadControlPresentValueCopy
}

func (m *_BACnetConstructedDataLoadControlPresentValue) String() string {
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
