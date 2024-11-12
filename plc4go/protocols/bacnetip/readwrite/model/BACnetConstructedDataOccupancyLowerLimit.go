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

// BACnetConstructedDataOccupancyLowerLimit is the corresponding interface of BACnetConstructedDataOccupancyLowerLimit
type BACnetConstructedDataOccupancyLowerLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetOccupancyLowerLimit returns OccupancyLowerLimit (property field)
	GetOccupancyLowerLimit() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataOccupancyLowerLimit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataOccupancyLowerLimit()
	// CreateBuilder creates a BACnetConstructedDataOccupancyLowerLimitBuilder
	CreateBACnetConstructedDataOccupancyLowerLimitBuilder() BACnetConstructedDataOccupancyLowerLimitBuilder
}

// _BACnetConstructedDataOccupancyLowerLimit is the data-structure of this message
type _BACnetConstructedDataOccupancyLowerLimit struct {
	BACnetConstructedDataContract
	OccupancyLowerLimit BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataOccupancyLowerLimit = (*_BACnetConstructedDataOccupancyLowerLimit)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataOccupancyLowerLimit)(nil)

// NewBACnetConstructedDataOccupancyLowerLimit factory function for _BACnetConstructedDataOccupancyLowerLimit
func NewBACnetConstructedDataOccupancyLowerLimit(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, occupancyLowerLimit BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOccupancyLowerLimit {
	if occupancyLowerLimit == nil {
		panic("occupancyLowerLimit of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataOccupancyLowerLimit must not be nil")
	}
	_result := &_BACnetConstructedDataOccupancyLowerLimit{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		OccupancyLowerLimit:           occupancyLowerLimit,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataOccupancyLowerLimitBuilder is a builder for BACnetConstructedDataOccupancyLowerLimit
type BACnetConstructedDataOccupancyLowerLimitBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(occupancyLowerLimit BACnetApplicationTagUnsignedInteger) BACnetConstructedDataOccupancyLowerLimitBuilder
	// WithOccupancyLowerLimit adds OccupancyLowerLimit (property field)
	WithOccupancyLowerLimit(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataOccupancyLowerLimitBuilder
	// WithOccupancyLowerLimitBuilder adds OccupancyLowerLimit (property field) which is build by the builder
	WithOccupancyLowerLimitBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataOccupancyLowerLimitBuilder
	// Build builds the BACnetConstructedDataOccupancyLowerLimit or returns an error if something is wrong
	Build() (BACnetConstructedDataOccupancyLowerLimit, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataOccupancyLowerLimit
}

// NewBACnetConstructedDataOccupancyLowerLimitBuilder() creates a BACnetConstructedDataOccupancyLowerLimitBuilder
func NewBACnetConstructedDataOccupancyLowerLimitBuilder() BACnetConstructedDataOccupancyLowerLimitBuilder {
	return &_BACnetConstructedDataOccupancyLowerLimitBuilder{_BACnetConstructedDataOccupancyLowerLimit: new(_BACnetConstructedDataOccupancyLowerLimit)}
}

type _BACnetConstructedDataOccupancyLowerLimitBuilder struct {
	*_BACnetConstructedDataOccupancyLowerLimit

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataOccupancyLowerLimitBuilder) = (*_BACnetConstructedDataOccupancyLowerLimitBuilder)(nil)

func (b *_BACnetConstructedDataOccupancyLowerLimitBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataOccupancyLowerLimitBuilder) WithMandatoryFields(occupancyLowerLimit BACnetApplicationTagUnsignedInteger) BACnetConstructedDataOccupancyLowerLimitBuilder {
	return b.WithOccupancyLowerLimit(occupancyLowerLimit)
}

func (b *_BACnetConstructedDataOccupancyLowerLimitBuilder) WithOccupancyLowerLimit(occupancyLowerLimit BACnetApplicationTagUnsignedInteger) BACnetConstructedDataOccupancyLowerLimitBuilder {
	b.OccupancyLowerLimit = occupancyLowerLimit
	return b
}

func (b *_BACnetConstructedDataOccupancyLowerLimitBuilder) WithOccupancyLowerLimitBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataOccupancyLowerLimitBuilder {
	builder := builderSupplier(b.OccupancyLowerLimit.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.OccupancyLowerLimit, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataOccupancyLowerLimitBuilder) Build() (BACnetConstructedDataOccupancyLowerLimit, error) {
	if b.OccupancyLowerLimit == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'occupancyLowerLimit' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataOccupancyLowerLimit.deepCopy(), nil
}

func (b *_BACnetConstructedDataOccupancyLowerLimitBuilder) MustBuild() BACnetConstructedDataOccupancyLowerLimit {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataOccupancyLowerLimitBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataOccupancyLowerLimitBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataOccupancyLowerLimitBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataOccupancyLowerLimitBuilder().(*_BACnetConstructedDataOccupancyLowerLimitBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataOccupancyLowerLimitBuilder creates a BACnetConstructedDataOccupancyLowerLimitBuilder
func (b *_BACnetConstructedDataOccupancyLowerLimit) CreateBACnetConstructedDataOccupancyLowerLimitBuilder() BACnetConstructedDataOccupancyLowerLimitBuilder {
	if b == nil {
		return NewBACnetConstructedDataOccupancyLowerLimitBuilder()
	}
	return &_BACnetConstructedDataOccupancyLowerLimitBuilder{_BACnetConstructedDataOccupancyLowerLimit: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOccupancyLowerLimit) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOccupancyLowerLimit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OCCUPANCY_LOWER_LIMIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOccupancyLowerLimit) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyLowerLimit) GetOccupancyLowerLimit() BACnetApplicationTagUnsignedInteger {
	return m.OccupancyLowerLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyLowerLimit) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetOccupancyLowerLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOccupancyLowerLimit(structType any) BACnetConstructedDataOccupancyLowerLimit {
	if casted, ok := structType.(BACnetConstructedDataOccupancyLowerLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOccupancyLowerLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOccupancyLowerLimit) GetTypeName() string {
	return "BACnetConstructedDataOccupancyLowerLimit"
}

func (m *_BACnetConstructedDataOccupancyLowerLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (occupancyLowerLimit)
	lengthInBits += m.OccupancyLowerLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOccupancyLowerLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataOccupancyLowerLimit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataOccupancyLowerLimit BACnetConstructedDataOccupancyLowerLimit, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOccupancyLowerLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOccupancyLowerLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	occupancyLowerLimit, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "occupancyLowerLimit", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'occupancyLowerLimit' field"))
	}
	m.OccupancyLowerLimit = occupancyLowerLimit

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), occupancyLowerLimit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOccupancyLowerLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOccupancyLowerLimit")
	}

	return m, nil
}

func (m *_BACnetConstructedDataOccupancyLowerLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOccupancyLowerLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOccupancyLowerLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOccupancyLowerLimit")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "occupancyLowerLimit", m.GetOccupancyLowerLimit(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'occupancyLowerLimit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOccupancyLowerLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOccupancyLowerLimit")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOccupancyLowerLimit) IsBACnetConstructedDataOccupancyLowerLimit() {}

func (m *_BACnetConstructedDataOccupancyLowerLimit) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataOccupancyLowerLimit) deepCopy() *_BACnetConstructedDataOccupancyLowerLimit {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataOccupancyLowerLimitCopy := &_BACnetConstructedDataOccupancyLowerLimit{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.OccupancyLowerLimit),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataOccupancyLowerLimitCopy
}

func (m *_BACnetConstructedDataOccupancyLowerLimit) String() string {
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
