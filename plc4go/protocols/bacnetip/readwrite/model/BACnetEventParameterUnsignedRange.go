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

// BACnetEventParameterUnsignedRange is the corresponding interface of BACnetEventParameterUnsignedRange
type BACnetEventParameterUnsignedRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetLowLimit returns LowLimit (property field)
	GetLowLimit() BACnetContextTagUnsignedInteger
	// GetHighLimit returns HighLimit (property field)
	GetHighLimit() BACnetContextTagUnsignedInteger
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// IsBACnetEventParameterUnsignedRange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventParameterUnsignedRange()
	// CreateBuilder creates a BACnetEventParameterUnsignedRangeBuilder
	CreateBACnetEventParameterUnsignedRangeBuilder() BACnetEventParameterUnsignedRangeBuilder
}

// _BACnetEventParameterUnsignedRange is the data-structure of this message
type _BACnetEventParameterUnsignedRange struct {
	BACnetEventParameterContract
	OpeningTag BACnetOpeningTag
	TimeDelay  BACnetContextTagUnsignedInteger
	LowLimit   BACnetContextTagUnsignedInteger
	HighLimit  BACnetContextTagUnsignedInteger
	ClosingTag BACnetClosingTag
}

var _ BACnetEventParameterUnsignedRange = (*_BACnetEventParameterUnsignedRange)(nil)
var _ BACnetEventParameterRequirements = (*_BACnetEventParameterUnsignedRange)(nil)

// NewBACnetEventParameterUnsignedRange factory function for _BACnetEventParameterUnsignedRange
func NewBACnetEventParameterUnsignedRange(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, lowLimit BACnetContextTagUnsignedInteger, highLimit BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag) *_BACnetEventParameterUnsignedRange {
	if openingTag == nil {
		panic("openingTag of type BACnetOpeningTag for BACnetEventParameterUnsignedRange must not be nil")
	}
	if timeDelay == nil {
		panic("timeDelay of type BACnetContextTagUnsignedInteger for BACnetEventParameterUnsignedRange must not be nil")
	}
	if lowLimit == nil {
		panic("lowLimit of type BACnetContextTagUnsignedInteger for BACnetEventParameterUnsignedRange must not be nil")
	}
	if highLimit == nil {
		panic("highLimit of type BACnetContextTagUnsignedInteger for BACnetEventParameterUnsignedRange must not be nil")
	}
	if closingTag == nil {
		panic("closingTag of type BACnetClosingTag for BACnetEventParameterUnsignedRange must not be nil")
	}
	_result := &_BACnetEventParameterUnsignedRange{
		BACnetEventParameterContract: NewBACnetEventParameter(peekedTagHeader),
		OpeningTag:                   openingTag,
		TimeDelay:                    timeDelay,
		LowLimit:                     lowLimit,
		HighLimit:                    highLimit,
		ClosingTag:                   closingTag,
	}
	_result.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventParameterUnsignedRangeBuilder is a builder for BACnetEventParameterUnsignedRange
type BACnetEventParameterUnsignedRangeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, lowLimit BACnetContextTagUnsignedInteger, highLimit BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag) BACnetEventParameterUnsignedRangeBuilder
	// WithOpeningTag adds OpeningTag (property field)
	WithOpeningTag(BACnetOpeningTag) BACnetEventParameterUnsignedRangeBuilder
	// WithOpeningTagBuilder adds OpeningTag (property field) which is build by the builder
	WithOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterUnsignedRangeBuilder
	// WithTimeDelay adds TimeDelay (property field)
	WithTimeDelay(BACnetContextTagUnsignedInteger) BACnetEventParameterUnsignedRangeBuilder
	// WithTimeDelayBuilder adds TimeDelay (property field) which is build by the builder
	WithTimeDelayBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterUnsignedRangeBuilder
	// WithLowLimit adds LowLimit (property field)
	WithLowLimit(BACnetContextTagUnsignedInteger) BACnetEventParameterUnsignedRangeBuilder
	// WithLowLimitBuilder adds LowLimit (property field) which is build by the builder
	WithLowLimitBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterUnsignedRangeBuilder
	// WithHighLimit adds HighLimit (property field)
	WithHighLimit(BACnetContextTagUnsignedInteger) BACnetEventParameterUnsignedRangeBuilder
	// WithHighLimitBuilder adds HighLimit (property field) which is build by the builder
	WithHighLimitBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterUnsignedRangeBuilder
	// WithClosingTag adds ClosingTag (property field)
	WithClosingTag(BACnetClosingTag) BACnetEventParameterUnsignedRangeBuilder
	// WithClosingTagBuilder adds ClosingTag (property field) which is build by the builder
	WithClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterUnsignedRangeBuilder
	// Build builds the BACnetEventParameterUnsignedRange or returns an error if something is wrong
	Build() (BACnetEventParameterUnsignedRange, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventParameterUnsignedRange
}

