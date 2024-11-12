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

// BACnetTimerStateChangeValueNull is the corresponding interface of BACnetTimerStateChangeValueNull
type BACnetTimerStateChangeValueNull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetTimerStateChangeValue
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetApplicationTagNull
	// IsBACnetTimerStateChangeValueNull is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetTimerStateChangeValueNull()
	// CreateBuilder creates a BACnetTimerStateChangeValueNullBuilder
	CreateBACnetTimerStateChangeValueNullBuilder() BACnetTimerStateChangeValueNullBuilder
}

// _BACnetTimerStateChangeValueNull is the data-structure of this message
type _BACnetTimerStateChangeValueNull struct {
	BACnetTimerStateChangeValueContract
	NullValue BACnetApplicationTagNull
}

var _ BACnetTimerStateChangeValueNull = (*_BACnetTimerStateChangeValueNull)(nil)
var _ BACnetTimerStateChangeValueRequirements = (*_BACnetTimerStateChangeValueNull)(nil)

// NewBACnetTimerStateChangeValueNull factory function for _BACnetTimerStateChangeValueNull
func NewBACnetTimerStateChangeValueNull(peekedTagHeader BACnetTagHeader, nullValue BACnetApplicationTagNull, objectTypeArgument BACnetObjectType) *_BACnetTimerStateChangeValueNull {
	if nullValue == nil {
		panic("nullValue of type BACnetApplicationTagNull for BACnetTimerStateChangeValueNull must not be nil")
	}
	_result := &_BACnetTimerStateChangeValueNull{
		BACnetTimerStateChangeValueContract: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
		NullValue:                           nullValue,
	}
	_result.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetTimerStateChangeValueNullBuilder is a builder for BACnetTimerStateChangeValueNull
type BACnetTimerStateChangeValueNullBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(nullValue BACnetApplicationTagNull) BACnetTimerStateChangeValueNullBuilder
	// WithNullValue adds NullValue (property field)
	WithNullValue(BACnetApplicationTagNull) BACnetTimerStateChangeValueNullBuilder
	// WithNullValueBuilder adds NullValue (property field) which is build by the builder
	WithNullValueBuilder(func(BACnetApplicationTagNullBuilder) BACnetApplicationTagNullBuilder) BACnetTimerStateChangeValueNullBuilder
	// Build builds the BACnetTimerStateChangeValueNull or returns an error if something is wrong
	Build() (BACnetTimerStateChangeValueNull, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetTimerStateChangeValueNull
}

// NewBACnetTimerStateChangeValueNullBuilder() creates a BACnetTimerStateChangeValueNullBuilder
func NewBACnetTimerStateChangeValueNullBuilder() BACnetTimerStateChangeValueNullBuilder {
	return &_BACnetTimerStateChangeValueNullBuilder{_BACnetTimerStateChangeValueNull: new(_BACnetTimerStateChangeValueNull)}
}

type _BACnetTimerStateChangeValueNullBuilder struct {
	*_BACnetTimerStateChangeValueNull

	parentBuilder *_BACnetTimerStateChangeValueBuilder

	err *utils.MultiError
}

var _ (BACnetTimerStateChangeValueNullBuilder) = (*_BACnetTimerStateChangeValueNullBuilder)(nil)

func (b *_BACnetTimerStateChangeValueNullBuilder) setParent(contract BACnetTimerStateChangeValueContract) {
	b.BACnetTimerStateChangeValueContract = contract
}

func (b *_BACnetTimerStateChangeValueNullBuilder) WithMandatoryFields(nullValue BACnetApplicationTagNull) BACnetTimerStateChangeValueNullBuilder {
	return b.WithNullValue(nullValue)
}

func (b *_BACnetTimerStateChangeValueNullBuilder) WithNullValue(nullValue BACnetApplicationTagNull) BACnetTimerStateChangeValueNullBuilder {
	b.NullValue = nullValue
	return b
}

func (b *_BACnetTimerStateChangeValueNullBuilder) WithNullValueBuilder(builderSupplier func(BACnetApplicationTagNullBuilder) BACnetApplicationTagNullBuilder) BACnetTimerStateChangeValueNullBuilder {
	builder := builderSupplier(b.NullValue.CreateBACnetApplicationTagNullBuilder())
	var err error
	b.NullValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagNullBuilder failed"))
	}
	return b
}

