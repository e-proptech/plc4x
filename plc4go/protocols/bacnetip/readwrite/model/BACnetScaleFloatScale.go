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

// BACnetScaleFloatScale is the corresponding interface of BACnetScaleFloatScale
type BACnetScaleFloatScale interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetScale
	// GetFloatScale returns FloatScale (property field)
	GetFloatScale() BACnetContextTagReal
	// IsBACnetScaleFloatScale is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetScaleFloatScale()
	// CreateBuilder creates a BACnetScaleFloatScaleBuilder
	CreateBACnetScaleFloatScaleBuilder() BACnetScaleFloatScaleBuilder
}

// _BACnetScaleFloatScale is the data-structure of this message
type _BACnetScaleFloatScale struct {
	BACnetScaleContract
	FloatScale BACnetContextTagReal
}

var _ BACnetScaleFloatScale = (*_BACnetScaleFloatScale)(nil)
var _ BACnetScaleRequirements = (*_BACnetScaleFloatScale)(nil)

// NewBACnetScaleFloatScale factory function for _BACnetScaleFloatScale
func NewBACnetScaleFloatScale(peekedTagHeader BACnetTagHeader, floatScale BACnetContextTagReal) *_BACnetScaleFloatScale {
	if floatScale == nil {
		panic("floatScale of type BACnetContextTagReal for BACnetScaleFloatScale must not be nil")
	}
	_result := &_BACnetScaleFloatScale{
		BACnetScaleContract: NewBACnetScale(peekedTagHeader),
		FloatScale:          floatScale,
	}
	_result.BACnetScaleContract.(*_BACnetScale)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetScaleFloatScaleBuilder is a builder for BACnetScaleFloatScale
type BACnetScaleFloatScaleBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(floatScale BACnetContextTagReal) BACnetScaleFloatScaleBuilder
	// WithFloatScale adds FloatScale (property field)
	WithFloatScale(BACnetContextTagReal) BACnetScaleFloatScaleBuilder
	// WithFloatScaleBuilder adds FloatScale (property field) which is build by the builder
	WithFloatScaleBuilder(func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetScaleFloatScaleBuilder
	// Build builds the BACnetScaleFloatScale or returns an error if something is wrong
	Build() (BACnetScaleFloatScale, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetScaleFloatScale
}

// NewBACnetScaleFloatScaleBuilder() creates a BACnetScaleFloatScaleBuilder
func NewBACnetScaleFloatScaleBuilder() BACnetScaleFloatScaleBuilder {
	return &_BACnetScaleFloatScaleBuilder{_BACnetScaleFloatScale: new(_BACnetScaleFloatScale)}
}

type _BACnetScaleFloatScaleBuilder struct {
	*_BACnetScaleFloatScale

	parentBuilder *_BACnetScaleBuilder

	err *utils.MultiError
}

var _ (BACnetScaleFloatScaleBuilder) = (*_BACnetScaleFloatScaleBuilder)(nil)

func (b *_BACnetScaleFloatScaleBuilder) setParent(contract BACnetScaleContract) {
	b.BACnetScaleContract = contract
}

func (b *_BACnetScaleFloatScaleBuilder) WithMandatoryFields(floatScale BACnetContextTagReal) BACnetScaleFloatScaleBuilder {
	return b.WithFloatScale(floatScale)
}

func (b *_BACnetScaleFloatScaleBuilder) WithFloatScale(floatScale BACnetContextTagReal) BACnetScaleFloatScaleBuilder {
	b.FloatScale = floatScale
	return b
}

func (b *_BACnetScaleFloatScaleBuilder) WithFloatScaleBuilder(builderSupplier func(BACnetContextTagRealBuilder) BACnetContextTagRealBuilder) BACnetScaleFloatScaleBuilder {
	builder := builderSupplier(b.FloatScale.CreateBACnetContextTagRealBuilder())
	var err error
	b.FloatScale, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagRealBuilder failed"))
	}
	return b
}

func (b *_BACnetScaleFloatScaleBuilder) Build() (BACnetScaleFloatScale, error) {
	if b.FloatScale == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'floatScale' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetScaleFloatScale.deepCopy(), nil
}

func (b *_BACnetScaleFloatScaleBuilder) MustBuild() BACnetScaleFloatScale {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetScaleFloatScaleBuilder) Done() BACnetScaleBuilder {
	return b.parentBuilder
}

func (b *_BACnetScaleFloatScaleBuilder) buildForBACnetScale() (BACnetScale, error) {
	return b.Build()
}

func (b *_BACnetScaleFloatScaleBuilder) DeepCopy() any {
	_copy := b.CreateBACnetScaleFloatScaleBuilder().(*_BACnetScaleFloatScaleBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetScaleFloatScaleBuilder creates a BACnetScaleFloatScaleBuilder
func (b *_BACnetScaleFloatScale) CreateBACnetScaleFloatScaleBuilder() BACnetScaleFloatScaleBuilder {
	if b == nil {
		return NewBACnetScaleFloatScaleBuilder()
	}
	return &_BACnetScaleFloatScaleBuilder{_BACnetScaleFloatScale: b.deepCopy()}
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

func (m *_BACnetScaleFloatScale) GetParent() BACnetScaleContract {
	return m.BACnetScaleContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetScaleFloatScale) GetFloatScale() BACnetContextTagReal {
	return m.FloatScale
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetScaleFloatScale(structType any) BACnetScaleFloatScale {
	if casted, ok := structType.(BACnetScaleFloatScale); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetScaleFloatScale); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetScaleFloatScale) GetTypeName() string {
	return "BACnetScaleFloatScale"
}

func (m *_BACnetScaleFloatScale) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetScaleContract.(*_BACnetScale).getLengthInBits(ctx))

	// Simple field (floatScale)
	lengthInBits += m.FloatScale.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetScaleFloatScale) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetScaleFloatScale) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetScale) (__bACnetScaleFloatScale BACnetScaleFloatScale, err error) {
	m.BACnetScaleContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetScaleFloatScale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetScaleFloatScale")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	floatScale, err := ReadSimpleField[BACnetContextTagReal](ctx, "floatScale", ReadComplex[BACnetContextTagReal](BACnetContextTagParseWithBufferProducer[BACnetContextTagReal]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_REAL)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'floatScale' field"))
	}
	m.FloatScale = floatScale

	if closeErr := readBuffer.CloseContext("BACnetScaleFloatScale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetScaleFloatScale")
	}

	return m, nil
}

func (m *_BACnetScaleFloatScale) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetScaleFloatScale) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetScaleFloatScale"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetScaleFloatScale")
		}

		if err := WriteSimpleField[BACnetContextTagReal](ctx, "floatScale", m.GetFloatScale(), WriteComplex[BACnetContextTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'floatScale' field")
		}

		if popErr := writeBuffer.PopContext("BACnetScaleFloatScale"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetScaleFloatScale")
		}
		return nil
	}
	return m.BACnetScaleContract.(*_BACnetScale).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetScaleFloatScale) IsBACnetScaleFloatScale() {}

func (m *_BACnetScaleFloatScale) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetScaleFloatScale) deepCopy() *_BACnetScaleFloatScale {
	if m == nil {
		return nil
	}
	_BACnetScaleFloatScaleCopy := &_BACnetScaleFloatScale{
		m.BACnetScaleContract.(*_BACnetScale).deepCopy(),
		utils.DeepCopy[BACnetContextTagReal](m.FloatScale),
	}
	m.BACnetScaleContract.(*_BACnetScale)._SubType = m
	return _BACnetScaleFloatScaleCopy
}

func (m *_BACnetScaleFloatScale) String() string {
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
