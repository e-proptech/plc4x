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

// BACnetPropertyStatesExtendedValue is the corresponding interface of BACnetPropertyStatesExtendedValue
type BACnetPropertyStatesExtendedValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPropertyStates
	// GetExtendedValue returns ExtendedValue (property field)
	GetExtendedValue() BACnetContextTagUnsignedInteger
	// IsBACnetPropertyStatesExtendedValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesExtendedValue()
	// CreateBuilder creates a BACnetPropertyStatesExtendedValueBuilder
	CreateBACnetPropertyStatesExtendedValueBuilder() BACnetPropertyStatesExtendedValueBuilder
}

// _BACnetPropertyStatesExtendedValue is the data-structure of this message
type _BACnetPropertyStatesExtendedValue struct {
	BACnetPropertyStatesContract
	ExtendedValue BACnetContextTagUnsignedInteger
}

var _ BACnetPropertyStatesExtendedValue = (*_BACnetPropertyStatesExtendedValue)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesExtendedValue)(nil)

// NewBACnetPropertyStatesExtendedValue factory function for _BACnetPropertyStatesExtendedValue
func NewBACnetPropertyStatesExtendedValue(peekedTagHeader BACnetTagHeader, extendedValue BACnetContextTagUnsignedInteger) *_BACnetPropertyStatesExtendedValue {
	if extendedValue == nil {
		panic("extendedValue of type BACnetContextTagUnsignedInteger for BACnetPropertyStatesExtendedValue must not be nil")
	}
	_result := &_BACnetPropertyStatesExtendedValue{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		ExtendedValue:                extendedValue,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPropertyStatesExtendedValueBuilder is a builder for BACnetPropertyStatesExtendedValue
type BACnetPropertyStatesExtendedValueBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(extendedValue BACnetContextTagUnsignedInteger) BACnetPropertyStatesExtendedValueBuilder
	// WithExtendedValue adds ExtendedValue (property field)
	WithExtendedValue(BACnetContextTagUnsignedInteger) BACnetPropertyStatesExtendedValueBuilder
	// WithExtendedValueBuilder adds ExtendedValue (property field) which is build by the builder
	WithExtendedValueBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetPropertyStatesExtendedValueBuilder
	// Build builds the BACnetPropertyStatesExtendedValue or returns an error if something is wrong
	Build() (BACnetPropertyStatesExtendedValue, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPropertyStatesExtendedValue
}

// NewBACnetPropertyStatesExtendedValueBuilder() creates a BACnetPropertyStatesExtendedValueBuilder
func NewBACnetPropertyStatesExtendedValueBuilder() BACnetPropertyStatesExtendedValueBuilder {
	return &_BACnetPropertyStatesExtendedValueBuilder{_BACnetPropertyStatesExtendedValue: new(_BACnetPropertyStatesExtendedValue)}
}

type _BACnetPropertyStatesExtendedValueBuilder struct {
	*_BACnetPropertyStatesExtendedValue

	parentBuilder *_BACnetPropertyStatesBuilder

	err *utils.MultiError
}

var _ (BACnetPropertyStatesExtendedValueBuilder) = (*_BACnetPropertyStatesExtendedValueBuilder)(nil)

func (b *_BACnetPropertyStatesExtendedValueBuilder) setParent(contract BACnetPropertyStatesContract) {
	b.BACnetPropertyStatesContract = contract
}

func (b *_BACnetPropertyStatesExtendedValueBuilder) WithMandatoryFields(extendedValue BACnetContextTagUnsignedInteger) BACnetPropertyStatesExtendedValueBuilder {
	return b.WithExtendedValue(extendedValue)
}

func (b *_BACnetPropertyStatesExtendedValueBuilder) WithExtendedValue(extendedValue BACnetContextTagUnsignedInteger) BACnetPropertyStatesExtendedValueBuilder {
	b.ExtendedValue = extendedValue
	return b
}

func (b *_BACnetPropertyStatesExtendedValueBuilder) WithExtendedValueBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetPropertyStatesExtendedValueBuilder {
	builder := builderSupplier(b.ExtendedValue.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.ExtendedValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetPropertyStatesExtendedValueBuilder) Build() (BACnetPropertyStatesExtendedValue, error) {
	if b.ExtendedValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'extendedValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPropertyStatesExtendedValue.deepCopy(), nil
}

func (b *_BACnetPropertyStatesExtendedValueBuilder) MustBuild() BACnetPropertyStatesExtendedValue {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetPropertyStatesExtendedValueBuilder) Done() BACnetPropertyStatesBuilder {
	return b.parentBuilder
}

func (b *_BACnetPropertyStatesExtendedValueBuilder) buildForBACnetPropertyStates() (BACnetPropertyStates, error) {
	return b.Build()
}

func (b *_BACnetPropertyStatesExtendedValueBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPropertyStatesExtendedValueBuilder().(*_BACnetPropertyStatesExtendedValueBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPropertyStatesExtendedValueBuilder creates a BACnetPropertyStatesExtendedValueBuilder
func (b *_BACnetPropertyStatesExtendedValue) CreateBACnetPropertyStatesExtendedValueBuilder() BACnetPropertyStatesExtendedValueBuilder {
	if b == nil {
		return NewBACnetPropertyStatesExtendedValueBuilder()
	}
	return &_BACnetPropertyStatesExtendedValueBuilder{_BACnetPropertyStatesExtendedValue: b.deepCopy()}
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

func (m *_BACnetPropertyStatesExtendedValue) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesExtendedValue) GetExtendedValue() BACnetContextTagUnsignedInteger {
	return m.ExtendedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesExtendedValue(structType any) BACnetPropertyStatesExtendedValue {
	if casted, ok := structType.(BACnetPropertyStatesExtendedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesExtendedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesExtendedValue) GetTypeName() string {
	return "BACnetPropertyStatesExtendedValue"
}

func (m *_BACnetPropertyStatesExtendedValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (extendedValue)
	lengthInBits += m.ExtendedValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesExtendedValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesExtendedValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesExtendedValue BACnetPropertyStatesExtendedValue, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesExtendedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesExtendedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	extendedValue, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "extendedValue", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(peekedTagNumber), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'extendedValue' field"))
	}
	m.ExtendedValue = extendedValue

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesExtendedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesExtendedValue")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesExtendedValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesExtendedValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesExtendedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesExtendedValue")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "extendedValue", m.GetExtendedValue(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'extendedValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesExtendedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesExtendedValue")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesExtendedValue) IsBACnetPropertyStatesExtendedValue() {}

func (m *_BACnetPropertyStatesExtendedValue) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPropertyStatesExtendedValue) deepCopy() *_BACnetPropertyStatesExtendedValue {
	if m == nil {
		return nil
	}
	_BACnetPropertyStatesExtendedValueCopy := &_BACnetPropertyStatesExtendedValue{
		m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).deepCopy(),
		m.ExtendedValue.DeepCopy().(BACnetContextTagUnsignedInteger),
	}
	m.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = m
	return _BACnetPropertyStatesExtendedValueCopy
}

func (m *_BACnetPropertyStatesExtendedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
