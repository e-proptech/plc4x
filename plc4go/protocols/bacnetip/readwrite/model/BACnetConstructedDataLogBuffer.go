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

// BACnetConstructedDataLogBuffer is the corresponding interface of BACnetConstructedDataLogBuffer
type BACnetConstructedDataLogBuffer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetFloorText returns FloorText (property field)
	GetFloorText() []BACnetLogRecord
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// IsBACnetConstructedDataLogBuffer is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLogBuffer()
	// CreateBuilder creates a BACnetConstructedDataLogBufferBuilder
	CreateBACnetConstructedDataLogBufferBuilder() BACnetConstructedDataLogBufferBuilder
}

// _BACnetConstructedDataLogBuffer is the data-structure of this message
type _BACnetConstructedDataLogBuffer struct {
	BACnetConstructedDataContract
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	FloorText            []BACnetLogRecord
}

var _ BACnetConstructedDataLogBuffer = (*_BACnetConstructedDataLogBuffer)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLogBuffer)(nil)

// NewBACnetConstructedDataLogBuffer factory function for _BACnetConstructedDataLogBuffer
func NewBACnetConstructedDataLogBuffer(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, numberOfDataElements BACnetApplicationTagUnsignedInteger, floorText []BACnetLogRecord, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLogBuffer {
	_result := &_BACnetConstructedDataLogBuffer{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		NumberOfDataElements:          numberOfDataElements,
		FloorText:                     floorText,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataLogBufferBuilder is a builder for BACnetConstructedDataLogBuffer
type BACnetConstructedDataLogBufferBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(floorText []BACnetLogRecord) BACnetConstructedDataLogBufferBuilder
	// WithNumberOfDataElements adds NumberOfDataElements (property field)
	WithOptionalNumberOfDataElements(BACnetApplicationTagUnsignedInteger) BACnetConstructedDataLogBufferBuilder
	// WithOptionalNumberOfDataElementsBuilder adds NumberOfDataElements (property field) which is build by the builder
	WithOptionalNumberOfDataElementsBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataLogBufferBuilder
	// WithFloorText adds FloorText (property field)
	WithFloorText(...BACnetLogRecord) BACnetConstructedDataLogBufferBuilder
	// Build builds the BACnetConstructedDataLogBuffer or returns an error if something is wrong
	Build() (BACnetConstructedDataLogBuffer, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataLogBuffer
}

// NewBACnetConstructedDataLogBufferBuilder() creates a BACnetConstructedDataLogBufferBuilder
func NewBACnetConstructedDataLogBufferBuilder() BACnetConstructedDataLogBufferBuilder {
	return &_BACnetConstructedDataLogBufferBuilder{_BACnetConstructedDataLogBuffer: new(_BACnetConstructedDataLogBuffer)}
}

type _BACnetConstructedDataLogBufferBuilder struct {
	*_BACnetConstructedDataLogBuffer

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataLogBufferBuilder) = (*_BACnetConstructedDataLogBufferBuilder)(nil)

func (b *_BACnetConstructedDataLogBufferBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataLogBufferBuilder) WithMandatoryFields(floorText []BACnetLogRecord) BACnetConstructedDataLogBufferBuilder {
	return b.WithFloorText(floorText...)
}

func (b *_BACnetConstructedDataLogBufferBuilder) WithOptionalNumberOfDataElements(numberOfDataElements BACnetApplicationTagUnsignedInteger) BACnetConstructedDataLogBufferBuilder {
	b.NumberOfDataElements = numberOfDataElements
	return b
}

func (b *_BACnetConstructedDataLogBufferBuilder) WithOptionalNumberOfDataElementsBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConstructedDataLogBufferBuilder {
	builder := builderSupplier(b.NumberOfDataElements.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.NumberOfDataElements, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataLogBufferBuilder) WithFloorText(floorText ...BACnetLogRecord) BACnetConstructedDataLogBufferBuilder {
	b.FloorText = floorText
	return b
}

func (b *_BACnetConstructedDataLogBufferBuilder) Build() (BACnetConstructedDataLogBuffer, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataLogBuffer.deepCopy(), nil
}

func (b *_BACnetConstructedDataLogBufferBuilder) MustBuild() BACnetConstructedDataLogBuffer {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataLogBufferBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataLogBufferBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataLogBufferBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataLogBufferBuilder().(*_BACnetConstructedDataLogBufferBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataLogBufferBuilder creates a BACnetConstructedDataLogBufferBuilder
func (b *_BACnetConstructedDataLogBuffer) CreateBACnetConstructedDataLogBufferBuilder() BACnetConstructedDataLogBufferBuilder {
	if b == nil {
		return NewBACnetConstructedDataLogBufferBuilder()
	}
	return &_BACnetConstructedDataLogBufferBuilder{_BACnetConstructedDataLogBuffer: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLogBuffer) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLogBuffer) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOG_BUFFER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLogBuffer) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLogBuffer) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataLogBuffer) GetFloorText() []BACnetLogRecord {
	return m.FloorText
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLogBuffer) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.GetNumberOfDataElements()
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLogBuffer(structType any) BACnetConstructedDataLogBuffer {
	if casted, ok := structType.(BACnetConstructedDataLogBuffer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLogBuffer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLogBuffer) GetTypeName() string {
	return "BACnetConstructedDataLogBuffer"
}

func (m *_BACnetConstructedDataLogBuffer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.FloorText) > 0 {
		for _, element := range m.FloorText {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataLogBuffer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLogBuffer) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLogBuffer BACnetConstructedDataLogBuffer, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLogBuffer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLogBuffer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	zero, err := ReadVirtualField[uint64](ctx, "zero", (*uint64)(nil), uint64(0))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'zero' field"))
	}
	_ = zero

	var numberOfDataElements BACnetApplicationTagUnsignedInteger
	_numberOfDataElements, err := ReadOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer), bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'numberOfDataElements' field"))
	}
	if _numberOfDataElements != nil {
		numberOfDataElements = *_numberOfDataElements
		m.NumberOfDataElements = numberOfDataElements
	}

	floorText, err := ReadTerminatedArrayField[BACnetLogRecord](ctx, "floorText", ReadComplex[BACnetLogRecord](BACnetLogRecordParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'floorText' field"))
	}
	m.FloorText = floorText

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLogBuffer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLogBuffer")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLogBuffer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLogBuffer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLogBuffer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLogBuffer")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		if err := WriteOptionalField[BACnetApplicationTagUnsignedInteger](ctx, "numberOfDataElements", GetRef(m.GetNumberOfDataElements()), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'numberOfDataElements' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "floorText", m.GetFloorText(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'floorText' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLogBuffer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLogBuffer")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLogBuffer) IsBACnetConstructedDataLogBuffer() {}

func (m *_BACnetConstructedDataLogBuffer) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataLogBuffer) deepCopy() *_BACnetConstructedDataLogBuffer {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataLogBufferCopy := &_BACnetConstructedDataLogBuffer{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.NumberOfDataElements),
		utils.DeepCopySlice[BACnetLogRecord, BACnetLogRecord](m.FloorText),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataLogBufferCopy
}

func (m *_BACnetConstructedDataLogBuffer) String() string {
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