// NewBACnetEventParameterUnsignedRangeBuilder() creates a BACnetEventParameterUnsignedRangeBuilder
func NewBACnetEventParameterUnsignedRangeBuilder() BACnetEventParameterUnsignedRangeBuilder {
	return &_BACnetEventParameterUnsignedRangeBuilder{_BACnetEventParameterUnsignedRange: new(_BACnetEventParameterUnsignedRange)}
}

type _BACnetEventParameterUnsignedRangeBuilder struct {
	*_BACnetEventParameterUnsignedRange

	parentBuilder *_BACnetEventParameterBuilder

	err *utils.MultiError
}

var _ (BACnetEventParameterUnsignedRangeBuilder) = (*_BACnetEventParameterUnsignedRangeBuilder)(nil)

func (b *_BACnetEventParameterUnsignedRangeBuilder) setParent(contract BACnetEventParameterContract) {
	b.BACnetEventParameterContract = contract
}

func (b *_BACnetEventParameterUnsignedRangeBuilder) WithMandatoryFields(openingTag BACnetOpeningTag, timeDelay BACnetContextTagUnsignedInteger, lowLimit BACnetContextTagUnsignedInteger, highLimit BACnetContextTagUnsignedInteger, closingTag BACnetClosingTag) BACnetEventParameterUnsignedRangeBuilder {
	return b.WithOpeningTag(openingTag).WithTimeDelay(timeDelay).WithLowLimit(lowLimit).WithHighLimit(highLimit).WithClosingTag(closingTag)
}

func (b *_BACnetEventParameterUnsignedRangeBuilder) WithOpeningTag(openingTag BACnetOpeningTag) BACnetEventParameterUnsignedRangeBuilder {
	b.OpeningTag = openingTag
	return b
}

