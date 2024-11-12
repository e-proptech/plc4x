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

// BACnetConstructedDataLocalDate is the corresponding interface of BACnetConstructedDataLocalDate
type BACnetConstructedDataLocalDate interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetLocalDate returns LocalDate (property field)
	GetLocalDate() BACnetApplicationTagDate
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDate
	// IsBACnetConstructedDataLocalDate is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLocalDate()
	// CreateBuilder creates a BACnetConstructedDataLocalDateBuilder
	CreateBACnetConstructedDataLocalDateBuilder() BACnetConstructedDataLocalDateBuilder
}

// _BACnetConstructedDataLocalDate is the data-structure of this message
type _BACnetConstructedDataLocalDate struct {
	BACnetConstructedDataContract
	LocalDate BACnetApplicationTagDate
}

var _ BACnetConstructedDataLocalDate = (*_BACnetConstructedDataLocalDate)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLocalDate)(nil)

// NewBACnetConstructedDataLocalDate factory function for _BACnetConstructedDataLocalDate
func NewBACnetConstructedDataLocalDate(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, localDate BACnetApplicationTagDate, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLocalDate {
	if localDate == nil {
		panic("localDate of type BACnetApplicationTagDate for BACnetConstructedDataLocalDate must not be nil")
	}
	_result := &_BACnetConstructedDataLocalDate{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		LocalDate:                     localDate,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLocalDateBuilder is a builder for BACnetConstructedDataLocalDate
type BACnetConstructedDataLocalDateBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(localDate BACnetApplicationTagDate) BACnetConstructedDataLocalDateBuilder
	// WithLocalDate adds LocalDate (property field)
	WithLocalDate(BACnetApplicationTagDate) BACnetConstructedDataLocalDateBuilder
	// WithLocalDateBuilder adds LocalDate (property field) which is build by the builder
	WithLocalDateBuilder(func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetConstructedDataLocalDateBuilder
	// Build builds the BACnetConstructedDataLocalDate or returns an error if something is wrong
	Build() (BACnetConstructedDataLocalDate, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLocalDate
}

// NewBACnetConstructedDataLocalDateBuilder() creates a BACnetConstructedDataLocalDateBuilder
func NewBACnetConstructedDataLocalDateBuilder() BACnetConstructedDataLocalDateBuilder {
	return &_BACnetConstructedDataLocalDateBuilder{_BACnetConstructedDataLocalDate: new(_BACnetConstructedDataLocalDate)}
}

type _BACnetConstructedDataLocalDateBuilder struct {
	*_BACnetConstructedDataLocalDate

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLocalDateBuilder) = (*_BACnetConstructedDataLocalDateBuilder)(nil)

func (b *_BACnetConstructedDataLocalDateBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataLocalDateBuilder) WithMandatoryFields(localDate BACnetApplicationTagDate) BACnetConstructedDataLocalDateBuilder {
	return b.WithLocalDate(localDate)
}

func (b *_BACnetConstructedDataLocalDateBuilder) WithLocalDate(localDate BACnetApplicationTagDate) BACnetConstructedDataLocalDateBuilder {
	b.LocalDate = localDate
	return b
}

func (b *_BACnetConstructedDataLocalDateBuilder) WithLocalDateBuilder(builderSupplier func(BACnetApplicationTagDateBuilder) BACnetApplicationTagDateBuilder) BACnetConstructedDataLocalDateBuilder {
	builder := builderSupplier(b.LocalDate.CreateBACnetApplicationTagDateBuilder())
	var err error
	b.LocalDate, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagDateBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLocalDateBuilder) Build() (BACnetConstructedDataLocalDate, error) {
	if b.LocalDate == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'localDate' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLocalDate.deepCopy(), nil
}

func (b *_BACnetConstructedDataLocalDateBuilder) MustBuild() BACnetConstructedDataLocalDate {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataLocalDateBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLocalDateBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLocalDateBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLocalDateBuilder().(*_BACnetConstructedDataLocalDateBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLocalDateBuilder creates a BACnetConstructedDataLocalDateBuilder
func (b *_BACnetConstructedDataLocalDate) CreateBACnetConstructedDataLocalDateBuilder() BACnetConstructedDataLocalDateBuilder {
	if b == nil {
		return NewBACnetConstructedDataLocalDateBuilder()
	}
	return &_BACnetConstructedDataLocalDateBuilder{_BACnetConstructedDataLocalDate: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLocalDate) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLocalDate) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOCAL_DATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLocalDate) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLocalDate) GetLocalDate() BACnetApplicationTagDate {
	return m.LocalDate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLocalDate) GetActualValue() BACnetApplicationTagDate {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagDate(m.GetLocalDate())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLocalDate(structType any) BACnetConstructedDataLocalDate {
	if casted, ok := structType.(BACnetConstructedDataLocalDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLocalDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLocalDate) GetTypeName() string {
	return "BACnetConstructedDataLocalDate"
}

func (m *_BACnetConstructedDataLocalDate) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (localDate)
	lengthInBits += m.LocalDate.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLocalDate) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLocalDate) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLocalDate BACnetConstructedDataLocalDate, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLocalDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLocalDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	localDate, err := ReadSimpleField[BACnetApplicationTagDate](ctx, "localDate", ReadComplex[BACnetApplicationTagDate](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagDate](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'localDate' field"))
	}
	m.LocalDate = localDate

	actualValue, err := ReadVirtualField[BACnetApplicationTagDate](ctx, "actualValue", (*BACnetApplicationTagDate)(nil), localDate)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLocalDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLocalDate")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLocalDate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLocalDate) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLocalDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLocalDate")
		}

		if err := WriteSimpleField[BACnetApplicationTagDate](ctx, "localDate", m.GetLocalDate(), WriteComplex[BACnetApplicationTagDate](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'localDate' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLocalDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLocalDate")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLocalDate) IsBACnetConstructedDataLocalDate() {}

func (m *_BACnetConstructedDataLocalDate) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLocalDate) deepCopy() *_BACnetConstructedDataLocalDate {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLocalDateCopy := &_BACnetConstructedDataLocalDate{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagDate](m.LocalDate),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLocalDateCopy
}

func (m *_BACnetConstructedDataLocalDate) String() string {
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
