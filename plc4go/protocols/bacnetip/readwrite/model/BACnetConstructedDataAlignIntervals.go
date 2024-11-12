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

// BACnetConstructedDataAlignIntervals is the corresponding interface of BACnetConstructedDataAlignIntervals
type BACnetConstructedDataAlignIntervals interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetAlignIntervals returns AlignIntervals (property field)
	GetAlignIntervals() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataAlignIntervals is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAlignIntervals()
	// CreateBuilder creates a BACnetConstructedDataAlignIntervalsBuilder
	CreateBACnetConstructedDataAlignIntervalsBuilder() BACnetConstructedDataAlignIntervalsBuilder
}

// _BACnetConstructedDataAlignIntervals is the data-structure of this message
type _BACnetConstructedDataAlignIntervals struct {
	BACnetConstructedDataContract
	AlignIntervals BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataAlignIntervals = (*_BACnetConstructedDataAlignIntervals)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAlignIntervals)(nil)

// NewBACnetConstructedDataAlignIntervals factory function for _BACnetConstructedDataAlignIntervals
func NewBACnetConstructedDataAlignIntervals(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, alignIntervals BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAlignIntervals {
	if alignIntervals == nil {
		panic("alignIntervals of type BACnetApplicationTagBoolean for BACnetConstructedDataAlignIntervals must not be nil")
	}
	_result := &_BACnetConstructedDataAlignIntervals{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AlignIntervals:                alignIntervals,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataAlignIntervalsBuilder is a builder for BACnetConstructedDataAlignIntervals
type BACnetConstructedDataAlignIntervalsBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(alignIntervals BACnetApplicationTagBoolean) BACnetConstructedDataAlignIntervalsBuilder
	// WithAlignIntervals adds AlignIntervals (property field)
	WithAlignIntervals(BACnetApplicationTagBoolean) BACnetConstructedDataAlignIntervalsBuilder
	// WithAlignIntervalsBuilder adds AlignIntervals (property field) which is build by the builder
	WithAlignIntervalsBuilder(func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataAlignIntervalsBuilder
	// Build builds the BACnetConstructedDataAlignIntervals or returns an error if something is wrong
	Build() (BACnetConstructedDataAlignIntervals, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataAlignIntervals
}

// NewBACnetConstructedDataAlignIntervalsBuilder() creates a BACnetConstructedDataAlignIntervalsBuilder
func NewBACnetConstructedDataAlignIntervalsBuilder() BACnetConstructedDataAlignIntervalsBuilder {
	return &_BACnetConstructedDataAlignIntervalsBuilder{_BACnetConstructedDataAlignIntervals: new(_BACnetConstructedDataAlignIntervals)}
}

type _BACnetConstructedDataAlignIntervalsBuilder struct {
	*_BACnetConstructedDataAlignIntervals

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataAlignIntervalsBuilder) = (*_BACnetConstructedDataAlignIntervalsBuilder)(nil)

func (b *_BACnetConstructedDataAlignIntervalsBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataAlignIntervalsBuilder) WithMandatoryFields(alignIntervals BACnetApplicationTagBoolean) BACnetConstructedDataAlignIntervalsBuilder {
	return b.WithAlignIntervals(alignIntervals)
}

func (b *_BACnetConstructedDataAlignIntervalsBuilder) WithAlignIntervals(alignIntervals BACnetApplicationTagBoolean) BACnetConstructedDataAlignIntervalsBuilder {
	b.AlignIntervals = alignIntervals
	return b
}

func (b *_BACnetConstructedDataAlignIntervalsBuilder) WithAlignIntervalsBuilder(builderSupplier func(BACnetApplicationTagBooleanBuilder) BACnetApplicationTagBooleanBuilder) BACnetConstructedDataAlignIntervalsBuilder {
	builder := builderSupplier(b.AlignIntervals.CreateBACnetApplicationTagBooleanBuilder())
	var err error
	b.AlignIntervals, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagBooleanBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataAlignIntervalsBuilder) Build() (BACnetConstructedDataAlignIntervals, error) {
	if b.AlignIntervals == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'alignIntervals' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataAlignIntervals.deepCopy(), nil
}

func (b *_BACnetConstructedDataAlignIntervalsBuilder) MustBuild() BACnetConstructedDataAlignIntervals {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataAlignIntervalsBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataAlignIntervalsBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataAlignIntervalsBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataAlignIntervalsBuilder().(*_BACnetConstructedDataAlignIntervalsBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataAlignIntervalsBuilder creates a BACnetConstructedDataAlignIntervalsBuilder
func (b *_BACnetConstructedDataAlignIntervals) CreateBACnetConstructedDataAlignIntervalsBuilder() BACnetConstructedDataAlignIntervalsBuilder {
	if b == nil {
		return NewBACnetConstructedDataAlignIntervalsBuilder()
	}
	return &_BACnetConstructedDataAlignIntervalsBuilder{_BACnetConstructedDataAlignIntervals: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAlignIntervals) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAlignIntervals) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALIGN_INTERVALS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAlignIntervals) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAlignIntervals) GetAlignIntervals() BACnetApplicationTagBoolean {
	return m.AlignIntervals
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAlignIntervals) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetAlignIntervals())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAlignIntervals(structType any) BACnetConstructedDataAlignIntervals {
	if casted, ok := structType.(BACnetConstructedDataAlignIntervals); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAlignIntervals); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAlignIntervals) GetTypeName() string {
	return "BACnetConstructedDataAlignIntervals"
}

func (m *_BACnetConstructedDataAlignIntervals) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (alignIntervals)
	lengthInBits += m.AlignIntervals.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAlignIntervals) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAlignIntervals) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAlignIntervals BACnetConstructedDataAlignIntervals, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAlignIntervals"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAlignIntervals")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	alignIntervals, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "alignIntervals", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alignIntervals' field"))
	}
	m.AlignIntervals = alignIntervals

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), alignIntervals)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAlignIntervals"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAlignIntervals")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAlignIntervals) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAlignIntervals) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAlignIntervals"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAlignIntervals")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "alignIntervals", m.GetAlignIntervals(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'alignIntervals' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAlignIntervals"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAlignIntervals")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAlignIntervals) IsBACnetConstructedDataAlignIntervals() {}

func (m *_BACnetConstructedDataAlignIntervals) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataAlignIntervals) deepCopy() *_BACnetConstructedDataAlignIntervals {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataAlignIntervalsCopy := &_BACnetConstructedDataAlignIntervals{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagBoolean](m.AlignIntervals),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataAlignIntervalsCopy
}

func (m *_BACnetConstructedDataAlignIntervals) String() string {
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
