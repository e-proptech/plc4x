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

// BACnetConstructedDataOccupancyUpperLimit is the corresponding interface of BACnetConstructedDataOccupancyUpperLimit
type BACnetConstructedDataOccupancyUpperLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetOccupancyUpperLimit returns OccupancyUpperLimit (property field)
	GetOccupancyUpperLimit() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataOccupancyUpperLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataOccupancyUpperLimit()
	// CreateBuilder creates a BACnetConstructedDataOccupancyUpperLimitBuilder
	CreateBACnetConstructedDataOccupancyUpperLimitBuilder() BACnetConstructedDataOccupancyUpperLimitBuilder
}

// _BACnetConstructedDataOccupancyUpperLimit is the data-structure of this message
type _BACnetConstructedDataOccupancyUpperLimit struct {
	BACnetConstructedDataContract
	OccupancyUpperLimit BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataOccupancyUpperLimit = (*_BACnetConstructedDataOccupancyUpperLimit)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataOccupancyUpperLimit)(nil)

// NewBACnetConstructedDataOccupancyUpperLimit factory function for _BACnetConstructedDataOccupancyUpperLimit
func NewBACnetConstructedDataOccupancyUpperLimit(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, occupancyUpperLimit BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOccupancyUpperLimit {
	if occupancyUpperLimit == nil {
		panic("occupancyUpperLimit of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataOccupancyUpperLimit must not be nil")
	}
	_result := &_BACnetConstructedDataOccupancyUpperLimit{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		OccupancyUpperLimit:           occupancyUpperLimit,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataOccupancyUpperLimitBuilder is a builder for BACnetConstructedDataOccupancyUpperLimit
type BACnetConstructedDataOccupancyUpperLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(occupancyUpperLimit BACnetApplicationTagUnsignedInteger) BACnetConstructedDataOccupancyUpperLimitBuilder
	// WithOccupancyUpperLimit adds OccupancyUpperLimit (property field)
	WithOccupancyUpperLimit(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataOccupancyUpperLimitBuilder
	// WithOccupancyUpperLimitBuilder adds OccupancyUpperLimit (property field) which is build by the builder
	WithOccupancyUpperLimitBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataOccupancyUpperLimitBuilder
	// Done is used to finish work on this child and return (or create one if none) to the parent builder
	Done() BACnetConstructedDataBuilder
	// Build builds the BACnetConstructedDataOccupancyUpperLimit or returns an error if something is wrong
	Build() (BACnetConstructedDataOccupancyUpperLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataOccupancyUpperLimit
}

// NewBACnetConstructedDataOccupancyUpperLimitBuilder() creates a BACnetConstructedDataOccupancyUpperLimitBuilder
func NewBACnetConstructedDataOccupancyUpperLimitBuilder() BACnetConstructedDataOccupancyUpperLimitBuilder {
	return &_BACnetConstructedDataOccupancyUpperLimitBuilder{_BACnetConstructedDataOccupancyUpperLimit: new(_BACnetConstructedDataOccupancyUpperLimit)}
}

type _BACnetConstructedDataOccupancyUpperLimitBuilder struct {
	*_BACnetConstructedDataOccupancyUpperLimit

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataOccupancyUpperLimitBuilder) = (*_BACnetConstructedDataOccupancyUpperLimitBuilder)(nil)

func (b *_BACnetConstructedDataOccupancyUpperLimitBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
	contract.(*_BACnetConstructedData)._SubType = b._BACnetConstructedDataOccupancyUpperLimit
}

func (b *_BACnetConstructedDataOccupancyUpperLimitBuilder) WithMandatoryFields(occupancyUpperLimit BACnetApplicationTagUnsignedInteger) BACnetConstructedDataOccupancyUpperLimitBuilder {
	return b.WithOccupancyUpperLimit(occupancyUpperLimit)
}

func (b *_BACnetConstructedDataOccupancyUpperLimitBuilder) WithOccupancyUpperLimit(occupancyUpperLimit BACnetApplicationTagUnsignedInteger) BACnetConstructedDataOccupancyUpperLimitBuilder {
	b.OccupancyUpperLimit = occupancyUpperLimit
	return b
}

func (b *_BACnetConstructedDataOccupancyUpperLimitBuilder) WithOccupancyUpperLimitBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataOccupancyUpperLimitBuilder {
	builder := builderSupplier(b.OccupancyUpperLimit.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.OccupancyUpperLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataOccupancyUpperLimitBuilder) Build() (BACnetConstructedDataOccupancyUpperLimit, error) {
	if b.OccupancyUpperLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'occupancyUpperLimit' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataOccupancyUpperLimit.deepCopy(), nil
}

func (b *_BACnetConstructedDataOccupancyUpperLimitBuilder) MustBuild() BACnetConstructedDataOccupancyUpperLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetConstructedDataOccupancyUpperLimitBuilder) Done() BACnetConstructedDataBuilder {
	if b.parentBuilder == nil {
		b.parentBuilder = NewBACnetConstructedDataBuilder().(*_BACnetConstructedDataBuilder)
	}
	return b.parentBuilder
}

func (b *_BACnetConstructedDataOccupancyUpperLimitBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataOccupancyUpperLimitBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataOccupancyUpperLimitBuilder().(*_BACnetConstructedDataOccupancyUpperLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataOccupancyUpperLimitBuilder creates a BACnetConstructedDataOccupancyUpperLimitBuilder
func (b *_BACnetConstructedDataOccupancyUpperLimit) CreateBACnetConstructedDataOccupancyUpperLimitBuilder() BACnetConstructedDataOccupancyUpperLimitBuilder {
	if b == nil {
		return NewBACnetConstructedDataOccupancyUpperLimitBuilder()
	}
	return &_BACnetConstructedDataOccupancyUpperLimitBuilder{_BACnetConstructedDataOccupancyUpperLimit: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOccupancyUpperLimit) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OCCUPANCY_UPPER_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOccupancyUpperLimit) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyUpperLimit) GetOccupancyUpperLimit() BACnetApplicationTagUnsignedInteger {
	return m.OccupancyUpperLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyUpperLimit) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetOccupancyUpperLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOccupancyUpperLimit(structType any) BACnetConstructedDataOccupancyUpperLimit {
	if casted, ok := structType.(BACnetConstructedDataOccupancyUpperLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOccupancyUpperLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) GetTypeName() string {
	return "BACnetConstructedDataOccupancyUpperLimit"
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (occupancyUpperLimit)
	lengthInBits += m.OccupancyUpperLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataOccupancyUpperLimit BACnetConstructedDataOccupancyUpperLimit, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOccupancyUpperLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOccupancyUpperLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	occupancyUpperLimit, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "occupancyUpperLimit", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'occupancyUpperLimit' field"))
	}
	m.OccupancyUpperLimit = occupancyUpperLimit

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), occupancyUpperLimit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOccupancyUpperLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOccupancyUpperLimit")
	}

	return m, nil
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOccupancyUpperLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOccupancyUpperLimit")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "occupancyUpperLimit", m.GetOccupancyUpperLimit(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'occupancyUpperLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOccupancyUpperLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOccupancyUpperLimit")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) IsBACnetConstructedDataOccupancyUpperLimit() {}

func (m *_BACnetConstructedDataOccupancyUpperLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) deepCopy() *_BACnetConstructedDataOccupancyUpperLimit {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataOccupancyUpperLimitCopy := &_BACnetConstructedDataOccupancyUpperLimit{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.OccupancyUpperLimit),
	}
	_BACnetConstructedDataOccupancyUpperLimitCopy.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataOccupancyUpperLimitCopy
}

func (m *_BACnetConstructedDataOccupancyUpperLimit) String() string {
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
