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

// BACnetConstructedDataLightingOutputPresentValue is the corresponding interface of BACnetConstructedDataLightingOutputPresentValue
type BACnetConstructedDataLightingOutputPresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataLightingOutputPresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLightingOutputPresentValue()
	// CreateBuilder creates a BACnetConstructedDataLightingOutputPresentValueBuilder
	CreateBACnetConstructedDataLightingOutputPresentValueBuilder() BACnetConstructedDataLightingOutputPresentValueBuilder
}

// _BACnetConstructedDataLightingOutputPresentValue is the data-structure of this message
type _BACnetConstructedDataLightingOutputPresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetApplicationTagReal
}

var _ BACnetConstructedDataLightingOutputPresentValue = (*_BACnetConstructedDataLightingOutputPresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLightingOutputPresentValue)(nil)

// NewBACnetConstructedDataLightingOutputPresentValue factory function for _BACnetConstructedDataLightingOutputPresentValue
func NewBACnetConstructedDataLightingOutputPresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLightingOutputPresentValue {
	if presentValue == nil {
		panic("presentValue of type BACnetApplicationTagReal for BACnetConstructedDataLightingOutputPresentValue must not be nil")
	}
	_result := &_BACnetConstructedDataLightingOutputPresentValue{
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

// BACnetConstructedDataLightingOutputPresentValueBuilder is a builder for BACnetConstructedDataLightingOutputPresentValue
type BACnetConstructedDataLightingOutputPresentValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(presentValue BACnetApplicationTagReal) BACnetConstructedDataLightingOutputPresentValueBuilder
	// WithPresentValue adds PresentValue (property field)
	WithPresentValue(BACnetApplicationTagReal) BACnetConstructedDataLightingOutputPresentValueBuilder
	// WithPresentValueBuilder adds PresentValue (property field) which is build by the builder
	WithPresentValueBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataLightingOutputPresentValueBuilder
	// Build builds the BACnetConstructedDataLightingOutputPresentValue or returns an error if something is wrong
	Build() (BACnetConstructedDataLightingOutputPresentValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLightingOutputPresentValue
}

// NewBACnetConstructedDataLightingOutputPresentValueBuilder() creates a BACnetConstructedDataLightingOutputPresentValueBuilder
func NewBACnetConstructedDataLightingOutputPresentValueBuilder() BACnetConstructedDataLightingOutputPresentValueBuilder {
	return &_BACnetConstructedDataLightingOutputPresentValueBuilder{_BACnetConstructedDataLightingOutputPresentValue: new(_BACnetConstructedDataLightingOutputPresentValue)}
}

type _BACnetConstructedDataLightingOutputPresentValueBuilder struct {
	*_BACnetConstructedDataLightingOutputPresentValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLightingOutputPresentValueBuilder) = (*_BACnetConstructedDataLightingOutputPresentValueBuilder)(nil)

func (b *_BACnetConstructedDataLightingOutputPresentValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataLightingOutputPresentValueBuilder) WithMandatoryFields(presentValue BACnetApplicationTagReal) BACnetConstructedDataLightingOutputPresentValueBuilder {
	return b.WithPresentValue(presentValue)
}

func (b *_BACnetConstructedDataLightingOutputPresentValueBuilder) WithPresentValue(presentValue BACnetApplicationTagReal) BACnetConstructedDataLightingOutputPresentValueBuilder {
	b.PresentValue = presentValue
	return b
}

func (b *_BACnetConstructedDataLightingOutputPresentValueBuilder) WithPresentValueBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataLightingOutputPresentValueBuilder {
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

func (b *_BACnetConstructedDataLightingOutputPresentValueBuilder) Build() (BACnetConstructedDataLightingOutputPresentValue, error) {
	if b.PresentValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'presentValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLightingOutputPresentValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataLightingOutputPresentValueBuilder) MustBuild() BACnetConstructedDataLightingOutputPresentValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataLightingOutputPresentValueBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLightingOutputPresentValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLightingOutputPresentValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLightingOutputPresentValueBuilder().(*_BACnetConstructedDataLightingOutputPresentValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLightingOutputPresentValueBuilder creates a BACnetConstructedDataLightingOutputPresentValueBuilder
func (b *_BACnetConstructedDataLightingOutputPresentValue) CreateBACnetConstructedDataLightingOutputPresentValueBuilder() BACnetConstructedDataLightingOutputPresentValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataLightingOutputPresentValueBuilder()
	}
	return &_BACnetConstructedDataLightingOutputPresentValueBuilder{_BACnetConstructedDataLightingOutputPresentValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLightingOutputPresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_LIGHTING_OUTPUT
}

func (m *_BACnetConstructedDataLightingOutputPresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLightingOutputPresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLightingOutputPresentValue) GetPresentValue() BACnetApplicationTagReal {
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

func (m *_BACnetConstructedDataLightingOutputPresentValue) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLightingOutputPresentValue(structType any) BACnetConstructedDataLightingOutputPresentValue {
	if casted, ok := structType.(BACnetConstructedDataLightingOutputPresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLightingOutputPresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLightingOutputPresentValue) GetTypeName() string {
	return "BACnetConstructedDataLightingOutputPresentValue"
}

func (m *_BACnetConstructedDataLightingOutputPresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLightingOutputPresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLightingOutputPresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLightingOutputPresentValue BACnetConstructedDataLightingOutputPresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLightingOutputPresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLightingOutputPresentValue")
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

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLightingOutputPresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLightingOutputPresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLightingOutputPresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLightingOutputPresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLightingOutputPresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLightingOutputPresentValue")
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

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLightingOutputPresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLightingOutputPresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLightingOutputPresentValue) IsBACnetConstructedDataLightingOutputPresentValue() {
}

func (m *_BACnetConstructedDataLightingOutputPresentValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLightingOutputPresentValue) deepCopy() *_BACnetConstructedDataLightingOutputPresentValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLightingOutputPresentValueCopy := &_BACnetConstructedDataLightingOutputPresentValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.PresentValue.DeepCopy().(BACnetApplicationTagReal),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLightingOutputPresentValueCopy
}

func (m *_BACnetConstructedDataLightingOutputPresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
