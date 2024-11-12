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

// BACnetConstructedDataTimeDelayNormal is the corresponding interface of BACnetConstructedDataTimeDelayNormal
type BACnetConstructedDataTimeDelayNormal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetTimeDelayNormal returns TimeDelayNormal (property field)
	GetTimeDelayNormal() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataTimeDelayNormal is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataTimeDelayNormal()
	// CreateBuilder creates a BACnetConstructedDataTimeDelayNormalBuilder
	CreateBACnetConstructedDataTimeDelayNormalBuilder() BACnetConstructedDataTimeDelayNormalBuilder
}

// _BACnetConstructedDataTimeDelayNormal is the data-structure of this message
type _BACnetConstructedDataTimeDelayNormal struct {
	BACnetConstructedDataContract
	TimeDelayNormal BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataTimeDelayNormal = (*_BACnetConstructedDataTimeDelayNormal)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataTimeDelayNormal)(nil)

// NewBACnetConstructedDataTimeDelayNormal factory function for _BACnetConstructedDataTimeDelayNormal
func NewBACnetConstructedDataTimeDelayNormal(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, timeDelayNormal BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimeDelayNormal {
	if timeDelayNormal == nil {
		panic("timeDelayNormal of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataTimeDelayNormal must not be nil")
	}
	_result := &_BACnetConstructedDataTimeDelayNormal{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		TimeDelayNormal:               timeDelayNormal,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataTimeDelayNormalBuilder is a builder for BACnetConstructedDataTimeDelayNormal
type BACnetConstructedDataTimeDelayNormalBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(timeDelayNormal BACnetApplicationTagUnsignedInteger) BACnetConstructedDataTimeDelayNormalBuilder
	// WithTimeDelayNormal adds TimeDelayNormal (property field)
	WithTimeDelayNormal(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataTimeDelayNormalBuilder
	// WithTimeDelayNormalBuilder adds TimeDelayNormal (property field) which is build by the builder
	WithTimeDelayNormalBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataTimeDelayNormalBuilder
	// Build builds the BACnetConstructedDataTimeDelayNormal or returns an error if something is wrong
	Build() (BACnetConstructedDataTimeDelayNormal, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataTimeDelayNormal
}

// NewBACnetConstructedDataTimeDelayNormalBuilder() creates a BACnetConstructedDataTimeDelayNormalBuilder
func NewBACnetConstructedDataTimeDelayNormalBuilder() BACnetConstructedDataTimeDelayNormalBuilder {
	return &_BACnetConstructedDataTimeDelayNormalBuilder{_BACnetConstructedDataTimeDelayNormal: new(_BACnetConstructedDataTimeDelayNormal)}
}

type _BACnetConstructedDataTimeDelayNormalBuilder struct {
	*_BACnetConstructedDataTimeDelayNormal

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataTimeDelayNormalBuilder) = (*_BACnetConstructedDataTimeDelayNormalBuilder)(nil)

func (b *_BACnetConstructedDataTimeDelayNormalBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataTimeDelayNormalBuilder) WithMandatoryFields(timeDelayNormal BACnetApplicationTagUnsignedInteger) BACnetConstructedDataTimeDelayNormalBuilder {
	return b.WithTimeDelayNormal(timeDelayNormal)
}

func (b *_BACnetConstructedDataTimeDelayNormalBuilder) WithTimeDelayNormal(timeDelayNormal BACnetApplicationTagUnsignedInteger) BACnetConstructedDataTimeDelayNormalBuilder {
	b.TimeDelayNormal = timeDelayNormal
	return b
}

func (b *_BACnetConstructedDataTimeDelayNormalBuilder) WithTimeDelayNormalBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataTimeDelayNormalBuilder {
	builder := builderSupplier(b.TimeDelayNormal.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.TimeDelayNormal, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataTimeDelayNormalBuilder) Build() (BACnetConstructedDataTimeDelayNormal, error) {
	if b.TimeDelayNormal == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'timeDelayNormal' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataTimeDelayNormal.deepCopy(), nil
}

func (b *_BACnetConstructedDataTimeDelayNormalBuilder) MustBuild() BACnetConstructedDataTimeDelayNormal {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataTimeDelayNormalBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataTimeDelayNormalBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataTimeDelayNormalBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataTimeDelayNormalBuilder().(*_BACnetConstructedDataTimeDelayNormalBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataTimeDelayNormalBuilder creates a BACnetConstructedDataTimeDelayNormalBuilder
func (b *_BACnetConstructedDataTimeDelayNormal) CreateBACnetConstructedDataTimeDelayNormalBuilder() BACnetConstructedDataTimeDelayNormalBuilder {
	if b == nil {
		return NewBACnetConstructedDataTimeDelayNormalBuilder()
	}
	return &_BACnetConstructedDataTimeDelayNormalBuilder{_BACnetConstructedDataTimeDelayNormal: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimeDelayNormal) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTimeDelayNormal) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_DELAY_NORMAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimeDelayNormal) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimeDelayNormal) GetTimeDelayNormal() BACnetApplicationTagUnsignedInteger {
	return m.TimeDelayNormal
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimeDelayNormal) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetTimeDelayNormal())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimeDelayNormal(structType any) BACnetConstructedDataTimeDelayNormal {
	if casted, ok := structType.(BACnetConstructedDataTimeDelayNormal); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeDelayNormal); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimeDelayNormal) GetTypeName() string {
	return "BACnetConstructedDataTimeDelayNormal"
}

func (m *_BACnetConstructedDataTimeDelayNormal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (timeDelayNormal)
	lengthInBits += m.TimeDelayNormal.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimeDelayNormal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataTimeDelayNormal) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataTimeDelayNormal BACnetConstructedDataTimeDelayNormal, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeDelayNormal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimeDelayNormal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	timeDelayNormal, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "timeDelayNormal", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'timeDelayNormal' field"))
	}
	m.TimeDelayNormal = timeDelayNormal

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), timeDelayNormal)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeDelayNormal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimeDelayNormal")
	}

	return m, nil
}

func (m *_BACnetConstructedDataTimeDelayNormal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimeDelayNormal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeDelayNormal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimeDelayNormal")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "timeDelayNormal", m.GetTimeDelayNormal(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'timeDelayNormal' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeDelayNormal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimeDelayNormal")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimeDelayNormal) IsBACnetConstructedDataTimeDelayNormal() {}

func (m *_BACnetConstructedDataTimeDelayNormal) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataTimeDelayNormal) deepCopy() *_BACnetConstructedDataTimeDelayNormal {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataTimeDelayNormalCopy := &_BACnetConstructedDataTimeDelayNormal{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.TimeDelayNormal),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataTimeDelayNormalCopy
}

func (m *_BACnetConstructedDataTimeDelayNormal) String() string {
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
