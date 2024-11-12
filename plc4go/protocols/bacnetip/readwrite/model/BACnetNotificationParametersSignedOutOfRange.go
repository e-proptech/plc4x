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

// BACnetNotificationParametersSignedOutOfRange is the corresponding interface of BACnetNotificationParametersSignedOutOfRange
type BACnetNotificationParametersSignedOutOfRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetExceedingValue returns ExceedingValue (property field)
	GetExceedingValue() BACnetContextTagSignedInteger
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetDeadband returns Deadband (property field)
	GetDeadband() BACnetContextTagUnsignedInteger
	// GetExceededLimit returns ExceededLimit (property field)
	GetExceededLimit() BACnetContextTagSignedInteger
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
	// IsBACnetNotificationParametersSignedOutOfRange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersSignedOutOfRange()
	// CreateBuilder creates a BACnetNotificationParametersSignedOutOfRangeBuilder
	CreateBACnetNotificationParametersSignedOutOfRangeBuilder() BACnetNotificationParametersSignedOutOfRangeBuilder
}

// _BACnetNotificationParametersSignedOutOfRange is the data-structure of this message
type _BACnetNotificationParametersSignedOutOfRange struct {
	BACnetNotificationParametersContract
	InnerOpeningTag BACnetOpeningTag
	ExceedingValue  BACnetContextTagSignedInteger
	StatusFlags     BACnetStatusFlagsTagged
	Deadband        BACnetContextTagUnsignedInteger
	ExceededLimit   BACnetContextTagSignedInteger
	InnerClosingTag BACnetClosingTag
}

var _ BACnetNotificationParametersSignedOutOfRange = (*_BACnetNotificationParametersSignedOutOfRange)(nil)
var _ BACnetNotificationParametersRequirements = (*_BACnetNotificationParametersSignedOutOfRange)(nil)

