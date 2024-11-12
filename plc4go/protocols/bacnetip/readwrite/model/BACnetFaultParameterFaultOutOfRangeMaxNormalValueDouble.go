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

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	// GetDoubleValue returns DoubleValue (property field)
	GetDoubleValue() BACnetApplicationTagDouble
	// IsBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble()
	// CreateBuilder creates a BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder
	CreateBACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder() BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder
}

// _BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble struct {
	BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract
	DoubleValue BACnetApplicationTagDouble
}

var _ BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble = (*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble)(nil)
var _ BACnetFaultParameterFaultOutOfRangeMaxNormalValueRequirements = (*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble)(nil)

// NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble factory function for _BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble
func NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, doubleValue BACnetApplicationTagDouble, tagNumber uint8) *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble {
	if doubleValue == nil {
		panic("doubleValue of type BACnetApplicationTagDouble for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble must not be nil")
	}
	_result := &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble{
		BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract: NewBACnetFaultParameterFaultOutOfRangeMaxNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		DoubleValue: doubleValue,
	}
	_result.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder is a builder for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(doubleValue BACnetApplicationTagDouble) BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder
	// WithDoubleValue adds DoubleValue (property field)
	WithDoubleValue(BACnetApplicationTagDouble) BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder
	// WithDoubleValueBuilder adds DoubleValue (property field) which is build by the builder
	WithDoubleValueBuilder(func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder
	// Build builds the BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble or returns an error if something is wrong
	Build() (BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble
}

// NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder() creates a BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder
func NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder() BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder {
	return &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder{_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble: new(_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble)}
}

type _BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder struct {
	*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble

	parentBuilder *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueBuilder

	err *utils.MultiError
}

var _ (BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder) = (*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder)(nil)

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder) setParent(contract BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract) {
	b.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract = contract
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder) WithMandatoryFields(doubleValue BACnetApplicationTagDouble) BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder {
	return b.WithDoubleValue(doubleValue)
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder) WithDoubleValue(doubleValue BACnetApplicationTagDouble) BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder {
	b.DoubleValue = doubleValue
	return b
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder) WithDoubleValueBuilder(builderSupplier func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder {
	builder := builderSupplier(b.DoubleValue.CreateBACnetApplicationTagDoubleBuilder())
	var err error
	b.DoubleValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDoubleBuilder failed"))
	}
	return b
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder) Build() (BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble, error) {
	if b.DoubleValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'doubleValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble.deepCopy(), nil
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder) MustBuild() BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder) Done() BACnetFaultParameterFaultOutOfRangeMaxNormalValueBuilder {
	return b.parentBuilder
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder) buildForBACnetFaultParameterFaultOutOfRangeMaxNormalValue() (BACnetFaultParameterFaultOutOfRangeMaxNormalValue, error) {
	return b.Build()
}

func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder) DeepCopy() any {
	_copy := b.CreateBACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder().(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder creates a BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder
func (b *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) CreateBACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder() BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder {
	if b == nil {
		return NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder()
	}
	return &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleBuilder{_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble: b.deepCopy()}
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

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetParent() BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract {
	return m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetDoubleValue() BACnetApplicationTagDouble {
	return m.DoubleValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble(structType any) BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue).getLengthInBits(ctx))

	// Simple field (doubleValue)
	lengthInBits += m.DoubleValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetFaultParameterFaultOutOfRangeMaxNormalValue, tagNumber uint8) (__bACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble, err error) {
	m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doubleValue, err := ReadSimpleField[BACnetApplicationTagDouble](ctx, "doubleValue", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doubleValue' field"))
	}
	m.DoubleValue = doubleValue

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble")
	}

	return m, nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble")
		}

		if err := WriteSimpleField[BACnetApplicationTagDouble](ctx, "doubleValue", m.GetDoubleValue(), WriteComplex[BACnetApplicationTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'doubleValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble")
		}
		return nil
	}
	return m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) IsBACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble() {
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) deepCopy() *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble {
	if m == nil {
		return nil
	}
	_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleCopy := &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble{
		m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagDouble](m.DoubleValue),
	}
	m.BACnetFaultParameterFaultOutOfRangeMaxNormalValueContract.(*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue)._SubType = m
	return _BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleCopy
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble) String() string {
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
