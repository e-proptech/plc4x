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

// BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber is the corresponding interface of BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber
type BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConfirmedServiceRequestReadRangeRange
	// GetReferenceSequenceNumber returns ReferenceSequenceNumber (property field)
	GetReferenceSequenceNumber() BACnetApplicationTagUnsignedInteger
	// GetCount returns Count (property field)
	GetCount() BACnetApplicationTagSignedInteger
	// IsBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber()
	// CreateBuilder creates a BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder
	CreateBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder() BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder
}

// _BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber is the data-structure of this message
type _BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber struct {
	BACnetConfirmedServiceRequestReadRangeRangeContract
	ReferenceSequenceNumber BACnetApplicationTagUnsignedInteger
	Count                   BACnetApplicationTagSignedInteger
}

var _ BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber = (*_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber)(nil)
var _ BACnetConfirmedServiceRequestReadRangeRangeRequirements = (*_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber)(nil)

// NewBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber factory function for _BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber
func NewBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber(peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag, referenceSequenceNumber BACnetApplicationTagUnsignedInteger, count BACnetApplicationTagSignedInteger) *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber {
	if referenceSequenceNumber == nil {
		panic("referenceSequenceNumber of type BACnetApplicationTagUnsignedInteger for BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber must not be nil")
	}
	if count == nil {
		panic("count of type BACnetApplicationTagSignedInteger for BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber must not be nil")
	}
	_result := &_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber{
		BACnetConfirmedServiceRequestReadRangeRangeContract: NewBACnetConfirmedServiceRequestReadRangeRange(peekedTagHeader, openingTag, closingTag),
		ReferenceSequenceNumber:                             referenceSequenceNumber,
		Count:                                               count,
	}
	_result.BACnetConfirmedServiceRequestReadRangeRangeContract.(*_BACnetConfirmedServiceRequestReadRangeRange)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder is a builder for BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber
type BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(referenceSequenceNumber BACnetApplicationTagUnsignedInteger, count BACnetApplicationTagSignedInteger) BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder
	// WithReferenceSequenceNumber adds ReferenceSequenceNumber (property field)
	WithReferenceSequenceNumber(BACnetApplicationTagUnsignedInteger) BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder
	// WithReferenceSequenceNumberBuilder adds ReferenceSequenceNumber (property field) which is build by the builder
	WithReferenceSequenceNumberBuilder(func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder
	// WithCount adds Count (property field)
	WithCount(BACnetApplicationTagSignedInteger) BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder
	// WithCountBuilder adds Count (property field) which is build by the builder
	WithCountBuilder(func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder
	// Build builds the BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber or returns an error if something is wrong
	Build() (BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber
}

// NewBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder() creates a BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder
func NewBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder() BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder {
	return &_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder{_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber: new(_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber)}
}

type _BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder struct {
	*_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber

	parentBuilder *_BACnetConfirmedServiceRequestReadRangeRangeBuilder

	err *utils.MultiError
}

var _ (BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder) = (*_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder)(nil)

func (b *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder) setParent(contract BACnetConfirmedServiceRequestReadRangeRangeContract) {
	b.BACnetConfirmedServiceRequestReadRangeRangeContract = contract
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder) WithMandatoryFields(referenceSequenceNumber BACnetApplicationTagUnsignedInteger, count BACnetApplicationTagSignedInteger) BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder {
	return b.WithReferenceSequenceNumber(referenceSequenceNumber).WithCount(count)
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder) WithReferenceSequenceNumber(referenceSequenceNumber BACnetApplicationTagUnsignedInteger) BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder {
	b.ReferenceSequenceNumber = referenceSequenceNumber
	return b
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder) WithReferenceSequenceNumberBuilder(builderSupplier func(BACnetApplicationTagUnsignedIntegerBuilder) BACnetApplicationTagUnsignedIntegerBuilder) BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder {
	builder := builderSupplier(b.ReferenceSequenceNumber.CreateBACnetApplicationTagUnsignedIntegerBuilder())
	var err error
	b.ReferenceSequenceNumber, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagUnsignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder) WithCount(count BACnetApplicationTagSignedInteger) BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder {
	b.Count = count
	return b
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder) WithCountBuilder(builderSupplier func(BACnetApplicationTagSignedIntegerBuilder) BACnetApplicationTagSignedIntegerBuilder) BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder {
	builder := builderSupplier(b.Count.CreateBACnetApplicationTagSignedIntegerBuilder())
	var err error
	b.Count, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetApplicationTagSignedIntegerBuilder failed"))
	}
	return b
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder) Build() (BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber, error) {
	if b.ReferenceSequenceNumber == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'referenceSequenceNumber' not set"))
	}
	if b.Count == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'count' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber.deepCopy(), nil
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder) MustBuild() BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder) Done() BACnetConfirmedServiceRequestReadRangeRangeBuilder {
	return b.parentBuilder
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder) buildForBACnetConfirmedServiceRequestReadRangeRange() (BACnetConfirmedServiceRequestReadRangeRange, error) {
	return b.Build()
}