// NewBACnetNotificationParametersSignedOutOfRange factory function for _BACnetNotificationParametersSignedOutOfRange
func NewBACnetNotificationParametersSignedOutOfRange(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, innerOpeningTag BACnetOpeningTag, exceedingValue BACnetContextTagSignedInteger, statusFlags BACnetStatusFlagsTagged, deadband BACnetContextTagUnsignedInteger, exceededLimit BACnetContextTagSignedInteger, innerClosingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersSignedOutOfRange {
	if innerOpeningTag == nil {
		panic("innerOpeningTag of type BACnetOpeningTag for BACnetNotificationParametersSignedOutOfRange must not be nil")
	}
	if exceedingValue == nil {
		panic("exceedingValue of type BACnetContextTagSignedInteger for BACnetNotificationParametersSignedOutOfRange must not be nil")
	}
	if statusFlags == nil {
		panic("statusFlags of type BACnetStatusFlagsTagged for BACnetNotificationParametersSignedOutOfRange must not be nil")
	}
	if deadband == nil {
		panic("deadband of type BACnetContextTagUnsignedInteger for BACnetNotificationParametersSignedOutOfRange must not be nil")
	}
	if exceededLimit == nil {
		panic("exceededLimit of type BACnetContextTagSignedInteger for BACnetNotificationParametersSignedOutOfRange must not be nil")
	}
	if innerClosingTag == nil {
		panic("innerClosingTag of type BACnetClosingTag for BACnetNotificationParametersSignedOutOfRange must not be nil")
	}
	_result := &_BACnetNotificationParametersSignedOutOfRange{
		BACnetNotificationParametersContract: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
		InnerOpeningTag:                      innerOpeningTag,
		ExceedingValue:                       exceedingValue,
		StatusFlags:                          statusFlags,
		Deadband:                             deadband,
		ExceededLimit:                        exceededLimit,
		InnerClosingTag:                      innerClosingTag,
	}
	_result.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetNotificationParametersSignedOutOfRangeBuilder is a builder for BACnetNotificationParametersSignedOutOfRange
type BACnetNotificationParametersSignedOutOfRangeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(innerOpeningTag BACnetOpeningTag, exceedingValue BACnetContextTagSignedInteger, statusFlags BACnetStatusFlagsTagged, deadband BACnetContextTagUnsignedInteger, exceededLimit BACnetContextTagSignedInteger, innerClosingTag BACnetClosingTag) BACnetNotificationParametersSignedOutOfRangeBuilder
	// WithInnerOpeningTag adds InnerOpeningTag (property field)
	WithInnerOpeningTag(BACnetOpeningTag) BACnetNotificationParametersSignedOutOfRangeBuilder
	// WithInnerOpeningTagBuilder adds InnerOpeningTag (property field) which is build by the builder
	WithInnerOpeningTagBuilder(func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersSignedOutOfRangeBuilder
	// WithExceedingValue adds ExceedingValue (property field)
	WithExceedingValue(BACnetContextTagSignedInteger) BACnetNotificationParametersSignedOutOfRangeBuilder
	// WithExceedingValueBuilder adds ExceedingValue (property field) which is build by the builder
	WithExceedingValueBuilder(func(BACnetContextTagSignedIntegerBuilder) BACnetContextTagSignedIntegerBuilder) BACnetNotificationParametersSignedOutOfRangeBuilder
	// WithStatusFlags adds StatusFlags (property field)
	WithStatusFlags(BACnetStatusFlagsTagged) BACnetNotificationParametersSignedOutOfRangeBuilder
	// WithStatusFlagsBuilder adds StatusFlags (property field) which is build by the builder
	WithStatusFlagsBuilder(func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetNotificationParametersSignedOutOfRangeBuilder
	// WithDeadband adds Deadband (property field)
	WithDeadband(BACnetContextTagUnsignedInteger) BACnetNotificationParametersSignedOutOfRangeBuilder
	// WithDeadbandBuilder adds Deadband (property field) which is build by the builder
	WithDeadbandBuilder(func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetNotificationParametersSignedOutOfRangeBuilder
	// WithExceededLimit adds ExceededLimit (property field)
	WithExceededLimit(BACnetContextTagSignedInteger) BACnetNotificationParametersSignedOutOfRangeBuilder
	// WithExceededLimitBuilder adds ExceededLimit (property field) which is build by the builder
	WithExceededLimitBuilder(func(BACnetContextTagSignedIntegerBuilder) BACnetContextTagSignedIntegerBuilder) BACnetNotificationParametersSignedOutOfRangeBuilder
	// WithInnerClosingTag adds InnerClosingTag (property field)
	WithInnerClosingTag(BACnetClosingTag) BACnetNotificationParametersSignedOutOfRangeBuilder
	// WithInnerClosingTagBuilder adds InnerClosingTag (property field) which is build by the builder
	WithInnerClosingTagBuilder(func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersSignedOutOfRangeBuilder
	// Build builds the BACnetNotificationParametersSignedOutOfRange or returns an error if something is wrong
	Build() (BACnetNotificationParametersSignedOutOfRange, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetNotificationParametersSignedOutOfRange
}

// NewBACnetNotificationParametersSignedOutOfRangeBuilder() creates a BACnetNotificationParametersSignedOutOfRangeBuilder
func NewBACnetNotificationParametersSignedOutOfRangeBuilder() BACnetNotificationParametersSignedOutOfRangeBuilder {
	return &_BACnetNotificationParametersSignedOutOfRangeBuilder{_BACnetNotificationParametersSignedOutOfRange: new(_BACnetNotificationParametersSignedOutOfRange)}
}

type _BACnetNotificationParametersSignedOutOfRangeBuilder struct {
	*_BACnetNotificationParametersSignedOutOfRange

	parentBuilder *_BACnetNotificationParametersBuilder

	err *utils.MultiError
}

var _ (BACnetNotificationParametersSignedOutOfRangeBuilder) = (*_BACnetNotificationParametersSignedOutOfRangeBuilder)(nil)

func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) setParent(contract BACnetNotificationParametersContract) {
	b.BACnetNotificationParametersContract = contract
}

func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) WithMandatoryFields(innerOpeningTag BACnetOpeningTag, exceedingValue BACnetContextTagSignedInteger, statusFlags BACnetStatusFlagsTagged, deadband BACnetContextTagUnsignedInteger, exceededLimit BACnetContextTagSignedInteger, innerClosingTag BACnetClosingTag) BACnetNotificationParametersSignedOutOfRangeBuilder {
	return b.WithInnerOpeningTag(innerOpeningTag).WithExceedingValue(exceedingValue).WithStatusFlags(statusFlags).WithDeadband(deadband).WithExceededLimit(exceededLimit).WithInnerClosingTag(innerClosingTag)
}

func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) WithInnerOpeningTag(innerOpeningTag BACnetOpeningTag) BACnetNotificationParametersSignedOutOfRangeBuilder {
	b.InnerOpeningTag = innerOpeningTag
	return b
}

func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) WithInnerOpeningTagBuilder(builderSupplier func(BACnetOpeningTagBuilder) BACnetOpeningTagBuilder) BACnetNotificationParametersSignedOutOfRangeBuilder {
	builder := builderSupplier(b.InnerOpeningTag.CreateBACnetOpeningTagBuilder())
	var err error
	b.InnerOpeningTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetOpeningTagBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) WithExceedingValue(exceedingValue BACnetContextTagSignedInteger) BACnetNotificationParametersSignedOutOfRangeBuilder {
	b.ExceedingValue = exceedingValue
	return b
}

func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) WithExceedingValueBuilder(builderSupplier func(BACnetContextTagSignedIntegerBuilder) BACnetContextTagSignedIntegerBuilder) BACnetNotificationParametersSignedOutOfRangeBuilder {
	builder := builderSupplier(b.ExceedingValue.CreateBACnetContextTagSignedIntegerBuilder())
	var err error
	b.ExceedingValue, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) WithStatusFlags(statusFlags BACnetStatusFlagsTagged) BACnetNotificationParametersSignedOutOfRangeBuilder {
	b.StatusFlags = statusFlags
	return b
}

func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) WithStatusFlagsBuilder(builderSupplier func(BACnetStatusFlagsTaggedBuilder) BACnetStatusFlagsTaggedBuilder) BACnetNotificationParametersSignedOutOfRangeBuilder {
	builder := builderSupplier(b.StatusFlags.CreateBACnetStatusFlagsTaggedBuilder())
	var err error
	b.StatusFlags, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetStatusFlagsTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) WithDeadband(deadband BACnetContextTagUnsignedInteger) BACnetNotificationParametersSignedOutOfRangeBuilder {
	b.Deadband = deadband
	return b
}

