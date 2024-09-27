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

// BACnetConstructedDataAnalogValuePresentValue is the corresponding interface of BACnetConstructedDataAnalogValuePresentValue
type BACnetConstructedDataAnalogValuePresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataAnalogValuePresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAnalogValuePresentValue()
	// CreateBuilder creates a BACnetConstructedDataAnalogValuePresentValueBuilder
	CreateBACnetConstructedDataAnalogValuePresentValueBuilder() BACnetConstructedDataAnalogValuePresentValueBuilder
}

// _BACnetConstructedDataAnalogValuePresentValue is the data-structure of this message
type _BACnetConstructedDataAnalogValuePresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetApplicationTagReal
}

var _ BACnetConstructedDataAnalogValuePresentValue = (*_BACnetConstructedDataAnalogValuePresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAnalogValuePresentValue)(nil)

// NewBACnetConstructedDataAnalogValuePresentValue factory function for _BACnetConstructedDataAnalogValuePresentValue
func NewBACnetConstructedDataAnalogValuePresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAnalogValuePresentValue {
	if presentValue == nil {
		panic("presentValue of type BACnetApplicationTagReal for BACnetConstructedDataAnalogValuePresentValue must not be nil")
	}
	_result := &_BACnetConstructedDataAnalogValuePresentValue{
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

// BACnetConstructedDataAnalogValuePresentValueBuilder is a builder for BACnetConstructedDataAnalogValuePresentValue
type BACnetConstructedDataAnalogValuePresentValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(presentValue BACnetApplicationTagReal) BACnetConstructedDataAnalogValuePresentValueBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(BACnetApplicationTagReal) BACnetConstructedDataAnalogValuePresentValueBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataAnalogValuePresentValueBuilder
	// Build builds the BACnetConstructedDataAnalogValuePresentValue or returns an error if something is wrong
	Build() (BACnetConstructedDataAnalogValuePresentValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAnalogValuePresentValue
}

// NewBACnetConstructedDataAnalogValuePresentValueBuilder() creates a BACnetConstructedDataAnalogValuePresentValueBuilder
func NewBACnetConstructedDataAnalogValuePresentValueBuilder() BACnetConstructedDataAnalogValuePresentValueBuilder {
	return &_BACnetConstructedDataAnalogValuePresentValueBuilder{_BACnetConstructedDataAnalogValuePresentValue: new(_BACnetConstructedDataAnalogValuePresentValue)}
}

type _BACnetConstructedDataAnalogValuePresentValueBuilder struct {
	*_BACnetConstructedDataAnalogValuePresentValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAnalogValuePresentValueBuilder) = (*_BACnetConstructedDataAnalogValuePresentValueBuilder)(nil)

func (b *_BACnetConstructedDataAnalogValuePresentValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataAnalogValuePresentValueBuilder) WithMandatoryFields(presentValue BACnetApplicationTagReal) BACnetConstructedDataAnalogValuePresentValueBuilder {
	return b.WithPresentValue(presentValue)
}

func (b *_BACnetConstructedDataAnalogValuePresentValueBuilder) WithPresentValue(presentValue BACnetApplicationTagReal) BACnetConstructedDataAnalogValuePresentValueBuilder {
	b.PresentValue = presentValue
	return b
}

func (b *_BACnetConstructedDataAnalogValuePresentValueBuilder) WithPresentValueBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataAnalogValuePresentValueBuilder {
	builder := builderSupplier(b.PresentValue.CreateBACnetApplicationTagRealBuilder())
	var err error
	b.PresentValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAnalogValuePresentValueBuilder) Build() (BACnetConstructedDataAnalogValuePresentValue, error) {
	if b.PresentValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAnalogValuePresentValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataAnalogValuePresentValueBuilder) MustBuild() BACnetConstructedDataAnalogValuePresentValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataAnalogValuePresentValueBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAnalogValuePresentValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAnalogValuePresentValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAnalogValuePresentValueBuilder().(*_BACnetConstructedDataAnalogValuePresentValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAnalogValuePresentValueBuilder creates a BACnetConstructedDataAnalogValuePresentValueBuilder
func (b *_BACnetConstructedDataAnalogValuePresentValue) CreateBACnetConstructedDataAnalogValuePresentValueBuilder() BACnetConstructedDataAnalogValuePresentValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataAnalogValuePresentValueBuilder()
	}
	return &_BACnetConstructedDataAnalogValuePresentValueBuilder{_BACnetConstructedDataAnalogValuePresentValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAnalogValuePresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ANALOG_VALUE
}

func (m *_BACnetConstructedDataAnalogValuePresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAnalogValuePresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogValuePresentValue) GetPresentValue() BACnetApplicationTagReal {
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

func (m *_BACnetConstructedDataAnalogValuePresentValue) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAnalogValuePresentValue(structType any) BACnetConstructedDataAnalogValuePresentValue {
	if casted, ok := structType.(BACnetConstructedDataAnalogValuePresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAnalogValuePresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAnalogValuePresentValue) GetTypeName() string {
	return "BACnetConstructedDataAnalogValuePresentValue"
}

func (m *_BACnetConstructedDataAnalogValuePresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAnalogValuePresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAnalogValuePresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAnalogValuePresentValue BACnetConstructedDataAnalogValuePresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAnalogValuePresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAnalogValuePresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "presentValue", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAnalogValuePresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAnalogValuePresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAnalogValuePresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAnalogValuePresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAnalogValuePresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAnalogValuePresentValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAnalogValuePresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAnalogValuePresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAnalogValuePresentValue) IsBACnetConstructedDataAnalogValuePresentValue() {
}

func (m *_BACnetConstructedDataAnalogValuePresentValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAnalogValuePresentValue) deepCopy() *_BACnetConstructedDataAnalogValuePresentValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAnalogValuePresentValueCopy := &_BACnetConstructedDataAnalogValuePresentValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.PresentValue.DeepCopy().(BACnetApplicationTagReal),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAnalogValuePresentValueCopy
}

func (m *_BACnetConstructedDataAnalogValuePresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
