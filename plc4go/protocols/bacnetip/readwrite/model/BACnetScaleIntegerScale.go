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

// BACnetScaleIntegerScale is the corresponding interface of BACnetScaleIntegerScale
type BACnetScaleIntegerScale interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetScale
	// GetIntegerScale returns IntegerScale (property field)
	GetIntegerScale() BACnetContextTagSignedInteger
	// IsBACnetScaleIntegerScale is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetScaleIntegerScale()
	// CreateBuilder creates a BACnetScaleIntegerScaleBuilder
	CreateBACnetScaleIntegerScaleBuilder() BACnetScaleIntegerScaleBuilder
}

// _BACnetScaleIntegerScale is the data-structure of this message
type _BACnetScaleIntegerScale struct {
	BACnetScaleContract
	IntegerScale BACnetContextTagSignedInteger
}

var _ BACnetScaleIntegerScale = (*_BACnetScaleIntegerScale)(nil)
var _ BACnetScaleRequirements = (*_BACnetScaleIntegerScale)(nil)

// NewBACnetScaleIntegerScale factory function for _BACnetScaleIntegerScale
func NewBACnetScaleIntegerScale(peekedTagHeader BACnetTagHeader, integerScale BACnetContextTagSignedInteger) *_BACnetScaleIntegerScale {
	if integerScale == nil {
		panic("integerScale of type BACnetContextTagSignedInteger for BACnetScaleIntegerScale must not be nil")
	}
	_result := &_BACnetScaleIntegerScale{
		BACnetScaleContract: NewBACnetScale(peekedTagHeader),
		IntegerScale:        integerScale,
	}
	_result.BACnetScaleContract.(*_BACnetScale)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetScaleIntegerScaleBuilder is a builder for BACnetScaleIntegerScale
type BACnetScaleIntegerScaleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(integerScale BACnetContextTagSignedInteger) BACnetScaleIntegerScaleBuilder
	// WithIntegerScale adds IntegerScale (property field)
	WithIntegerScale(BACnetContextTagSignedInteger) BACnetScaleIntegerScaleBuilder
	// WithIntegerScaleBuilder adds IntegerScale (property field) which is build by the builder
	WithIntegerScaleBuilder(func(BACnetContextTagSignedIntegerBuilder) BACnetContextTagSignedIntegerBuilder) BACnetScaleIntegerScaleBuilder
	// Build builds the BACnetScaleIntegerScale or returns an error if something is wrong
	Build() (BACnetScaleIntegerScale, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetScaleIntegerScale
}

// NewBACnetScaleIntegerScaleBuilder() creates a BACnetScaleIntegerScaleBuilder
func NewBACnetScaleIntegerScaleBuilder() BACnetScaleIntegerScaleBuilder {
	return &_BACnetScaleIntegerScaleBuilder{_BACnetScaleIntegerScale: new(_BACnetScaleIntegerScale)}
}

type _BACnetScaleIntegerScaleBuilder struct {
	*_BACnetScaleIntegerScale

	parentBuilder *_BACnetScaleBuilder

	err *utils.MultiError
}

var _ (BACnetScaleIntegerScaleBuilder) = (*_BACnetScaleIntegerScaleBuilder)(nil)

func (b *_BACnetScaleIntegerScaleBuilder) setParent(contract BACnetScaleContract) {
	b.BACnetScaleContract = contract
}

func (b *_BACnetScaleIntegerScaleBuilder) WithMandatoryFields(integerScale BACnetContextTagSignedInteger) BACnetScaleIntegerScaleBuilder {
	return b.WithIntegerScale(integerScale)
}

func (b *_BACnetScaleIntegerScaleBuilder) WithIntegerScale(integerScale BACnetContextTagSignedInteger) BACnetScaleIntegerScaleBuilder {
	b.IntegerScale = integerScale
	return b
}

func (b *_BACnetScaleIntegerScaleBuilder) WithIntegerScaleBuilder(builderSupplier func(BACnetContextTagSignedIntegerBuilder) BACnetContextTagSignedIntegerBuilder) BACnetScaleIntegerScaleBuilder {
	builder := builderSupplier(b.IntegerScale.CreateBACnetContextTagSignedIntegerBuilder())
	var err error
	b.IntegerScale, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetScaleIntegerScaleBuilder) Build() (BACnetScaleIntegerScale, error) {
	if b.IntegerScale == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'integerScale' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetScaleIntegerScale.deepCopy(), nil
}

func (b *_BACnetScaleIntegerScaleBuilder) MustBuild() BACnetScaleIntegerScale {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetScaleIntegerScaleBuilder) Done() BACnetScaleBuilder {
	return b.parentBuilder
}

func (b *_BACnetScaleIntegerScaleBuilder) buildForBACnetScale() (BACnetScale, error) {
	return b.Build()
}

func (b *_BACnetScaleIntegerScaleBuilder) DeepCopy() any {
	_copy := b.CreateBACnetScaleIntegerScaleBuilder().(*_BACnetScaleIntegerScaleBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetScaleIntegerScaleBuilder creates a BACnetScaleIntegerScaleBuilder
func (b *_BACnetScaleIntegerScale) CreateBACnetScaleIntegerScaleBuilder() BACnetScaleIntegerScaleBuilder {
	if b == nil {
		return NewBACnetScaleIntegerScaleBuilder()
	}
	return &_BACnetScaleIntegerScaleBuilder{_BACnetScaleIntegerScale: b.deepCopy()}
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

func (m *_BACnetScaleIntegerScale) GetParent() BACnetScaleContract {
	return m.BACnetScaleContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetScaleIntegerScale) GetIntegerScale() BACnetContextTagSignedInteger {
	return m.IntegerScale
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetScaleIntegerScale(structType any) BACnetScaleIntegerScale {
	if casted, ok := structType.(BACnetScaleIntegerScale); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetScaleIntegerScale); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetScaleIntegerScale) GetTypeName() string {
	return "BACnetScaleIntegerScale"
}

func (m *_BACnetScaleIntegerScale) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetScaleContract.(*_BACnetScale).getLengthInBits(ctx))

	// Simple field (integerScale)
	lengthInBits += m.IntegerScale.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetScaleIntegerScale) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetScaleIntegerScale) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetScale) (__bACnetScaleIntegerScale BACnetScaleIntegerScale, err error) {
	m.BACnetScaleContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetScaleIntegerScale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetScaleIntegerScale")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	integerScale, err := ReadSimpleField[BACnetContextTagSignedInteger](ctx, "integerScale", ReadComplex[BACnetContextTagSignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagSignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_SIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'integerScale' field"))
	}
	m.IntegerScale = integerScale

	if closeErr := readBuffer.CloseContext("BACnetScaleIntegerScale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetScaleIntegerScale")
	}

	return m, nil
}

func (m *_BACnetScaleIntegerScale) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetScaleIntegerScale) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetScaleIntegerScale"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetScaleIntegerScale")
		}

		if err := WriteSimpleField[BACnetContextTagSignedInteger](ctx, "integerScale", m.GetIntegerScale(), WriteComplex[BACnetContextTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'integerScale' field")
		}

		if popErr := writeBuffer.PopContext("BACnetScaleIntegerScale"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetScaleIntegerScale")
		}
		return nil
	}
	return m.BACnetScaleContract.(*_BACnetScale).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetScaleIntegerScale) IsBACnetScaleIntegerScale() {}

func (m *_BACnetScaleIntegerScale) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetScaleIntegerScale) deepCopy() *_BACnetScaleIntegerScale {
	if m == nil {
		return nil
	}
	_BACnetScaleIntegerScaleCopy := &_BACnetScaleIntegerScale{
		m.BACnetScaleContract.(*_BACnetScale).deepCopy(),
		utils.DeepCopy[BACnetContextTagSignedInteger](m.IntegerScale),
	}
	m.BACnetScaleContract.(*_BACnetScale)._SubType = m
	return _BACnetScaleIntegerScaleCopy
}

func (m *_BACnetScaleIntegerScale) String() string {
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