func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) WithDeadbandBuilder(builderSupplier func(BACnetContextTagUnsignedIntegerBuilder) BACnetContextTagUnsignedIntegerBuilder) BACnetNotificationParametersSignedOutOfRangeBuilder {
	builder := builderSupplier(b.Deadband.CreateBACnetContextTagUnsignedIntegerBuilder())
	var err error
	b.Deadband, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) WithExceededLimit(exceededLimit BACnetContextTagSignedInteger) BACnetNotificationParametersSignedOutOfRangeBuilder {
	b.ExceededLimit = exceededLimit
	return b
}

func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) WithExceededLimitBuilder(builderSupplier func(BACnetContextTagSignedIntegerBuilder) BACnetContextTagSignedIntegerBuilder) BACnetNotificationParametersSignedOutOfRangeBuilder {
	builder := builderSupplier(b.ExceededLimit.CreateBACnetContextTagSignedIntegerBuilder())
	var err error
	b.ExceededLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetContextTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) WithInnerClosingTag(innerClosingTag BACnetClosingTag) BACnetNotificationParametersSignedOutOfRangeBuilder {
	b.InnerClosingTag = innerClosingTag
	return b
}

func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) WithInnerClosingTagBuilder(builderSupplier func(BACnetClosingTagBuilder) BACnetClosingTagBuilder) BACnetNotificationParametersSignedOutOfRangeBuilder {
	builder := builderSupplier(b.InnerClosingTag.CreateBACnetClosingTagBuilder())
	var err error
	b.InnerClosingTag, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetClosingTagBuilder failed"))
	}
	return b
}

func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) Build() (BACnetNotificationParametersSignedOutOfRange, error) {
	if b.InnerOpeningTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerOpeningTag' not set"))
	}
	if b.ExceedingValue == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'exceedingValue' not set"))
	}
	if b.StatusFlags == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'statusFlags' not set"))
	}
	if b.Deadband == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'deadband' not set"))
	}
	if b.ExceededLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'exceededLimit' not set"))
	}
	if b.InnerClosingTag == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'innerClosingTag' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetNotificationParametersSignedOutOfRange.deepCopy(), nil
}

func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) MustBuild() BACnetNotificationParametersSignedOutOfRange {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) Done() BACnetNotificationParametersBuilder {
	return b.parentBuilder
}

func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) buildForBACnetNotificationParameters() (BACnetNotificationParameters, error) {
	return b.Build()
}

