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

// BACnetConstructedDataAdjustValue is the corresponding interface of BACnetConstructedDataAdjustValue
type BACnetConstructedDataAdjustValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetAdjustValue returns AdjustValue (property field)
	GetAdjustValue() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
	// IsBACnetConstructedDataAdjustValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAdjustValue()
	// CreateBuilder creates a BACnetConstructedDataAdjustValueBuilder
	CreateBACnetConstructedDataAdjustValueBuilder() BACnetConstructedDataAdjustValueBuilder
}

// _BACnetConstructedDataAdjustValue is the data-structure of this message
type _BACnetConstructedDataAdjustValue struct {
	BACnetConstructedDataContract
	AdjustValue BACnetApplicationTagSignedInteger
}

var _ BACnetConstructedDataAdjustValue = (*_BACnetConstructedDataAdjustValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAdjustValue)(nil)

// NewBACnetConstructedDataAdjustValue factory function for _BACnetConstructedDataAdjustValue
func NewBACnetConstructedDataAdjustValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, adjustValue BACnetApplicationTagSignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAdjustValue {
	if adjustValue == nil {
		panic("adjustValue of type BACnetApplicationTagSignedInteger for BACnetConstructedDataAdjustValue must not be nil")
	}
	_result := &_BACnetConstructedDataAdjustValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AdjustValue:                   adjustValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAdjustValueBuilder is a builder for BACnetConstructedDataAdjustValue
type BACnetConstructedDataAdjustValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(adjustValue BACnetApplicationTagSignedInteger) BACnetConstructedDataAdjustValueBuilder
	// WithAdjustValue adds AdjustValue (property field)
	WithAdjustValue(BACnetApplicationTagSignedInteger) BACnetConstructedDataAdjustValueBuilder
	// WithAdjustValueBuilder adds AdjustValue (property field) which is build by the builder
	WithAdjustValueBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataAdjustValueBuilder
	// Build builds the BACnetConstructedDataAdjustValue or returns an error if something is wrong
	Build() (BACnetConstructedDataAdjustValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAdjustValue
}

// NewBACnetConstructedDataAdjustValueBuilder() creates a BACnetConstructedDataAdjustValueBuilder
func NewBACnetConstructedDataAdjustValueBuilder() BACnetConstructedDataAdjustValueBuilder {
	return &_BACnetConstructedDataAdjustValueBuilder{_BACnetConstructedDataAdjustValue: new(_BACnetConstructedDataAdjustValue)}
}

type _BACnetConstructedDataAdjustValueBuilder struct {
	*_BACnetConstructedDataAdjustValue

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAdjustValueBuilder) = (*_BACnetConstructedDataAdjustValueBuilder)(nil)

func (b *_BACnetConstructedDataAdjustValueBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataAdjustValueBuilder) WithMandatoryFields(adjustValue BACnetApplicationTagSignedInteger) BACnetConstructedDataAdjustValueBuilder {
	return b.WithAdjustValue(adjustValue)
}

func (b *_BACnetConstructedDataAdjustValueBuilder) WithAdjustValue(adjustValue BACnetApplicationTagSignedInteger) BACnetConstructedDataAdjustValueBuilder {
	b.AdjustValue = adjustValue
	return b
}

func (b *_BACnetConstructedDataAdjustValueBuilder) WithAdjustValueBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConstructedDataAdjustValueBuilder {
	builder := builderSupplier(b.AdjustValue.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	b.AdjustValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAdjustValueBuilder) Build() (BACnetConstructedDataAdjustValue, error) {
	if b.AdjustValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'adjustValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAdjustValue.deepCopy(), nil
}

func (b *_BACnetConstructedDataAdjustValueBuilder) MustBuild() BACnetConstructedDataAdjustValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataAdjustValueBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAdjustValueBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAdjustValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAdjustValueBuilder().(*_BACnetConstructedDataAdjustValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAdjustValueBuilder creates a BACnetConstructedDataAdjustValueBuilder
func (b *_BACnetConstructedDataAdjustValue) CreateBACnetConstructedDataAdjustValueBuilder() BACnetConstructedDataAdjustValueBuilder {
	if b == nil {
		return NewBACnetConstructedDataAdjustValueBuilder()
	}
	return &_BACnetConstructedDataAdjustValueBuilder{_BACnetConstructedDataAdjustValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAdjustValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAdjustValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ADJUST_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAdjustValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAdjustValue) GetAdjustValue() BACnetApplicationTagSignedInteger {
	return m.AdjustValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAdjustValue) GetActualValue() BACnetApplicationTagSignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagSignedInteger(m.GetAdjustValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAdjustValue(structType any) BACnetConstructedDataAdjustValue {
	if casted, ok := structType.(BACnetConstructedDataAdjustValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAdjustValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAdjustValue) GetTypeName() string {
	return "BACnetConstructedDataAdjustValue"
}

func (m *_BACnetConstructedDataAdjustValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (adjustValue)
	lengthInBits += m.AdjustValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAdjustValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAdjustValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAdjustValue BACnetConstructedDataAdjustValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAdjustValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAdjustValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	adjustValue, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "adjustValue", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'adjustValue' field"))
	}
	m.AdjustValue = adjustValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagSignedInteger](ctx, "actualValue", (*BACnetApplicationTagSignedInteger)(nil), adjustValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAdjustValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAdjustValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAdjustValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAdjustValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAdjustValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAdjustValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "adjustValue", m.GetAdjustValue(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'adjustValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAdjustValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAdjustValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAdjustValue) IsBACnetConstructedDataAdjustValue() {}

func (m *_BACnetConstructedDataAdjustValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAdjustValue) deepCopy() *_BACnetConstructedDataAdjustValue {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAdjustValueCopy := &_BACnetConstructedDataAdjustValue{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.AdjustValue.DeepCopy().(BACnetApplicationTagSignedInteger),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAdjustValueCopy
}

func (m *_BACnetConstructedDataAdjustValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
