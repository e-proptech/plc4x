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

// BACnetConstructedDataAnalogOutputMaxPresValue is the corresponding interface of BACnetConstructedDataAnalogOutputMaxPresValue
type BACnetConstructedDataAnalogOutputMaxPresValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMaxPresValue returns MaxPresValue (property field)
	GetMaxPresValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataAnalogOutputMaxPresValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAnalogOutputMaxPresValue()
	// CreateBuilder creates a BACnetConstructedDataAnalogOutputMaxPresValueBuilder
	CreateBACnetConstructedDataAnalogOutputMaxPresValueBuilder() BACnetConstructedDataAnalogOutputMaxPresValueBuilder
}

// _BACnetConstructedDataAnalogOutputMaxPresValue is the data-structure of this message
type _BACnetConstructedDataAnalogOutputMaxPresValue struct {
	BACnetConstructedDataContract
	MaxPresValue BACnetApplicationTagReal
}

var _ BACnetConstructedDataAnalogOutputMaxPresValue = (*_BACnetConstructedDataAnalogOutputMaxPresValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAnalogOutputMaxPresValue)(nil)

// NewBACnetConstructedDataAnalogOutputMaxPresValue factory function for _BACnetConstructedDataAnalogOutputMaxPresValue
func NewBACnetConstructedDataAnalogOutputMaxPresValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, maxPresValue BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAnalogOutputMaxPresValue {
	if maxPresValue == nil {
		panic("maxPresValue of type BACnetApplicationTagReal for BACnetConstructedDataAnalogOutputMaxPresValue must not be nil")
	}
	_result := &_BACnetConstructedDataAnalogOutputMaxPresValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaxPresValue:                  maxPresValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAnalogOutputMaxPresValueBuilder is a builder for BACnetConstructedDataAnalogOutputMaxPresValue
type BACnetConstructedDataAnalogOutputMaxPresValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(maxPresValue BACnetApplicationTagReal) BACnetConstructedDataAnalogOutputMaxPresValueBuilder
	// WithMaxPresValue adds MaxPresValue (property field)
	WithMaxPresValue(BACnetApplicationTagReal) BACnetConstructedDataAnalogOutputMaxPresValueBuilder
	// WithMaxPresValueBuilder adds MaxPresValue (property field) which is build by the builder
	WithMaxPresValueBuilder(func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataAnalogOutputMaxPresValueBuilder
	// Build builds the BACnetConstructedDataAnalogOutputMaxPresValue or returns an error if something is wrong
	Build() (BACnetConstructedDataAnalogOutputMaxPresValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAnalogOutputMaxPresValue
}

// NewBACnetConstructedDataAnalogOutputMaxPresValueBuilder() creates a BACnetConstructedDataAnalogOutputMaxPresValueBuilder
func NewBACnetConstructedDataAnalogOutputMaxPresValueBuilder() BACnetConstructedDataAnalogOutputMaxPresValueBuilder {
	return &_BACnetConstructedDataAnalogOutputMaxPresValueBuilder{_BACnetConstructedDataAnalogOutputMaxPresValue: new(_BACnetConstructedDataAnalogOutputMaxPresValue)}
}

type _BACnetConstructedDataAnalogOutputMaxPresValueBuilder struct {
	*_BACnetConstructedDataAnalogOutputMaxPresValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAnalogOutputMaxPresValueBuilder) = (*_BACnetConstructedDataAnalogOutputMaxPresValueBuilder)(nil)

func (b *_BACnetConstructedDataAnalogOutputMaxPresValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataAnalogOutputMaxPresValueBuilder) WithMandatoryFields(maxPresValue BACnetApplicationTagReal) BACnetConstructedDataAnalogOutputMaxPresValueBuilder {
	return b.WithMaxPresValue(maxPresValue)
}

func (b *_BACnetConstructedDataAnalogOutputMaxPresValueBuilder) WithMaxPresValue(maxPresValue BACnetApplicationTagReal) BACnetConstructedDataAnalogOutputMaxPresValueBuilder {
	b.MaxPresValue = maxPresValue
	return b
}

func (b *_BACnetConstructedDataAnalogOutputMaxPresValueBuilder) WithMaxPresValueBuilder(builderSupplier func(BACnetApplicationTagRealBuilder) BACnetApplicationTagRealBuilder) BACnetConstructedDataAnalogOutputMaxPresValueBuilder {
	builder := builderSupplier(b.MaxPresValue.CreateBACnetApplicationTagRealBuilder())
	var err error
	b.MaxPresValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAnalogOutputMaxPresValueBuilder) Build() (BACnetConstructedDataAnalogOutputMaxPresValue, error) {
	if b.MaxPresValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'maxPresValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAnalogOutputMaxPresValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataAnalogOutputMaxPresValueBuilder) MustBuild() BACnetConstructedDataAnalogOutputMaxPresValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataAnalogOutputMaxPresValueBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAnalogOutputMaxPresValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAnalogOutputMaxPresValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAnalogOutputMaxPresValueBuilder().(*_BACnetConstructedDataAnalogOutputMaxPresValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAnalogOutputMaxPresValueBuilder creates a BACnetConstructedDataAnalogOutputMaxPresValueBuilder
func (b *_BACnetConstructedDataAnalogOutputMaxPresValue) CreateBACnetConstructedDataAnalogOutputMaxPresValueBuilder() BACnetConstructedDataAnalogOutputMaxPresValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataAnalogOutputMaxPresValueBuilder()
	}
	return &_BACnetConstructedDataAnalogOutputMaxPresValueBuilder{_BACnetConstructedDataAnalogOutputMaxPresValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_ANALOG_OUTPUT
}

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_PRES_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) GetMaxPresValue() BACnetApplicationTagReal {
	return m.MaxPresValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetMaxPresValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAnalogOutputMaxPresValue(structType any) BACnetConstructedDataAnalogOutputMaxPresValue {
	if casted, ok := structType.(BACnetConstructedDataAnalogOutputMaxPresValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAnalogOutputMaxPresValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) GetTypeName() string {
	return "BACnetConstructedDataAnalogOutputMaxPresValue"
}

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (maxPresValue)
	lengthInBits += m.MaxPresValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAnalogOutputMaxPresValue BACnetConstructedDataAnalogOutputMaxPresValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAnalogOutputMaxPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAnalogOutputMaxPresValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maxPresValue, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "maxPresValue", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxPresValue' field"))
	}
	m.MaxPresValue = maxPresValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), maxPresValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAnalogOutputMaxPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAnalogOutputMaxPresValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAnalogOutputMaxPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAnalogOutputMaxPresValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "maxPresValue", m.GetMaxPresValue(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxPresValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAnalogOutputMaxPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAnalogOutputMaxPresValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) IsBACnetConstructedDataAnalogOutputMaxPresValue() {
}

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) deepCopy() *_BACnetConstructedDataAnalogOutputMaxPresValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAnalogOutputMaxPresValueCopy := &_BACnetConstructedDataAnalogOutputMaxPresValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagReal](m.MaxPresValue),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAnalogOutputMaxPresValueCopy
}

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) String() string {
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