func (b *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder().(*_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder creates a BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder
func (b *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) CreateBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder() BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder {
	if b == nil {
		return NewBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder()
	}
	return &_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberBuilder{_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber: b.deepCopy()}
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

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetParent() BACnetConfirmedServiceRequestReadRangeRangeContract {
	return m.BACnetConfirmedServiceRequestReadRangeRangeContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetReferenceSequenceNumber() BACnetApplicationTagUnsignedInteger {
	return m.ReferenceSequenceNumber
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetCount() BACnetApplicationTagSignedInteger {
	return m.Count
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber(structType any) BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber {
	if casted, ok := structType.(BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber"
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConfirmedServiceRequestReadRangeRangeContract.(*_BACnetConfirmedServiceRequestReadRangeRange).getLengthInBits(ctx))

	// Simple field (referenceSequenceNumber)
	lengthInBits += m.ReferenceSequenceNumber.GetLengthInBits(ctx)

	// Simple field (count)
	lengthInBits += m.Count.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConfirmedServiceRequestReadRangeRange) (__bACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber, err error) {
	m.BACnetConfirmedServiceRequestReadRangeRangeContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	referenceSequenceNumber, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "referenceSequenceNumber", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referenceSequenceNumber' field"))
	}
	m.ReferenceSequenceNumber = referenceSequenceNumber

	count, err := ReadSimpleField[BACnetApplicationTagSignedInteger](ctx, "count", ReadComplex[BACnetApplicationTagSignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagSignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'count' field"))
	}
	m.Count = count

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber")
	}

	return m, nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "referenceSequenceNumber", m.GetReferenceSequenceNumber(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'referenceSequenceNumber' field")
		}

		if err := WriteSimpleField[BACnetApplicationTagSignedInteger](ctx, "count", m.GetCount(), WriteComplex[BACnetApplicationTagSignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'count' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber")
		}
		return nil
	}
	return m.BACnetConfirmedServiceRequestReadRangeRangeContract.(*_BACnetConfirmedServiceRequestReadRangeRange).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) IsBACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber() {
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) deepCopy() *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber {
	if m == nil {
		return nil
	}
	_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberCopy := &_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber{
		m.BACnetConfirmedServiceRequestReadRangeRangeContract.(*_BACnetConfirmedServiceRequestReadRangeRange).deepCopy(),
		utils.DeepCopy[BACnetApplicationTagUnsignedInteger](m.ReferenceSequenceNumber),
		utils.DeepCopy[BACnetApplicationTagSignedInteger](m.Count),
	}
	m.BACnetConfirmedServiceRequestReadRangeRangeContract.(*_BACnetConfirmedServiceRequestReadRangeRange)._SubType = m
	return _BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumberCopy
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeBySequenceNumber) String() string {
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
