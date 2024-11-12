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

// BACnetChannelValueDouble is the corresponding interface of BACnetChannelValueDouble
type BACnetChannelValueDouble interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetChannelValue
	// GetDoubleValue returns DoubleValue (property field)
	GetDoubleValue() BACnetApplicationTagDouble
	// IsBACnetChannelValueDouble is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetChannelValueDouble()
	// CreateBuilder creates a BACnetChannelValueDoubleBuilder
	CreateBACnetChannelValueDoubleBuilder() BACnetChannelValueDoubleBuilder
}

// _BACnetChannelValueDouble is the data-structure of this message
type _BACnetChannelValueDouble struct {
	BACnetChannelValueContract
	DoubleValue BACnetApplicationTagDouble
}

var _ BACnetChannelValueDouble = (*_BACnetChannelValueDouble)(nil)
var _ BACnetChannelValueRequirements = (*_BACnetChannelValueDouble)(nil)

// NewBACnetChannelValueDouble factory function for _BACnetChannelValueDouble
func NewBACnetChannelValueDouble(peekedTagHeader BACnetTagHeader, doubleValue BACnetApplicationTagDouble) *_BACnetChannelValueDouble {
	if doubleValue == nil {
		panic("doubleValue of type BACnetApplicationTagDouble for BACnetChannelValueDouble must not be nil")
	}
	_result := &_BACnetChannelValueDouble{
		BACnetChannelValueContract: NewBACnetChannelValue(peekedTagHeader),
		DoubleValue:                doubleValue,
	}
	_result.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetChannelValueDoubleBuilder is a builder for BACnetChannelValueDouble
type BACnetChannelValueDoubleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(doubleValue BACnetApplicationTagDouble) BACnetChannelValueDoubleBuilder
	// WithDoubleValue adds DoubleValue (property field)
	WithDoubleValue(BACnetApplicationTagDouble) BACnetChannelValueDoubleBuilder
	// WithDoubleValueBuilder adds DoubleValue (property field) which is build by the builder
	WithDoubleValueBuilder(func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetChannelValueDoubleBuilder
	// Build builds the BACnetChannelValueDouble or returns an error if something is wrong
	Build() (BACnetChannelValueDouble, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetChannelValueDouble
}

// NewBACnetChannelValueDoubleBuilder() creates a BACnetChannelValueDoubleBuilder
func NewBACnetChannelValueDoubleBuilder() BACnetChannelValueDoubleBuilder {
	return &_BACnetChannelValueDoubleBuilder{_BACnetChannelValueDouble: new(_BACnetChannelValueDouble)}
}

type _BACnetChannelValueDoubleBuilder struct {
	*_BACnetChannelValueDouble

	parentBuilder *_BACnetChannelValueBuilder

	err *utils.MultiError
}

var _ (BACnetChannelValueDoubleBuilder) = (*_BACnetChannelValueDoubleBuilder)(nil)

func (b *_BACnetChannelValueDoubleBuilder) setParent(contract BACnetChannelValueContract) {
	b.BACnetChannelValueContract = contract
}

func (b *_BACnetChannelValueDoubleBuilder) WithMandatoryFields(doubleValue BACnetApplicationTagDouble) BACnetChannelValueDoubleBuilder {
	return b.WithDoubleValue(doubleValue)
}

func (b *_BACnetChannelValueDoubleBuilder) WithDoubleValue(doubleValue BACnetApplicationTagDouble) BACnetChannelValueDoubleBuilder {
	b.DoubleValue = doubleValue
	return b
}

func (b *_BACnetChannelValueDoubleBuilder) WithDoubleValueBuilder(builderSupplier func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetChannelValueDoubleBuilder {
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

func (b *_BACnetChannelValueDoubleBuilder) Build() (BACnetChannelValueDouble, error) {
	if b.DoubleValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'doubleValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetChannelValueDouble.deepCopy(), nil
}

func (b *_BACnetChannelValueDoubleBuilder) MustBuild() BACnetChannelValueDouble {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetChannelValueDoubleBuilder) Done() BACnetChannelValueBuilder {
	return b.parentBuilder
}

func (b *_BACnetChannelValueDoubleBuilder) buildForBACnetChannelValue() (BACnetChannelValue, error) {
	return b.Build()
}

func (b *_BACnetChannelValueDoubleBuilder) DeepCopy() any {
	_copy := b.CreateBACnetChannelValueDoubleBuilder().(*_BACnetChannelValueDoubleBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetChannelValueDoubleBuilder creates a BACnetChannelValueDoubleBuilder
func (b *_BACnetChannelValueDouble) CreateBACnetChannelValueDoubleBuilder() BACnetChannelValueDoubleBuilder {
	if b == nil {
		return NewBACnetChannelValueDoubleBuilder()
	}
	return &_BACnetChannelValueDoubleBuilder{_BACnetChannelValueDouble: b.deepCopy()}
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

func (m *_BACnetChannelValueDouble) GetParent() BACnetChannelValueContract {
	return m.BACnetChannelValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueDouble) GetDoubleValue() BACnetApplicationTagDouble {
	return m.DoubleValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueDouble(structType any) BACnetChannelValueDouble {
	if casted, ok := structType.(BACnetChannelValueDouble); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueDouble); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueDouble) GetTypeName() string {
	return "BACnetChannelValueDouble"
}

func (m *_BACnetChannelValueDouble) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetChannelValueContract.(*_BACnetChannelValue).getLengthInBits(ctx))

	// Simple field (doubleValue)
	lengthInBits += m.DoubleValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetChannelValueDouble) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetChannelValueDouble) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetChannelValue) (__bACnetChannelValueDouble BACnetChannelValueDouble, err error) {
	m.BACnetChannelValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueDouble"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueDouble")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doubleValue, err := ReadSimpleField[BACnetApplicationTagDouble](ctx, "doubleValue", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doubleValue' field"))
	}
	m.DoubleValue = doubleValue

	if closeErr := readBuffer.CloseContext("BACnetChannelValueDouble"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueDouble")
	}

	return m, nil
}

func (m *_BACnetChannelValueDouble) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueDouble) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueDouble"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueDouble")
		}

		if err := WriteSimpleField[BACnetApplicationTagDouble](ctx, "doubleValue", m.GetDoubleValue(), WriteComplex[BACnetApplicationTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'doubleValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetChannelValueDouble"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueDouble")
		}
		return nil
	}
	return m.BACnetChannelValueContract.(*_BACnetChannelValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetChannelValueDouble) IsBACnetChannelValueDouble() {}

func (m *_BACnetChannelValueDouble) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetChannelValueDouble) deepCopy() *_BACnetChannelValueDouble {
	if m == nil {
		return nil
	}
	_BACnetChannelValueDoubleCopy := &_BACnetChannelValueDouble{
		m.BACnetChannelValueContract.(*_BACnetChannelValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagDouble](m.DoubleValue),
	}
	m.BACnetChannelValueContract.(*_BACnetChannelValue)._SubType = m
	return _BACnetChannelValueDoubleCopy
}

func (m *_BACnetChannelValueDouble) String() string {
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
