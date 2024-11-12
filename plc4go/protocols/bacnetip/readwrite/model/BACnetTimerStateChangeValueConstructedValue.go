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

// BACnetTimerStateChangeValueConstructedValue is the corresponding interface of BACnetTimerStateChangeValueConstructedValue
type BACnetTimerStateChangeValueConstructedValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetTimerStateChangeValue
	// GetConstructedValue returns ConstructedValue (property field)
	GetConstructedValue() BACnetConstructedData
	// IsBACnetTimerStateChangeValueConstructedValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimerStateChangeValueConstructedValue()
	// CreateBuilder creates a BACnetTimerStateChangeValueConstructedValueBuilder
	CreateBACnetTimerStateChangeValueConstructedValueBuilder() BACnetTimerStateChangeValueConstructedValueBuilder
}

// _BACnetTimerStateChangeValueConstructedValue is the data-structure of this message
type _BACnetTimerStateChangeValueConstructedValue struct {
	BACnetTimerStateChangeValueContract
	ConstructedValue BACnetConstructedData
}

var _ BACnetTimerStateChangeValueConstructedValue = (*_BACnetTimerStateChangeValueConstructedValue)(nil)
var _ BACnetTimerStateChangeValueRequirements = (*_BACnetTimerStateChangeValueConstructedValue)(nil)

// NewBACnetTimerStateChangeValueConstructedValue factory function for _BACnetTimerStateChangeValueConstructedValue
func NewBACnetTimerStateChangeValueConstructedValue(peekedTagHeader BACnetTagHeader, constructedValue BACnetConstructedData, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueConstructedValue {
	if constructedValue == nil {
		panic("constructedValue of type BACnetConstructedData for BACnetTimerStateChangeValueConstructedValue must not be nil")
	}
	_result := &_BACnetTimerStateChangeValueConstructedValue{
		BACnetTimerStateChangeValueContract: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
		ConstructedValue:                    constructedValue,
	}
	_result.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTimerStateChangeValueConstructedValueBuilder is a builder for BACnetTimerStateChangeValueConstructedValue
type BACnetTimerStateChangeValueConstructedValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(constructedValue BACnetConstructedData) BACnetTimerStateChangeValueConstructedValueBuilder
	// WithConstructedValue adds ConstructedValue (property field)
	WithConstructedValue(BACnetConstructedData) BACnetTimerStateChangeValueConstructedValueBuilder
	// WithConstructedValueBuilder adds ConstructedValue (property field) which is build by the builder
	WithConstructedValueBuilder(func(BACnetConstructedDataBuilder) BACnetConstructedDataBuilder) BACnetTimerStateChangeValueConstructedValueBuilder
	// Build builds the BACnetTimerStateChangeValueConstructedValue or returns an error if something is wrong
	Build() (BACnetTimerStateChangeValueConstructedValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTimerStateChangeValueConstructedValue
}

// NewBACnetTimerStateChangeValueConstructedValueBuilder() creates a BACnetTimerStateChangeValueConstructedValueBuilder
func NewBACnetTimerStateChangeValueConstructedValueBuilder() BACnetTimerStateChangeValueConstructedValueBuilder {
	return &_BACnetTimerStateChangeValueConstructedValueBuilder{_BACnetTimerStateChangeValueConstructedValue: new(_BACnetTimerStateChangeValueConstructedValue)}
}

type _BACnetTimerStateChangeValueConstructedValueBuilder struct {
	*_BACnetTimerStateChangeValueConstructedValue

	parentBuilder *_BACnetTimerStateChangeValueBuilder

	err *utils.MultiError
}

var _ (BACnetTimerStateChangeValueConstructedValueBuilder) = (*_BACnetTimerStateChangeValueConstructedValueBuilder)(nil)

func (b *_BACnetTimerStateChangeValueConstructedValueBuilder) setParent(contract BACnetTimerStateChangeValueContract) {
	b.BACnetTimerStateChangeValueContract = contract
}

func (b *_BACnetTimerStateChangeValueConstructedValueBuilder) WithMandatoryFields(constructedValue BACnetConstructedData) BACnetTimerStateChangeValueConstructedValueBuilder {
	return b.WithConstructedValue(constructedValue)
}

func (b *_BACnetTimerStateChangeValueConstructedValueBuilder) WithConstructedValue(constructedValue BACnetConstructedData) BACnetTimerStateChangeValueConstructedValueBuilder {
	b.ConstructedValue = constructedValue
	return b
}

func (b *_BACnetTimerStateChangeValueConstructedValueBuilder) WithConstructedValueBuilder(builderSupplier func(BACnetConstructedDataBuilder) BACnetConstructedDataBuilder) BACnetTimerStateChangeValueConstructedValueBuilder {
	builder := builderSupplier(b.ConstructedValue.CreateBACnetConstructedDataBuilder())
	var err error
	b.ConstructedValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetConstructedDataBuilder failed"))
	}
	return b
}

func (b *_BACnetTimerStateChangeValueConstructedValueBuilder) Build() (BACnetTimerStateChangeValueConstructedValue, error) {
	if b.ConstructedValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'constructedValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTimerStateChangeValueConstructedValue.deepCopy(), nil
}

func (b *_BACnetTimerStateChangeValueConstructedValueBuilder) MustBuild() BACnetTimerStateChangeValueConstructedValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetTimerStateChangeValueConstructedValueBuilder) Done() BACnetTimerStateChangeValueBuilder {
	return b.parentBuilder
}