func (b *_BACnetEventParameterUnsignedRangeBuilder) WithOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetEventParameterUnsignedRangeBuilder {
	builder := builderSupplier(b.OpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.OpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterUnsignedRangeBuilder) WithTimeDelay(timeDelay BACnetContextTagUnsignedInteger) BACnetEventParameterUnsignedRangeBuilder {
	b.TimeDelay = timeDelay
	return b
}

func (b *_BACnetEventParameterUnsignedRangeBuilder) WithTimeDelayBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterUnsignedRangeBuilder {
	builder := builderSupplier(b.TimeDelay.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.TimeDelay, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterUnsignedRangeBuilder) WithLowLimit(lowLimit BACnetContextTagUnsignedInteger) BACnetEventParameterUnsignedRangeBuilder {
	b.LowLimit = lowLimit
	return b
}

func (b *_BACnetEventParameterUnsignedRangeBuilder) WithLowLimitBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterUnsignedRangeBuilder {
	builder := builderSupplier(b.LowLimit.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.LowLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterUnsignedRangeBuilder) WithHighLimit(highLimit BACnetContextTagUnsignedInteger) BACnetEventParameterUnsignedRangeBuilder {
	b.HighLimit = highLimit
	return b
}

func (b *_BACnetEventParameterUnsignedRangeBuilder) WithHighLimitBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetEventParameterUnsignedRangeBuilder {
	builder := builderSupplier(b.HighLimit.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.HighLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterUnsignedRangeBuilder) WithClosingTag(closingTag BACnetClosingTag) BACnetEventParameterUnsignedRangeBuilder {
	b.ClosingTag = closingTag
	return b
}

func (b *_BACnetEventParameterUnsignedRangeBuilder) WithClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetEventParameterUnsignedRangeBuilder {
	builder := builderSupplier(b.ClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.ClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetEventParameterUnsignedRangeBuilder) Build() (BACnetEventParameterUnsignedRange, error) {
	if b.OpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'openingTag' not set"))
	}
	if b.TimeDelay == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeDelay' not set"))
	}
	if b.LowLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'lowLimit' not set"))
	}
	if b.HighLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'highLimit' not set"))
	}
	if b.ClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'closingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetEventParameterUnsignedRange.deepCopy(), nil
}

func (b *_BACnetEventParameterUnsignedRangeBuilder) MustBuild() BACnetEventParameterUnsignedRange {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetEventParameterUnsignedRangeBuilder) Done() BACnetEventParameterBuilder {
	return b.parentBuilder
}

func (b *_BACnetEventParameterUnsignedRangeBuilder) buildForBACnetEventParameter() (BACnetEventParameter, error) {
	return b.Build()
}

func (b *_BACnetEventParameterUnsignedRangeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetEventParameterUnsignedRangeBuilder().(*_BACnetEventParameterUnsignedRangeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetEventParameterUnsignedRangeBuilder creates a BACnetEventParameterUnsignedRangeBuilder
func (b *_BACnetEventParameterUnsignedRange) CreateBACnetEventParameterUnsignedRangeBuilder() BACnetEventParameterUnsignedRangeBuilder {
	if b == nil {
		return NewBACnetEventParameterUnsignedRangeBuilder()
	}
	return &_BACnetEventParameterUnsignedRangeBuilder{_BACnetEventParameterUnsignedRange: b.deepCopy()}
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

func (m *_BACnetEventParameterUnsignedRange) GetParent() BACnetEventParameterContract {
	return m.BACnetEventParameterContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterUnsignedRange) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterUnsignedRange) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterUnsignedRange) GetLowLimit() BACnetContextTagUnsignedInteger {
	return m.LowLimit
}

func (m *_BACnetEventParameterUnsignedRange) GetHighLimit() BACnetContextTagUnsignedInteger {
	return m.HighLimit
}

func (m *_BACnetEventParameterUnsignedRange) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterUnsignedRange(structType any) BACnetEventParameterUnsignedRange {
	if casted, ok := structType.(BACnetEventParameterUnsignedRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterUnsignedRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterUnsignedRange) GetTypeName() string {
	return "BACnetEventParameterUnsignedRange"
}

func (m *_BACnetEventParameterUnsignedRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventParameterContract.(*_BACnetEventParameter).getLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits(ctx)

	// Simple field (lowLimit)
	lengthInBits += m.LowLimit.GetLengthInBits(ctx)

	// Simple field (highLimit)
	lengthInBits += m.HighLimit.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventParameterUnsignedRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventParameterUnsignedRange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventParameter) (__bACnetEventParameterUnsignedRange BACnetEventParameterUnsignedRange, err error) {
	m.BACnetEventParameterContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterUnsignedRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterUnsignedRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	openingTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "openingTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(uint8(11))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'openingTag' field"))
	}
	m.OpeningTag = openingTag

	timeDelay, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelay' field"))
	}
	m.TimeDelay = timeDelay

	lowLimit, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "lowLimit", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(1)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lowLimit' field"))
	}
	m.LowLimit = lowLimit

	highLimit, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "highLimit", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'highLimit' field"))
	}
	m.HighLimit = highLimit

	closingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "closingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(uint8(11))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'closingTag' field"))
	}
	m.ClosingTag = closingTag

	if closeErr := readBuffer.CloseContext("BACnetEventParameterUnsignedRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterUnsignedRange")
	}

	return m, nil
}

func (m *_BACnetEventParameterUnsignedRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterUnsignedRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterUnsignedRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterUnsignedRange")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "openingTag", m.GetOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'openingTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "timeDelay", m.GetTimeDelay(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDelay' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "lowLimit", m.GetLowLimit(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'lowLimit' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "highLimit", m.GetHighLimit(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'highLimit' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "closingTag", m.GetClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterUnsignedRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterUnsignedRange")
		}
		return nil
	}
	return m.BACnetEventParameterContract.(*_BACnetEventParameter).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventParameterUnsignedRange) IsBACnetEventParameterUnsignedRange() {}

func (m *_BACnetEventParameterUnsignedRange) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventParameterUnsignedRange) deepCopy() *_BACnetEventParameterUnsignedRange {
	if m == nil {
		return nil
	}
	_BACnetEventParameterUnsignedRangeCopy := &_BACnetEventParameterUnsignedRange{
		m.BACnetEventParameterContract.(*_BACnetEventParameter).deepCopy(),
		utils.DeepCopy[BACnetOpeningTag](m.OpeningTag),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.TimeDelay),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.LowLimit),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.HighLimit),
		utils.DeepCopy[BACnetClosingTag](m.ClosingTag),
	}
	m.BACnetEventParameterContract.(*_BACnetEventParameter)._SubType = m
	return _BACnetEventParameterUnsignedRangeCopy
}

func (m *_BACnetEventParameterUnsignedRange) String() string {
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