func (b *_BACnetNotificationParametersSignedOutOfRangeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetNotificationParametersSignedOutOfRangeBuilder().(*_BACnetNotificationParametersSignedOutOfRangeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetNotificationParametersSignedOutOfRangeBuilder creates a BACnetNotificationParametersSignedOutOfRangeBuilder
func (b *_BACnetNotificationParametersSignedOutOfRange) CreateBACnetNotificationParametersSignedOutOfRangeBuilder() BACnetNotificationParametersSignedOutOfRangeBuilder {
	if b == nil {
		return NewBACnetNotificationParametersSignedOutOfRangeBuilder()
	}
	return &_BACnetNotificationParametersSignedOutOfRangeBuilder{_BACnetNotificationParametersSignedOutOfRange: b.deepCopy()}
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

func (m *_BACnetNotificationParametersSignedOutOfRange) GetParent() BACnetNotificationParametersContract {
	return m.BACnetNotificationParametersContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersSignedOutOfRange) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersSignedOutOfRange) GetExceedingValue() BACnetContextTagSignedInteger {
	return m.ExceedingValue
}

func (m *_BACnetNotificationParametersSignedOutOfRange) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersSignedOutOfRange) GetDeadband() BACnetContextTagUnsignedInteger {
	return m.Deadband
}

func (m *_BACnetNotificationParametersSignedOutOfRange) GetExceededLimit() BACnetContextTagSignedInteger {
	return m.ExceededLimit
}

func (m *_BACnetNotificationParametersSignedOutOfRange) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersSignedOutOfRange(structType any) BACnetNotificationParametersSignedOutOfRange {
	if casted, ok := structType.(BACnetNotificationParametersSignedOutOfRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersSignedOutOfRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersSignedOutOfRange) GetTypeName() string {
	return "BACnetNotificationParametersSignedOutOfRange"
}

func (m *_BACnetNotificationParametersSignedOutOfRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).getLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (exceedingValue)
	lengthInBits += m.ExceedingValue.GetLengthInBits(ctx)

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits(ctx)

	// Simple field (deadband)
	lengthInBits += m.Deadband.GetLengthInBits(ctx)

	// Simple field (exceededLimit)
	lengthInBits += m.ExceededLimit.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersSignedOutOfRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersSignedOutOfRange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParameters, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (__bACnetNotificationParametersSignedOutOfRange BACnetNotificationParametersSignedOutOfRange, err error) {
	m.BACnetNotificationParametersContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersSignedOutOfRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersSignedOutOfRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	innerOpeningTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerOpeningTag' field"))
	}
	m.InnerOpeningTag = innerOpeningTag

	exceedingValue, err := ReadSimpleField[BACnetContextTagSignedInteger](ctx, "exceedingValue", ReadComplex[BACnetContextTagSignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagSignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_SIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'exceedingValue' field"))
	}
	m.ExceedingValue = exceedingValue

	statusFlags, err := ReadSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", ReadComplex[BACnetStatusFlagsTagged](BACnetStatusFlagsTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusFlags' field"))
	}
	m.StatusFlags = statusFlags

	deadband, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "deadband", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deadband' field"))
	}
	m.Deadband = deadband

	exceededLimit, err := ReadSimpleField[BACnetContextTagSignedInteger](ctx, "exceededLimit", ReadComplex[BACnetContextTagSignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagSignedInteger]((uint8)(uint8(3)), (BACnetDataType)(BACnetDataType_SIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'exceededLimit' field"))
	}
	m.ExceededLimit = exceededLimit

	innerClosingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "innerClosingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerClosingTag' field"))
	}
	m.InnerClosingTag = innerClosingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersSignedOutOfRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersSignedOutOfRange")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersSignedOutOfRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersSignedOutOfRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersSignedOutOfRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersSignedOutOfRange")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", m.GetInnerOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerOpeningTag' field")
		}

		if err := WriteSimpleField[BACnetContextTagSignedInteger](ctx, "exceedingValue", m.GetExceedingValue(), WriteComplex[BACnetContextTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'exceedingValue' field")
		}

		if err := WriteSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", m.GetStatusFlags(), WriteComplex[BACnetStatusFlagsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusFlags' field")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "deadband", m.GetDeadband(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deadband' field")
		}

		if err := WriteSimpleField[BACnetContextTagSignedInteger](ctx, "exceededLimit", m.GetExceededLimit(), WriteComplex[BACnetContextTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'exceededLimit' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "innerClosingTag", m.GetInnerClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersSignedOutOfRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersSignedOutOfRange")
		}
		return nil
	}
	return m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersSignedOutOfRange) IsBACnetNotificationParametersSignedOutOfRange() {
}

func (m *_BACnetNotificationParametersSignedOutOfRange) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetNotificationParametersSignedOutOfRange) deepCopy() *_BACnetNotificationParametersSignedOutOfRange {
	if m == nil {
		return nil
	}
	_BACnetNotificationParametersSignedOutOfRangeCopy := &_BACnetNotificationParametersSignedOutOfRange{
		m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).deepCopy(),
		utils.DeepCopy[BACnetOpeningTag](m.InnerOpeningTag),
		utils.DeepCopy[BACnetContextTagSignedInteger](m.ExceedingValue),
		utils.DeepCopy[BACnetStatusFlagsTagged](m.StatusFlags),
		utils.DeepCopy[BACnetContextTagUnsignedInteger](m.Deadband),
		utils.DeepCopy[BACnetContextTagSignedInteger](m.ExceededLimit),
		utils.DeepCopy[BACnetClosingTag](m.InnerClosingTag),
	}
	m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = m
	return _BACnetNotificationParametersSignedOutOfRangeCopy
}

func (m *_BACnetNotificationParametersSignedOutOfRange) String() string {
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
