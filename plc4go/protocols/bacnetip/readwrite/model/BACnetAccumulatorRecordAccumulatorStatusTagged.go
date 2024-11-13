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

// BACnetAccumulatorRecordAccumulatorStatusTagged is the corresponding interface of BACnetAccumulatorRecordAccumulatorStatusTagged
type BACnetAccumulatorRecordAccumulatorStatusTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetAccumulatorRecordAccumulatorStatus
	// IsBACnetAccumulatorRecordAccumulatorStatusTagged is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetAccumulatorRecordAccumulatorStatusTagged()
	// CreateBuilder creates a BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder
	CreateBACnetAccumulatorRecordAccumulatorStatusTaggedBuilder() BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder
}

// _BACnetAccumulatorRecordAccumulatorStatusTagged is the data-structure of this message
type _BACnetAccumulatorRecordAccumulatorStatusTagged struct {
	Header BACnetTagHeader
	Value  BACnetAccumulatorRecordAccumulatorStatus

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

var _ BACnetAccumulatorRecordAccumulatorStatusTagged = (*_BACnetAccumulatorRecordAccumulatorStatusTagged)(nil)

// NewBACnetAccumulatorRecordAccumulatorStatusTagged factory function for _BACnetAccumulatorRecordAccumulatorStatusTagged
func NewBACnetAccumulatorRecordAccumulatorStatusTagged(header BACnetTagHeader, value BACnetAccumulatorRecordAccumulatorStatus, tagNumber uint8, tagClass TagClass) *_BACnetAccumulatorRecordAccumulatorStatusTagged {
	if header == nil {
		panic("header of type BACnetTagHeader for BACnetAccumulatorRecordAccumulatorStatusTagged must not be nil")
	}
	return &_BACnetAccumulatorRecordAccumulatorStatusTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder is a builder for BACnetAccumulatorRecordAccumulatorStatusTagged
type BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(header BACnetTagHeader, value BACnetAccumulatorRecordAccumulatorStatus) BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder
	// WithHeader adds Header (property field)
	WithHeader(BACnetTagHeader) BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder
	// WithHeaderBuilder adds Header (property field) which is build by the builder
	WithHeaderBuilder(func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder
	// WithValue adds Value (property field)
	WithValue(BACnetAccumulatorRecordAccumulatorStatus) BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder
	// WithArgTagNumber sets a parser argument
	WithArgTagNumber(uint8) BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder
	// WithArgTagClass sets a parser argument
	WithArgTagClass(TagClass) BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder
	// Build builds the BACnetAccumulatorRecordAccumulatorStatusTagged or returns an error if something is wrong
	Build() (BACnetAccumulatorRecordAccumulatorStatusTagged, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetAccumulatorRecordAccumulatorStatusTagged
}

// NewBACnetAccumulatorRecordAccumulatorStatusTaggedBuilder() creates a BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder
func NewBACnetAccumulatorRecordAccumulatorStatusTaggedBuilder() BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder {
	return &_BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder{_BACnetAccumulatorRecordAccumulatorStatusTagged: new(_BACnetAccumulatorRecordAccumulatorStatusTagged)}
}

type _BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder struct {
	*_BACnetAccumulatorRecordAccumulatorStatusTagged

	err *utils.MultiError
}

var _ (BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder) = (*_BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder)(nil)

func (b *_BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder) WithMandatoryFields(header BACnetTagHeader, value BACnetAccumulatorRecordAccumulatorStatus) BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder {
	return b.WithHeader(header).WithValue(value)
}

func (b *_BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder) WithHeader(header BACnetTagHeader) BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder {
	b.Header = header
	return b
}

func (b *_BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder) WithHeaderBuilder(builderSupplier func(BACnetTagHeaderBuilder) BACnetTagHeaderBuilder) BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder {
	builder := builderSupplier(b.Header.CreateBACnetTagHeaderBuilder())
	var err error
	b.Header, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetTagHeaderBuilder failed"))
	}
	return b
}

func (b *_BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder) WithValue(value BACnetAccumulatorRecordAccumulatorStatus) BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder {
	b.Value = value
	return b
}

func (b *_BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder) WithArgTagNumber(tagNumber uint8) BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder {
	b.TagNumber = tagNumber
	return b
}
func (b *_BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder) WithArgTagClass(tagClass TagClass) BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder {
	b.TagClass = tagClass
	return b
}

func (b *_BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder) Build() (BACnetAccumulatorRecordAccumulatorStatusTagged, error) {
	if b.Header == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'header' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetAccumulatorRecordAccumulatorStatusTagged.deepCopy(), nil
}

func (b *_BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder) MustBuild() BACnetAccumulatorRecordAccumulatorStatusTagged {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

func (b *_BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder) DeepCopy() any {
	_copy := b.CreateBACnetAccumulatorRecordAccumulatorStatusTaggedBuilder().(*_BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetAccumulatorRecordAccumulatorStatusTaggedBuilder creates a BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder
func (b *_BACnetAccumulatorRecordAccumulatorStatusTagged) CreateBACnetAccumulatorRecordAccumulatorStatusTaggedBuilder() BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder {
	if b == nil {
		return NewBACnetAccumulatorRecordAccumulatorStatusTaggedBuilder()
	}
	return &_BACnetAccumulatorRecordAccumulatorStatusTaggedBuilder{_BACnetAccumulatorRecordAccumulatorStatusTagged: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccumulatorRecordAccumulatorStatusTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetAccumulatorRecordAccumulatorStatusTagged) GetValue() BACnetAccumulatorRecordAccumulatorStatus {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetAccumulatorRecordAccumulatorStatusTagged(structType any) BACnetAccumulatorRecordAccumulatorStatusTagged {
	if casted, ok := structType.(BACnetAccumulatorRecordAccumulatorStatusTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccumulatorRecordAccumulatorStatusTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccumulatorRecordAccumulatorStatusTagged) GetTypeName() string {
	return "BACnetAccumulatorRecordAccumulatorStatusTagged"
}

func (m *_BACnetAccumulatorRecordAccumulatorStatusTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetAccumulatorRecordAccumulatorStatusTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccumulatorRecordAccumulatorStatusTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetAccumulatorRecordAccumulatorStatusTagged, error) {
	return BACnetAccumulatorRecordAccumulatorStatusTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetAccumulatorRecordAccumulatorStatusTaggedParseWithBufferProducer(tagNumber uint8, tagClass TagClass) func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccumulatorRecordAccumulatorStatusTagged, error) {
	return func(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccumulatorRecordAccumulatorStatusTagged, error) {
		return BACnetAccumulatorRecordAccumulatorStatusTaggedParseWithBuffer(ctx, readBuffer, tagNumber, tagClass)
	}
}

func BACnetAccumulatorRecordAccumulatorStatusTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetAccumulatorRecordAccumulatorStatusTagged, error) {
	v, err := (&_BACnetAccumulatorRecordAccumulatorStatusTagged{TagNumber: tagNumber, TagClass: tagClass}).parse(ctx, readBuffer, tagNumber, tagClass)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (m *_BACnetAccumulatorRecordAccumulatorStatusTagged) parse(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (__bACnetAccumulatorRecordAccumulatorStatusTagged BACnetAccumulatorRecordAccumulatorStatusTagged, err error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAccumulatorRecordAccumulatorStatusTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccumulatorRecordAccumulatorStatusTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	header, err := ReadSimpleField[BACnetTagHeader](ctx, "header", ReadComplex[BACnetTagHeader](BACnetTagHeaderParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'header' field"))
	}
	m.Header = header

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{Message: "tagnumber doesn't match"})
	}

	value, err := ReadManualField[BACnetAccumulatorRecordAccumulatorStatus](ctx, "value", readBuffer, EnsureType[BACnetAccumulatorRecordAccumulatorStatus](ReadEnumGenericFailing(ctx, readBuffer, header.GetActualLength(), BACnetAccumulatorRecordAccumulatorStatus_NORMAL)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	m.Value = value

	if closeErr := readBuffer.CloseContext("BACnetAccumulatorRecordAccumulatorStatusTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccumulatorRecordAccumulatorStatusTagged")
	}

	return m, nil
}

func (m *_BACnetAccumulatorRecordAccumulatorStatusTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccumulatorRecordAccumulatorStatusTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAccumulatorRecordAccumulatorStatusTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccumulatorRecordAccumulatorStatusTagged")
	}

	if err := WriteSimpleField[BACnetTagHeader](ctx, "header", m.GetHeader(), WriteComplex[BACnetTagHeader](writeBuffer)); err != nil {
		return errors.Wrap(err, "Error serializing 'header' field")
	}

	if err := WriteManualField[BACnetAccumulatorRecordAccumulatorStatus](ctx, "value", func(ctx context.Context) error { return WriteEnumGeneric(ctx, writeBuffer, m.GetValue()) }, writeBuffer); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAccumulatorRecordAccumulatorStatusTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccumulatorRecordAccumulatorStatusTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetAccumulatorRecordAccumulatorStatusTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetAccumulatorRecordAccumulatorStatusTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetAccumulatorRecordAccumulatorStatusTagged) IsBACnetAccumulatorRecordAccumulatorStatusTagged() {
}

func (m *_BACnetAccumulatorRecordAccumulatorStatusTagged) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetAccumulatorRecordAccumulatorStatusTagged) deepCopy() *_BACnetAccumulatorRecordAccumulatorStatusTagged {
	if m == nil {
		return nil
	}
	_BACnetAccumulatorRecordAccumulatorStatusTaggedCopy := &_BACnetAccumulatorRecordAccumulatorStatusTagged{
		utils.DeepCopy[BACnetTagHeader](m.Header),
		m.Value,
		m.TagNumber,
		m.TagClass,
	}
	return _BACnetAccumulatorRecordAccumulatorStatusTaggedCopy
}

func (m *_BACnetAccumulatorRecordAccumulatorStatusTagged) String() string {
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