func (b *_BACnetTimerStateChangeValueConstructedValueBuilder) buildForBACnetTimerStateChangeValue() (BACnetTimerStateChangeValue, error) {
	return b.Build()
}

func (b *_BACnetTimerStateChangeValueConstructedValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTimerStateChangeValueConstructedValueBuilder().(*_BACnetTimerStateChangeValueConstructedValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTimerStateChangeValueConstructedValueBuilder creates a BACnetTimerStateChangeValueConstructedValueBuilder
func (b *_BACnetTimerStateChangeValueConstructedValue) CreateBACnetTimerStateChangeValueConstructedValueBuilder() BACnetTimerStateChangeValueConstructedValueBuilder {
	if b == nil {
		return NewBACnetTimerStateChangeValueConstructedValueBuilder()
	}
	return &_BACnetTimerStateChangeValueConstructedValueBuilder{_BACnetTimerStateChangeValueConstructedValue: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetTimerStateChangeValueConstructedValue) GetParent() BACnetTimerStateChangeValueContract {
	return m.BACnetTimerStateChangeValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueConstructedValue) GetConstructedValue() BACnetConstructedData {
	return m.ConstructedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueConstructedValue(structType any) BACnetTimerStateChangeValueConstructedValue {
	if casted, ok := structType.(BACnetTimerStateChangeValueConstructedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueConstructedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueConstructedValue) GetTypeName() string {
	return "BACnetTimerStateChangeValueConstructedValue"
}

func (m *_BACnetTimerStateChangeValueConstructedValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).getLengthInBits(ctx))

	// Simple field (constructedValue)
	lengthInBits += m.ConstructedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueConstructedValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetTimerStateChangeValueConstructedValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetTimerStateChangeValue, objectTypeArgument BACnetObjectType) (__bACnetTimerStateChangeValueConstructedValue BACnetTimerStateChangeValueConstructedValue, err error) {
	m.BACnetTimerStateChangeValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueConstructedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueConstructedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	constructedValue, err := ReadSimpleField[BACnetConstructedData](ctx, "constructedValue", ReadComplex[BACnetConstructedData](BACnetConstructedDataParseWithBufferProducer[BACnetConstructedData]((uint8)(uint8(1)), (BACnetObjectType)(objectTypeArgument), (BACnetPropertyIdentifier)(BACnetPropertyIdentifier_VENDOR_PROPRIETARY_VALUE), (BACnetTagPayloadUnsignedInteger)(nil)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'constructedValue' field"))
	}
	m.ConstructedValue = constructedValue

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueConstructedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueConstructedValue")
	}

	return m, nil
}

func (m *_BACnetTimerStateChangeValueConstructedValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueConstructedValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueConstructedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueConstructedValue")
		}

		if err := WriteSimpleField[BACnetConstructedData](ctx, "constructedValue", m.GetConstructedValue(), WriteComplex[BACnetConstructedData](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'constructedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueConstructedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueConstructedValue")
		}
		return nil
	}
	return m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueConstructedValue) IsBACnetTimerStateChangeValueConstructedValue() {
}

func (m *_BACnetTimerStateChangeValueConstructedValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTimerStateChangeValueConstructedValue) deepCopy() *_BACnetTimerStateChangeValueConstructedValue {
	if m == nil {
		return nil
	}
	_BACnetTimerStateChangeValueConstructedValueCopy := &_BACnetTimerStateChangeValueConstructedValue{
		m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).deepCopy(),
		utils.DeepCopy[BACnetConstructedData](m.ConstructedValue),
	}
	m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = m
	return _BACnetTimerStateChangeValueConstructedValueCopy
}

func (m *_BACnetTimerStateChangeValueConstructedValue) String() string {
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
