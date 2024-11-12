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

// BACnetConstructedDataMaxSegmentsAccepted is the corresponding interface of BACnetConstructedDataMaxSegmentsAccepted
type BACnetConstructedDataMaxSegmentsAccepted interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetMaxSegmentsAccepted returns MaxSegmentsAccepted (property field)
	GetMaxSegmentsAccepted() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataMaxSegmentsAccepted is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMaxSegmentsAccepted()
	// CreateBuilder creates a BACnetConstructedDataMaxSegmentsAcceptedBuilder
	CreateBACnetConstructedDataMaxSegmentsAcceptedBuilder() BACnetConstructedDataMaxSegmentsAcceptedBuilder
}

// _BACnetConstructedDataMaxSegmentsAccepted is the data-structure of this message
type _BACnetConstructedDataMaxSegmentsAccepted struct {
	BACnetConstructedDataContract
	MaxSegmentsAccepted BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataMaxSegmentsAccepted = (*_BACnetConstructedDataMaxSegmentsAccepted)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMaxSegmentsAccepted)(nil)

// NewBACnetConstructedDataMaxSegmentsAccepted factory function for _BACnetConstructedDataMaxSegmentsAccepted
func NewBACnetConstructedDataMaxSegmentsAccepted(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, maxSegmentsAccepted BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMaxSegmentsAccepted {
	if maxSegmentsAccepted == nil {
		panic("maxSegmentsAccepted of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataMaxSegmentsAccepted must not be nil")
	}
	_result := &_BACnetConstructedDataMaxSegmentsAccepted{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaxSegmentsAccepted:           maxSegmentsAccepted,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataMaxSegmentsAcceptedBuilder is a builder for BACnetConstructedDataMaxSegmentsAccepted
type BACnetConstructedDataMaxSegmentsAcceptedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(maxSegmentsAccepted BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMaxSegmentsAcceptedBuilder
	// WithMaxSegmentsAccepted adds MaxSegmentsAccepted (property field)
	WithMaxSegmentsAccepted(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMaxSegmentsAcceptedBuilder
	// WithMaxSegmentsAcceptedBuilder adds MaxSegmentsAccepted (property field) which is build by the builder
	WithMaxSegmentsAcceptedBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataMaxSegmentsAcceptedBuilder
	// Build builds the BACnetConstructedDataMaxSegmentsAccepted or returns an error if something is wrong
	Build() (BACnetConstructedDataMaxSegmentsAccepted, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataMaxSegmentsAccepted
}

// NewBACnetConstructedDataMaxSegmentsAcceptedBuilder() creates a BACnetConstructedDataMaxSegmentsAcceptedBuilder
func NewBACnetConstructedDataMaxSegmentsAcceptedBuilder() BACnetConstructedDataMaxSegmentsAcceptedBuilder {
	return &_BACnetConstructedDataMaxSegmentsAcceptedBuilder{_BACnetConstructedDataMaxSegmentsAccepted: new(_BACnetConstructedDataMaxSegmentsAccepted)}
}

type _BACnetConstructedDataMaxSegmentsAcceptedBuilder struct {
	*_BACnetConstructedDataMaxSegmentsAccepted

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataMaxSegmentsAcceptedBuilder) = (*_BACnetConstructedDataMaxSegmentsAcceptedBuilder)(nil)

func (b *_BACnetConstructedDataMaxSegmentsAcceptedBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataMaxSegmentsAcceptedBuilder) WithMandatoryFields(maxSegmentsAccepted BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMaxSegmentsAcceptedBuilder {
	return b.WithMaxSegmentsAccepted(maxSegmentsAccepted)
}

func (b *_BACnetConstructedDataMaxSegmentsAcceptedBuilder) WithMaxSegmentsAccepted(maxSegmentsAccepted BACnetApplicationTagUnsignedInteger) BACnetConstructedDataMaxSegmentsAcceptedBuilder {
	b.MaxSegmentsAccepted = maxSegmentsAccepted
	return b
}

func (b *_BACnetConstructedDataMaxSegmentsAcceptedBuilder) WithMaxSegmentsAcceptedBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataMaxSegmentsAcceptedBuilder {
	builder := builderSupplier(b.MaxSegmentsAccepted.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.MaxSegmentsAccepted, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataMaxSegmentsAcceptedBuilder) Build() (BACnetConstructedDataMaxSegmentsAccepted, error) {
	if b.MaxSegmentsAccepted == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'maxSegmentsAccepted' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataMaxSegmentsAccepted.deepCopy(), nil
}

func (b *_BACnetConstructedDataMaxSegmentsAcceptedBuilder) MustBuild() BACnetConstructedDataMaxSegmentsAccepted {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataMaxSegmentsAcceptedBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataMaxSegmentsAcceptedBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataMaxSegmentsAcceptedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataMaxSegmentsAcceptedBuilder().(*_BACnetConstructedDataMaxSegmentsAcceptedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataMaxSegmentsAcceptedBuilder creates a BACnetConstructedDataMaxSegmentsAcceptedBuilder
func (b *_BACnetConstructedDataMaxSegmentsAccepted) CreateBACnetConstructedDataMaxSegmentsAcceptedBuilder() BACnetConstructedDataMaxSegmentsAcceptedBuilder {
	if b == nil {
		return NewBACnetConstructedDataMaxSegmentsAcceptedBuilder()
	}
	return &_BACnetConstructedDataMaxSegmentsAcceptedBuilder{_BACnetConstructedDataMaxSegmentsAccepted: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMaxSegmentsAccepted) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMaxSegmentsAccepted) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAX_SEGMENTS_ACCEPTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMaxSegmentsAccepted) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMaxSegmentsAccepted) GetMaxSegmentsAccepted() BACnetApplicationTagUnsignedInteger {
	return m.MaxSegmentsAccepted
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMaxSegmentsAccepted) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetMaxSegmentsAccepted())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMaxSegmentsAccepted(structType any) BACnetConstructedDataMaxSegmentsAccepted {
	if casted, ok := structType.(BACnetConstructedDataMaxSegmentsAccepted); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaxSegmentsAccepted); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMaxSegmentsAccepted) GetTypeName() string {
	return "BACnetConstructedDataMaxSegmentsAccepted"
}

func (m *_BACnetConstructedDataMaxSegmentsAccepted) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (maxSegmentsAccepted)
	lengthInBits += m.MaxSegmentsAccepted.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMaxSegmentsAccepted) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMaxSegmentsAccepted) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMaxSegmentsAccepted BACnetConstructedDataMaxSegmentsAccepted, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaxSegmentsAccepted"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaxSegmentsAccepted")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maxSegmentsAccepted, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maxSegmentsAccepted", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maxSegmentsAccepted' field"))
	}
	m.MaxSegmentsAccepted = maxSegmentsAccepted

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), maxSegmentsAccepted)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaxSegmentsAccepted"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaxSegmentsAccepted")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMaxSegmentsAccepted) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMaxSegmentsAccepted) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaxSegmentsAccepted"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaxSegmentsAccepted")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "maxSegmentsAccepted", m.GetMaxSegmentsAccepted(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maxSegmentsAccepted' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaxSegmentsAccepted"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaxSegmentsAccepted")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMaxSegmentsAccepted) IsBACnetConstructedDataMaxSegmentsAccepted() {}

func (m *_BACnetConstructedDataMaxSegmentsAccepted) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataMaxSegmentsAccepted) deepCopy() *_BACnetConstructedDataMaxSegmentsAccepted {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataMaxSegmentsAcceptedCopy := &_BACnetConstructedDataMaxSegmentsAccepted{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.MaxSegmentsAccepted),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataMaxSegmentsAcceptedCopy
}

func (m *_BACnetConstructedDataMaxSegmentsAccepted) String() string {
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