func (b *_BACnetTimerStateChangeValueNullBuilder) Build() (BACnetTimerStateChangeValueNull, error) {
	if b.NullValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'nullValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetTimerStateChangeValueNull.deepCopy(), nil
}

func (b *_BACnetTimerStateChangeValueNullBuilder) MustBuild() BACnetTimerStateChangeValueNull {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetTimerStateChangeValueNullBuilder) Done() BACnetTimerStateChangeValueBuilder {
	return b.parentBuilder
}

func (b *_BACnetTimerStateChangeValueNullBuilder) buildForBACnetTimerStateChangeValue() (BACnetTimerStateChangeValue, error) {
	return b.Build()
}

func (b *_BACnetTimerStateChangeValueNullBuilder) DeepCopy() any {
	_copy := b.CreateBACnetTimerStateChangeValueNullBuilder().(*_BACnetTimerStateChangeValueNullBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetTimerStateChangeValueNullBuilder creates a BACnetTimerStateChangeValueNullBuilder
func (b *_BACnetTimerStateChangeValueNull) CreateBACnetTimerStateChangeValueNullBuilder() BACnetTimerStateChangeValueNullBuilder {
	if b == nil {
		return NewBACnetTimerStateChangeValueNullBuilder()
	}
	return &_BACnetTimerStateChangeValueNullBuilder{_BACnetTimerStateChangeValueNull: b.deepCopy()}
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

func (m *_BACnetTimerStateChangeValueNull) GetParent() BACnetTimerStateChangeValueContract {
	return m.BACnetTimerStateChangeValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueNull) GetNullValue() BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueNull(structType any) BACnetTimerStateChangeValueNull {
	if casted, ok := structType.(BACnetTimerStateChangeValueNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueNull) GetTypeName() string {
	return "BACnetTimerStateChangeValueNull"
}

func (m *_BACnetTimerStateChangeValueNull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).getLengthInBits(ctx))

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetTimerStateChangeValueNull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetTimerStateChangeValueNull) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetTimerStateChangeValue, objectTypeArgument BACnetObjectType) (__bACnetTimerStateChangeValueNull BACnetTimerStateChangeValueNull, err error) {
	m.BACnetTimerStateChangeValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nullValue, err := ReadSimpleField[BACnetApplicationTagNull](ctx, "nullValue", ReadComplex[BACnetApplicationTagNull](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagNull](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nullValue' field"))
	}
	m.NullValue = nullValue

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueNull")
	}

	return m, nil
}

func (m *_BACnetTimerStateChangeValueNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueNull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueNull")
		}

		if err := WriteSimpleField[BACnetApplicationTagNull](ctx, "nullValue", m.GetNullValue(), WriteComplex[BACnetApplicationTagNull](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueNull")
		}
		return nil
	}
	return m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetTimerStateChangeValueNull) IsBACnetTimerStateChangeValueNull() {}

func (m *_BACnetTimerStateChangeValueNull) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetTimerStateChangeValueNull) deepCopy() *_BACnetTimerStateChangeValueNull {
	if m == nil {
		return nil
	}
	_BACnetTimerStateChangeValueNullCopy := &_BACnetTimerStateChangeValueNull{
		m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagNull](m.NullValue),
	}
	m.BACnetTimerStateChangeValueContract.(*_BACnetTimerStateChangeValue)._SubType = m
	return _BACnetTimerStateChangeValueNullCopy
}

func (m *_BACnetTimerStateChangeValueNull) String() string {
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
