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

// BACnetPriorityValueNull is the corresponding interface of BACnetPriorityValueNull
type BACnetPriorityValueNull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetPriorityValue
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetApplicationTagNull
	// IsBACnetPriorityValueNull is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPriorityValueNull()
	// CreateBuilder creates a BACnetPriorityValueNullBuilder
	CreateBACnetPriorityValueNullBuilder() BACnetPriorityValueNullBuilder
}

// _BACnetPriorityValueNull is the data-structure of this message
type _BACnetPriorityValueNull struct {
	BACnetPriorityValueContract
	NullValue BACnetApplicationTagNull
}

var _ BACnetPriorityValueNull = (*_BACnetPriorityValueNull)(nil)
var _ BACnetPriorityValueRequirements = (*_BACnetPriorityValueNull)(nil)

// NewBACnetPriorityValueNull factory function for _BACnetPriorityValueNull
func NewBACnetPriorityValueNull(peekedTagHeader BACnetTagHeader, nullValue BACnetApplicationTagNull, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueNull {
	if nullValue == nil {
		panic("nullValue of type BACnetApplicationTagNull for BACnetPriorityValueNull must not be nil")
	}
	_result := &_BACnetPriorityValueNull{
		BACnetPriorityValueContract: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
		NullValue:                   nullValue,
	}
	_result.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetPriorityValueNullBuilder is a builder for BACnetPriorityValueNull
type BACnetPriorityValueNullBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(nullValue BACnetApplicationTagNull) BACnetPriorityValueNullBuilder
	// WithNullValue adds NullValue (property field)
	WithNullValue(BACnetApplicationTagNull) BACnetPriorityValueNullBuilder
	// WithNullValueBuilder adds NullValue (property field) which is build by the builder
	WithNullValueBuilder(func(BACnetApplicationTagNullBuilder) BACnetApplicationTagNullBuilder) BACnetPriorityValueNullBuilder
	// Build builds the BACnetPriorityValueNull or returns an error if something is wrong
	Build() (BACnetPriorityValueNull, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetPriorityValueNull
}

// NewBACnetPriorityValueNullBuilder() creates a BACnetPriorityValueNullBuilder
func NewBACnetPriorityValueNullBuilder() BACnetPriorityValueNullBuilder {
	return &_BACnetPriorityValueNullBuilder{_BACnetPriorityValueNull: new(_BACnetPriorityValueNull)}
}

type _BACnetPriorityValueNullBuilder struct {
	*_BACnetPriorityValueNull

	parentBuilder *_BACnetPriorityValueBuilder

	err *utils.MultiError
}

var _ (BACnetPriorityValueNullBuilder) = (*_BACnetPriorityValueNullBuilder)(nil)

func (b *_BACnetPriorityValueNullBuilder) setParent(contract BACnetPriorityValueContract) {
	b.BACnetPriorityValueContract = contract
}

func (b *_BACnetPriorityValueNullBuilder) WithMandatoryFields(nullValue BACnetApplicationTagNull) BACnetPriorityValueNullBuilder {
	return b.WithNullValue(nullValue)
}

func (b *_BACnetPriorityValueNullBuilder) WithNullValue(nullValue BACnetApplicationTagNull) BACnetPriorityValueNullBuilder {
	b.NullValue = nullValue
	return b
}

func (b *_BACnetPriorityValueNullBuilder) WithNullValueBuilder(builderSupplier func(BACnetApplicationTagNullBuilder) BACnetApplicationTagNullBuilder) BACnetPriorityValueNullBuilder {
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

func (b *_BACnetPriorityValueNullBuilder) Build() (BACnetPriorityValueNull, error) {
	if b.NullValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'nullValue' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetPriorityValueNull.deepCopy(), nil
}

func (b *_BACnetPriorityValueNullBuilder) MustBuild() BACnetPriorityValueNull {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetPriorityValueNullBuilder) Done() BACnetPriorityValueBuilder {
	return b.parentBuilder
}

func (b *_BACnetPriorityValueNullBuilder) buildForBACnetPriorityValue() (BACnetPriorityValue, error) {
	return b.Build()
}

func (b *_BACnetPriorityValueNullBuilder) DeepCopy() any {
	_copy := b.CreateBACnetPriorityValueNullBuilder().(*_BACnetPriorityValueNullBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetPriorityValueNullBuilder creates a BACnetPriorityValueNullBuilder
func (b *_BACnetPriorityValueNull) CreateBACnetPriorityValueNullBuilder() BACnetPriorityValueNullBuilder {
	if b == nil {
		return NewBACnetPriorityValueNullBuilder()
	}
	return &_BACnetPriorityValueNullBuilder{_BACnetPriorityValueNull: b.deepCopy()}
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

func (m *_BACnetPriorityValueNull) GetParent() BACnetPriorityValueContract {
	return m.BACnetPriorityValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueNull) GetNullValue() BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueNull(structType any) BACnetPriorityValueNull {
	if casted, ok := structType.(BACnetPriorityValueNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueNull) GetTypeName() string {
	return "BACnetPriorityValueNull"
}

func (m *_BACnetPriorityValueNull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPriorityValueContract.(*_BACnetPriorityValue).getLengthInBits(ctx))

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPriorityValueNull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPriorityValueNull) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPriorityValue, objectTypeArgument BACnetObjectType) (__bACnetPriorityValueNull BACnetPriorityValueNull, err error) {
	m.BACnetPriorityValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	nullValue, err := ReadSimpleField[BACnetApplicationTagNull](ctx, "nullValue", ReadComplex[BACnetApplicationTagNull](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagNull](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'nullValue' field"))
	}
	m.NullValue = nullValue

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueNull")
	}

	return m, nil
}

func (m *_BACnetPriorityValueNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueNull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueNull")
		}

		if err := WriteSimpleField[BACnetApplicationTagNull](ctx, "nullValue", m.GetNullValue(), WriteComplex[BACnetApplicationTagNull](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'nullValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueNull")
		}
		return nil
	}
	return m.BACnetPriorityValueContract.(*_BACnetPriorityValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueNull) IsBACnetPriorityValueNull() {}

func (m *_BACnetPriorityValueNull) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetPriorityValueNull) deepCopy() *_BACnetPriorityValueNull {
	if m == nil {
		return nil
	}
	_BACnetPriorityValueNullCopy := &_BACnetPriorityValueNull{
		m.BACnetPriorityValueContract.(*_BACnetPriorityValue).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagNull](m.NullValue),
	}
	m.BACnetPriorityValueContract.(*_BACnetPriorityValue)._SubType = m
	return _BACnetPriorityValueNullCopy
}

func (m *_BACnetPriorityValueNull) String() string {
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
