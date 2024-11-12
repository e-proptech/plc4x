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

// BACnetPriorityValueDouble is the corresponding interface of BACnetPriorityValueDouble
type BACnetPriorityValueDouble interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPriorityValue
	// GetDoubleValue returns DoubleValue (property field)
	GetDoubleValue() BACnetApplicationTagDouble
	// IsBACnetPriorityValueDouble is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPriorityValueDouble()
	// CreateBuilder creates a BACnetPriorityValueDoubleBuilder
	CreateBACnetPriorityValueDoubleBuilder() BACnetPriorityValueDoubleBuilder
}

// _BACnetPriorityValueDouble is the data-structure of this message
type _BACnetPriorityValueDouble struct {
	BACnetPriorityValueContract
	DoubleValue BACnetApplicationTagDouble
}

var _ BACnetPriorityValueDouble = (*_BACnetPriorityValueDouble)(nil)
var _ BACnetPriorityValueRequirements = (*_BACnetPriorityValueDouble)(nil)

// NewBACnetPriorityValueDouble factory function for _BACnetPriorityValueDouble
func NewBACnetPriorityValueDouble(peekedTagHeader BACnetTagHeader, doubleValue BACnetApplicationTagDouble, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueDouble {
	if doubleValue == nil {
		panic("doubleValue of type BACnetApplicationTagDouble for BACnetPriorityValueDouble must not be nil")
	}
	_result := &_BACnetPriorityValueDouble{
		BACnetPriorityValueContract: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
		DoubleValue:                 doubleValue,
	}
	_result.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPriorityValueDoubleBuilder is a builder for BACnetPriorityValueDouble
type BACnetPriorityValueDoubleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(doubleValue BACnetApplicationTagDouble) BACnetPriorityValueDoubleBuilder
	// WithDoubleValue adds DoubleValue (property field)
	WithDoubleValue(BACnetApplicationTagDouble) BACnetPriorityValueDoubleBuilder
	// WithDoubleValueBuilder adds DoubleValue (property field) which is build by the builder
	WithDoubleValueBuilder(func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetPriorityValueDoubleBuilder
	// Build builds the BACnetPriorityValueDouble or returns an error if something is wrong
	Build() (BACnetPriorityValueDouble, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPriorityValueDouble
}

// NewBACnetPriorityValueDoubleBuilder() creates a BACnetPriorityValueDoubleBuilder
func NewBACnetPriorityValueDoubleBuilder() BACnetPriorityValueDoubleBuilder {
	return &_BACnetPriorityValueDoubleBuilder{_BACnetPriorityValueDouble: new(_BACnetPriorityValueDouble)}
}

type _BACnetPriorityValueDoubleBuilder struct {
	*_BACnetPriorityValueDouble

	parentBuilder *_BACnetPriorityValueBuilder

	err *utils.MultiError
}

var _ (BACnetPriorityValueDoubleBuilder) = (*_BACnetPriorityValueDoubleBuilder)(nil)

func (b *_BACnetPriorityValueDoubleBuilder) setParent(contract BACnetPriorityValueContract) {
	b.BACnetPriorityValueContract = contract
}

func (b *_BACnetPriorityValueDoubleBuilder) WithMandatoryFields(doubleValue BACnetApplicationTagDouble) BACnetPriorityValueDoubleBuilder {
	return b.WithDoubleValue(doubleValue)
}

func (b *_BACnetPriorityValueDoubleBuilder) WithDoubleValue(doubleValue BACnetApplicationTagDouble) BACnetPriorityValueDoubleBuilder {
	b.DoubleValue = doubleValue
	return b
}

func (b *_BACnetPriorityValueDoubleBuilder) WithDoubleValueBuilder(builderSupplier func(BACnetApplicationTagDoubleBuilder) BACnetApplicationTagDoubleBuilder) BACnetPriorityValueDoubleBuilder {
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

func (b *_BACnetPriorityValueDoubleBuilder) Build() (BACnetPriorityValueDouble, error) {
	if b.DoubleValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'doubleValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPriorityValueDouble.deepCopy(), nil
}

func (b *_BACnetPriorityValueDoubleBuilder) MustBuild() BACnetPriorityValueDouble {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetPriorityValueDoubleBuilder) Done() BACnetPriorityValueBuilder {
	return b.parentBuilder
}

func (b *_BACnetPriorityValueDoubleBuilder) buildForBACnetPriorityValue() (BACnetPriorityValue, error) {
	return b.Build()
}

func (b *_BACnetPriorityValueDoubleBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPriorityValueDoubleBuilder().(*_BACnetPriorityValueDoubleBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPriorityValueDoubleBuilder creates a BACnetPriorityValueDoubleBuilder
func (b *_BACnetPriorityValueDouble) CreateBACnetPriorityValueDoubleBuilder() BACnetPriorityValueDoubleBuilder {
	if b == nil {
		return NewBACnetPriorityValueDoubleBuilder()
	}
	return &_BACnetPriorityValueDoubleBuilder{_BACnetPriorityValueDouble: b.deepCopy()}
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

func (m *_BACnetPriorityValueDouble) GetParent() BACnetPriorityValueContract {
	return m.BACnetPriorityValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueDouble) GetDoubleValue() BACnetApplicationTagDouble {
	return m.DoubleValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueDouble(structType any) BACnetPriorityValueDouble {
	if casted, ok := structType.(BACnetPriorityValueDouble); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueDouble); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueDouble) GetTypeName() string {
	return "BACnetPriorityValueDouble"
}

func (m *_BACnetPriorityValueDouble) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPriorityValueContract.(*_BACnetPriorityValue).getLengthInBits(ctx))

	// Simple field (doubleValue)
	lengthInBits += m.DoubleValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPriorityValueDouble) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPriorityValueDouble) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPriorityValue, objectTypeArgument BACnetObjectType) (__bACnetPriorityValueDouble BACnetPriorityValueDouble, err error) {
	m.BACnetPriorityValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueDouble"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueDouble")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	doubleValue, err := ReadSimpleField[BACnetApplicationTagDouble](ctx, "doubleValue", ReadComplex[BACnetApplicationTagDouble](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDouble](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'doubleValue' field"))
	}
	m.DoubleValue = doubleValue

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueDouble"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueDouble")
	}

	return m, nil
}

func (m *_BACnetPriorityValueDouble) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueDouble) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueDouble"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueDouble")
		}

		if err := WriteSimpleField[BACnetApplicationTagDouble](ctx, "doubleValue", m.GetDoubleValue(), WriteComplex[BACnetApplicationTagDouble](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'doubleValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueDouble"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueDouble")
		}
		return nil
	}
	return m.BACnetPriorityValueContract.(*_BACnetPriorityValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueDouble) IsBACnetPriorityValueDouble() {}

func (m *_BACnetPriorityValueDouble) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPriorityValueDouble) deepCopy() *_BACnetPriorityValueDouble {
	if m == nil {
		return nil
	}
	_BACnetPriorityValueDoubleCopy := &_BACnetPriorityValueDouble{
		m.BACnetPriorityValueContract.(*_BACnetPriorityValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagDouble](m.DoubleValue),
	}
	m.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = m
	return _BACnetPriorityValueDoubleCopy
}

func (m *_BACnetPriorityValueDouble) String() string {
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
